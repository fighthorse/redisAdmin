package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/fighthorse/redisAdmin/component/log"
	"github.com/gin-gonic/gin"
)

func AccessLogging(c *gin.Context) {
	defer func(begin time.Time) {
		requestTime := float64(time.Since(begin)) / float64(time.Second)
		r := c.Request
		url := r.RequestURI
		if FilterUrl(url) {
			return
		}
		accesslog := log.Fields{
			"host":                 strings.Split(r.Host, ":")[0],
			"clientip":             strings.Split(r.RemoteAddr, ":")[0],
			"request_method":       r.Method,
			"request_url":          url,
			"status":               c.Writer.Status(),
			"http_user_agent":      r.UserAgent(),
			"request_time":         requestTime,
			"http_x_forwarded_for": GetIP(r),
		}
		log.AccessLog(r.Context(), accesslog)
	}(time.Now())
	c.Next()
}

// GetIP 获取连接ip
func GetIP(r *http.Request) string {
	// 先从HTTP_X_CLUSTER_CLIENT_IP获取
	ip := r.Header.Get("HTTP_X_CLUSTER_CLIENT_IP")
	if ip == "" {
		ip = r.Header.Get("HTTP_CLIENT_IP")
		if ip == "" {
			ip = r.Header.Get("HTTP_X_FORWARDED_FOR")
			if ip == "" {
				ip = r.Header.Get("X-FORWARDED-FOR")
				if ip == "" {
					ip = strings.Split(r.RemoteAddr, ":")[0]
				}
			}
		}
	}
	return strings.Split(ip, ",")[0]
}
