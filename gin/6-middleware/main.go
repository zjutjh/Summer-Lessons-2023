package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func m1(c *gin.Context) {
	fmt.Println("这是一个中间件")

	// 中间件也可以返回响应
	c.JSON(200, gin.H{
		"msg": "中间件",
	})

	// 中断请求，不执行后续内容
	// c.Abort()
}

func main() {
	r := gin.Default()

	// 将 m1 作为一个中间件注册到该路由
	r.GET("/", m1, func(c *gin.Context) {
		fmt.Println("这是处理函数")

		c.JSON(200, gin.H{
			"msg": "处理函数",
		})
	})

	r.Run()
}

// // Next()
// func m1(c *gin.Context) {
//     fmt.Println("m1 in...")

//     // 跳到下一个 handler
//     c.Next()

//     fmt.Println("m1 out...")
// }

// func m2(c *gin.Context) {
//     fmt.Println("m2 in...")

//     // 跳到下一个 handler
//     c.Next()

//     fmt.Println("m2 out...")
// }

// func main() {
//     r := gin.Default()

//     r.GET("/", m1, m2, func(c *gin.Context) {
//         fmt.Println("/ in...")

//         fmt.Println("/ out...")
//     })

//     r.Run()
// }

// // Set() 和 Get()
// type UserInfo struct {
//     Name string
//     Age  int
// }

// func m1(c *gin.Context) {
//     user := UserInfo{
//         Name: "XiMo",
//         Age:  3,
//     }

//     // 将 user 以键值对的形式存储到 Context 中
//     c.Set("user", user)
// }

// func main() {
//     r := gin.Default()

//     r.GET("/", m1, func(c *gin.Context) {
//         // 根据 key 返回 value（any 类型）和一个 bool 值
//         // bool 值用来判断是否有这个 key
//         user, ok := c.Get("user")
//         fmt.Println(gin.H{
//             "userExist": ok,
//         })

//         // 断言，会返回断言类型的值和一个 bool 值
//         // bool 值用来判断是否断言成功
//         // 断言失败，返回的值为断言类型的零值
//         _user, _ok := user.(UserInfo)
//         fmt.Println(gin.H{
//             "type assertion": _ok,
//         })

//         fmt.Println(_user)
//     })

//     r.Run()
// }
