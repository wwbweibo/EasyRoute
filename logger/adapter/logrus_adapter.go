package adapter

import "github.com/sirupsen/logrus"

type LogrusAdapter struct {
}

func (l LogrusAdapter) Fatal(format string, param ...interface{}) {
	logrus.Fatalf(format, param...)
}

func (l LogrusAdapter) Error(format string, param ...interface{}) {
	logrus.Errorf(format, param...)
}

func (l LogrusAdapter) Warning(format string, param ...interface{}) {
	logrus.Warningf(format, param...)
}

func (l LogrusAdapter) Info(format string, param ...interface{}) {
	logrus.Infof(format, param...)
}

func (l LogrusAdapter) Debug(format string, param ...interface{}) {
	logrus.Debugf(format, param...)
}
