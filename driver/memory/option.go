package memory

import (
	"github.com/webhippie/workers/log"
)

type Option func(*Memory)

func WithLogger(logger log.Logger) Option {
	return func(m *Memory) {
		m.logger = logger
	}
}

func WithConfig(config *Config) Option {
	return func(m *Memory) {
		m.config = config.WithDefaults()
	}
}

func WithDefaultConfig() Option {
	return func(m *Memory) {
		config := &Config{}
		m.config = config.WithDefaults()
	}
}
