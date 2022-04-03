package httpclient

import (
	"net"
	"net/http"
	"strconv"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// 注入请求操作，如http 头注入
type InjectRequest func(*http.Request)

func InjectSidecarHeader(req *http.Request, ServiceName string) {
	req.Header.Set("X-Meshclient", selfServerName)
	req.Header.Set("X-Meshservice", ServiceName)
}

func InjectTrace(name string, tracer opentracing.Tracer, span opentracing.Span, req *http.Request) {
	span.SetOperationName(name)
	ext.HTTPMethod.Set(span, req.Method)
	ext.HTTPUrl.Set(span, req.URL.String())
	host, portString, err := net.SplitHostPort(req.URL.Host)
	if err == nil {
		ext.PeerHostname.Set(span, host)
		if port, err := strconv.Atoi(portString); err != nil {
			ext.PeerPort.Set(span, uint16(port))
		}
	} else {
		ext.PeerHostname.Set(span, req.URL.Host)
	}
	ext.SpanKindRPCClient.Set(span)
	tracer.Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
}
