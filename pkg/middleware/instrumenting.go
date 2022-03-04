package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	httpRequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{Name: "app_request_totals", Help: "http request count"},
		[]string{"url", "status"})
	httpRequestDuration = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "app_request_duration_milliseconds",
			Help:       "http request duration",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"url", "status"})

	httpSlowCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{Name: "app_slow_request_totals", Help: "http slow request count"},
		[]string{"url"})

	httpRequestCountWithoutUrl = prometheus.NewCounterVec(
		prometheus.CounterOpts{Name: "app_request_all_totals", Help: "http request count without url"},
		[]string{"client"})
)

func init() {
	prometheus.MustRegister(httpRequestCount)
	prometheus.MustRegister(httpRequestDuration)
	prometheus.MustRegister(httpSlowCount)
	prometheus.MustRegister(httpRequestCountWithoutUrl)
}

func Instrument(c *gin.Context) {
	defer func(begin time.Time) {
		r := c.Request
		url := r.RequestURI
		if FilterUrl(url) {
			return
		}
		w := c.Writer
		values := []string{r.URL.Path, strconv.Itoa(w.Status())}
		httpRequestCount.WithLabelValues(values...).Inc()
		timeElapsed := float64(time.Since(begin)) / float64(time.Millisecond)
		// 超过1s需要记录慢请求
		if timeElapsed > 1000 {
			httpSlowCount.WithLabelValues(r.URL.Path).Inc()
		}
		httpRequestDuration.WithLabelValues(values...).Observe(timeElapsed)
		httpRequestCountWithoutUrl.WithLabelValues(GetMeshClient(r)).Inc()
	}(time.Now())

	c.Next()
}

func GetMeshClient(r *http.Request) string {
	mc := r.Header.Get("x-client")
	if mc == "" {
		mc = "/"
	}
	return mc
}
