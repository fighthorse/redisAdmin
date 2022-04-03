package controller

import (
	"github.com/fighthorse/redisAdmin/controller/amap"
	"github.com/fighthorse/redisAdmin/controller/hc"
	"github.com/fighthorse/redisAdmin/controller/login"
	"github.com/fighthorse/redisAdmin/controller/redis"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) error {

	// 标记状态
	r.Static("/assets", "./assets")
	r.StaticFile("/favicon.ico", "./assets/favicon.ico")

	// 注册路由 服务
	hc.RegisterHttp(r)

	redis.RegisterHttp(r)
	login.RegisterHttp(r)
	amap.RegisterHttp(r)

	return nil
}
