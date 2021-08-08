package mylogger

import "strings"

//代表日志级别
type Level uint64

const(
	DebugLevel Level = iota
	InfoLevel
	WarningLevel
	ErrorLevel
	FatalLevel
)


type Logger interface {
	Debug(format string,args ...interface{})
	Info(format string,args ...interface{})
	Warn(format string,args ...interface{})
	Error(format string,args ...interface{})
	Fatal(format string,args ...interface{})
	Close()

}









//根据传的level 获得字符串
func getLevelStr(level Level)string {
	switch level {
	case DebugLevel:
		return "Debug"
	case InfoLevel:
		return "INFO"
	case WarningLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case FatalLevel:
		return "FATAL"
	default:
		return "DEBUG"
	}
}
func getLevelByStr(levelStr string)Level{
	levelStr = strings.ToLower(levelStr)
	switch levelStr {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warn":
		return WarningLevel
	case "error":
		return ErrorLevel
	case "fatal":
		return FatalLevel
	default:
		return DebugLevel
	}
}