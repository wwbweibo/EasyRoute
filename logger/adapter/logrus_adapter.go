package adapter

import "github.com/sirupsen/logrus"

type LogrusAdapter struct {
}

func (l LogrusAdapter) Error(message string, err ...error) {
	logrus.Error(message, err)
}

func (l LogrusAdapter) Warning(message string, err ...error) {
	logrus.Warning(message, err)
}

func (l LogrusAdapter) Info(message string, err ...error) {
	logrus.Info(message, err)
}

func (l LogrusAdapter) Debug(message string, err ...error) {
	logrus.Debug(message, err)
}
