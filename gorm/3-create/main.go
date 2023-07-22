package main

import (
	"fmt"
	"gorm/database"
)

type User struct {
	ID   uint
	Name string
	Age  int
	Sex  string
}

func main() {
	// 数据库初始化（连接）
	database.Init()
	// 自动迁移建表
	database.DB.AutoMigrate(&User{})

	user := User{Name: "XiMo", Age: 3, Sex: "male"}

	result := database.DB.Create(&user) // 通过数据的指针来创建

	fmt.Println(user.ID)             // 返回插入记录的主键
	fmt.Println(result.Error)        // 返回 error
	fmt.Println(result.RowsAffected) // 返回插入记录的条数

	// users := []User{
	// 	{Name: "惜寞", Age: 3, Sex: "male"},
	// 	{Name: "青鸟", Age: 19, Sex: "male"},
	// }

	// result := database.DB.Create(&users) // 传入切片的指针

	// fmt.Println(result.Error)        // 返回 error
	// fmt.Println(result.RowsAffected) // 返回插入记录的条数

	// // 创建记录并为指定的字段分配值：
	// user := User{Name: "竹林", Age: 19, Sex: "male"}

	// database.DB.Select("Name", "Age").Create(&user)

	// // 创建记录并忽略要省略的传递字段的值：
	// user := User{Name: "木木", Age: 19, Sex: "female"}

	// database.DB.Omit("Name", "Age").Create(&user)
}
