package logger

var log = newLog()

func Error(message string, err ...error) {
	log.logger.Error(message, err...)
}

func Warning(message string, err ...error) {
	log.logger.Warning(message, err...)
}

func Info(message string, err ...error) {
	log.logger.Info(message, err...)
}

func Debug(message string, err ...error) {
	log.logger.Debug(message, err...)
}

func WithLogger(logger Logger) {
	log.logger = logger
}
