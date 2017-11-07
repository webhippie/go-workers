package workers

var (
	defaultIdentifier    = "local"
	defaultNamespace     = "workers"
	defaultPollIntervall = 15
)

type Config struct {
	Identifier    string
	Namespace     string
	PollIntervall int
}

func (c *Config) WithDefaults() *Config {
	if c.Identifier == "" {
		c.Identifier = defaultIdentifier
	}

	if c.Namespace == "" {
		c.Namespace = defaultNamespace
	}

	if c.PollIntervall == 0 {
		c.PollIntervall = defaultPollIntervall
	}

	return c
}
