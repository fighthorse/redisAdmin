package httpclient

import (
	"sync"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/sony/gobreaker"
)

var (
	circuitBreaker      = map[string]*gobreaker.CircuitBreaker{}
	circuitBreakerMutex = &sync.RWMutex{}
)

type settings struct {
	Name                string  `mapstructure:"name"`
	MaxRequests         uint32  `mapstructure:"maxRequests"`
	Interval            uint32  `mapstructure:"interval"`
	Timeout             uint32  `mapstructure:"timeout"`
	FailureRatio        float64 `mapstructure:"failureRatio"`
	ConsecutiveFailures uint32  `mapstructure:"consecutiveFailures"`
}

// 10s 连续失败100次熔断，1s后继续放量，如果失败立马进入开启状态
var defaultSetting = &settings{MaxRequests: 100, Interval: 10, Timeout: 1, FailureRatio: 1, ConsecutiveFailures: 100}

// InitCircuitBreaker 初始化熔断配置
// {
//    "default":{
//        "maxRequests":1,
//        "interval":1,
//        "timeout":10,
//        "failureRatio":1,
//        "consecutiveFailures":100
//    }
//}
func InitCircuitBreaker(c []map[string]interface{}) {
	// 从config Remote中获取熔断配置
	d := make(map[string]*gobreaker.CircuitBreaker)
	if c == nil {
		return
	}

	for k := range c {
		s := &settings{}
		if err := mapstructure.Decode(k, s); err == nil {
			d[s.Name] = newCircuitBreaker(s.Name, s)
		}
	}
	circuitBreaker = d
	httpGobreakGauge.Reset()
}

// GetCircuitBreaker returns a CircuitBreaker by name
func GetCircuitBreaker(name string) *gobreaker.CircuitBreaker {
	circuitBreakerMutex.RLock()
	cb, ok := circuitBreaker[name]
	if !ok {
		circuitBreakerMutex.RUnlock()
		circuitBreakerMutex.Lock()
		defer circuitBreakerMutex.Unlock()
		// because we released the rlock before we obtained the exclusive lock,
		// we need to double check that some other thread didn't beat us to
		// creation.
		if cb, ok := circuitBreaker[name]; ok {
			return cb
		}
		cb = newCircuitBreaker(name, defaultSetting)
		circuitBreaker[name] = cb
	} else {
		circuitBreakerMutex.RUnlock()
	}

	return cb
}

func newCircuitBreaker(name string, s *settings) *gobreaker.CircuitBreaker {
	return gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:          name,
		MaxRequests:   s.MaxRequests,
		Interval:      time.Duration(s.Interval) * time.Second,
		Timeout:       time.Duration(s.Timeout) * time.Second,
		OnStateChange: onStateChange,
		ReadyToTrip:   readyToTrip(s.FailureRatio, s.ConsecutiveFailures),
	})
}

func readyToTrip(failureRatio float64, consecutiveFailures uint32) func(counts gobreaker.Counts) bool {
	return func(counts gobreaker.Counts) bool {
		ratio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.ConsecutiveFailures >= consecutiveFailures && ratio >= failureRatio
	}
}

func onStateChange(name string, from gobreaker.State, to gobreaker.State) {
	switch to {
	case gobreaker.StateOpen:
		httpGobreakGauge.WithLabelValues(name).Set(1)

	case gobreaker.StateClosed:
		httpGobreakGauge.WithLabelValues(name).Set(0)

	case gobreaker.StateHalfOpen:
		httpGobreakGauge.WithLabelValues(name).Set(2)
	}
}
