package redis

var (
	defaultHost    = "localhost"
	defaultPort    = 6379
	defaultTimeout = 30
	defaultMaxIdle = 3
)

type Config struct {
	Host     string
	Port     int
	Password string
	Database int
	Timeout  int
	MaxIdle  int
}

func (c *Config) WithDefaults() *Config {
	if c.Host == "" {
		c.Host = defaultHost
	}

	if c.Port == 0 {
		c.Port = defaultPort
	}

	if c.Timeout == 0 {
		c.Timeout = defaultTimeout
	}

	if c.MaxIdle == 0 {
		c.MaxIdle = defaultMaxIdle
	}

	return c
}
