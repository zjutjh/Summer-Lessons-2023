package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Raw(c *gin.Context) {
	// GetRawData 实际上是去获取 request.body 中的内容
	// 它返回两个参数，一个是获取到的 []byte 类型的 body 数据，另一个是 error 类型
	// 这里忽略了 error 的处理
	data, _ := c.GetRawData()

	// 打印 data 可以看到传过来的原始数据
	fmt.Println(data)

	// 将 []byte 转成 string 类型可以看它实际传过来的内容
	fmt.Println(string(data))

	// 注：下面是对 JSON 数据类型的处理
	// 我们可以通过 json 包里的 Unmarshal 来对 JSON 数据类型进行反序列化
	// 就是我前几节课所说的用 JSON 中的数据去给结构体和 map 类型变量赋值

	// 1.结构体，tag 在反序列化的时候依旧可用，会根据 tag 将对应的值赋给对应的键
	type UserInfo struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		Sex  string `json:"sex"`
	}

	var userStruct UserInfo

	// JSON 反序列化，会返回一个 error，这里忽略了对其的处理
	_ = json.Unmarshal(data, &userStruct)

	c.JSON(200, userStruct)

	// // 2.map
	// var userMap map[string]interface{}

	// _ = json.Unmarshal(data, &userMap)

	// //获取 JSON 中的 key，注意使用 ["key"] 获取
	// // name := userMap["name"]
	// // age := userMap["age"]
	// // sex := userMap["sex"]

	// c.JSON(200, userMap)
}

func main() {
	r := gin.Default()

	r.POST("/raw", Raw)

	r.Run()
}
