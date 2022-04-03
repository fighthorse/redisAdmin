package hc

import "github.com/gin-gonic/gin"

func RegisterHttp(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "OK")
	})
	r.GET("/index", Index)
	r.GET("/hc", Hc)
	r.GET("/metrics", Metrics)
	r.GET("/service", Service)
	r.GET("/reload", Reload)
	r.GET("/stop", Stop)
	r.GET("/pid", Pid)
}
