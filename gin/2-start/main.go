package main

import (
	"fmt"

	// 将 gin 引入到代码中
	"github.com/gin-gonic/gin"
)

// SayHello 函数是一个处理 HTTP 请求的回调函数
// 接受一个 gin 封装过的规定参数，即上下文对象 Context，它是所有请求处理器（处理请求的函数或方法）的入口参数
// Context 包含了 Request 和 ResponseWriter 两个参数，用于获取前端请求信息和返回响应
// 本质上是对于 Request 和 Response 的封装，提供了丰富的方法用于获取当前请求的上下文信息以及返回响应
func SayHello(c *gin.Context) {
	// 200 表示 HTTP 响应状态码（<=> http.StatusOK）
	// 使用 Context 的 String 函数将 "Hello 精弘!" 这句话以纯文本（字符串）的形式返回给前端
	// 实际上是对返回响应的封装
	c.String(200, "Hello 精弘!")
}

func main() {
	// gin.Default 函数会生成一个默认的 Engine （路由引擎）对象（集成了 Logger 和 Recovery 两个中间件，中间件后面的课会讲）
	// Engine 是 Gin 框架最重要的数据结构，它是 Gin 框架的入口，本质上是一个 Http Handler
	// 它是一个用于处理 HTTP 请求的对象，维护了一张路由表，将不同的 HTTP 请求路径映射到不同的处理函数上
	r := gin.Default()

	// r.GET 函数接受两个参数，一个是路径，即前端请求的 URL，另一个是回调函数，用于处理前端发送的请求
	// r.GET 函数将 /hello 路径添加到了 r 的路由表中，将对 /hello 路径的 GET 请求映射到 SayHello 函数上
	// 表示前端给后端的 /hello 路由发送一个 HTTP 的 GET 请求时
	// 后端会执行后面的 SayHello 函数，对前端的请求做出一个响应，给前端返回一个 "Hello 精弘!" 的字符串
	r.GET("/hello", SayHello)
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello 精弘!")
	})

	// 启动服务并监听,只接受一个 ip:port 格式的 string 参数，表示服务运行的 ip 地址和端口号
	// ":8080" 是简写，省略了 ip，表示监听本地所有 ip （如 127.0.0.1）的 8080 端口，接收并处理 HTTP 请求
	// 在浏览器（前端）访问 127.0.0.1:8080/hello 就可以看到 SayHello 函数做出的响应
	// 等价于 r.Run()，将 port 端口也省略，默认为 8080 端口
	err := r.Run(":8080")

	// 错误处理，如果错误不为 nil 则在终端打印错误
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}
