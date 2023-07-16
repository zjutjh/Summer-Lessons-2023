package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Any 可以匹配所有 HTTP 请求
	r.Any("/any", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "Any",
		})
		// switch c.Request.Method {
		// case "GET":
		// 	c.JSON(200, gin.H{
		// 		"msg": "ANY_GET",
		// 	})
		// case "POST":
		// 	c.JSON(200, gin.H{
		// 		"msg": "ANY_POST",
		// 	})
		// case "PUT":
		// 	c.JSON(200, gin.H{
		// 		"msg": "ANY_PUT",
		// 	})
		// case "DELETE":
		// 	c.JSON(200, gin.H{
		// 		"msg": "ANY_DELETE",
		// 	})
		// default:
		// 	// HEAD PATCH...
		// 	c.JSON(200, gin.H{
		// 		"msg": "ANY_OTHER",
		// 	})
		// }
	})

	// NoRoute 用于处理找不到路由的情况，即当没有匹配到任何定义的路由时执行的处理函数
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"msg": "NoRoute",
		})
	})

	// NoMethod 用于处理找不到路由的情况，即当没有匹配到任何定义的路由时执行的处理函数
	r.NoMethod(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"msg": "NoMethod",
		})
	})

	// 参数路由
	// 匹配参数
	r.GET("/path_1/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"id": c.Param("id"),
		})
	})

	// 匹配任意路径，只要是以 /path_2 开头
	r.GET("/path_2/*action", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"action": c.Param("action"),
		})
	})

	// r.POST("/user/register", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"msg": "/user/register",
	// 	})
	// })
	// r.POST("/user/login/email", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"msg": "/user/login/email",
	// 	})
	// })
	// r.POST("/user/login/phone", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"msg": "/user/login/phone",
	// 	})
	// })
	// r.POST("/user/login/password", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"msg": "/user/login/password",
	// 	})
	// })
	// r.POST("/user/exit", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"msg": "/user/exit",
	// 	})
	// })

	// 路由组 将共同的 URL 前缀给抽离出来
	user := r.Group("/user")
	{
		user.POST("/register", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "/user/register",
			})
		})

		// 路由组支持嵌套
		login := user.Group("/login")
		{
			login.POST("/email", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"msg": "/user/login/email",
				})
			})
			login.POST("/phone", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"msg": "/user/login/phone",
				})
			})
			login.POST("/password", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"msg": "/user/login/password",
				})
			})
		}

		user.POST("/exit", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "/user/exit",
			})
		})
	}

	r.Run()
}
