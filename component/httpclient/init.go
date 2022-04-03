package httpclient

import (
	"github.com/fighthorse/redisAdmin/component/conf"
	"github.com/mitchellh/mapstructure"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	remoteCallErrorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{Name: "remote_call_error_count", Help: "remote call error count"},
		[]string{"url", "status"})
	circuitBreakerCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{Name: "circuit_breaker_count", Help: "circuit breaker count"},
		[]string{"url"})
	remoteCallRequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{Name: "app_remote_call_request_totals", Help: "http remote call request count"},
		[]string{"name", "url"})
	remoteCallCodeErrorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{Name: "app_remote_call_api_code_err", Help: "http remote call api code request err count"},
		[]string{"service", "url", "code", "msg"})
	httpGobreakGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{Name: "circuit_breaker_details", Help: "http circuit breaker details"},
		[]string{"url"})
)

type HttpServer struct {
	SelfServiceName string                 `yaml:"self_service_name"`
	CloseBreaker    bool                   `yaml:"close_breaker"`
	BreakerCfg      map[string]interface{} `yaml:"breaker_cfg"`
	ChildServer     map[string]interface{} `yaml:"child_server"`
}

func init() {
	prometheus.MustRegister(remoteCallErrorCount)
	prometheus.MustRegister(circuitBreakerCount)
	prometheus.MustRegister(remoteCallRequestCount)
	prometheus.MustRegister(remoteCallCodeErrorCount)
	prometheus.MustRegister(httpGobreakGauge)
}

func Init(cfg conf.HttpServer) {
	InitSelfService(cfg.SelfServiceName, cfg.CloseBreaker)
}

// selfName 本身服务名称
// closeBreaker 子服务熔断开启 默认开启 ，true 关闭
func InitSelfService(selfName string, closeBreaker bool) {
	selfServerName = selfName
	isDisableCircuitBreaker = closeBreaker
}

func InitChildService(cfg []map[string]interface{}) {
	for _, v := range cfg {
		item := &Server{}
		if err := mapstructure.Decode(v, item); err == nil {
			childServer[item.Name] = item
		}
	}
}
