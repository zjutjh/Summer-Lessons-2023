package main

import (
	"gorm/database"
)

type User struct {
	ID   uint
	Name string
	Age  int
}

func main() {
	// 数据库初始化连接
	database.Init()
	// 自动迁移建表
	database.DB.AutoMigrate(&User{})

	// 造数据
	users := []User{
		{Name: "XiMo", Age: 3},
		{Name: "XiMo", Age: 0},
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

	// var user User
	// var users []User

	// // 根据主键查询第一条记录
	// database.DB.First(&user)
	// fmt.Println(user)

	// // 获取一条记录，没有指定排序字段
	// database.DB.Take(&user)
	// fmt.Println(user)

	// // 根据主键查询最后一条记录
	// database.DB.Last(&user)
	// fmt.Println(user)

	// // 查询指定的某条记录(仅当主键为整型时可用)
	// database.DB.First(&user, 8)
	// fmt.Println(user)

	// // 查询所有的记录
	// database.DB.Find(&users)
	// fmt.Println(users)

	// // string 条件
	// // 获取第一条匹配的记录
	// database.DB.Where("name = ?", "XiMo").First(&user)
	// fmt.Println(user)

	// // 获取所有匹配的记录
	// database.DB.Where("name <> ?", "XiMo").Find(&users)
	// fmt.Println(users)

	// // 使用 struct 查询时，GORM 将仅使用非零字段进行查询
	// // 这意味着如果你的字段的值为 `0` `''` `false` 或其他零值时，将不会用于构建查询条件
	// database.DB.Debug().Where(&User{Name: "XiMo", Age: 0}).Find(&users)
	// fmt.Println(users)

	// // 要在查询条件中包含零值，可以使用 map，它会将所有键值作为查询条件包含在内
	// database.DB.Debug().Where(map[string]interface{}{"Name": "XiMo", "Age": 0}).Find(&users)
	// fmt.Println(users)

	// // 指定结构体查询字段
	// // 如果指定了字段，即使没有赋值也会用零值作用于查询条件的构建
	// database.DB.Debug().Where(&User{Name: "XiMo"}, "name", "Age").Find(&users)
	// fmt.Println(users)

	// // 如果没有指定，即使赋值了也不会用于查询条件的构建
	// database.DB.Debug().Where(&User{Name: "XiMo"}, "Age").Find(&users)
	// fmt.Println(users)

	// // 选择特定字段
	// // Select 指定要从数据库中检索的字段
	// database.DB.Select("name", "age").Find(&users)
	// fmt.Println(users)

	// // 排序
	// // 按年龄降序，姓名升序
	// database.DB.Order("age desc, name").Find(&users)
	// fmt.Println(users)

	// // Limit & Offset
	// // 指定检索的最大记录数为 3
	// database.DB.Limit(3).Find(&users)
	// fmt.Println(users)

	// // -1 表示不限制最大记录数
	// database.DB.Limit(-1).Find(&users)
	// fmt.Println(users)

	// // 指定检索的最大记录数为 5，并在返回记录前往后偏移五个位置
	// database.DB.Limit(5).Offset(5).Find(&users)

	// // 组合使用
	// database.DB.Where("age < 8").Order("age desc").Limit(3).Offset(2).Find(&users)
	// fmt.Println(users)
}
