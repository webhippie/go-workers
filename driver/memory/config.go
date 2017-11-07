package memory

// Config represents the memory driver configuration.
type Config struct {
}

// WithDefaults initializes useful default config values.
func (c *Config) WithDefaults() *Config {
	return c
}
