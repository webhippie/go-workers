package workers

import (
	"fmt"
	"runtime"
	"time"

	"github.com/go-kit/kit/log/level"
)

type MiddlewareLogging struct{}

func (l *MiddlewareLogging) Call(queue string, message *Msg, next func() bool) (acknowledge bool) {
	prefix := fmt.Sprint(queue, " JID-", message.Jid())

	start := time.Now()

	level.Debug(Logger).Log(
		"msg", "job started",
		"prefix", prefix,
		"args", message.Args().ToJson(),
	)

	defer func() {
		if e := recover(); e != nil {
			buf := make([]byte, 4096)
			buf = buf[:runtime.Stack(buf, false)]

			level.Error(Logger).Log(
				"msg", "job paniced",
				"prefix", prefix,
				"duration", time.Since(start),
				"err", e,
				"stack", buf,
			)

			panic(e)
		}
	}()

	acknowledge = next()

	level.Debug(Logger).Log(
		"msg", "job done",
		"prefix", prefix,
		"duration", time.Since(start),
	)

	return
}
