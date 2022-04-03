package amap

import "github.com/gin-gonic/gin"

func RegisterHttp(r *gin.Engine) {

	r.GET("/amap/weather", Weather)

}
