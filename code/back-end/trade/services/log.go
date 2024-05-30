package services

import (
	"fmt"
	"log"
	"os"
	"runtime"
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
		logger: log.New(os.Stdout, "["+logName+"]: ", log.Ldate|log.Ltime),
		level:  level}
}

func (ml *ServicesLogger) Debug(message string, v ...interface{}) {
	if ml.level >= DEBUG {
		_, callerFile, _, _ := runtime.Caller(1)
		msg := fmt.Sprintf(message, v...)
		ml.logger.Printf(" %s [Debug]: %s\n", callerFile, msg)
	}
}

func (ml *ServicesLogger) Info(message string, v ...any) {
	if ml.level >= INFO {
		_, callerFile, _, _ := runtime.Caller(1)
		msg := fmt.Sprintf(message, v...)
		ml.logger.Printf(" %s [Log]: %s\n", callerFile, msg)
	}
}

func (ml *ServicesLogger) Warning(message string, v ...interface{}) {
	if ml.level >= WARNING {
		_, callerFile, _, _ := runtime.Caller(1)
		msg := fmt.Sprintf(message, v...)
		ml.logger.Printf(" %s [Warning]: %s\n", callerFile, msg)
	}
}

func (ml *ServicesLogger) Error(message string, v ...any) {
	if ml.level >= ERROR {
		_, callerFile, _, _ := runtime.Caller(1)
		msg := fmt.Sprintf(message, v...)
		ml.logger.Printf(" %s [Error]: %s\n", callerFile, msg)
	}
}

var (
	CUST                  = NewLogger("CUST", ERROR)
	FairLaunchDebugLogger = NewLogger("FLDL", ERROR)
	FEE                   = NewLogger("FEE", ERROR)
)
