package mylog

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type LogLevel uint16

const (
	UNKNOW LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

func parseLogLevel(s string) (LogLevel, error) {
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
		err := errors.New("无效日志级别")
		return UNKNOW, err
	}
}

// Logger 结构体
type Logger struct {
	Level LogLevel
}

// NewLogger 构造函数
func NewLogger(levelStr string) Logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		Level: level,
	}
}

func (l Logger) Debug(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [Debug] %s\n", now.Format("2006-01-02 00:00:00"), msg)
}
func (l Logger) Trace(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [Trace] %s\n", now.Format("2006-01-02 00:00:00"), msg)
}
func (l Logger) Info(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [Info] %s\n", now.Format("2006-01-02 00:00:00"), msg)
}
func (l Logger) Warning(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [Warning] %s\n", now.Format("2006-01-02 00:00:00"), msg)
}
func (l Logger) Error(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [Error] %s\n", now.Format("2006-01-02 00:00:00"), msg)
}
func (l Logger) Fatal(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [Fatal] %s\n", now.Format("2006-01-02 00:00:00"), msg)
}
