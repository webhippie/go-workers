package redis

var (
	defaultHost    = "localhost:6379"
	defaultPort    = 6379
	defaultTimeout = 30
	defaultMaxIdle = 3
)

// Config represents the redis driver configuration.
type Config struct {
	Host     string
	Password string
	Database int
	Timeout  int
	MaxIdle  int
}

// WithDefaults initializes useful default config values.
func (c *Config) WithDefaults() *Config {
	if c.Host == "" {
		c.Host = defaultHost
	}

	if c.Timeout == 0 {
		c.Timeout = defaultTimeout
	}

	if c.MaxIdle == 0 {
		c.MaxIdle = defaultMaxIdle
	}

	return c
}
