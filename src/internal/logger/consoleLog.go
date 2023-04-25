// package logger consoleLog 输出至屏幕
package logger

import (
	"fmt"
	"time"
)

// ConsoleLogger 日志结构体
type ConsoleLogger struct {
	Level MyLogLevel
}

// NewConsoleLog 构造函数
func NewConsoleLog(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}

	return ConsoleLogger{
		level,
	}
}

func (l ConsoleLogger) enable(logLevel MyLogLevel) bool {
	return l.Level <= logLevel
}

func (l ConsoleLogger) log(lv MyLogLevel, msgFormat string, a ...interface{}) {
	if l.enable(DEBUG) {
		msg := fmt.Sprintf(msgFormat, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Printf("[%s] [%s] [%s : %s : %d] %s \n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
	}
}

func (l ConsoleLogger) Debug(msgFormat string, a ...interface{}) {
	l.log(DEBUG, msgFormat, a...)
}

func (l ConsoleLogger) Trace(msgFormat string, a ...interface{}) {
	l.log(TRACE, msgFormat, a...)
}

func (l ConsoleLogger) Info(msgFormat string, a ...interface{}) {
	l.log(INFO, msgFormat, a...)
}

func (l ConsoleLogger) Warning(msgFormat string, a ...interface{}) {
	l.log(WARNING, msgFormat, a...)
}

func (l ConsoleLogger) Error(msgFormat string, a ...interface{}) {
	l.log(ERROR, msgFormat, a...)
}

func (l ConsoleLogger) Fatal(msgFormat string, a ...interface{}) {
	l.log(FATAL, msgFormat, a...)
}
