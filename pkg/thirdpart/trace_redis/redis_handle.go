package trace_redis

import (
	"context"
	"time"

	"github.com/fighthorse/redisAdmin/pkg/log"
	goredis "github.com/go-redis/redis"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	redisErrorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{Name: "redis_error_count", Help: "redis error count"},
		[]string{"schema"})
)

// TODO: We should use redis clusters config map instead of local DefaultMgr!!!
var (
	DefaultMgr *Manager
)

func init() {
	prometheus.MustRegister(redisErrorCount)
	DefaultMgr = NewManager(nil)
}

type RedisClient struct {
	*Client
	schema string
	addr   string
	db     int
}

func (c *RedisClient) handleCmdErr(ctx context.Context, cmd goredis.Cmder) {
	if err := cmd.Err(); err != nil {
		if err == goredis.Nil {
			return
		}

		logContext := log.Fields{
			"schema": c.schema,
			"addr":   c.addr,
			"db":     c.db,
			"args":   cmd.Args(),
			"err":    err.Error(),
		}
		log.Error(ctx, "redis_error", logContext)
		redisErrorCount.WithLabelValues(c.schema).Inc()
	}
}

func (c *RedisClient) Get(ctx context.Context, key string) (string, error) {
	// 使用 trace client
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.Get(key)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) Set(ctx context.Context, key string, iface interface{}, i time.Duration) (string, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.Set(key, iface, i)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) HMGet(ctx context.Context, key string, vals ...string) ([]interface{}, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.HMGet(key, vals...)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) HMSet(ctx context.Context, key string, mdata map[string]interface{}) (string, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.HMSet(key, mdata)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) Del(ctx context.Context, vals ...string) (int64, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.Del(vals...)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) Exists(ctx context.Context, vals ...string) (int64, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.Exists(vals...)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) Expire(ctx context.Context, key string, i time.Duration) (bool, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.Expire(key, i)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) SIsMember(ctx context.Context, key string, iface interface{}) (bool, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.SIsMember(key, iface)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) SAdd(ctx context.Context, key string, vals ...interface{}) (int64, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.SAdd(key, vals...)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) SDiff(ctx context.Context, vals ...string) ([]string, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.SDiff(vals...)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) SetBit(ctx context.Context, key string, i int64, i1 int) (int64, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.SetBit(key, i, i1)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) GetBit(ctx context.Context, key string, i int64) (int64, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.GetBit(key, i)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) LPush(ctx context.Context, key string, vals ...interface{}) (int64, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.LPush(key, vals...)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) RPop(ctx context.Context, key string) (string, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.RPop(key)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) Incr(ctx context.Context, key string) (int64, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.Incr(key)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) IncrBy(ctx context.Context, key string, val int64) (int64, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.IncrBy(key, val)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) HDel(ctx context.Context, key string, vals ...string) (int64, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.HDel(key, vals...)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) HExists(ctx context.Context, key string, val string) (bool, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.HExists(key, val)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) HGet(ctx context.Context, key string, val string) (string, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.HGet(key, val)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) HSet(ctx context.Context, key string, val string, iface interface{}) (bool, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.HSet(key, val, iface)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) TTL(ctx context.Context, key string) (time.Duration, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.TTL(key)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.HGetAll(key)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) SRandMember(ctx context.Context, key string) (string, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.SRandMember(key)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) SRandMemberN(ctx context.Context, key string, i int64) ([]string, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.SRandMemberN(key, i)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) SCard(ctx context.Context, key string) (int64, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.SCard(key)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) MGet(ctx context.Context, vals ...string) ([]interface{}, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.MGet(vals...)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) MSet(ctx context.Context, vals ...interface{}) (string, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.MSet(vals...)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) BitCount(ctx context.Context, key string, pos ...int64) (int64, error) {
	var bc *goredis.BitCount
	if len(pos) == 2 {
		bc = &goredis.BitCount{
			Start: pos[0],
			End:   pos[1],
		}
	}
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.BitCount(key, bc)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) Decr(ctx context.Context, key string) (int64, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.Decr(key)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) HIncrBy(ctx context.Context, key string, val string, i int64) (int64, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.HIncrBy(key, val, i)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) HLen(ctx context.Context, key string) (int64, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.HLen(key)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) LLen(ctx context.Context, key string) (int64, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.LLen(key)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) LPop(ctx context.Context, key string) (string, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.LPop(key)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) RPush(ctx context.Context, key string, vals ...interface{}) (int64, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.RPush(key, vals...)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) SetNX(ctx context.Context, key string, iface interface{}, i time.Duration) (bool, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.SetNX(key, iface, i)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) SMembers(ctx context.Context, key string) ([]string, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.SMembers(key)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) SRem(ctx context.Context, key string, vals ...interface{}) (int64, error) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	cmd := tc.SRem(key, vals...)
	c.handleCmdErr(ctx, cmd)

	return cmd.Result()
}

func (c *RedisClient) Pipeline(ctx context.Context) (iface goredis.Pipeliner) {
	tc := c.Client.Trace(ctx, opentracing.GlobalTracer())
	return tc.Pipeline()
}
