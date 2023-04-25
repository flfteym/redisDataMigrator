package logger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

type MyLogLevel uint16

// Logger 日志接口
type Logger interface {
	Debug(msgFormat string, a ...interface{})
	Trace(msgFormat string, a ...interface{})
	Info(msgFormat string, a ...interface{})
	Warning(msgFormat string, a ...interface{})
	Error(msgFormat string, a ...interface{})
	Fatal(msgFormat string, a ...interface{})
}

const (
	UNKNOWN MyLogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

func parseLogLevel(s string) (MyLogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("invalid logger level")
		return UNKNOWN, err
	}
}

func getLogString(lv MyLogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case FATAL:
		return "FATAL"
	}
	return "DEBUG"
}

func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime.Coller() failed")
	}

	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	funcName = strings.Split(funcName, ".")[1]
	return funcName, fileName, lineNo
}
