package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jishulangcom/go-config"
	"github.com/jmoiron/sqlx"
	"os"
)

var DB *sqlx.DB

func NewDB(conf *config.MysqlCnfDto) *sqlx.DB {
	var err error
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", conf.User, conf.Pwd, conf.Host, conf.Port, conf.DbName, conf.Charset)
	db, err := sqlx.Open("mysql", url)
	if err != nil {
		panic("mysql 连接失败")
		os.Exit(1)
	}

	if err = db.Ping(); err != nil {
		panic("mysql 连接失败")
		os.Exit(1)
	}

	DB = db
	return db
}

func CloseDB() {
	DB.Close()
}
