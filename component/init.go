package component

import (
	"github.com/fighthorse/redisAdmin/component/conf"
	"github.com/fighthorse/redisAdmin/component/httpclient"
	"github.com/fighthorse/redisAdmin/component/log"
	"github.com/fighthorse/redisAdmin/component/thirdpart/trace_redis"
	"github.com/fighthorse/redisAdmin/component/trace"
)

func InitComponent() {
	// redis cfg
	trace_redis.InitCfg(conf.GConfig.Redis)
	// http cfg
	httpclient.Init(conf.GConfig.HttpServer)
	httpclient.InitCircuitBreaker(conf.GConfig.HttpBreaker)
	httpclient.InitChildService(conf.GConfig.ChildServer)
	//trace
	trace.Init()
	// log
	log.Init()
}
