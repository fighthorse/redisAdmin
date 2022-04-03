package httpclient

import (
	"net/http"
	"sync"
)

var (
	newOnce   sync.Once
	newClient *http.Client
)

var (
	selfServerName          string // 本身服务名称
	isDisableCircuitBreaker bool   // 子服务熔断开启 默认开启 ，true 关闭
	timeOutCfg              = map[string]float64{}
	childServer             = map[string]*Server{}
)

func GetSelfServiceName() string {
	return selfServerName
}

func New(name string) *Client {
	// transport TODO
	newOnce.Do(func() {
		newClient = &http.Client{
			Transport: http.DefaultTransport,
		}
	})

	cfg, ok := childServer[name]
	if !ok {
		panic("http_server <" + name + "> service cfg not found! ")
	}

	timeOutCfg[name] = cfg.Timeout

	return &Client{
		client:               newClient,
		host:                 cfg.Url,
		name:                 name,
		discoveryServiceName: cfg.DiscoveryServiceName,
		discoveryTag:         cfg.DiscoveryTag,
	}
}
