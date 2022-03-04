package trace_redis

import (
	"fmt"
	"runtime"

	"github.com/fighthorse/redisAdmin/pkg/conf"
)

// Global redis manager
var (
	RedisMgr  *Manager
	marooning = ManagerConfig{}
)

const (
	MaxDialTimeout  = 1000 // millisecond
	MaxReadTimeout  = 1000 // millisecond
	MaxWriteTimeout = 3000 // millisecond
	MaxPoolSize     = 1024
	MaxPoolTimeout  = 2 // second
	MinIdleConns    = 3
	MaxRetries      = 1
)

// A config of go redis
type Config struct {
	Network              string `yaml:"network"`
	Addr                 string `yaml:"addr"`
	Passwd               string `yaml:"password"`
	DB                   int    `yaml:"database"`
	DialTimeout          int    `yaml:"dial_timeout"`
	ReadTimeout          int    `yaml:"read_timeout"`
	WriteTimeout         int    `yaml:"write_timeout"`
	PoolSize             int    `yaml:"pool_size"`
	PoolTimeout          int    `yaml:"pool_timeout"`
	MinIdleConns         int    `yaml:"min_idle_conns"`
	MaxRetries           int    `yaml:"max_retries"`
	TraceIncludeNotFound bool   `yaml:"trace_include_not_found"`
}

// Name returns client name of the config
func (c *Config) Name() string {
	return fmt.Sprintf("%s(%s/%d)", c.Network, c.Addr, c.DB)
}

// FillWithDefaults apply default values for fields with invalid values.
func (c *Config) FillWithDefaults() {
	maxCPU := runtime.NumCPU()

	if c.DialTimeout <= 0 || c.DialTimeout > MaxDialTimeout*maxCPU {
		c.DialTimeout = MaxDialTimeout
	}

	if c.ReadTimeout <= 0 || c.ReadTimeout > MaxReadTimeout*maxCPU {
		c.ReadTimeout = MaxReadTimeout
	}

	if c.WriteTimeout <= 0 || c.WriteTimeout > MaxWriteTimeout*maxCPU {
		c.WriteTimeout = MaxWriteTimeout
	}

	if c.PoolSize <= 0 {
		c.PoolSize = 10 * maxCPU
	}

	if c.PoolTimeout <= 0 || c.PoolTimeout > MaxPoolTimeout*maxCPU {
		c.PoolTimeout = MaxPoolTimeout
	}

	if c.MinIdleConns <= 0 || c.MinIdleConns > MinIdleConns*maxCPU {
		c.MinIdleConns = MinIdleConns
	}

	if c.MaxRetries < 0 || c.MaxRetries > MaxRetries*maxCPU {
		c.MaxRetries = MaxRetries
	}
}

// A ManagerConfig defines a list of redis config with its name
type ManagerConfig map[string]*Config

func convertToConfig(c conf.Redis) *Config {
	ret := &Config{}
	fn := func(s float64) int {
		return int(s * 1000)
	}
	ret.Addr = c.Addr
	ret.Passwd = c.Pwd
	ret.DB = int(c.Db)
	ret.DialTimeout = fn(c.DialTimeout)
	ret.ReadTimeout = fn(c.ReadTimeout)
	ret.WriteTimeout = fn(c.WriteTimeout)
	ret.PoolSize = int(c.PoolSize)
	ret.MinIdleConns = int(c.MinIdleConns)
	ret.MaxRetries = int(c.MaxRetries)
	return ret
}

func InitCfg(cfg []conf.Redis) {
	if cfg == nil || len(cfg) <= 0 {
		return
	}
	for _, v := range cfg {
		c := convertToConfig(v)
		marooning[v.Name] = c
	}
	RedisMgr = NewManager(&marooning)
}

func AddCfg(v conf.Redis) {
	c := convertToConfig(v)
	marooning[v.Name] = c
	RedisMgr = NewManager(&marooning)
}

func ListCfg() map[string]interface{} {
	out := RedisMgr.List(&marooning)
	return out
}

func NewClient(schema string) *RedisClient {
	client, err := RedisMgr.NewClient(schema)
	if err != nil {
		if err == ErrNotFoundConfig {
			panic(fmt.Errorf("redis no such config section: %s", schema))
		}
		panic(fmt.Errorf("new redis client error: %s", err.Error()))
	}

	addr := marooning[schema].Addr
	db := marooning[schema].DB
	return &RedisClient{client, schema, addr, db}
}

func NewClientDb(schema string, db int) *RedisClient {
	client, err := RedisMgr.NewClient(schema)
	if err != nil {
		if err == ErrNotFoundConfig {
			panic(fmt.Errorf("redis no such config section: %s", schema))
		}
		panic(fmt.Errorf("new redis client error: %s", err.Error()))
	}
	clientNew, err := client.Select(db)
	addr := marooning[schema].Addr
	return &RedisClient{clientNew, schema, addr, db}
}
