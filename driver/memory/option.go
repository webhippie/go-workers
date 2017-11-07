package memory

import (
	"github.com/webhippie/workers/log"
)

// Option is a simple return value for the configuration chain.
type Option func(*Memory)

// WithLogger injects a logger into the memory driver.
func WithLogger(logger log.Logger) Option {
	return func(m *Memory) {
		m.logger = logger
	}
}

// WithConfig injects a config into the memory driver.
func WithConfig(config *Config) Option {
	return func(m *Memory) {
		m.config = config.WithDefaults()
	}
}

// WithDefaultConfig injects a default config into the memory driver.
func WithDefaultConfig() Option {
	return func(m *Memory) {
		config := &Config{}
		m.config = config.WithDefaults()
	}
}
