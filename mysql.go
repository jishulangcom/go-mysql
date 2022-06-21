package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/jishulangcom/go-config"
)

//单例，开启事务时用这个db(models.db)db.Begin(),db.Rollback(),db.Commit()
var DB *gorm.DB

//连接数据库，创建表
func NewDB(conf *config.MysqlCnfDto) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", conf.User, conf.Pwd, conf.Host, conf.Port, conf.DbName, conf.Charset)
	// database_connect_string := "root:root@/jsl-share?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	//打开数据库链接，返回db实例
	DB, err = gorm.Open("mysql", url)
	if err != nil {
		panic("mysql 连接失败")
	}

	// DB.SingularTable(true) //取消建表时的表名自动添加“s”
}

func CloseDB() {
	DB.Close()
}