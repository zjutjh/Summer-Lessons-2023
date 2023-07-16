package main

import (
	"github.com/gin-gonic/gin"
)

func Query(c *gin.Context) {
	// 访问 /query?name=ximo&age=3&sex=
	// Query 获取key对应的值，不存在返回空字符串
	nameQuery := c.Query("name")
	ageQuery := c.Query("age")
	sexQuery := c.Query("sex")
	organizationQuery := c.Query("organization")
	c.JSON(200, gin.H{
		"nameQuery":         nameQuery,
		"ageQuery":          ageQuery,
		"sexQuery":          sexQuery,
		"organizationQuery": organizationQuery,
	})

	// // 访问 /query?name=ximo&age=3&sex=
	// // DefaultQuery key不存在时返回一个默认值
	// organizationDefaultQuery := c.DefaultQuery("organization", "精弘网络")
	// sexDefaultQuery := c.DefaultQuery("sex", "male")
	// nameDefaultQuery := c.DefaultQuery("name", "惜寞")
	// c.JSON(200, gin.H{
	// 	"organizationDefaultQuery": organizationDefaultQuery,
	// 	"sexDefaultQuery":          sexDefaultQuery,
	// 	"nameDefaultQuery":         nameDefaultQuery,
	// })

	// // 访问 /query?name=ximo&age=3&sex=
	// // GetQuery 获取key对应的值，并且返回bool标识，标识成功或者失败
	// nameGetQuery, nameExist := c.GetQuery("name")
	// sexGetQuery, sexExist := c.GetQuery("sex")
	// organizationGetQuery, orgaorganizationExist := c.GetQuery("organization")
	// c.JSON(200, gin.H{
	// 	"nameGetQuery":          nameGetQuery,
	// 	"nameExist":             nameExist,
	// 	"sexGetQuery":           sexGetQuery,
	// 	"sexExist":              sexExist,
	// 	"organizationGetQuery":  organizationGetQuery,
	// 	"orgaorganizationExist": orgaorganizationExist,
	// })

	// // 访问 /query?name=ximo&age=3&sex=&hobby=code&hobby=sleep&hobby=
	// // QueryArray
	// hobbyQuery := c.Query("hobby")
	// hobbyQueryArray := c.QueryArray("hobby")
	// nameQueryArray := c.QueryArray("name")
	// sexQueryArray := c.QueryArray("sex")
	// organizationQueryArray := c.QueryArray("organization")
	// c.JSON(200, gin.H{
	// 	"hobbyQuery":             hobbyQuery,
	// 	"hobbyQueryArray":        hobbyQueryArray,
	// 	"nameQueryArray":         nameQueryArray,
	// 	"sexQueryArray":          sexQueryArray,
	// 	"organizationQueryArray": organizationQueryArray,
	// })

	// // 访问 /query?name=ximo&age=3&sex=&hobby=code&hobby=sleep&hobby=
	// // GetQueryArray
	// hobbyGetQueryArray, hobbyExist := c.GetQueryArray("hobby")
	// nameGetQueryArray, nameExist := c.GetQueryArray("name")
	// sexGetQueryArray, sexExist := c.GetQueryArray("sex")
	// organizationGetQueryArray, orgaorganizationExist := c.GetQueryArray("organization")
	// c.JSON(200, gin.H{
	// 	"hobbyGetQueryArray":        hobbyGetQueryArray,
	// 	"hobbyExist":                hobbyExist,
	// 	"nameGetQueryArray: ":       nameGetQueryArray,
	// 	"nameExist":                 nameExist,
	// 	"sexGetQueryArray":          sexGetQueryArray,
	// 	"sexExist":                  sexExist,
	// 	"organizationGetQueryArray": organizationGetQueryArray,
	// 	"orgaorganizationExist":     orgaorganizationExist,
	// })

	// // 访问 /query?user[name]=ximo&user[age]=3&user[sex]=&user[hobby]=code&user[hobby]=sleep
	// // QueryMap
	// userQueryMap := c.QueryMap("user")
	// adminQueryMap := c.QueryMap("admin")
	// c.JSON(200, gin.H{
	// 	"userQueryMap":  userQueryMap,
	// 	"adminQueryMap": adminQueryMap,
	// })

	// // 访问 /query?user[name]=ximo&user[age]=3&user[sex]=&user[hobby]=code&user[hobby]=sleep
	// // GetQueryMap
	// userGetQueryMap, userExist := c.GetQueryMap("user")
	// adminGetQueryMap, adminExist := c.GetQueryMap("admin")
	// c.JSON(200, gin.H{
	// 	"userGetQueryMap":  userGetQueryMap,
	// 	"userExist":        userExist,
	// 	"adminGetQueryMap": adminGetQueryMap,
	// 	"adminExist":       adminExist,
	// })
}

func main() {
	r := gin.Default()

	r.GET("/query", Query)

	r.Run(":8080")
}
