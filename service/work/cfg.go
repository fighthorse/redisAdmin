package work

import (
	"github.com/fighthorse/redisAdmin/pkg/conf"
	"github.com/fighthorse/redisAdmin/pkg/redis"
	"github.com/fighthorse/redisAdmin/pkg/thirdpart/trace_redis"
	"github.com/fighthorse/redisAdmin/protos"
	"github.com/gin-gonic/gin"
)

func AddRedisCfg(c *gin.Context, data protos.AddCfgReq) (interface{}, error) {
	v := conf.Redis{
		Name: data.Name,
		Addr: data.Addr,
		Pwd:  data.Pwd,
	}
	trace_redis.AddCfg(v)

	redis.LoadOthersNew(data.Name)

	d := trace_redis.ListCfg()
	return d, nil
}

func ListRedisCfg(c *gin.Context) (map[string]interface{}, error) {
	d := trace_redis.ListCfg()
	return d, nil
}
