package workers

// Driver defines the interface for queue drivers.
type Driver interface {
	Ping() error
	Close() error
}
