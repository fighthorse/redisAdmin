package redis

import (
	"context"
	"time"

	"github.com/fighthorse/redisAdmin/pkg/thirdpart/trace_redis"
)

type redisInstance struct {
	Client *trace_redis.RedisClient
	name   string
}

func (r *redisInstance) Select(ctx context.Context, db int) (*redisInstance, error) {
	cfg := &redisInstance{
		Client: trace_redis.NewClientDb(r.name, db),
		name:   r.name,
	}
	return cfg, nil
}

func (r *redisInstance) GetKey(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key)
}

func (r *redisInstance) HMSetKey(ctx context.Context, key string, mdata map[string]interface{}) (string, error) {
	return r.Client.HMSet(ctx, key, mdata)
}

func (r *redisInstance) HMGetKey(ctx context.Context, key string, vals ...string) ([]interface{}, error) {
	return r.Client.HMGet(ctx, key, vals...)
}

func (r *redisInstance) SIsMember(ctx context.Context, key string, iface interface{}) (bool, error) {
	return r.Client.SIsMember(ctx, key, iface)
}

func (r *redisInstance) SAdd(ctx context.Context, key string, vals ...interface{}) (int64, error) {
	return r.Client.SAdd(ctx, key, vals...)
}

func (r *redisInstance) SDiff(ctx context.Context, vals ...string) ([]string, error) {
	return r.Client.SDiff(ctx, vals...)
}

func (r *redisInstance) SetBit(ctx context.Context, key string, i int64, i1 int) (int64, error) {
	return r.Client.SetBit(ctx, key, i, i1)
}

func (r *redisInstance) LPush(ctx context.Context, key string, vals ...interface{}) (int64, error) {
	return r.Client.LPush(ctx, key, vals...)
}

func (r *redisInstance) Incr(ctx context.Context, key string) (int64, error) {
	return r.Client.Incr(ctx, key)
}

func (r *redisInstance) HDel(ctx context.Context, key string, vals ...string) (int64, error) {
	return r.Client.HDel(ctx, key, vals...)
}

func (r *redisInstance) HExists(ctx context.Context, key string, val string) (bool, error) {
	return r.Client.HExists(ctx, key, val)
}

func (r *redisInstance) HGet(ctx context.Context, key string, val string) (string, error) {
	return r.Client.HGet(ctx, key, val)
}

func (r *redisInstance) HSet(ctx context.Context, key string, val string, iface interface{}) (bool, error) {
	return r.Client.HSet(ctx, key, val, iface)
}
func (r *redisInstance) TTL(ctx context.Context, key string) (time.Duration, error) {
	return r.Client.TTL(ctx, key)
}

func (r *redisInstance) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return r.Client.HGetAll(ctx, key)
}
