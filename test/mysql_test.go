package test

import (
	"github.com/jishulangcom/go-config"
	"github.com/jishulangcom/go-mysql"
	"testing"
)


func Test(t *testing.T) {
	conf := config.MysqlCnfDto{
		Host:         "127.0.0.1",
		Port:         "3306",
		User:         "root",
		Pwd:          "root",
		DbName:       "test",
		Charset:      "utf8mb4",
		Debug:        true,
	}
	mysql.NewDB(&conf)
	defer mysql.CloseDB()

	mysql.DB.SingularTable(true)
}
