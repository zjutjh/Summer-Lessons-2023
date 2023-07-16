package main

import (
	"github.com/gin-gonic/gin"
)

func Json(c *gin.Context) {
	// 1.使用结构体，可以灵活利用 tag 对字段进行"换名"
	type UserInfo struct {
		Name     string `json:"username"`
		Age      int    `json:"age"`
		Password string `json:"-"` // 忽略该字段
	}

	user := UserInfo{
		Name:     "XiMo",
		Age:      3,
		Password: "123456",
	}

	// c.JSON 实际上是将结构体和 map 类型的变量进行序列化（将对象转换为JSON格式的字符串的过程）
	// 注意 UserInfo.Name 在序列化中变成了 "username"，UserInfo.Age 变成了 "age"
	// "-" 表示忽略该字段，所以 UserInfo.Password 在序列化的时候会被忽略
	// 响应将返回：{"username": "XiMo", "age": 3}
	c.JSON(200, user)

	// 2.使用 map
	userMap := map[string]any{
		"name": "XiMo",
		"age":  3,
	}

	c.JSON(200, userMap)

	// 3.使用 gin.H
	// gin.H 实际上是 map[string]any 的一种快捷方式
	c.JSON(200, gin.H{
		"name": "XiMo",
		"age":  3,
	})
}

func main() {
	r := gin.Default()

	r.GET("/json", Json)

	r.Run()
}
