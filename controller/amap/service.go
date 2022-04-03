package amap

import (
	"github.com/fighthorse/redisAdmin/internal/service/amap"
	"github.com/gin-gonic/gin"
)

func Weather(c *gin.Context) {
	ip, _ := c.GetQuery("location_ip")
	data := amap.GetIpWeather(c, ip)
	c.JSON(200, gin.H{"code": 0, "message": "ok", "data": data})
}
