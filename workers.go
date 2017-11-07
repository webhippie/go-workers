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
	w := &Workers{}

	for _, opt := range opts {
		opt(w)
	}

	if w.config == nil {
		return nil, ErrNoConfigProvided
	}

	if w.driver == nil {
		return nil, ErrNoDriverProvided
	}

	if w.logger == nil {
		w.logger = log.NewGokit(os.Stdout)
	}

	return w, nil
}

type Workers struct {
	config *Config
	logger log.Logger
	driver Driver
}

func (w *Workers) Start() error {
	if err := w.driver.Ping(); err != nil {
		return err
	}

	return nil
}

func (w *Workers) Stop() error {
	return nil
}
