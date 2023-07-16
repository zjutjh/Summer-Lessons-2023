package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type UserInfo struct {
	Name string `form:"name_form" uri:"name_uri" json:"name_form" binding:"required"` // binding:"required" tag 表示该属性值不能为空
	Age  int    `form:"age_form" uri:"age_uri" json:"age_form"`
	Sex  string `form:"sex_form" uri:"sex" json:"sex_form"`
}

func main() {
	r := gin.Default()

	// 对应反射 tag form
	// 绑定 Query 示例
	// 1. /query?name_form=ximo&age_form=3&sex_form=male 正常响应
	// 2. /query?name_form=ximo&age_form=3&sex_form= 或 /query?name_form=ximo&age_form=3 正常响应
	// 3. /query?age_form=3&sex_form=male 或 /query?name_form=&age_form=3&sex_form=male 返回 error，binding:"required" tag 表示 Name 属性值不能为空
	r.GET("/query", func(c *gin.Context) {
		var user UserInfo

		// 根据请求的 Content-type 自动识别请求数据类型并利用反射机制自动提取请求中的参数到结构体中
		// 会返回一个 error 参数
		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, user)
	})

	// 对应反射 tag uri
	// 绑定 Param 的示例 /param/ximo/3/male
	r.POST("/param/:name_uri/:age_uri/:sex_uri", func(c *gin.Context) {
		var user UserInfo
		c.ShouldBindWith(&user, binding.JSON)
		err := c.ShouldBindUri(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, user)
	})

	// 对应反射 tag json
	// 绑定 JSON 的示例
	// 1. {"user_json": "ximo", "age_json": 3, "sex_json": male} 正常响应
	// 2. {"user_json": "ximo", "age_json": 3} 或 {"user_json": "ximo", "age_json": 3, "sex_json":""} 正常响应
	// 3. {"age_json": 3, "sex_json": male} 或 {"user_json": "", "age_json": 3, "sex_json": male}返回 error，binding:"required" tag 表示 Name 属性值不能为空
	r.POST("/json", func(c *gin.Context) {
		var user UserInfo

		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, user)
	})

	// 对应反射 tag form（和 Query 相同）
	// 绑定 form 表单示例
	// 1. name_form=ximo&age_form=3&sex_form=male 正常响应
	// 2. name_form=ximo&age_form=3&sex_form= 或 name_form=ximo&age_form=3 正常响应
	// 3. age_form=3&sex_form=male 或 name_form=&age_form=3&sex_form=male 返回 error，binding:"required" tag 表示 Name 属性值不能为空
	r.POST("/post-form", func(c *gin.Context) {
		var user UserInfo

		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, user)
	})

	r.Run()
}
