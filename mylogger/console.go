package mylogger

import (
	"fmt"
	"os"
	"time"
)

//网终端记录日志
type ConsoleLogger struct {
	level Level
}

func NewConsoleLogger(levelStr string)(*ConsoleLogger){
	f:= &ConsoleLogger{
		level: getLevelByStr(levelStr),
	}
	return f
}


func (f *ConsoleLogger)Debug(format string,args ...interface{}){
	f.log(DebugLevel,format,args)
}

func (f *ConsoleLogger)Info(format string,args ...interface{}){
	f.log(InfoLevel,format,args)
}

func (f *ConsoleLogger)Warn(format string,args ...interface{}){
	f.log(WarningLevel,format,args)
}

func (f *ConsoleLogger)Error(format string,args ...interface{}){
	f.log(ErrorLevel,format,args)
}
func (f *ConsoleLogger)Fatal(format string,args ...interface{}){
	f.log(FatalLevel,format,args)
}
func (f *ConsoleLogger)Close(format string,args ...interface{}){
	f.log(FatalLevel,format,args)
}


func (f *ConsoleLogger)log(level Level,format string,args ...interface{}){
	if f.level > level{
		return
	}
	msg := fmt.Sprintf(format,args...)
	nowStr:= time.Now().Format("2006-01-02 15:04:05.000")
	fileName,line,funcName := GetCallerInfo(3)
	levelStr:=getLevelStr(level)
	logMsg:=fmt.Sprintf("[%s][%s:%d][%s][%s]%s",
		nowStr,fileName,line,funcName,levelStr,msg)
	fmt.Fprintln(os.Stdout,logMsg)

}
