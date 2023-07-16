package main

import (
	"github.com/gin-gonic/gin"
)

func PostForm(c *gin.Context) {
	// PostForm 获取key对应的值，不存在返回空字符串
	// name=ximo&age=3&sex=
	namePostForm := c.PostForm("name")
	agePostForm := c.PostForm("age")
	sexPostForm := c.PostForm("sex")
	organizationPostForm := c.PostForm("organization")
	c.JSON(200, gin.H{
		"namePostForm":         namePostForm,
		"agePostForm":          agePostForm,
		"sexPostForm":          sexPostForm,
		"organizationPostForm": organizationPostForm,
	})

	// // DefaultPostForm key不存在时返回一个默认值
	// // name=ximo&age=3&sex=
	// organizationDefaultPostForm := c.DefaultPostForm("organization", "精弘网络")
	// sexDefaultPostForm := c.DefaultPostForm("sex", "male")
	// nameDefaultPostForm := c.DefaultPostForm("name", "惜寞")
	// c.JSON(200, gin.H{
	// 	"organizationPostForm":        organizationPostForm,
	// 	"organizationDefaultPostForm": organizationDefaultPostForm,
	// 	"sexPostForm":                 sexPostForm,
	// 	"sexDefaultPostForm":          sexDefaultPostForm,
	// 	"namePostForm":                namePostForm,
	// 	"nameDefaultPostForm":         nameDefaultPostForm,
	// })

	// // GetPostForm 获取key对应的值，并且返回bool标识，标识成功或者失败
	// // name=ximo&age=3&sex=
	// nameGetPostForm, nameExist := c.GetPostForm("name")
	// sexGetPostForm, sexExist := c.GetPostForm("sex")
	// organizationGetPostForm, orgaorganizationExist := c.GetPostForm("organization")
	// c.JSON(200, gin.H{
	// 	"nameGetPostForm":         nameGetPostForm,
	// 	"nameExist":               nameExist,
	// 	"sexGetPostForm":          sexGetPostForm,
	// 	"sexExist":                sexExist,
	// 	"organizationGetPostForm": organizationGetPostForm,
	// 	"orgaorganizationExist":   orgaorganizationExist,
	// })

	// // PostFormArray 获取key对应的值，值是一个字符串数组，不存在返回空字符串数组
	// // name=ximo&age=3&sex=&hobby=code&hobby=sleep&hobby=
	// hobbyPostForm := c.PostForm("hobby")
	// hobbyPostFormArray := c.PostFormArray("hobby")
	// namePostFormArray := c.PostFormArray("name")
	// sexPostFormArray := c.PostFormArray("sex")
	// organizationPostFormArray := c.PostFormArray("organization")
	// c.JSON(200, gin.H{
	// 	"hobbyPostForm":             hobbyPostForm,
	// 	"hobbyPostFormArray":        hobbyPostFormArray,
	// 	"namePostFormArray":         namePostFormArray,
	// 	"sexPostFormArray":          sexPostFormArray,
	// 	"organizationPostFormArray": organizationPostFormArray,
	// })

	// // GetPostFormArray 获取key对应的值，并且返回bool标识，标识成功或者失败
	// // name=ximo&age=3&sex=&hobby=code&hobby=sleep&hobby=
	// hobbyGetPostFormArray, hobbyExist := c.GetPostFormArray("hobby")
	// nameGetPostFormArray, nameExist := c.GetPostFormArray("name")
	// sexGetPostFormArray, sexExist := c.GetPostFormArray("sex")
	// organizationGetPostFormArray, orgaorganizationExist := c.GetPostFormArray("organization")
	// c.JSON(200, gin.H{
	// 	"hobbyGetPostFormArray":        hobbyGetPostFormArray,
	// 	"hobbyExist":                   hobbyExist,
	// 	"nameGetPostFormArray: ":       nameGetPostFormArray,
	// 	"nameExist":                    nameExist,
	// 	"sexGetPostFormArray":          sexGetPostFormArray,
	// 	"sexExist":                     sexExist,
	// 	"organizationGetPostFormArray": organizationGetPostFormArray,
	// 	"orgaorganizationExist":        orgaorganizationExist,
	// })

	// // PostFormMap 获取key对应的值，值是一个字符串map[string]string，不存在返回空
	// // user[name]=ximo&user[age]=3&user[sex]=&user[hobby]=code&user[hobby]=sleep
	// userPostFormMap := c.PostFormMap("user")
	// adminPostFormMap := c.PostFormMap("admin")
	// c.JSON(200, gin.H{
	// 	"userPostFormMap":  userPostFormMap,
	// 	"adminPostFormMap": adminPostFormMap,
	// })

	// // GetPostFormMap 获取key对应的值，值是一个字符串map[string]string，并且返回bool标识，标识成功或者失败
	// // user[name]=ximo&user[age]=3&user[sex]=&user[hobby]=code&user[hobby]=sleep
	// userGetPostFormMap, userExist := c.GetPostFormMap("user")
	// adminGetPostFormMap, adminExist := c.GetPostFormMap("admin")
	// c.JSON(200, gin.H{
	// 	"userGetPostFormMap":  userGetPostFormMap,
	// 	"userExist":           userExist,
	// 	"adminGetPostFormMap": adminGetPostFormMap,
	// 	"adminExist":          adminExist,
	// })
}

func main() {
	r := gin.Default()

	r.POST("/post-form", PostForm)

	r.Run(":8080")
}
