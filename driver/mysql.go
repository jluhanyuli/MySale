package driver

import "database/sql"

var db *sql.DB //声明一个全局的 db 变量

func InitMySQL() (db *sql.DB,err error) {
	dsn := "root:password@tcp(127.0.0.1:3306)/blog"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	return
}