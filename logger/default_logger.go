package logger

import "fmt"

type DefaultLogger struct {
}

func (logger *DefaultLogger) Error(message string, err ...error) {
	fmt.Printf("error: %s %s", message, err)
}

func (logger *DefaultLogger) Warning(message string, err ...error) {
	fmt.Printf("warning: %s %s", message, err)
}

func (logger *DefaultLogger) Info(message string, err ...error) {
	fmt.Printf("info: %s %s", message, err)
}

func (logger *DefaultLogger) Debug(message string, err ...error) {
	fmt.Printf("debug: %s %s", message, err)
}
