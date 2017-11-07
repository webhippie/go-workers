package redis

import (
	"errors"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/webhippie/workers/log"
)

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

	return r
}

type Redis struct {
	config *Config
	logger log.Logger
	handle *redis.Pool
}

func (r Redis) Connect() error {
	r.handle = &redis.Pool{
		MaxIdle:     r.config.MaxIdle,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", net.JoinHostPort(r.config.Host, strconv.Itoa(r.config.Port)))

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
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			_, err := conn.Do("PING")

			return err
		},
	}

	for i := 0; i < r.config.Timeout; i++ {
		conn := r.handle.Get()
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
