package trace_redis

import "errors"

// errors
var (
	ErrNotFoundConfig = errors.New("no config found")
	ErrInvalidConfig  = errors.New("named config is not a valid *Config type")
)
