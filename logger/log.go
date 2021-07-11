package logger

type logger struct {
	logger Logger
}

func newLog() *logger {
	return &logger{
		logger: &DefaultLogger{},
	}
}

func (log *logger) Error(message string, err ...error) {
	log.logger.Error(message, err...)
}

func (log *logger) Warning(message string, err ...error) {
	log.logger.Warning(message, err...)
}

func (log *logger) Info(message string, err ...error) {
	log.logger.Info(message, err...)
}

func (log *logger) Debug(message string, err ...error) {
	log.logger.Debug(message, err...)
}

func (log *logger) WithLogger(logger Logger) {
	log.logger = logger
}
