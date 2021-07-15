package logger

// Logger is Logger interface
type Logger interface {
	Fatal(format string, param ...interface{})
	Error(format string, param ...interface{})
	Warning(format string, param ...interface{})
	Info(format string, param ...interface{})
	Debug(format string, param ...interface{})
}
