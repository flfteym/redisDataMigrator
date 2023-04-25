// package logger fileLog 输出至文件
package logger

import (
	"fmt"
	"os"
	"path"
	"time"
)

type FileLogger struct {
	Level       MyLogLevel
	filePath    string
	fileName    string
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
}

func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}

	fl := &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}

	err = fl.initFile()
	if err != nil {
		panic(err)
	}
	return fl
}

func (l *FileLogger) initFile() error {
	fullFileName := path.Join(l.filePath, l.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Open log file failed, err: %v\n", err)
		return err
	}

	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Open err log file failed, err: %v\n", err)
		return err
	}

	l.fileObj = fileObj
	l.errFileObj = errFileObj
	return nil
}

func (l *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("Get file info failed, err: %v\n", err)
		return false
	}
	return fileInfo.Size() >= l.maxFileSize
}

// 切割日志文件
func (l *FileLogger) splitFile(file *os.File) (*os.File, error) {
	// 备份
	nowStr := time.Now().Format("20060102150405000")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("Get file info failed, err: %v\n", err)
		return nil, err
	}

	logName := path.Join(l.filePath, fileInfo.Name())
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr)
	err = file.Close()
	if err != nil {
		return nil, err
	}
	err = os.Rename(logName, newLogName)
	if err != nil {
		return nil, err
	}

	// 打开新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Open new log file fail err = %v\n", err)
		return nil, err
	}

	return fileObj, nil
}

func (l *FileLogger) log(lv MyLogLevel, format string, a ...interface{}) {
	if l.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		if l.checkSize(l.fileObj) {
			newFile, err := l.splitFile(l.fileObj)
			if err != nil {
				return
			}
			l.fileObj = newFile
		}

		_, err := fmt.Fprintf(l.fileObj, "[%s] [%s] [%s : %s : %d] %s \n", now.Format("2006-01-02 15:04:05 "), getLogString(lv), fileName, funcName, lineNo, msg)
		if err != nil {
			panic(err)
		}

		if lv >= ERROR {
			if l.checkSize(l.errFileObj) {
				newFile, err := l.splitFile(l.errFileObj)
				if err != nil {
					return
				}
				l.errFileObj = newFile
			}

			_, err := fmt.Fprintf(l.errFileObj, "[%s] [%s] [%s : %s : %d] %s \n", now.Format("2006-01-02 15:04:05 "), getLogString(lv), fileName, funcName, lineNo, msg)
			if err != nil {
				panic(err)
			}
		}
	}
}

func (l *FileLogger) enable(logLevel MyLogLevel) bool {
	return l.Level <= logLevel
}

func (l *FileLogger) Debug(format string, a ...interface{}) {
	if l.enable(DEBUG) {
		l.log(DEBUG, format, a...)
	}
}

func (l *FileLogger) Trace(format string, a ...interface{}) {
	if l.enable(TRACE) {
		l.log(TRACE, format, a...)
	}
}

func (l *FileLogger) Info(format string, a ...interface{}) {
	if l.enable(TRACE) {
		l.log(INFO, format, a...)
	}
}

func (l *FileLogger) Warning(format string, a ...interface{}) {
	if l.enable(WARNING) {
		l.log(WARNING, format, a...)
	}
}

func (l *FileLogger) Error(format string, a ...interface{}) {
	if l.enable(ERROR) {
		l.log(ERROR, format, a...)
	}
}

func (l *FileLogger) Fatal(format string, a ...interface{}) {
	if l.enable(FATAL) {
		l.log(FATAL, format, a...)
	}
}

func (l *FileLogger) close() {
	err := l.fileObj.Close()
	if err != nil {
		panic(err)
	}
	err = l.errFileObj.Close()
	if err != nil {
		panic(err)
	}
}
