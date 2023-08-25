# String
```go
r.GET("/hello", func(c *gin.Context) {
        c.String(200, "Hello 精弘!")
    })
```
上节课我们讲的是 c.String 给前端返回一个字符串，但在实际的前后端分离的开发过程中，直接使用 string 进行前后端数据传输并不便利，所以我们往往会选择使用 JSON 的数据格式进行前后端的数据传输

# JSON
+ JSON 数据类型
	+ JSON(JavaScript Object Notation, JS 对象简谱) 是一种轻量级的数据交换格式，用它可以来表示各种各样复杂的数据，如对象，数组，集合，以及集合的集合等数据
	+ JSON 实际上就是一串字符串，只不过元素会使用特定的符号标注。 {} 双括号表示对象，[] 中括号表示数组，”” 双引号内是属性或值，: 冒号表示后者是前者的值(这个值可以是字符串、数字、也可以是另一个数组或对象)。
+ 一些常见的 JSON 格式
	+ 一个JSON对象——JSONObject
		+ `{"name":"XiMo", "age":3}`
		+ `{"name":"XiMo", "age":3，"address":{"city":HangZhou", "country":"China"}}`
	+ 一个JSON数组——JSONArray
		+ `["XiMo", "惜寞"]`
		+ `[{"name":"XiMo", "age":3}, {"name":"惜寞", "age":4}]`
		+ `[{"name":"XiMo", "age":3, "address":{"city":"HangZhou", "country":"China"}}, {"name":"惜寞", "age": 4, "address":{"city":"JiaXing", "country":"China"}}]`
	+ 可以通过[可视化](http://www.esjson.com/jsonviewer.html)将 JSON 数据类型格式化来查看，结构清晰，并且内容相同，字符串形式只是将空格回车给去掉了而已
	+ 当然，数组可以包含对象，在对象中也可以包含数组
+ 为什么普遍选择 JSON 用于前后端数据的传输
	+ 采用完全独立于任何程序语言的文本格式，使JSON成为理想的数据交换语言
	+ 易于人阅读和编写，键值对类型的数据结构具有良好的可读性
	+ 数据格式比较简单, 格式都是压缩的，占用带宽小，能有效地提升网络传输效率
	+ 易于解析，前端可以很方便的进行 JSON 数据的读取
	+ JSON 格式能够直接为后端代码使用，大大简化了前后端的代码开发量，但是完成的任务不变，且易于维护

和返回 string 一样，Gin 在 Context 里面也给我们封装了返回 JSON 的方法，下面是一个简单的 Gin 返回 JSON 的示例代码：
```go
package main

import (
    "github.com/gin-gonic/gin"
)

func Json(c *gin.Context) {
    // 1.使用结构体，可以灵活利用 tag 对字段进行"换名"
    type UserInfo struct {
        Name     string `json:"username"`
        Age      int    `json:"age"`
        Password string `json:"-"` // 忽略该字段
    }

    user := UserInfo{
        Name:     "XiMo",
        Age:      3,
        Password: "123456",
    }

    // c.JSON 实际上是将结构体和 map 类型的变量进行序列化（将对象转换为JSON格式的字符串的过程）
    // 注意 UserInfo.Name 在序列化中变成了 "username"，UserInfo.Age 变成了 "age"
    // "-" 表示忽略该字段，所以 UserInfo.Password 在序列化的时候会被忽略
    // 响应将返回：{"username": "XiMo", "age": 3}
    c.JSON(200, user)

    // // 2.使用 map
    // userMap := map[string]any{
    //     "name": "XiMo",
    //     "age":  3,
    // }

    // c.JSON(200, userMap)

    // // 3.使用 gin.H
    // // gin.H 实际上是 map[string]any 的一种快捷方式
    // c.JSON(200, gin.H{
    //     "name": "XiMo",
    //     "age":  3,
    // })
}

func main() {
    r := gin.Default()

    r.GET("/json", Json)

    r.Run()
}
```
其实除了 JSON 和 string 类型的返回，Gin 还给我们提供了响应 XML、YAML、HTML 等的方式，但由于使用较少，所以这里不多涉及，感兴趣的可以自行查阅