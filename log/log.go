package log

// Logger defines a common logging interface.
type Logger interface {
	Debug(msg string, keyvals ...interface{})
	Error(msg string, keyvals ...interface{})
	Info(msg string, keyvals ...interface{})
}
