package redis

import (
	"errors"
	"os"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/webhippie/workers/log"
)

// New initializes a new memory driver.
func New(opts ...Option) *Redis {
	r := &Redis{}

	for _, opt := range opts {
		opt(r)
	}

	if r.config == nil {
		config := &Config{}
		r.config = config.WithDefaults()
	}

	if r.logger == nil {
		r.logger = log.NewGokit(os.Stdout)
	}

	if r.pool == nil {
		r.pool = &redis.Pool{
			MaxIdle:     r.config.MaxIdle,
			IdleTimeout: 180 * time.Second,
			Dial: func() (redis.Conn, error) {
				conn, err := redis.Dial("tcp", r.config.Host)

				if err != nil {
					return nil, err
				}

				if r.config.Password != "" {
					if _, err := conn.Do("AUTH", r.config.Password); err != nil {
						conn.Close()
						return nil, err
					}
				}

				if r.config.Database >= 0 {
					if _, err := conn.Do("SELECT", r.config.Database); err != nil {
						conn.Close()
						return nil, err
					}
				}

				return conn, err
			},
		}
	}

	return r
}

// Redis defines a simple struct with all required actions.
type Redis struct {
	config *Config
	logger log.Logger
	pool   *redis.Pool
}

// Ping implements the driver interface.
func (r Redis) Ping() error {
	for i := 0; i < r.config.Timeout; i++ {
		conn := r.pool.Get()
		defer conn.Close()

		if conn.Err() == nil {
			_, err := conn.Do("PING")

			if err == nil {
				return nil
			}
		}

		r.logger.Info("redis ping failed, retry in 1s")
		time.Sleep(time.Second)
	}

	return errors.New("redis ping attempts failed")
}

// Close implements the driver interface.
func (r Redis) Close() error {
	return r.pool.Close()
}
