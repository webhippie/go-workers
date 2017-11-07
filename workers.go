package workers

import (
	"errors"
	"os"

	"github.com/webhippie/workers/log"
)

var (
	// ErrNoConfigProvided defines an error if no config is provided.
	ErrNoConfigProvided = errors.New("no configuration provided")

	// ErrNoDriverProvided defines an error if no driver is provided.
	ErrNoDriverProvided = errors.New("no driver provided")
)

// New initializes a new worker.
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

// Workers defines a simple struct with all required actions.
type Workers struct {
	config *Config
	logger log.Logger
	driver Driver
}

// Start simply starts the worker process.
func (w *Workers) Start() error {
	return w.driver.Ping()
}

// Stop simply stops the worker process.
func (w *Workers) Stop() error {
	return w.driver.Close()
}
