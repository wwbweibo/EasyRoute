package logger

var log = newLog()

func Error(format string, param ...interface{}) {
	log.logger.Error(format, param...)
}

func Warning(format string, param ...interface{}) {
	log.logger.Warning(format, param...)
}

func Info(format string, param ...interface{}) {
	log.logger.Info(format, param...)
}

func Debug(format string, param ...interface{}) {
	log.logger.Debug(format, param...)
}

func WithLogger(logger Logger) {
	log.logger = logger
}
