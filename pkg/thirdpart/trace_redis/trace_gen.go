package trace_redis

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// A TraceClient wraps redis with tracer.
type TraceClient struct {
	*Client

	tracer          opentracing.Tracer
	spanCtx         opentracing.SpanContext
	includeNotFound bool
}

func (c *TraceClient) Append(key string, val string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_append", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.Append(key, val)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("Append %v %v", key, val))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("Append %v", key))
	}

	return
}

func (c *TraceClient) BLPop(i time.Duration, vals ...string) (cmd *redis.StringSliceCmd) {
	span := c.tracer.StartSpan("redis_blpop", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.BLPop(i, vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("BLPop %v %v", i, vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("BLPop %v", i))
	}

	return
}

func (c *TraceClient) BRPop(i time.Duration, vals ...string) (cmd *redis.StringSliceCmd) {
	span := c.tracer.StartSpan("redis_brpop", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.BRPop(i, vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("BRPop %v %v", i, vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("BRPop %v", i))
	}

	return
}

func (c *TraceClient) BRPopLPush(key string, val string, i time.Duration) (cmd *redis.StringCmd) {
	span := c.tracer.StartSpan("redis_brpoplpush", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.BRPopLPush(key, val, i)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("BRPopLPush %v %v %v", key, val, i))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("BRPopLPush %v", key))
	}

	return
}

func (c *TraceClient) BZPopMax(i time.Duration, vals ...string) (cmd *redis.ZWithKeyCmd) {
	span := c.tracer.StartSpan("redis_bzpopmax", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.BZPopMax(i, vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("BZPopMax %v %v", i, vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("BZPopMax %v", i))
	}

	return
}

func (c *TraceClient) BZPopMin(i time.Duration, vals ...string) (cmd *redis.ZWithKeyCmd) {
	span := c.tracer.StartSpan("redis_bzpopmin", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.BZPopMin(i, vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("BZPopMin %v %v", i, vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("BZPopMin %v", i))
	}

	return
}

func (c *TraceClient) BgRewriteAOF() (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_bgrewriteaof", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.BgRewriteAOF()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "BgRewriteAOF")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "BgRewriteAOF")
	}

	return
}

func (c *TraceClient) BgSave() (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_bgsave", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.BgSave()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "BgSave")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "BgSave")
	}

	return
}

func (c *TraceClient) BitCount(key string, ptr *redis.BitCount) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_bitcount", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.BitCount(key, ptr)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("BitCount %v %v", key, ptr))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("BitCount %v", key))
	}

	return
}

func (c *TraceClient) BitOpAnd(key string, vals ...string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_bitopand", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.BitOpAnd(key, vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("BitOpAnd %v %v", key, vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("BitOpAnd %v", key))
	}

	return
}

func (c *TraceClient) BitOpNot(key string, val string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_bitopnot", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.BitOpNot(key, val)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("BitOpNot %v %v", key, val))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("BitOpNot %v", key))
	}

	return
}

func (c *TraceClient) BitOpOr(key string, vals ...string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_bitopor", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.BitOpOr(key, vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("BitOpOr %v %v", key, vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("BitOpOr %v", key))
	}

	return
}

func (c *TraceClient) BitOpXor(key string, vals ...string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_bitopxor", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.BitOpXor(key, vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("BitOpXor %v %v", key, vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("BitOpXor %v", key))
	}

	return
}

func (c *TraceClient) BitPos(key string, i int64, vals ...int64) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_bitpos", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.BitPos(key, i, vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("BitPos %v %v %v", key, i, vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("BitPos %v", key))
	}

	return
}

func (c *TraceClient) ClientGetName() (cmd *redis.StringCmd) {
	span := c.tracer.StartSpan("redis_clientgetname", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ClientGetName()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "ClientGetName")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "ClientGetName")
	}

	return
}

func (c *TraceClient) ClientID() (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_clientid", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ClientID()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "ClientID")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "ClientID")
	}

	return
}

func (c *TraceClient) ClientKill(key string) (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_clientkill", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ClientKill(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ClientKill %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ClientKill %v", key))
	}

	return
}

func (c *TraceClient) ClientKillByFilter(vals ...string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_clientkillbyfilter", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ClientKillByFilter(vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ClientKillByFilter %v", vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ClientKillByFilter %v", vals))
	}

	return
}

func (c *TraceClient) ClientList() (cmd *redis.StringCmd) {
	span := c.tracer.StartSpan("redis_clientlist", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ClientList()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "ClientList")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "ClientList")
	}

	return
}

func (c *TraceClient) ClientPause(i time.Duration) (cmd *redis.BoolCmd) {
	span := c.tracer.StartSpan("redis_clientpause", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ClientPause(i)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ClientPause %v", i))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ClientPause %v", i))
	}

	return
}

func (c *TraceClient) ClientUnblock(i int64) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_clientunblock", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ClientUnblock(i)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ClientUnblock %v", i))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ClientUnblock %v", i))
	}

	return
}

func (c *TraceClient) ClientUnblockWithError(i int64) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_clientunblockwitherror", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ClientUnblockWithError(i)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ClientUnblockWithError %v", i))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ClientUnblockWithError %v", i))
	}

	return
}

func (c *TraceClient) Close() (err error) {
	span := c.tracer.StartSpan("redis_close", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	err = c.Client.Close()
	if err != nil {
		ext.DBStatement.Set(span, "Close")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "Close")
	}

	return
}

func (c *TraceClient) ClusterAddSlots(vals ...int) (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_clusteraddslots", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ClusterAddSlots(vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ClusterAddSlots %v", vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ClusterAddSlots %v", vals))
	}

	return
}

func (c *TraceClient) ClusterAddSlotsRange(i int, i1 int) (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_clusteraddslotsrange", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ClusterAddSlotsRange(i, i1)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ClusterAddSlotsRange %v %v", i, i1))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ClusterAddSlotsRange %v", i))
	}

	return
}

func (c *TraceClient) ClusterCountFailureReports(key string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_clustercountfailurereports", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ClusterCountFailureReports(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ClusterCountFailureReports %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ClusterCountFailureReports %v", key))
	}

	return
}

func (c *TraceClient) ClusterCountKeysInSlot(i int) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_clustercountkeysinslot", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ClusterCountKeysInSlot(i)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ClusterCountKeysInSlot %v", i))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ClusterCountKeysInSlot %v", i))
	}

	return
}

func (c *TraceClient) ClusterDelSlots(vals ...int) (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_clusterdelslots", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ClusterDelSlots(vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ClusterDelSlots %v", vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ClusterDelSlots %v", vals))
	}

	return
}

func (c *TraceClient) ClusterDelSlotsRange(i int, i1 int) (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_clusterdelslotsrange", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ClusterDelSlotsRange(i, i1)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ClusterDelSlotsRange %v %v", i, i1))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ClusterDelSlotsRange %v", i))
	}

	return
}

func (c *TraceClient) ClusterFailover() (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_clusterfailover", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ClusterFailover()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "ClusterFailover")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "ClusterFailover")
	}

	return
}

func (c *TraceClient) ClusterForget(key string) (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_clusterforget", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ClusterForget(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ClusterForget %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ClusterForget %v", key))
	}

	return
}

func (c *TraceClient) ClusterGetKeysInSlot(i int, i1 int) (cmd *redis.StringSliceCmd) {
	span := c.tracer.StartSpan("redis_clustergetkeysinslot", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ClusterGetKeysInSlot(i, i1)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ClusterGetKeysInSlot %v %v", i, i1))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ClusterGetKeysInSlot %v", i))
	}

	return
}

func (c *TraceClient) ClusterInfo() (cmd *redis.StringCmd) {
	span := c.tracer.StartSpan("redis_clusterinfo", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ClusterInfo()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "ClusterInfo")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "ClusterInfo")
	}

	return
}

func (c *TraceClient) ClusterKeySlot(key string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_clusterkeyslot", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ClusterKeySlot(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ClusterKeySlot %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ClusterKeySlot %v", key))
	}

	return
}

func (c *TraceClient) ClusterMeet(key string, val string) (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_clustermeet", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ClusterMeet(key, val)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ClusterMeet %v %v", key, val))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ClusterMeet %v", key))
	}

	return
}

func (c *TraceClient) ClusterNodes() (cmd *redis.StringCmd) {
	span := c.tracer.StartSpan("redis_clusternodes", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ClusterNodes()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "ClusterNodes")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "ClusterNodes")
	}

	return
}

func (c *TraceClient) ClusterReplicate(key string) (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_clusterreplicate", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ClusterReplicate(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ClusterReplicate %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ClusterReplicate %v", key))
	}

	return
}

func (c *TraceClient) ClusterResetHard() (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_clusterresethard", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ClusterResetHard()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "ClusterResetHard")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "ClusterResetHard")
	}

	return
}

func (c *TraceClient) ClusterResetSoft() (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_clusterresetsoft", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ClusterResetSoft()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "ClusterResetSoft")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "ClusterResetSoft")
	}

	return
}

func (c *TraceClient) ClusterSaveConfig() (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_clustersaveconfig", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ClusterSaveConfig()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "ClusterSaveConfig")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "ClusterSaveConfig")
	}

	return
}

func (c *TraceClient) ClusterSlaves(key string) (cmd *redis.StringSliceCmd) {
	span := c.tracer.StartSpan("redis_clusterslaves", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ClusterSlaves(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ClusterSlaves %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ClusterSlaves %v", key))
	}

	return
}

func (c *TraceClient) ClusterSlots() (cmd *redis.ClusterSlotsCmd) {
	span := c.tracer.StartSpan("redis_clusterslots", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ClusterSlots()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "ClusterSlots")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "ClusterSlots")
	}

	return
}

func (c *TraceClient) Command() (cmd *redis.CommandsInfoCmd) {
	span := c.tracer.StartSpan("redis_command", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.Command()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "Command")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "Command")
	}

	return
}

func (c *TraceClient) ConfigGet(key string) (cmd *redis.SliceCmd) {
	span := c.tracer.StartSpan("redis_configget", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ConfigGet(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ConfigGet %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ConfigGet %v", key))
	}

	return
}

func (c *TraceClient) ConfigResetStat() (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_configresetstat", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ConfigResetStat()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "ConfigResetStat")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "ConfigResetStat")
	}

	return
}

func (c *TraceClient) ConfigRewrite() (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_configrewrite", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ConfigRewrite()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "ConfigRewrite")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "ConfigRewrite")
	}

	return
}

func (c *TraceClient) ConfigSet(key string, val string) (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_configset", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ConfigSet(key, val)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ConfigSet %v %v", key, val))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ConfigSet %v", key))
	}

	return
}

func (c *TraceClient) Context() (iface context.Context) {
	span := c.tracer.StartSpan("redis_context", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	iface = c.Client.Context()

	return
}

func (c *TraceClient) DBSize() (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_dbsize", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.DBSize()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "DBSize")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "DBSize")
	}

	return
}

func (c *TraceClient) DbSize() (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_dbsize", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.DbSize()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "DbSize")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "DbSize")
	}

	return
}

func (c *TraceClient) DebugObject(key string) (cmd *redis.StringCmd) {
	span := c.tracer.StartSpan("redis_debugobject", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.DebugObject(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("DebugObject %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("DebugObject %v", key))
	}

	return
}

func (c *TraceClient) Decr(key string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_decr", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.Decr(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("Decr %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("Decr %v", key))
	}

	return
}

func (c *TraceClient) DecrBy(key string, i int64) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_decrby", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.DecrBy(key, i)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("DecrBy %v %v", key, i))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("DecrBy %v", key))
	}

	return
}

func (c *TraceClient) Del(vals ...string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_del", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.Del(vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("Del %v", vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("Del %v", vals))
	}

	return
}

func (c *TraceClient) Do(vals ...interface{}) (cmd *redis.Cmd) {
	span := c.tracer.StartSpan("redis_do", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.Do(vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("Do %v", vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("Do %v", vals))
	}

	return
}

func (c *TraceClient) Dump(key string) (cmd *redis.StringCmd) {
	span := c.tracer.StartSpan("redis_dump", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.Dump(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("Dump %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("Dump %v", key))
	}

	return
}

func (c *TraceClient) Echo(iface interface{}) (cmd *redis.StringCmd) {
	span := c.tracer.StartSpan("redis_echo", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.Echo(iface)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("Echo %v", iface))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("Echo %v", iface))
	}

	return
}

func (c *TraceClient) Eval(key string, vals []string, vals1 ...interface{}) (cmd *redis.Cmd) {
	span := c.tracer.StartSpan("redis_eval", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.Eval(key, vals, vals1...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("Eval %v %v %v", key, vals, vals1))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("Eval %v", key))
	}

	return
}

func (c *TraceClient) EvalSha(key string, vals []string, vals1 ...interface{}) (cmd *redis.Cmd) {
	span := c.tracer.StartSpan("redis_evalsha", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.EvalSha(key, vals, vals1...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("EvalSha %v %v %v", key, vals, vals1))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("EvalSha %v", key))
	}

	return
}

func (c *TraceClient) Exists(vals ...string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_exists", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.Exists(vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("Exists %v", vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("Exists %v", vals))
	}

	return
}

func (c *TraceClient) Expire(key string, i time.Duration) (cmd *redis.BoolCmd) {
	span := c.tracer.StartSpan("redis_expire", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.Expire(key, i)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("Expire %v %v", key, i))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("Expire %v", key))
	}

	return
}

func (c *TraceClient) ExpireAt(key string, t time.Time) (cmd *redis.BoolCmd) {
	span := c.tracer.StartSpan("redis_expireat", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ExpireAt(key, t)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ExpireAt %v %v", key, t))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ExpireAt %v", key))
	}

	return
}

func (c *TraceClient) FlushAll() (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_flushall", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.FlushAll()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "FlushAll")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "FlushAll")
	}

	return
}

func (c *TraceClient) FlushAllAsync() (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_flushallasync", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.FlushAllAsync()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "FlushAllAsync")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "FlushAllAsync")
	}

	return
}

func (c *TraceClient) FlushDB() (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_flushdb", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.FlushDB()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "FlushDB")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "FlushDB")
	}

	return
}

func (c *TraceClient) FlushDBAsync() (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_flushdbasync", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.FlushDBAsync()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "FlushDBAsync")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "FlushDBAsync")
	}

	return
}

func (c *TraceClient) FlushDb() (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_flushdb", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.FlushDb()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "FlushDb")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "FlushDb")
	}

	return
}

func (c *TraceClient) GeoAdd(key string, location ...*redis.GeoLocation) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_geoadd", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.GeoAdd(key, location...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("GeoAdd %v %v", key, location))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("GeoAdd %v", key))
	}

	return
}

func (c *TraceClient) GeoDist(key string, val string, val1 string, val2 string) (cmd *redis.FloatCmd) {
	span := c.tracer.StartSpan("redis_geodist", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.GeoDist(key, val, val1, val2)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("GeoDist %v %v %v %v", key, val, val1, val2))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("GeoDist %v", key))
	}

	return
}

func (c *TraceClient) GeoHash(key string, vals ...string) (cmd *redis.StringSliceCmd) {
	span := c.tracer.StartSpan("redis_geohash", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.GeoHash(key, vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("GeoHash %v %v", key, vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("GeoHash %v", key))
	}

	return
}

func (c *TraceClient) GeoPos(key string, vals ...string) (cmd *redis.GeoPosCmd) {
	span := c.tracer.StartSpan("redis_geopos", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.GeoPos(key, vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("GeoPos %v %v", key, vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("GeoPos %v", key))
	}

	return
}

func (c *TraceClient) GeoRadius(key string, f float64, f1 float64, ptr *redis.GeoRadiusQuery) (cmd *redis.GeoLocationCmd) {
	span := c.tracer.StartSpan("redis_georadius", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.GeoRadius(key, f, f1, ptr)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("GeoRadius %v %v %v %v", key, f, f1, ptr))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("GeoRadius %v", key))
	}

	return
}

func (c *TraceClient) GeoRadiusByMember(key string, val string, ptr *redis.GeoRadiusQuery) (cmd *redis.GeoLocationCmd) {
	span := c.tracer.StartSpan("redis_georadiusbymember", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.GeoRadiusByMember(key, val, ptr)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("GeoRadiusByMember %v %v %v", key, val, ptr))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("GeoRadiusByMember %v", key))
	}

	return
}

func (c *TraceClient) GeoRadiusByMemberRO(key string, val string, ptr *redis.GeoRadiusQuery) (cmd *redis.GeoLocationCmd) {
	span := c.tracer.StartSpan("redis_georadiusbymemberro", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.GeoRadiusByMemberRO(key, val, ptr)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("GeoRadiusByMemberRO %v %v %v", key, val, ptr))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("GeoRadiusByMemberRO %v", key))
	}

	return
}

func (c *TraceClient) GeoRadiusRO(key string, f float64, f1 float64, ptr *redis.GeoRadiusQuery) (cmd *redis.GeoLocationCmd) {
	span := c.tracer.StartSpan("redis_georadiusro", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.GeoRadiusRO(key, f, f1, ptr)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("GeoRadiusRO %v %v %v %v", key, f, f1, ptr))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("GeoRadiusRO %v", key))
	}

	return
}

func (c *TraceClient) Get(key string) (cmd *redis.StringCmd) {
	span := c.tracer.StartSpan("redis_get", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.Get(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("Get %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("Get %v", key))
	}

	return
}

func (c *TraceClient) GetBit(key string, i int64) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_getbit", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.GetBit(key, i)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("GetBit %v %v", key, i))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("GetBit %v", key))
	}

	return
}

func (c *TraceClient) GetRange(key string, i int64, i1 int64) (cmd *redis.StringCmd) {
	span := c.tracer.StartSpan("redis_getrange", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.GetRange(key, i, i1)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("GetRange %v %v %v", key, i, i1))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("GetRange %v", key))
	}

	return
}

func (c *TraceClient) GetSet(key string, iface interface{}) (cmd *redis.StringCmd) {
	span := c.tracer.StartSpan("redis_getset", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.GetSet(key, iface)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("GetSet %v %v", key, iface))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("GetSet %v", key))
	}

	return
}

func (c *TraceClient) HDel(key string, vals ...string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_hdel", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.HDel(key, vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("HDel %v %v", key, vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("HDel %v", key))
	}

	return
}

func (c *TraceClient) HExists(key string, val string) (cmd *redis.BoolCmd) {
	span := c.tracer.StartSpan("redis_hexists", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.HExists(key, val)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("HExists %v %v", key, val))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("HExists %v", key))
	}

	return
}

func (c *TraceClient) HGet(key string, val string) (cmd *redis.StringCmd) {
	span := c.tracer.StartSpan("redis_hget", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.HGet(key, val)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("HGet %v %v", key, val))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("HGet %v", key))
	}

	return
}

func (c *TraceClient) HGetAll(key string) (cmd *redis.StringStringMapCmd) {
	span := c.tracer.StartSpan("redis_hgetall", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.HGetAll(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("HGetAll %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("HGetAll %v", key))
	}

	return
}

func (c *TraceClient) HIncrBy(key string, val string, i int64) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_hincrby", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.HIncrBy(key, val, i)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("HIncrBy %v %v %v", key, val, i))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("HIncrBy %v", key))
	}

	return
}

func (c *TraceClient) HIncrByFloat(key string, val string, f float64) (cmd *redis.FloatCmd) {
	span := c.tracer.StartSpan("redis_hincrbyfloat", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.HIncrByFloat(key, val, f)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("HIncrByFloat %v %v %v", key, val, f))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("HIncrByFloat %v", key))
	}

	return
}

func (c *TraceClient) HKeys(key string) (cmd *redis.StringSliceCmd) {
	span := c.tracer.StartSpan("redis_hkeys", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.HKeys(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("HKeys %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("HKeys %v", key))
	}

	return
}

func (c *TraceClient) HLen(key string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_hlen", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.HLen(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("HLen %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("HLen %v", key))
	}

	return
}

func (c *TraceClient) HMGet(key string, vals ...string) (cmd *redis.SliceCmd) {
	span := c.tracer.StartSpan("redis_hmget", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.HMGet(key, vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("HMGet %v %v", key, vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("HMGet %v", key))
	}

	return
}

func (c *TraceClient) HMSet(key string, mdata map[string]interface{}) (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_hmset", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.HMSet(key, mdata)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("HMSet %v %v", key, mdata))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("HMSet %v", key))
	}

	return
}

func (c *TraceClient) HScan(key string, n uint64, val string, i int64) (cmd *redis.ScanCmd) {
	span := c.tracer.StartSpan("redis_hscan", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.HScan(key, n, val, i)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("HScan %v %v %v %v", key, n, val, i))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("HScan %v", key))
	}

	return
}

func (c *TraceClient) HSet(key string, val string, iface interface{}) (cmd *redis.BoolCmd) {
	span := c.tracer.StartSpan("redis_hset", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.HSet(key, val, iface)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("HSet %v %v %v", key, val, iface))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("HSet %v", key))
	}

	return
}

func (c *TraceClient) HSetNX(key string, val string, iface interface{}) (cmd *redis.BoolCmd) {
	span := c.tracer.StartSpan("redis_hsetnx", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.HSetNX(key, val, iface)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("HSetNX %v %v %v", key, val, iface))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("HSetNX %v", key))
	}

	return
}

func (c *TraceClient) HVals(key string) (cmd *redis.StringSliceCmd) {
	span := c.tracer.StartSpan("redis_hvals", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.HVals(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("HVals %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("HVals %v", key))
	}

	return
}

func (c *TraceClient) Incr(key string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_incr", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.Incr(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("Incr %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("Incr %v", key))
	}

	return
}

func (c *TraceClient) IncrBy(key string, i int64) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_incrby", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.IncrBy(key, i)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("IncrBy %v %v", key, i))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("IncrBy %v", key))
	}

	return
}

func (c *TraceClient) IncrByFloat(key string, f float64) (cmd *redis.FloatCmd) {
	span := c.tracer.StartSpan("redis_incrbyfloat", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.IncrByFloat(key, f)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("IncrByFloat %v %v", key, f))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("IncrByFloat %v", key))
	}

	return
}

func (c *TraceClient) Info(vals ...string) (cmd *redis.StringCmd) {
	span := c.tracer.StartSpan("redis_info", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.Info(vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("Info %v", vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("Info %v", vals))
	}

	return
}

func (c *TraceClient) Keys(key string) (cmd *redis.StringSliceCmd) {
	span := c.tracer.StartSpan("redis_keys", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.Keys(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("Keys %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("Keys %v", key))
	}

	return
}

func (c *TraceClient) LIndex(key string, i int64) (cmd *redis.StringCmd) {
	span := c.tracer.StartSpan("redis_lindex", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.LIndex(key, i)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("LIndex %v %v", key, i))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("LIndex %v", key))
	}

	return
}

func (c *TraceClient) LInsert(key string, val string, iface interface{}, iface1 interface{}) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_linsert", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.LInsert(key, val, iface, iface1)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("LInsert %v %v %v %v", key, val, iface, iface1))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("LInsert %v", key))
	}

	return
}

func (c *TraceClient) LInsertAfter(key string, iface interface{}, iface1 interface{}) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_linsertafter", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.LInsertAfter(key, iface, iface1)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("LInsertAfter %v %v %v", key, iface, iface1))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("LInsertAfter %v", key))
	}

	return
}

func (c *TraceClient) LInsertBefore(key string, iface interface{}, iface1 interface{}) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_linsertbefore", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.LInsertBefore(key, iface, iface1)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("LInsertBefore %v %v %v", key, iface, iface1))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("LInsertBefore %v", key))
	}

	return
}

func (c *TraceClient) LLen(key string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_llen", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.LLen(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("LLen %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("LLen %v", key))
	}

	return
}

func (c *TraceClient) LPop(key string) (cmd *redis.StringCmd) {
	span := c.tracer.StartSpan("redis_lpop", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.LPop(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("LPop %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("LPop %v", key))
	}

	return
}

func (c *TraceClient) LPush(key string, vals ...interface{}) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_lpush", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.LPush(key, vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("LPush %v %v", key, vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("LPush %v", key))
	}

	return
}

func (c *TraceClient) LPushX(key string, iface interface{}) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_lpushx", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.LPushX(key, iface)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("LPushX %v %v", key, iface))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("LPushX %v", key))
	}

	return
}

func (c *TraceClient) LRange(key string, i int64, i1 int64) (cmd *redis.StringSliceCmd) {
	span := c.tracer.StartSpan("redis_lrange", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.LRange(key, i, i1)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("LRange %v %v %v", key, i, i1))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("LRange %v", key))
	}

	return
}

func (c *TraceClient) LRem(key string, i int64, iface interface{}) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_lrem", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.LRem(key, i, iface)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("LRem %v %v %v", key, i, iface))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("LRem %v", key))
	}

	return
}

func (c *TraceClient) LSet(key string, i int64, iface interface{}) (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_lset", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.LSet(key, i, iface)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("LSet %v %v %v", key, i, iface))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("LSet %v", key))
	}

	return
}

func (c *TraceClient) LTrim(key string, i int64, i1 int64) (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_ltrim", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.LTrim(key, i, i1)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("LTrim %v %v %v", key, i, i1))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("LTrim %v", key))
	}

	return
}

func (c *TraceClient) LastSave() (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_lastsave", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.LastSave()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "LastSave")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "LastSave")
	}

	return
}

func (c *TraceClient) MGet(vals ...string) (cmd *redis.SliceCmd) {
	span := c.tracer.StartSpan("redis_mget", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.MGet(vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("MGet %v", vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("MGet %v", vals))
	}

	return
}

func (c *TraceClient) MSet(vals ...interface{}) (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_mset", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.MSet(vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("MSet %v", vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("MSet %v", vals))
	}

	return
}

func (c *TraceClient) MSetNX(vals ...interface{}) (cmd *redis.BoolCmd) {
	span := c.tracer.StartSpan("redis_msetnx", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.MSetNX(vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("MSetNX %v", vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("MSetNX %v", vals))
	}

	return
}

func (c *TraceClient) MemoryUsage(key string, vals ...int) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_memoryusage", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.MemoryUsage(key, vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("MemoryUsage %v %v", key, vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("MemoryUsage %v", key))
	}

	return
}

func (c *TraceClient) Migrate(key string, val string, val1 string, i int64, i1 time.Duration) (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_migrate", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.Migrate(key, val, val1, i, i1)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("Migrate %v %v %v %v %v", key, val, val1, i, i1))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("Migrate %v", key))
	}

	return
}

func (c *TraceClient) Move(key string, i int64) (cmd *redis.BoolCmd) {
	span := c.tracer.StartSpan("redis_move", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.Move(key, i)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("Move %v %v", key, i))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("Move %v", key))
	}

	return
}

func (c *TraceClient) ObjectEncoding(key string) (cmd *redis.StringCmd) {
	span := c.tracer.StartSpan("redis_objectencoding", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ObjectEncoding(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ObjectEncoding %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ObjectEncoding %v", key))
	}

	return
}

func (c *TraceClient) ObjectIdleTime(key string) (cmd *redis.DurationCmd) {
	span := c.tracer.StartSpan("redis_objectidletime", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ObjectIdleTime(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ObjectIdleTime %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ObjectIdleTime %v", key))
	}

	return
}

func (c *TraceClient) ObjectRefCount(key string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_objectrefcount", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ObjectRefCount(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ObjectRefCount %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ObjectRefCount %v", key))
	}

	return
}

func (c *TraceClient) Options() (options *redis.Options) {
	span := c.tracer.StartSpan("redis_options", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	options = c.Client.Options()

	return
}

func (c *TraceClient) PExpire(key string, i time.Duration) (cmd *redis.BoolCmd) {
	span := c.tracer.StartSpan("redis_pexpire", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.PExpire(key, i)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("PExpire %v %v", key, i))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("PExpire %v", key))
	}

	return
}

func (c *TraceClient) PExpireAt(key string, t time.Time) (cmd *redis.BoolCmd) {
	span := c.tracer.StartSpan("redis_pexpireat", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.PExpireAt(key, t)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("PExpireAt %v %v", key, t))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("PExpireAt %v", key))
	}

	return
}

func (c *TraceClient) PFAdd(key string, vals ...interface{}) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_pfadd", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.PFAdd(key, vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("PFAdd %v %v", key, vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("PFAdd %v", key))
	}

	return
}

func (c *TraceClient) PFCount(vals ...string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_pfcount", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.PFCount(vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("PFCount %v", vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("PFCount %v", vals))
	}

	return
}

func (c *TraceClient) PFMerge(key string, vals ...string) (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_pfmerge", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.PFMerge(key, vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("PFMerge %v %v", key, vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("PFMerge %v", key))
	}

	return
}

func (c *TraceClient) PSubscribe(vals ...string) (sub *redis.PubSub) {
	span := c.tracer.StartSpan("redis_psubscribe", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	sub = c.Client.PSubscribe(vals...)

	return
}

func (c *TraceClient) PTTL(key string) (cmd *redis.DurationCmd) {
	span := c.tracer.StartSpan("redis_pttl", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.PTTL(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("PTTL %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("PTTL %v", key))
	}

	return
}

func (c *TraceClient) Persist(key string) (cmd *redis.BoolCmd) {
	span := c.tracer.StartSpan("redis_persist", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.Persist(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("Persist %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("Persist %v", key))
	}

	return
}

func (c *TraceClient) Ping() (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_ping", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.Ping()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "Ping")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "Ping")
	}

	return
}

func (c *TraceClient) Pipeline() (iface redis.Pipeliner) {
	span := c.tracer.StartSpan("redis_pipeline", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	iface = c.Client.Pipeline()

	return
}

func (c *TraceClient) Pipelined(fn func(redis.Pipeliner) error) (vals []redis.Cmder, err error) {
	span := c.tracer.StartSpan("redis_pipelined", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	vals, err = c.Client.Pipelined(fn)
	if err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("Pipelined %p", fn))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", vals, err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("Pipelined %p", fn))
	}

	return
}

func (c *TraceClient) PoolStats() (stats *redis.PoolStats) {
	span := c.tracer.StartSpan("redis_poolstats", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	stats = c.Client.PoolStats()

	return
}

func (c *TraceClient) Process(iface redis.Cmder) (err error) {
	span := c.tracer.StartSpan("redis_process", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	err = c.Client.Process(iface)
	if err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("Process %v", iface))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("Process %v", iface))
	}

	return
}

func (c *TraceClient) PubSubChannels(key string) (cmd *redis.StringSliceCmd) {
	span := c.tracer.StartSpan("redis_pubsubchannels", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.PubSubChannels(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("PubSubChannels %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("PubSubChannels %v", key))
	}

	return
}

func (c *TraceClient) PubSubNumPat() (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_pubsubnumpat", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.PubSubNumPat()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "PubSubNumPat")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "PubSubNumPat")
	}

	return
}

func (c *TraceClient) PubSubNumSub(vals ...string) (cmd *redis.StringIntMapCmd) {
	span := c.tracer.StartSpan("redis_pubsubnumsub", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.PubSubNumSub(vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("PubSubNumSub %v", vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("PubSubNumSub %v", vals))
	}

	return
}

func (c *TraceClient) Publish(key string, iface interface{}) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_publish", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.Publish(key, iface)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("Publish %v %v", key, iface))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("Publish %v", key))
	}

	return
}

func (c *TraceClient) Quit() (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_quit", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.Quit()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "Quit")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "Quit")
	}

	return
}

func (c *TraceClient) RPop(key string) (cmd *redis.StringCmd) {
	span := c.tracer.StartSpan("redis_rpop", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.RPop(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("RPop %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("RPop %v", key))
	}

	return
}

func (c *TraceClient) RPopLPush(key string, val string) (cmd *redis.StringCmd) {
	span := c.tracer.StartSpan("redis_rpoplpush", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.RPopLPush(key, val)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("RPopLPush %v %v", key, val))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("RPopLPush %v", key))
	}

	return
}

func (c *TraceClient) RPush(key string, vals ...interface{}) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_rpush", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.RPush(key, vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("RPush %v %v", key, vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("RPush %v", key))
	}

	return
}

func (c *TraceClient) RPushX(key string, iface interface{}) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_rpushx", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.RPushX(key, iface)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("RPushX %v %v", key, iface))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("RPushX %v", key))
	}

	return
}

func (c *TraceClient) RandomKey() (cmd *redis.StringCmd) {
	span := c.tracer.StartSpan("redis_randomkey", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.RandomKey()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "RandomKey")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "RandomKey")
	}

	return
}

func (c *TraceClient) ReadOnly() (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_readonly", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ReadOnly()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "ReadOnly")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "ReadOnly")
	}

	return
}

func (c *TraceClient) ReadWrite() (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_readwrite", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ReadWrite()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "ReadWrite")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "ReadWrite")
	}

	return
}

func (c *TraceClient) Rename(key string, val string) (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_rename", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.Rename(key, val)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("Rename %v %v", key, val))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("Rename %v", key))
	}

	return
}

func (c *TraceClient) RenameNX(key string, val string) (cmd *redis.BoolCmd) {
	span := c.tracer.StartSpan("redis_renamenx", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.RenameNX(key, val)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("RenameNX %v %v", key, val))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("RenameNX %v", key))
	}

	return
}

func (c *TraceClient) Restore(key string, i time.Duration, val string) (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_restore", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.Restore(key, i, val)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("Restore %v %v %v", key, i, val))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("Restore %v", key))
	}

	return
}

func (c *TraceClient) RestoreReplace(key string, i time.Duration, val string) (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_restorereplace", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.RestoreReplace(key, i, val)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("RestoreReplace %v %v %v", key, i, val))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("RestoreReplace %v", key))
	}

	return
}

func (c *TraceClient) SAdd(key string, vals ...interface{}) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_sadd", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.SAdd(key, vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("SAdd %v %v", key, vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("SAdd %v", key))
	}

	return
}

func (c *TraceClient) SCard(key string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_scard", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.SCard(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("SCard %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("SCard %v", key))
	}

	return
}

func (c *TraceClient) SDiff(vals ...string) (cmd *redis.StringSliceCmd) {
	span := c.tracer.StartSpan("redis_sdiff", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.SDiff(vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("SDiff %v", vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("SDiff %v", vals))
	}

	return
}

func (c *TraceClient) SDiffStore(key string, vals ...string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_sdiffstore", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.SDiffStore(key, vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("SDiffStore %v %v", key, vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("SDiffStore %v", key))
	}

	return
}

func (c *TraceClient) SInter(vals ...string) (cmd *redis.StringSliceCmd) {
	span := c.tracer.StartSpan("redis_sinter", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.SInter(vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("SInter %v", vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("SInter %v", vals))
	}

	return
}

func (c *TraceClient) SInterStore(key string, vals ...string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_sinterstore", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.SInterStore(key, vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("SInterStore %v %v", key, vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("SInterStore %v", key))
	}

	return
}

func (c *TraceClient) SIsMember(key string, iface interface{}) (cmd *redis.BoolCmd) {
	span := c.tracer.StartSpan("redis_sismember", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.SIsMember(key, iface)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("SIsMember %v %v", key, iface))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("SIsMember %v", key))
	}

	return
}

func (c *TraceClient) SMembers(key string) (cmd *redis.StringSliceCmd) {
	span := c.tracer.StartSpan("redis_smembers", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.SMembers(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("SMembers %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("SMembers %v", key))
	}

	return
}

func (c *TraceClient) SMembersMap(key string) (cmd *redis.StringStructMapCmd) {
	span := c.tracer.StartSpan("redis_smembersmap", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.SMembersMap(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("SMembersMap %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("SMembersMap %v", key))
	}

	return
}

func (c *TraceClient) SMove(key string, val string, iface interface{}) (cmd *redis.BoolCmd) {
	span := c.tracer.StartSpan("redis_smove", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.SMove(key, val, iface)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("SMove %v %v %v", key, val, iface))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("SMove %v", key))
	}

	return
}

func (c *TraceClient) SPop(key string) (cmd *redis.StringCmd) {
	span := c.tracer.StartSpan("redis_spop", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.SPop(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("SPop %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("SPop %v", key))
	}

	return
}

func (c *TraceClient) SPopN(key string, i int64) (cmd *redis.StringSliceCmd) {
	span := c.tracer.StartSpan("redis_spopn", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.SPopN(key, i)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("SPopN %v %v", key, i))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("SPopN %v", key))
	}

	return
}

func (c *TraceClient) SRandMember(key string) (cmd *redis.StringCmd) {
	span := c.tracer.StartSpan("redis_srandmember", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.SRandMember(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("SRandMember %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("SRandMember %v", key))
	}

	return
}

func (c *TraceClient) SRandMemberN(key string, i int64) (cmd *redis.StringSliceCmd) {
	span := c.tracer.StartSpan("redis_srandmembern", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.SRandMemberN(key, i)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("SRandMemberN %v %v", key, i))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("SRandMemberN %v", key))
	}

	return
}

func (c *TraceClient) SRem(key string, vals ...interface{}) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_srem", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.SRem(key, vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("SRem %v %v", key, vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("SRem %v", key))
	}

	return
}

func (c *TraceClient) SScan(key string, n uint64, val string, i int64) (cmd *redis.ScanCmd) {
	span := c.tracer.StartSpan("redis_sscan", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.SScan(key, n, val, i)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("SScan %v %v %v %v", key, n, val, i))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("SScan %v", key))
	}

	return
}

func (c *TraceClient) SUnion(vals ...string) (cmd *redis.StringSliceCmd) {
	span := c.tracer.StartSpan("redis_sunion", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.SUnion(vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("SUnion %v", vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("SUnion %v", vals))
	}

	return
}

func (c *TraceClient) SUnionStore(key string, vals ...string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_sunionstore", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.SUnionStore(key, vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("SUnionStore %v %v", key, vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("SUnionStore %v", key))
	}

	return
}

func (c *TraceClient) Save() (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_save", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.Save()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "Save")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "Save")
	}

	return
}

func (c *TraceClient) Scan(n uint64, val string, i int64) (cmd *redis.ScanCmd) {
	span := c.tracer.StartSpan("redis_scan", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.Scan(n, val, i)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("Scan %v %v %v", n, val, i))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("Scan %v", n))
	}

	return
}

func (c *TraceClient) ScriptExists(vals ...string) (cmd *redis.BoolSliceCmd) {
	span := c.tracer.StartSpan("redis_scriptexists", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ScriptExists(vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ScriptExists %v", vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ScriptExists %v", vals))
	}

	return
}

func (c *TraceClient) ScriptFlush() (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_scriptflush", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ScriptFlush()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "ScriptFlush")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "ScriptFlush")
	}

	return
}

func (c *TraceClient) ScriptKill() (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_scriptkill", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ScriptKill()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "ScriptKill")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "ScriptKill")
	}

	return
}

func (c *TraceClient) ScriptLoad(key string) (cmd *redis.StringCmd) {
	span := c.tracer.StartSpan("redis_scriptload", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ScriptLoad(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ScriptLoad %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ScriptLoad %v", key))
	}

	return
}

func (c *TraceClient) Set(key string, iface interface{}, i time.Duration) (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_set", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.Set(key, iface, i)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("Set %v %v %v", key, iface, i))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("Set %v", key))
	}

	return
}

func (c *TraceClient) SetBit(key string, i int64, i1 int) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_setbit", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.SetBit(key, i, i1)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("SetBit %v %v %v", key, i, i1))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("SetBit %v", key))
	}

	return
}

func (c *TraceClient) SetLimiter(iface redis.Limiter) (client *redis.Client) {
	span := c.tracer.StartSpan("redis_setlimiter", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	client = c.Client.SetLimiter(iface)

	return
}

func (c *TraceClient) SetNX(key string, iface interface{}, i time.Duration) (cmd *redis.BoolCmd) {
	span := c.tracer.StartSpan("redis_setnx", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.SetNX(key, iface, i)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("SetNX %v %v %v", key, iface, i))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("SetNX %v", key))
	}

	return
}

func (c *TraceClient) SetRange(key string, i int64, val string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_setrange", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.SetRange(key, i, val)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("SetRange %v %v %v", key, i, val))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("SetRange %v", key))
	}

	return
}

func (c *TraceClient) SetXX(key string, iface interface{}, i time.Duration) (cmd *redis.BoolCmd) {
	span := c.tracer.StartSpan("redis_setxx", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.SetXX(key, iface, i)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("SetXX %v %v %v", key, iface, i))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("SetXX %v", key))
	}

	return
}

func (c *TraceClient) Shutdown() (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_shutdown", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.Shutdown()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "Shutdown")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "Shutdown")
	}

	return
}

func (c *TraceClient) ShutdownNoSave() (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_shutdownnosave", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ShutdownNoSave()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "ShutdownNoSave")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "ShutdownNoSave")
	}

	return
}

func (c *TraceClient) ShutdownSave() (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_shutdownsave", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ShutdownSave()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "ShutdownSave")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "ShutdownSave")
	}

	return
}

func (c *TraceClient) SlaveOf(key string, val string) (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_slaveof", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.SlaveOf(key, val)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("SlaveOf %v %v", key, val))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("SlaveOf %v", key))
	}

	return
}

func (c *TraceClient) SlowLog() {
	span := c.tracer.StartSpan("redis_slowlog", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	c.Client.SlowLog()

	return
}

func (c *TraceClient) Sort(key string, ptr *redis.Sort) (cmd *redis.StringSliceCmd) {
	span := c.tracer.StartSpan("redis_sort", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.Sort(key, ptr)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("Sort %v %v", key, ptr))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("Sort %v", key))
	}

	return
}

func (c *TraceClient) SortInterfaces(key string, ptr *redis.Sort) (cmd *redis.SliceCmd) {
	span := c.tracer.StartSpan("redis_sortinterfaces", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.SortInterfaces(key, ptr)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("SortInterfaces %v %v", key, ptr))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("SortInterfaces %v", key))
	}

	return
}

func (c *TraceClient) SortStore(key string, val string, ptr *redis.Sort) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_sortstore", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.SortStore(key, val, ptr)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("SortStore %v %v %v", key, val, ptr))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("SortStore %v", key))
	}

	return
}

func (c *TraceClient) StrLen(key string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_strlen", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.StrLen(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("StrLen %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("StrLen %v", key))
	}

	return
}

func (c *TraceClient) String() (ret string) {
	span := c.tracer.StartSpan("redis_string", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	ret = c.Client.String()

	return
}

func (c *TraceClient) Subscribe(vals ...string) (sub *redis.PubSub) {
	span := c.tracer.StartSpan("redis_subscribe", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	sub = c.Client.Subscribe(vals...)

	return
}

func (c *TraceClient) Sync() {
	span := c.tracer.StartSpan("redis_sync", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	c.Client.Sync()

	return
}

func (c *TraceClient) TTL(key string) (cmd *redis.DurationCmd) {
	span := c.tracer.StartSpan("redis_ttl", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.TTL(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("TTL %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("TTL %v", key))
	}

	return
}

func (c *TraceClient) Time() (cmd *redis.TimeCmd) {
	span := c.tracer.StartSpan("redis_time", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.Time()
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, "Time")

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, "Time")
	}

	return
}

func (c *TraceClient) Touch(vals ...string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_touch", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.Touch(vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("Touch %v", vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("Touch %v", vals))
	}

	return
}

func (c *TraceClient) TxPipeline() (iface redis.Pipeliner) {
	span := c.tracer.StartSpan("redis_txpipeline", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	iface = c.Client.TxPipeline()

	return
}

func (c *TraceClient) TxPipelined(fn func(redis.Pipeliner) error) (vals []redis.Cmder, err error) {
	span := c.tracer.StartSpan("redis_txpipelined", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	vals, err = c.Client.TxPipelined(fn)
	if err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("TxPipelined %p", fn))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", vals, err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("TxPipelined %p", fn))
	}

	return
}

func (c *TraceClient) Type(key string) (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_type", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.Type(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("Type %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("Type %v", key))
	}

	return
}

func (c *TraceClient) Unlink(vals ...string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_unlink", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.Unlink(vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("Unlink %v", vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("Unlink %v", vals))
	}

	return
}

func (c *TraceClient) Wait(i int, i1 time.Duration) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_wait", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.Wait(i, i1)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("Wait %v %v", i, i1))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("Wait %v", i))
	}

	return
}

func (c *TraceClient) Watch(fn func(*redis.Tx) error, vals ...string) (err error) {
	span := c.tracer.StartSpan("redis_watch", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	err = c.Client.Watch(fn, vals...)
	if err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("Watch %p %v", fn, vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("Watch %p", fn))
	}

	return
}

func (c *TraceClient) WithContext(iface context.Context) (client *redis.Client) {
	span := c.tracer.StartSpan("redis_withcontext", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	client = c.Client.WithContext(iface)

	return
}

func (c *TraceClient) WrapProcess(fn func(func(redis.Cmder) error) func(redis.Cmder) error) {
	span := c.tracer.StartSpan("redis_wrapprocess", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	c.Client.WrapProcess(fn)

	return
}

func (c *TraceClient) WrapProcessPipeline(fn func(func([]redis.Cmder) error) func([]redis.Cmder) error) {
	span := c.tracer.StartSpan("redis_wrapprocesspipeline", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	c.Client.WrapProcessPipeline(fn)

	return
}

func (c *TraceClient) XAck(key string, val string, vals ...string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_xack", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.XAck(key, val, vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("XAck %v %v %v", key, val, vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("XAck %v", key))
	}

	return
}

func (c *TraceClient) XAdd(ptr *redis.XAddArgs) (cmd *redis.StringCmd) {
	span := c.tracer.StartSpan("redis_xadd", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.XAdd(ptr)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("XAdd %v", ptr))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("XAdd %v", ptr))
	}

	return
}

func (c *TraceClient) XClaim(ptr *redis.XClaimArgs) (cmd *redis.XMessageSliceCmd) {
	span := c.tracer.StartSpan("redis_xclaim", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.XClaim(ptr)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("XClaim %v", ptr))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("XClaim %v", ptr))
	}

	return
}

func (c *TraceClient) XClaimJustID(ptr *redis.XClaimArgs) (cmd *redis.StringSliceCmd) {
	span := c.tracer.StartSpan("redis_xclaimjustid", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.XClaimJustID(ptr)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("XClaimJustID %v", ptr))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("XClaimJustID %v", ptr))
	}

	return
}

func (c *TraceClient) XDel(key string, vals ...string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_xdel", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.XDel(key, vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("XDel %v %v", key, vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("XDel %v", key))
	}

	return
}

func (c *TraceClient) XGroupCreate(key string, val string, val1 string) (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_xgroupcreate", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.XGroupCreate(key, val, val1)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("XGroupCreate %v %v %v", key, val, val1))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("XGroupCreate %v", key))
	}

	return
}

func (c *TraceClient) XGroupCreateMkStream(key string, val string, val1 string) (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_xgroupcreatemkstream", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.XGroupCreateMkStream(key, val, val1)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("XGroupCreateMkStream %v %v %v", key, val, val1))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("XGroupCreateMkStream %v", key))
	}

	return
}

func (c *TraceClient) XGroupDelConsumer(key string, val string, val1 string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_xgroupdelconsumer", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.XGroupDelConsumer(key, val, val1)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("XGroupDelConsumer %v %v %v", key, val, val1))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("XGroupDelConsumer %v", key))
	}

	return
}

func (c *TraceClient) XGroupDestroy(key string, val string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_xgroupdestroy", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.XGroupDestroy(key, val)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("XGroupDestroy %v %v", key, val))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("XGroupDestroy %v", key))
	}

	return
}

func (c *TraceClient) XGroupSetID(key string, val string, val1 string) (cmd *redis.StatusCmd) {
	span := c.tracer.StartSpan("redis_xgroupsetid", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.XGroupSetID(key, val, val1)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("XGroupSetID %v %v %v", key, val, val1))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("XGroupSetID %v", key))
	}

	return
}

func (c *TraceClient) XLen(key string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_xlen", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.XLen(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("XLen %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("XLen %v", key))
	}

	return
}

func (c *TraceClient) XPending(key string, val string) (cmd *redis.XPendingCmd) {
	span := c.tracer.StartSpan("redis_xpending", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.XPending(key, val)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("XPending %v %v", key, val))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("XPending %v", key))
	}

	return
}

func (c *TraceClient) XPendingExt(ptr *redis.XPendingExtArgs) (cmd *redis.XPendingExtCmd) {
	span := c.tracer.StartSpan("redis_xpendingext", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.XPendingExt(ptr)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("XPendingExt %v", ptr))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("XPendingExt %v", ptr))
	}

	return
}

func (c *TraceClient) XRange(key string, val string, val1 string) (cmd *redis.XMessageSliceCmd) {
	span := c.tracer.StartSpan("redis_xrange", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.XRange(key, val, val1)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("XRange %v %v %v", key, val, val1))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("XRange %v", key))
	}

	return
}

func (c *TraceClient) XRangeN(key string, val string, val1 string, i int64) (cmd *redis.XMessageSliceCmd) {
	span := c.tracer.StartSpan("redis_xrangen", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.XRangeN(key, val, val1, i)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("XRangeN %v %v %v %v", key, val, val1, i))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("XRangeN %v", key))
	}

	return
}

func (c *TraceClient) XRead(ptr *redis.XReadArgs) (cmd *redis.XStreamSliceCmd) {
	span := c.tracer.StartSpan("redis_xread", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.XRead(ptr)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("XRead %v", ptr))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("XRead %v", ptr))
	}

	return
}

func (c *TraceClient) XReadGroup(ptr *redis.XReadGroupArgs) (cmd *redis.XStreamSliceCmd) {
	span := c.tracer.StartSpan("redis_xreadgroup", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.XReadGroup(ptr)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("XReadGroup %v", ptr))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("XReadGroup %v", ptr))
	}

	return
}

func (c *TraceClient) XReadStreams(vals ...string) (cmd *redis.XStreamSliceCmd) {
	span := c.tracer.StartSpan("redis_xreadstreams", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.XReadStreams(vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("XReadStreams %v", vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("XReadStreams %v", vals))
	}

	return
}

func (c *TraceClient) XRevRange(key string, val string, val1 string) (cmd *redis.XMessageSliceCmd) {
	span := c.tracer.StartSpan("redis_xrevrange", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.XRevRange(key, val, val1)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("XRevRange %v %v %v", key, val, val1))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("XRevRange %v", key))
	}

	return
}

func (c *TraceClient) XRevRangeN(key string, val string, val1 string, i int64) (cmd *redis.XMessageSliceCmd) {
	span := c.tracer.StartSpan("redis_xrevrangen", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.XRevRangeN(key, val, val1, i)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("XRevRangeN %v %v %v %v", key, val, val1, i))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("XRevRangeN %v", key))
	}

	return
}

func (c *TraceClient) XTrim(key string, i int64) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_xtrim", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.XTrim(key, i)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("XTrim %v %v", key, i))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("XTrim %v", key))
	}

	return
}

func (c *TraceClient) XTrimApprox(key string, i int64) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_xtrimapprox", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.XTrimApprox(key, i)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("XTrimApprox %v %v", key, i))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("XTrimApprox %v", key))
	}

	return
}

func (c *TraceClient) ZAdd(key string, z ...redis.Z) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_zadd", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZAdd(key, z...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZAdd %v %v", key, z))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZAdd %v", key))
	}

	return
}

func (c *TraceClient) ZAddCh(key string, z ...redis.Z) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_zaddch", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZAddCh(key, z...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZAddCh %v %v", key, z))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZAddCh %v", key))
	}

	return
}

func (c *TraceClient) ZAddNX(key string, z ...redis.Z) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_zaddnx", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZAddNX(key, z...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZAddNX %v %v", key, z))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZAddNX %v", key))
	}

	return
}

func (c *TraceClient) ZAddNXCh(key string, z ...redis.Z) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_zaddnxch", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZAddNXCh(key, z...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZAddNXCh %v %v", key, z))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZAddNXCh %v", key))
	}

	return
}

func (c *TraceClient) ZAddXX(key string, z ...redis.Z) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_zaddxx", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZAddXX(key, z...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZAddXX %v %v", key, z))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZAddXX %v", key))
	}

	return
}

func (c *TraceClient) ZAddXXCh(key string, z ...redis.Z) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_zaddxxch", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZAddXXCh(key, z...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZAddXXCh %v %v", key, z))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZAddXXCh %v", key))
	}

	return
}

func (c *TraceClient) ZCard(key string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_zcard", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZCard(key)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZCard %v", key))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZCard %v", key))
	}

	return
}

func (c *TraceClient) ZCount(key string, val string, val1 string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_zcount", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZCount(key, val, val1)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZCount %v %v %v", key, val, val1))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZCount %v", key))
	}

	return
}

func (c *TraceClient) ZIncr(key string, v redis.Z) (cmd *redis.FloatCmd) {
	span := c.tracer.StartSpan("redis_zincr", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZIncr(key, v)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZIncr %v %v", key, v))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZIncr %v", key))
	}

	return
}

func (c *TraceClient) ZIncrBy(key string, f float64, val string) (cmd *redis.FloatCmd) {
	span := c.tracer.StartSpan("redis_zincrby", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZIncrBy(key, f, val)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZIncrBy %v %v %v", key, f, val))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZIncrBy %v", key))
	}

	return
}

func (c *TraceClient) ZIncrNX(key string, v redis.Z) (cmd *redis.FloatCmd) {
	span := c.tracer.StartSpan("redis_zincrnx", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZIncrNX(key, v)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZIncrNX %v %v", key, v))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZIncrNX %v", key))
	}

	return
}

func (c *TraceClient) ZIncrXX(key string, v redis.Z) (cmd *redis.FloatCmd) {
	span := c.tracer.StartSpan("redis_zincrxx", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZIncrXX(key, v)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZIncrXX %v %v", key, v))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZIncrXX %v", key))
	}

	return
}

func (c *TraceClient) ZInterStore(key string, v redis.ZStore, vals ...string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_zinterstore", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZInterStore(key, v, vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZInterStore %v %v %v", key, v, vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZInterStore %v", key))
	}

	return
}

func (c *TraceClient) ZLexCount(key string, val string, val1 string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_zlexcount", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZLexCount(key, val, val1)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZLexCount %v %v %v", key, val, val1))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZLexCount %v", key))
	}

	return
}

func (c *TraceClient) ZPopMax(key string, vals ...int64) (cmd *redis.ZSliceCmd) {
	span := c.tracer.StartSpan("redis_zpopmax", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZPopMax(key, vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZPopMax %v %v", key, vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZPopMax %v", key))
	}

	return
}

func (c *TraceClient) ZPopMin(key string, vals ...int64) (cmd *redis.ZSliceCmd) {
	span := c.tracer.StartSpan("redis_zpopmin", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZPopMin(key, vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZPopMin %v %v", key, vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZPopMin %v", key))
	}

	return
}

func (c *TraceClient) ZRange(key string, i int64, i1 int64) (cmd *redis.StringSliceCmd) {
	span := c.tracer.StartSpan("redis_zrange", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZRange(key, i, i1)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZRange %v %v %v", key, i, i1))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZRange %v", key))
	}

	return
}

func (c *TraceClient) ZRangeByLex(key string, v redis.ZRangeBy) (cmd *redis.StringSliceCmd) {
	span := c.tracer.StartSpan("redis_zrangebylex", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZRangeByLex(key, v)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZRangeByLex %v %v", key, v))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZRangeByLex %v", key))
	}

	return
}

func (c *TraceClient) ZRangeByScore(key string, v redis.ZRangeBy) (cmd *redis.StringSliceCmd) {
	span := c.tracer.StartSpan("redis_zrangebyscore", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZRangeByScore(key, v)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZRangeByScore %v %v", key, v))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZRangeByScore %v", key))
	}

	return
}

func (c *TraceClient) ZRangeByScoreWithScores(key string, v redis.ZRangeBy) (cmd *redis.ZSliceCmd) {
	span := c.tracer.StartSpan("redis_zrangebyscorewithscores", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZRangeByScoreWithScores(key, v)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZRangeByScoreWithScores %v %v", key, v))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZRangeByScoreWithScores %v", key))
	}

	return
}

func (c *TraceClient) ZRangeWithScores(key string, i int64, i1 int64) (cmd *redis.ZSliceCmd) {
	span := c.tracer.StartSpan("redis_zrangewithscores", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZRangeWithScores(key, i, i1)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZRangeWithScores %v %v %v", key, i, i1))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZRangeWithScores %v", key))
	}

	return
}

func (c *TraceClient) ZRank(key string, val string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_zrank", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZRank(key, val)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZRank %v %v", key, val))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZRank %v", key))
	}

	return
}

func (c *TraceClient) ZRem(key string, vals ...interface{}) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_zrem", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZRem(key, vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZRem %v %v", key, vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZRem %v", key))
	}

	return
}

func (c *TraceClient) ZRemRangeByLex(key string, val string, val1 string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_zremrangebylex", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZRemRangeByLex(key, val, val1)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZRemRangeByLex %v %v %v", key, val, val1))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZRemRangeByLex %v", key))
	}

	return
}

func (c *TraceClient) ZRemRangeByRank(key string, i int64, i1 int64) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_zremrangebyrank", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZRemRangeByRank(key, i, i1)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZRemRangeByRank %v %v %v", key, i, i1))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZRemRangeByRank %v", key))
	}

	return
}

func (c *TraceClient) ZRemRangeByScore(key string, val string, val1 string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_zremrangebyscore", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZRemRangeByScore(key, val, val1)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZRemRangeByScore %v %v %v", key, val, val1))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZRemRangeByScore %v", key))
	}

	return
}

func (c *TraceClient) ZRevRange(key string, i int64, i1 int64) (cmd *redis.StringSliceCmd) {
	span := c.tracer.StartSpan("redis_zrevrange", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZRevRange(key, i, i1)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZRevRange %v %v %v", key, i, i1))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZRevRange %v", key))
	}

	return
}

func (c *TraceClient) ZRevRangeByLex(key string, v redis.ZRangeBy) (cmd *redis.StringSliceCmd) {
	span := c.tracer.StartSpan("redis_zrevrangebylex", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZRevRangeByLex(key, v)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZRevRangeByLex %v %v", key, v))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZRevRangeByLex %v", key))
	}

	return
}

func (c *TraceClient) ZRevRangeByScore(key string, v redis.ZRangeBy) (cmd *redis.StringSliceCmd) {
	span := c.tracer.StartSpan("redis_zrevrangebyscore", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZRevRangeByScore(key, v)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZRevRangeByScore %v %v", key, v))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZRevRangeByScore %v", key))
	}

	return
}

func (c *TraceClient) ZRevRangeByScoreWithScores(key string, v redis.ZRangeBy) (cmd *redis.ZSliceCmd) {
	span := c.tracer.StartSpan("redis_zrevrangebyscorewithscores", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZRevRangeByScoreWithScores(key, v)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZRevRangeByScoreWithScores %v %v", key, v))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZRevRangeByScoreWithScores %v", key))
	}

	return
}

func (c *TraceClient) ZRevRangeWithScores(key string, i int64, i1 int64) (cmd *redis.ZSliceCmd) {
	span := c.tracer.StartSpan("redis_zrevrangewithscores", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZRevRangeWithScores(key, i, i1)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZRevRangeWithScores %v %v %v", key, i, i1))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZRevRangeWithScores %v", key))
	}

	return
}

func (c *TraceClient) ZRevRank(key string, val string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_zrevrank", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZRevRank(key, val)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZRevRank %v %v", key, val))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZRevRank %v", key))
	}

	return
}

func (c *TraceClient) ZScan(key string, n uint64, val string, i int64) (cmd *redis.ScanCmd) {
	span := c.tracer.StartSpan("redis_zscan", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZScan(key, n, val, i)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZScan %v %v %v %v", key, n, val, i))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZScan %v", key))
	}

	return
}

func (c *TraceClient) ZScore(key string, val string) (cmd *redis.FloatCmd) {
	span := c.tracer.StartSpan("redis_zscore", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZScore(key, val)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZScore %v %v", key, val))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZScore %v", key))
	}

	return
}

func (c *TraceClient) ZUnionStore(key string, v redis.ZStore, vals ...string) (cmd *redis.IntCmd) {
	span := c.tracer.StartSpan("redis_zunionstore", opentracing.ChildOf(c.spanCtx))
	defer span.Finish()

	ext.DBType.Set(span, "redis")
	ext.DBInstance.Set(span, strconv.Itoa(c.config.DB))
	ext.PeerService.Set(span, "redis")
	ext.PeerHostname.Set(span, c.config.Addr)
	ext.SpanKindRPCClient.Set(span)

	cmd = c.Client.ZUnionStore(key, v, vals...)
	if err := cmd.Err(); err != nil {
		ext.DBStatement.Set(span, fmt.Sprintf("ZUnionStore %v %v %v", key, v, vals))

		if c.includeNotFound || err != redis.Nil {
			ext.Error.Set(span, true)
			span.LogKV("event", "error", "message", err.Error())
		}
	} else {
		ext.DBStatement.Set(span, fmt.Sprintf("ZUnionStore %v", key))
	}

	return
}
