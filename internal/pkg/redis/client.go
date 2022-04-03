package redis

import (
	"context"
	"fmt"

	"github.com/fighthorse/redisAdmin/component/thirdpart/trace_redis"
)

var (
	Others = map[string]*redisInstance{}
)

func Init() {
	LoadOthersNew("base")
}

func Test() {
	_, _ = LoadOthersDB("base", 0).Client.Get(context.Background(), "test")
}

func LoadOthersNew(name string) {
	cfg := &redisInstance{}
	cfg.name = name
	cfg.Client = trace_redis.NewClient(cfg.name)
	Others[name] = cfg
}

func LoadOthersDB(name string, db int) *redisInstance {
	dbZero, ok := Others[name]
	if !ok {
		return nil
	}
	if db == 0 {
		return dbZero
	}
	nameNew := fmt.Sprintf("%s_%d", name, db)
	if ll, ok := Others[nameNew]; ok {
		return ll
	}
	ll, err := dbZero.Select(context.Background(), db)
	if err != nil {
		return nil
	}
	Others[nameNew] = ll
	return ll
}
