package memory

import (
	"os"

	"github.com/webhippie/workers/log"
)

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

type Memory struct {
	config *Config
	logger log.Logger
}

func (m Memory) Connect() error {
	return nil
}
