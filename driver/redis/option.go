package redis

import (
	"github.com/webhippie/workers/log"
)

type Option func(*Redis)

func WithLogger(logger log.Logger) Option {
	return func(r *Redis) {
		r.logger = logger
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
