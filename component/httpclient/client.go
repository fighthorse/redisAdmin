package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/fighthorse/redisAdmin/component/log"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	stdgobreaker "github.com/sony/gobreaker"
)

type Server struct {
	Name                 string  `mapstructure:"name"`
	Url                  string  `mapstructure:"url"`
	DiscoveryServiceName string  `mapstructure:"discovery_service_name"`
	DiscoveryTag         string  `mapstructure:"discovery_tag"`
	Timeout              float64 `mapstructure:"timeout"`
}

type Client struct {
	client               *http.Client
	host                 string
	name                 string
	discoveryServiceName string
	discoveryTag         string
	discoveryDC          string
}

// Error represents a http error.
type httpError struct {
	code int
	text string
}

// Code returns the http error code.
func (e *httpError) Code() int {
	return e.code
}

// Error returns the error message in string format.
func (e *httpError) Error() string {
	return e.text
}

type contentType int

const (
	urlencodedType contentType = iota
	jsonType
	xmlType
)

func (c *Client) doRequest(ctx context.Context, name string, req *http.Request) (resp *http.Response, err error) {
	// 创建客户端并发送请求
	client := c.client
	if false {
		client = http.DefaultClient
		client.Timeout = 30 * time.Second
	}

	// 获取超时配置
	timeout := getTimeout(ctx, c.name, name)
	if timeout > 0 {
		client.Timeout = 600 * time.Millisecond
	}

	return client.Do(req)
}

// Send 发送http请求
// reqBody:
//     map[string]interface{}, key为 __json__ 时代表发送json类型，key为 __xml__ 代表发送xml类型
func (c *Client) send(ctx context.Context, method string, reqPath string, reqBody map[string]interface{}, result interface{}, irs ...InjectRequest) (err error) {
	isAbsURI := strings.HasPrefix(reqPath, "http://") || strings.HasPrefix(reqPath, "https://")

	// 根据服务名获取ip
	reqUrl := reqPath
	if !isAbsURI {
		reqUrl = c.host + reqPath
	}

	req, err := newRequest(ctx, method, reqUrl, reqBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 以子服务名+path组成新唯一名称 如 auth + /info
	name := c.name + reqPath
	if isAbsURI {
		name = c.name + req.URL.Path
	}

	for _, o := range irs {
		o(req)
	}

	// 注入sidecar调用头
	if len(c.discoveryServiceName) > 0 {
		InjectSidecarHeader(req, c.discoveryServiceName)
	}

	// 注入trace span, 从请求中获取span.context
	if ctx != nil {
		if span := opentracing.SpanFromContext(ctx); span != nil {
			tracer := opentracing.GlobalTracer()
			span := tracer.StartSpan("call_remote_server", opentracing.ChildOf(span.Context()))
			defer func() {
				// 错误响应设置trace状态
				if err != nil {
					ext.Error.Set(span, true)
					span.LogKV("event", "error", "error.kind", "internal error", "message", err.Error())
					if e, ok := err.(*httpError); ok {
						ext.HTTPStatusCode.Set(span, uint16(e.Code()))
					}
				} else {
					ext.HTTPStatusCode.Set(span, 200)
				}
				span.Finish()
			}()

			// Add standard OpenTracing tags.
			// 添加trace请求注入
			InjectTrace(name, tracer, span, req)
		}
	}

	var (
		bodyBytes []byte
	)

	resp, err := c.doRequest(ctx, name, req)
	if err != nil {
		err = &httpError{
			code: 500,
			text: err.Error(),
		}
		goto ERREND
	}
	defer resp.Body.Close()

	// 记录请求调用次数
	remoteCallRequestCount.WithLabelValues(c.name, name).Inc()

	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = &httpError{
			code: resp.StatusCode,
			text: err.Error(),
		}
		goto ERREND
	}

	if resp.StatusCode > http.StatusPartialContent {
		err = &httpError{
			code: resp.StatusCode,
			text: string(bodyBytes),
		}
		goto ERREND
	}

	// 解码响应
	err = decodeResponse(bodyBytes, result, c.discoveryServiceName, reqPath)
	if err != nil {
		err = &httpError{
			code: resp.StatusCode,
			text: err.Error(),
		}
		goto ERREND
	}

	return

ERREND:
	// 记录子服务调用失败监控
	{
		var values []string
		if e, ok := err.(*httpError); ok {
			values = []string{name, strconv.Itoa(e.Code())}
		} else {
			values = []string{name, "200"}
		}
		remoteCallErrorCount.WithLabelValues(values...).Inc()
	}

	// 记录子服务调用失败日志
	{
		logContext := make(map[string]interface{}, 1)
		logContext["url"] = reqUrl
		logContext["name"] = c.name

		if e, ok := err.(*httpError); ok {
			logContext["status"] = e.Code()
		}

		logContext["err"] = err.Error()
		log.Error(ctx, "rpc_non_successful", logContext)
	}

	// 设置错误响应
	if result != nil {
		t := reflect.TypeOf(result)
		if t.Kind() == reflect.Ptr {
			v := reflect.ValueOf(result).Elem()
			if v.Kind() == reflect.Struct {
				var code int
				var msg string
				if e, ok := err.(*httpError); ok {
					code = e.Code()
					msg = e.Error()
				} else {
					code = 500
					msg = err.Error()
				}
				codeField := v.FieldByName("Code")
				if codeField.CanSet() {
					codeField.Set(reflect.ValueOf(code))
				}
				msgField := v.FieldByName("Message")
				if msgField.CanSet() {
					msgField.Set(reflect.ValueOf(msg))
				}
			}
		}
	}
	return
}

func (c *Client) Send(ctx context.Context, method string, path string, reqBody map[string]interface{}, result interface{}, irs ...InjectRequest) (err error) {
	f := func() (interface{}, error) {
		err = c.send(ctx, method, path, reqBody, result, irs...)
		return nil, err
	}

	// 以子服务名+path组成新唯一名称 如 auth + /info
	name := c.name + path

	// 获取熔断禁用状态，默认开启熔断
	if isDisableCircuitBreaker {
		_, err = f()
	} else {
		cb := GetCircuitBreaker(name)
		_, err = cb.Execute(f)
		// 熔断开启
		if err == stdgobreaker.ErrOpenState || err == stdgobreaker.ErrTooManyRequests {
			// 记录熔断次数
			circuitBreakerCount.WithLabelValues(name).Inc()

			logContext := make(map[string]interface{}, 1)
			logContext["url"] = name
			logContext["err"] = err.Error()
			log.Error(ctx, "rpc_non_successful", logContext)
			// 熔断code返回-10002
			if result != nil {
				t := reflect.TypeOf(result)
				if t.Kind() == reflect.Ptr {
					v := reflect.ValueOf(result).Elem()
					if v.Kind() == reflect.Struct {
						codeField := v.FieldByName("Code")
						if codeField.CanSet() {
							codeField.Set(reflect.ValueOf(-10002))
						}
						msgField := v.FieldByName("Message")
						if msgField.CanSet() {
							msgField.Set(reflect.ValueOf("服务器繁忙，请稍后再试"))
						}
					}
				}
			}
		}
	}
	return
}

func getContentType(in map[string]interface{}) contentType {
	if _, ok := in["__json__"]; ok {
		return jsonType
	}
	if _, ok := in["__xml__"]; ok {
		return xmlType
	}
	return urlencodedType
}

func newRequest(ctx context.Context, method string, path string, in map[string]interface{}) (req *http.Request, err error) {
	// 编码请求
	t := getContentType(in)
	data := encodeRequest(in, t)

	if strings.ToUpper(method) == "GET" {
		req, err = http.NewRequest(method, path, nil)
		if err != nil {
			return
		}
		req.URL.RawQuery = data

	} else {
		req, err = http.NewRequest(method, path, strings.NewReader(data))
		if err != nil {
			return
		}
	}

	// inject context
	if ctx != nil {
		*req = *req.WithContext(ctx)
	}

	switch t {
	case urlencodedType:
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	case jsonType:
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	case xmlType:
		req.Header.Set("Content-Type", "application/xml; charset=utf-8")
	}

	return
}

func encodeRequest(in map[string]interface{}, t contentType) (request string) {
	switch t {
	case urlencodedType:
		params := make(url.Values)
		for k, v := range in {
			params.Add(k, interface2String(v))
		}
		request = params.Encode()
	case jsonType:
		b, _ := json.Marshal(in["__json__"])
		request = string(b)
	case xmlType:
		request, _ = in["__xml__"].(string)
	}
	return
}

// out 为结构体或文本
func decodeResponse(bodyBytes []byte, out interface{}, serviceName, path string) (err error) {
	if out == nil {
		return
	}
	// 如果 out 是结构体
	v := reflect.Indirect(reflect.ValueOf(out))
	if v.Kind() == reflect.Struct {
		err := json.Unmarshal(bodyBytes, out)
		if err != nil {
			return err
		}
		if v.FieldByName("Code").Kind() == reflect.Int {
			code := v.FieldByName("Code").Int()
			msg := v.FieldByName("Message").String()
			if code != 0 {
				remoteCallCodeErrorCount.WithLabelValues(serviceName, path, fmt.Sprintf("%d", code), msg).Inc()
			}
		}
	} else {
		o, _ := out.(*[]byte)
		*o = bodyBytes
	}
	return
}

func interface2String(v interface{}) string {
	s := ""
	switch t := v.(type) {
	case map[string]interface{}, []int:
		b, _ := json.Marshal(v)
		s = string(b)
	case bool:
		if t {
			return "1"
		} else {
			return "0"
		}
	default:
		s = fmt.Sprintf("%v", v)
	}
	return s
}

func getTimeout(ctx context.Context, serverName string, name string) time.Duration {
	// 获取超时设置, 默认1s超时
	var t float64 = 1
	// 先获取特性配置，没有获取默认值
	var ok bool
	if t, ok = timeOutCfg[name]; !ok {
		// 获取子模块默认超时
		t = 0.6
	}
	return time.Duration(t*1000) * time.Millisecond
}
