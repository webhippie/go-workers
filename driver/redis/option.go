package redis

import (
	"github.com/garyburd/redigo/redis"
	"github.com/webhippie/workers/log"
)

type Option func(*Redis)

func WithLogger(logger log.Logger) Option {
	return func(r *Redis) {
		r.logger = logger
	}
}

func WithPool(pool *redis.Pool) Option {
	return func(r *Redis) {
		r.pool = pool
	}
}

func WithConfig(config *Config) Option {
	return func(r *Redis) {
		r.config = config.WithDefaults()
	}
}

func WithDefaultConfig() Option {
	return func(r *Redis) {
		config := &Config{}
		r.config = config.WithDefaults()
	}
}
