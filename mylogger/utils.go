package mylogger

import (
	"path"
	"runtime"
)

func GetCallerInfo(skip int)(fileName string,line int,funcName string){
	pc,fileName,line,ok:=runtime.Caller(skip)
	if !ok{
		return
	}
	fileName = path.Base(fileName)    //  a/b/c  拿到c
	funcName =runtime.FuncForPC(pc).Name()
	funcName =path.Base(fileName)
	return
}