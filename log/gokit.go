package log

import (
	"os"

	gokitlog "github.com/go-kit/kit/log"
	gokitlevel "github.com/go-kit/kit/log/level"
)

func NewGokit(target *os.File) Gokit {
	return Gokit{
		Logger: gokitlog.NewLogfmtLogger(gokitlog.NewSyncWriter(target)),
	}
}

type Gokit struct {
	Logger gokitlog.Logger
}

func (l Gokit) Debug(msg string, keyvals ...interface{}) {
	gokitlevel.Debug(gokitlog.With(l.Logger, keyvals)).Log("msg", msg)
}

func (l Gokit) Info(msg string, keyvals ...interface{}) {
	gokitlevel.Info(gokitlog.With(l.Logger, keyvals)).Log("msg", msg)
}

func (l Gokit) Error(msg string, keyvals ...interface{}) {
	gokitlevel.Error(gokitlog.With(l.Logger, keyvals)).Log("msg", msg)
}
