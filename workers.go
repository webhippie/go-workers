package workers

import (
	"errors"
	"os"

	"github.com/webhippie/workers/log"
)

var (
	ErrNoConfigProvided = errors.New("no configuration provided")
	ErrNoDriverProvided = errors.New("no driver provided")
)

func New(opts ...Option) (*Workers, error) {
	workers := &Workers{}

	for _, opt := range opts {
		opt(workers)
	}

	if workers.config == nil {
		return nil, ErrNoConfigProvided
	}

	if workers.driver == nil {
		return nil, ErrNoDriverProvided
	}

	if workers.logger == nil {
		workers.logger = log.NewGokit(os.Stdout)
	}

	if err := workers.driver.Connect(); err != nil {
		return nil, err
	}

	return workers, nil
}

type Workers struct {
	config *Config
	logger log.Logger
	driver Driver
}
