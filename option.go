package workers

import (
	"github.com/webhippie/workers/log"
)

// Option is a simple return value for the configuration chain.
type Option func(*Workers)

// WithLogger injects a logger into the workers.
func WithLogger(logger log.Logger) Option {
	return func(w *Workers) {
		w.logger = logger
	}
}

// WithDriver injects a driver into the workers.
func WithDriver(driver Driver) Option {
	return func(w *Workers) {
		w.driver = driver
	}
}

// WithConfig injects a config into the workers.
func WithConfig(config *Config) Option {
	return func(w *Workers) {
		w.config = config.WithDefaults()
	}
}

// WithDefaultConfig injects a default configuration into the workers.
func WithDefaultConfig() Option {
	return func(w *Workers) {
		config := &Config{}
		w.config = config.WithDefaults()
	}
}
