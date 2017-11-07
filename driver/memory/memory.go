package memory

import (
	"os"

	"github.com/webhippie/workers/log"
)

// New initializes a new memory driver.
func New(opts ...Option) *Memory {
	m := &Memory{}

	for _, opt := range opts {
		opt(m)
	}

	if m.config == nil {
		config := &Config{}
		m.config = config.WithDefaults()
	}

	if m.logger == nil {
		m.logger = log.NewGokit(os.Stdout)
	}

	return m
}

// Memory defines a simple struct with all required actions.
type Memory struct {
	config *Config
	logger log.Logger
}

// Ping implements the driver interface.
func (m Memory) Ping() error {
	return nil
}

// Close implements the driver interface.
func (m Memory) Close() error {
	return nil
}
