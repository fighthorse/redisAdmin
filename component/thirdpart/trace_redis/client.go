package trace_redis

import (
	"context"
	"sync"
	"time"

	"github.com/fighthorse/redisAdmin/component/log"
	goredis "github.com/go-redis/redis"
	"github.com/opentracing/opentracing-go"
)

// A Client wraps redis client with custom features
type Client struct {
	*goredis.Client

	mux    sync.RWMutex
	log    log.Logger
	config *Config
}

// New creates a new redis client with config given and a dummy logger.
func New(config *Config) (*Client, error) {
	return NewWithLogger(config, log.NewDummyLogger())
}

// NewWithLogger creates a new redis client with config and logger given.
func NewWithLogger(config *Config, log log.Logger) (*Client, error) {
	config.FillWithDefaults()

	client := goredis.NewClient(&goredis.Options{
		Network:      config.Network,
		Addr:         config.Addr,
		Password:     config.Passwd,
		DB:           config.DB,
		DialTimeout:  time.Duration(config.DialTimeout) * time.Millisecond,
		ReadTimeout:  time.Duration(config.ReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(config.WriteTimeout) * time.Millisecond,
		PoolSize:     config.PoolSize,
		PoolTimeout:  time.Duration(config.PoolTimeout) * time.Second,
		MinIdleConns: config.MinIdleConns,
		MaxRetries:   config.MaxRetries,
	})

	return &Client{
		Client: client,
		config: config,
		log:    log,
	}, nil
}

// Select changes db by coping out a new client.
//
// NOTE: There maybe a deadlock if internal invocations panic!!!
func (c *Client) Select(db int) (*Client, error) {
	c.mux.RLock()

	opts := c.Options()
	if opts.DB == db {
		c.mux.RUnlock()

		return c, nil
	}

	c.mux.RUnlock()

	// creates a new client
	c.mux.Lock()
	defer c.mux.Unlock()

	config := &Config{
		Network:      opts.Network,
		Addr:         opts.Addr,
		Passwd:       opts.Password,
		DB:           db,
		DialTimeout:  int(opts.DialTimeout / time.Millisecond),
		ReadTimeout:  int(opts.ReadTimeout / time.Millisecond),
		WriteTimeout: int(opts.WriteTimeout / time.Millisecond),
		PoolSize:     opts.PoolSize,
		PoolTimeout:  int(opts.PoolTimeout / time.Second),
		MinIdleConns: opts.MinIdleConns,
		MaxRetries:   opts.MaxRetries,
	}

	name := config.Name()

	// first, try loading a client from default manager
	client, err := DefaultMgr.NewClientWithLogger(name, c.log)
	if err == nil {
		return client, nil
	}
	c.log.Warnf("DefaultMgr.NewClient(%s): %v", name, err)

	// second, register new client with default manager
	DefaultMgr.Add(name, config)

	return DefaultMgr.NewClientWithLogger(name, c.log)
}

// Trace creates a new redis client with tracer.
func (c *Client) Trace(ctx context.Context, tracers ...opentracing.Tracer) *TraceClient {
	if ctx == nil {
		return c.TraceWithSpanContext(nil, tracers...)
	}

	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		return c.TraceWithSpanContext(span.Context(), tracers...)
	}

	return c.TraceWithSpanContext(nil, tracers...)
}

// ctx.Request => redisClient.TraceWithSpanContext(ctx.Request.Context())
func (c *Client) TraceWithSpanContext(ctx opentracing.SpanContext, tracers ...opentracing.Tracer) *TraceClient {
	var tracer opentracing.Tracer
	if len(tracers) > 0 {
		tracer = tracers[0]
	} else {
		tracer = opentracing.GlobalTracer()
	}

	return &TraceClient{
		Client:          c,
		tracer:          tracer,
		spanCtx:         ctx,
		includeNotFound: c.config.TraceIncludeNotFound,
	}
}
