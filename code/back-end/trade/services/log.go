package services

import (
	"fmt"
	"log"
	"os"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
)

type ServicesLogger struct {
	logger *log.Logger
	level  LogLevel
}

func NewLogger(logName string, level LogLevel) *ServicesLogger {
	return &ServicesLogger{
		logger: log.New(os.Stdout, "["+logName+"]: ", log.Ldate|log.Ltime|log.Lshortfile),
		level:  level}
}

func (ml *ServicesLogger) Debug(message string, v ...interface{}) {
	if ml.level >= DEBUG {
		msg := fmt.Sprintf(message, v...)
		ml.logger.Printf("[Debug]: %s\n", msg)
	}
}

func (ml *ServicesLogger) Info(message string, v ...interface{}) {
	if ml.level >= INFO {
		msg := fmt.Sprintf(message, v...)
		ml.logger.Printf("[Log]: %s\n", msg)
	}
}

func (ml *ServicesLogger) Warning(message string, v ...interface{}) {
	if ml.level >= WARNING {
		msg := fmt.Sprintf(message, v...)
		ml.logger.Printf("[Warning]: %s\n", msg)
	}
}

func (ml *ServicesLogger) Error(message string, v ...interface{}) {
	if ml.level >= ERROR {
		msg := fmt.Sprintf(message, v...)
		ml.logger.Printf("[Error]: %s\n", msg)
	}
}

var (
	CUST = NewLogger("CUST", ERROR)
)
