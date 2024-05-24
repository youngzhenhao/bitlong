package services

import (
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

func (ml *ServicesLogger) Debug(message string) {
	if ml.level >= DEBUG {
		ml.logger.Printf("[Debug]: %s\n", message)
	}
}

func (ml *ServicesLogger) Info(message string) {
	if ml.level >= INFO {
		ml.logger.Printf("[Log]: %s\n", message)
	}
}

func (ml *ServicesLogger) Warning(message string) {
	if ml.level >= WARNING {
		ml.logger.Printf("[Warning]: %s\n", message)
	}
}

func (ml *ServicesLogger) Error(message string) {
	if ml.level >= ERROR {
		ml.logger.Printf("[Error]: %s\n", message)
	}
}

var (
	CUST                  = NewLogger("CUST", ERROR)
	FairLaunchDebugLogger = NewLogger("FLDL", ERROR)
)
