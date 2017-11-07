package redis

import (
	"github.com/garyburd/redigo/redis"
	"github.com/webhippie/workers/log"
)

// Option is a simple return value for the configuration chain.
type Option func(*Redis)

// WithLogger injects a logger into the redis driver.
func WithLogger(logger log.Logger) Option {
	return func(r *Redis) {
		r.logger = logger
	}
}

// WithPool injects a pool into the redis driver.
func WithPool(pool *redis.Pool) Option {
	return func(r *Redis) {
		r.pool = pool
	}
}

// WithConfig injects a config into the redis driver.
func WithConfig(config *Config) Option {
	return func(r *Redis) {
		r.config = config.WithDefaults()
	}
}

// WithDefaultConfig injects a default config into the redis driver.
func WithDefaultConfig() Option {
	return func(r *Redis) {
		config := &Config{}
		r.config = config.WithDefaults()
	}
}
