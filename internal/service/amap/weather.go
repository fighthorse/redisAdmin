package amap

import (
	"github.com/fighthorse/redisAdmin/component/conf"
	"github.com/fighthorse/redisAdmin/internal/pkg/httpserver"
	"github.com/gin-gonic/gin"
)

func GetIpWeather(c *gin.Context, ip string) interface{} {
	if ip == "" {
		ip = c.ClientIP()
	}
	req := map[string]interface{}{
		"city":       "310106",
		"key":        conf.GConfig.AmapServer.Key,
		"extensions": "base",
		"output":     "JSON",
	}
	result, err := httpserver.Amap.GetWeatherInfo(c.Request.Context(), req)
	if err != nil {
		return err.Error()
	}
	return result
}
