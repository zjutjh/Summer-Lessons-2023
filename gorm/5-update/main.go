package main

import (
	"gorm/database"
	"time"
)

type User struct {
	ID        uint
	Name      string
	Age       int
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	// 数据库初始化连接
	database.Init()
	// 自动迁移建表
	database.DB.AutoMigrate(&User{})

	// 造数据
	users := []User{
		{Name: "XiMo", Age: 3, Active: true},
		{Name: "XiMo", Age: 0, Active: true},
		{Name: "惜寞", Age: 3, Active: false},
		{Name: "青鸟", Age: 8, Active: true},
		{Name: "竹林", Age: 5, Active: true},
		{Name: "木木", Age: 9, Active: false},
		{Name: "Sora", Age: 7, Active: true},
		{Name: "桑葚", Age: 114514, Active: true},
		{Name: "知更鸟", Age: 6, Active: false},
		{Name: "花仙", Age: 4, Active: true},
		{Name: "一一", Age: 10, Active: true},
	}

	database.DB.Create(&users)

	// var user User
	// database.DB.First(&user)

	// Save
	// Save 会保存所有的字段，即使字段是零值
	// user.Name = "XxxMoo"
	// user.Age = 0
	// database.DB.Save(&user)

	// 如果保存值不包含主键（或者主键值在数据库中不存在），将执行 Create 创建
	// database.DB.Save(&User{Name: "派大星", Age: 20, Active: true})

	// 主键存在则会执行更新保存所有字段
	// database.DB.Save(&User{ID: 1, Name: "XiMo", Age: 3, Active: false})

	// 更新单个列
	// 根据条件更新
	// database.DB.Model(&User{}).Where("active = ?", true).Update("name", "精弘")

	// 根据主键更新
	// database.DB.Model(&user).Update("name", "网络")

	// 更新多列
	// 用结构体更新，不会更新零值，要更新零值可以选择 map 或 选定字段
	// database.DB.Debug().Model(&user).Updates(User{Name: "XiMo", Age: 0, Active: false})

	// 用 map 更新
	// database.DB.Debug().Model(&user).Updates(map[string]interface{}{"name": "惜寞", "age": 0, "active": false})

	// 更新选定字段
	// database.DB.Debug().Model(&user).Select("Name", "Age").Updates(User{Name: "XxxMoo", Age: 0, Active: false})
}
