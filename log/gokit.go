package log

import (
	"os"

	gokitlog "github.com/go-kit/kit/log"
	gokitlevel "github.com/go-kit/kit/log/level"
)

// NewGokit initializes a new gokit logger with default options.
func NewGokit(target *os.File) Gokit {
	return Gokit{
		Logger: gokitlog.NewLogfmtLogger(gokitlog.NewSyncWriter(target)),
	}
}

// Gokit is a simple struct to implement the logger interface.
type Gokit struct {
	Logger gokitlog.Logger
}

// Debug implements the logger interface.
func (l Gokit) Debug(msg string, keyvals ...interface{}) {
	gokitlevel.Debug(gokitlog.With(l.Logger, keyvals)).Log("msg", msg)
}

// Info implements the logger interface.
func (l Gokit) Info(msg string, keyvals ...interface{}) {
	gokitlevel.Info(gokitlog.With(l.Logger, keyvals)).Log("msg", msg)
}

// Error implements the logger interface.
func (l Gokit) Error(msg string, keyvals ...interface{}) {
	gokitlevel.Error(gokitlog.With(l.Logger, keyvals)).Log("msg", msg)
}
