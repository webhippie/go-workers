package workers

import (
	"github.com/webhippie/workers/log"
)

type Option func(*Workers)

func WithLogger(logger log.Logger) Option {
	return func(w *Workers) {
		w.logger = logger
	}
}

func WithDriver(driver Driver) Option {
	return func(w *Workers) {
		w.driver = driver
	}
}

func WithConfig(config *Config) Option {
	return func(w *Workers) {
		w.config = config.WithDefaults()
	}
}

func WithDefaultConfig() Option {
	return func(w *Workers) {
		config := &Config{}
		w.config = config.WithDefaults()
	}
}
