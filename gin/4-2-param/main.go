package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 注意会匹配 /param/XiMo 但不会匹配 /param/ 或者 /param
	r.GET("/param/:name", func(c *gin.Context) {
		// 访问 /param/XiMo
		name := c.Param("name")

		// 不存在会返回空字符串
		age := c.Param("age")

		c.JSON(200, gin.H{
			"name": name,
			"age":  age,
		})
	})

	// 常见的用法有博客的归档，去查看某年某月写的文章
	r.GET("/blog/:year/:month", func(c *gin.Context) {
		year := c.Param("year")
		month := c.Param("month")

		c.JSON(200, gin.H{
			"year":  year,
			"month": month,
		})
	})

	r.Run(":8080")
}
