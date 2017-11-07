package workers

type Driver interface {
	Ping() error
}
