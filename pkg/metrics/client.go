package metrics

import "github.com/prometheus/client_golang/prometheus"

func init() {
	//---float---//
	prometheus.MustRegister(apiCount)
}
