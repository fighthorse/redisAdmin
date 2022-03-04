package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	apiCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_reward_count", Help: "user event stats"},
		[]string{"app", "msg", "code"})
)

func IncApiCount(app, msg, code string) {
	apiCount.WithLabelValues([]string{app, msg, code}...).Inc()
}
