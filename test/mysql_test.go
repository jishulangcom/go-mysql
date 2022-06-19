package test

import (
	"fmt"
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
		DbName:       "information_schema",
		Charset:      "utf8mb4",
		Debug:        true,
	}
	mysql.NewDB(&conf)
	defer mysql.CloseDB()

	//
	_, err := mysql.DB.Query("select * from CHARACTER_SETS")
	if err != nil {
		fmt.Println(err)
		return
	}
}
