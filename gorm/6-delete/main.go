package main

import (
	"gorm/database"

	"gorm.io/gorm"
)

type User struct {
	ID        uint
	Name      string
	Age       int
	DeletedAt gorm.DeletedAt
}

func main() {
	database.Init()
	database.DB.AutoMigrate(&User{})

	// 造数据
	users := []User{
		{Name: "惜寞", Age: 3},
		{Name: "青鸟", Age: 8},
		{Name: "竹林", Age: 5},
		{Name: "木木", Age: 9},
		{Name: "Sora", Age: 7},
		{Name: "桑葚", Age: 114514},
		{Name: "知更鸟", Age: 6},
		{Name: "花仙", Age: 4},
		{Name: "一一", Age: 10},
	}

	database.DB.Create(&users)

	// 利用结构体删除需要主键
	// database.DB.Delete(&User{Name: "惜寞", Age: 3})

	// 当你试图执行不带任何条件的批量删除时
	// GORM 将不会运行并返回 ErrMissingWhereClause 错误
	// database.DB.Delete(&User{})

	// 需要主键，否则也会返回 ErrMissingWhereClause 错误
	// database.DB.Delete(&[]User{{Name: "竹林"}, {Name: "木木"}})

	// 想要全局删除可以使用下面三种方法
	// 1.添加条件
	// database.DB.Where("1 = 1").Delete(&User{})

	// 2.原生 SQL
	// database.DB.Exec("DELETE FROM users")

	// 3.开启 AllowGlobalUpdate 模式
	// database.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&User{})

	// 软删除 gorm.DeletedAt
	// database.DB.Debug().Delete(&User{}, 6)

	// var user User
	// 软删除的记录在查询时会被忽略
	// database.DB.Debug().First(&user, 6)

	// 可以使用 Unscoped 来查询到被软删除的记录
	// database.DB.Unscoped().First(&user, 6)
	// fmt.Println(user)

	// 可以使用 Unscoped 来永久删除匹配的记录
	// database.DB.Unscoped().Delete(&User{}, 6)
}
