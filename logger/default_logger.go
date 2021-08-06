package logger

import (
	"fmt"
	llog "log"
)

type DefaultLogger struct {
}

func (logger *DefaultLogger) Fatal(format string, param ...interface{}) {
	llog.Fatalf("[FATAL] "+format+"\n", param...)
}

func (logger *DefaultLogger) Error(format string, param ...interface{}) {
	fmt.Printf("[Error] "+format+"\n", param...)
}

func (logger *DefaultLogger) Warning(format string, param ...interface{}) {
	fmt.Printf("[WARN] "+format+"\n", param...)
}

func (logger *DefaultLogger) Info(format string, param ...interface{}) {
	fmt.Printf("[INFO] "+format+"\n", param...)
}

func (logger *DefaultLogger) Debug(format string, param ...interface{}) {
	fmt.Printf("[DEBUG] "+format+"\n", param...)
}
