package logger

// Logger is Logger interface
type Logger interface {
	Error(message string, err ...error)
	Warning(message string, err ...error)
	Info(message string, err ...error)
	Debug(message string, err ...error)
}
