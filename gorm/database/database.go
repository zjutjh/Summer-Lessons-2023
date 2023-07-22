package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义一个全局变量
var DB *gorm.DB

func Init() {
	user := "root"
	pass := "123456"
	host := "127.0.0.1"
	port := "3306"
	DBname := "gorm_learn"

	// dsn data-source-name 告知数据库所在的位置以及数据库相关的属性
	// user:pass@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, DBname)

	// 连接数据库，获得 DB 类型实例，用于后面对数据库进行的操作
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, error = " + err.Error())
	}

	// 赋值给全局变量
	DB = db
}
