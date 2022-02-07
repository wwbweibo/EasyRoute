package log

type logger struct {
	logger Logger
}

func newLog() *logger {
	return &logger{
		logger: &DefaultLogger{},
	}
}

func (log *logger) Fatal(format string, param ...interface{}) {
	log.logger.Fatal(format, param...)
}

func (log *logger) Error(format string, param ...interface{}) {
	log.logger.Error(format, param...)
}

func (log *logger) Warning(format string, param ...interface{}) {
	log.logger.Warning(format, param...)
}

func (log *logger) Info(format string, param ...interface{}) {
	log.logger.Info(format, param...)
}

func (log *logger) Debug(format string, param ...interface{}) {
	log.logger.Debug(format, param...)
}
func (log *logger) WithLogger(logger Logger) {
	log.logger = logger
}
