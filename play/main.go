package main

import (
	"fmt"
	"os"

	gokit "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/webhippie/workers"
	"github.com/webhippie/workers/driver/redis"
	"github.com/webhippie/workers/log"
)

func main() {
	logger := gokit.NewLogfmtLogger(gokit.NewSyncWriter(os.Stdout))

	logger = gokit.WithPrefix(logger,
		"app", "workers",
		"ts", gokit.DefaultTimestampUTC,
	)

	worker, err := workers.New(
		workers.WithDefaultConfig(),
		workers.WithLogger(log.Gokit{logger}),
		workers.WithDriver(
			redis.New(
				redis.WithDefaultConfig(),
				redis.WithLogger(log.Gokit{logger}),
			),
		),
	)

	if err != nil {
		level.Error(logger).Log(
			"err", err,
		)
	} else {
		fmt.Printf("%v", worker)
	}
}
