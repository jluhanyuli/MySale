package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

//文件日志结构体
type FileLogger struct{
	level Level
	fileName string
	filepath string
	file *os.File
	errFile *os.File
	maxSize int64
}

func NewFileLogger(levelStr,fileName,filePath string)(*FileLogger){
	f:= &FileLogger{
		level: getLevelByStr(levelStr),
		fileName: fileName,
		filepath: filePath,
		maxSize: 10*1024*1024,
	}
	f.initFile()
	return f
}

func (f *FileLogger)initFile(){
	logName := path.Join(f.filepath,f.fileName)

	fileObj,err:=os.OpenFile(logName,os.O_CREATE|os.O_WRONLY|os.O_APPEND,0644)
	if err!=nil{
		panic(fmt.Errorf("open log file %s wrong",logName))
	}
	f.file=fileObj

	errLogName :=fmt.Sprintf("%s.err",logName)
	errFileObj,err:=os.OpenFile(errLogName,os.O_CREATE|os.O_WRONLY|os.O_APPEND,0644)
	if err!=nil{
		panic(fmt.Errorf("open log file %s wrong",logName))
	}
	f.errFile=errFileObj
}

func (f *FileLogger)Debug(format string,args ...interface{}){
	f.log(DebugLevel,format,args)
}

func (f *FileLogger)Info(format string,args ...interface{}){
	f.log(InfoLevel,format,args)
}

func (f *FileLogger)Warn(format string,args ...interface{}){
	f.log(WarningLevel,format,args)
}

func (f *FileLogger)Error(format string,args ...interface{}){
	f.log(ErrorLevel,format,args)
}
func (f *FileLogger)Fatal(format string,args ...interface{}){
	f.log(FatalLevel,format,args)
}
func (f *FileLogger)Close(){
	f.file.Close()
}

func (f *FileLogger)log(level Level,format string,args ...interface{}){
	if f.level>=level{
		return
	}
	msg := fmt.Sprintf(format,args...)
	nowStr:= time.Now().Format("2006-01-02 15:04:05.000")
	fileName,line,funcName := GetCallerInfo(3)
	levelStr:=getLevelStr(level)
	if f.checkNeedSplid(f.file){
		f.file=f.splitLogFile(f.file)
	}
	logMsg:=fmt.Sprintf("[%s][%s:%d][%s][%s]%s",
		nowStr,fileName,line,funcName,levelStr,msg)
	fmt.Fprintln(f.file,logMsg)
	//若有err或者fatal 记在err文件里
	if level>=ErrorLevel{
		if f.checkNeedSplid(f.file){
			f.errFile=f.splitLogFile(f.errFile)
		}
		fmt.Fprintln(f.errFile,logMsg)
	}
}

func (f *FileLogger)splitLogFile(file *os.File) *os.File {
	//写之前需要做一个检查，是否需要拆分文件
		fileName := f.file.Name()
		backupName := fmt.Sprintf("%s_%v.back", fileName, time.Now().Unix())
		//原文件备份与关闭
		f.file.Close()
		os.Rename(fileName, backupName)

		fileObj, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			panic(fmt.Errorf("open log file %s wrong", fileName))
		}
		return fileObj
}
func (f *FileLogger)checkNeedSplid(file *os.File)bool{
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	if fileSize >= f.maxSize {
		return  true
	}
	return false
}
