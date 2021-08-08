package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"server/mylogger"
)




var logger mylogger.Logger
func main(){
	//db,err:=driver.InitMySQL()
	//if err != nil {
	//	fmt.Print(err.Error())
	//}
	//defer db.Close()
	//err = db.Ping()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	logger =mylogger.NewFileLogger("info","./","xxx.log")
	//logger1 :=mylogger.NewConsoleLogger("info")
	defer logger.Close()
	logger.Debug("%s的 Hello World","lyh")
	logger.Info("%s的 Hello World","lyh")
	logger.Error("%s的 Hello World","lyh")


	fmt.Println("hello_world")
}

