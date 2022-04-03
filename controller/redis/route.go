package redis

import (
	"github.com/fighthorse/redisAdmin/component/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterHttp(r *gin.Engine) {

	redis := r.Group("/redis")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	redis.Use(middleware.TokenRequired)
	{
		redis.POST("/init", SearchInit)
		redis.POST("/search", Search)
		redis.GET("/searchKey", SearchKey)
		redis.POST("/searchKey", SearchKey)
		redis.POST("/searchNowKey", SearchNowKey)
		redis.GET("/info", Info)
		redis.POST("/handle", Handle)
		redis.POST("/addCfg", AddCfg)
		redis.POST("/getKey", GetKey)
	}

}
