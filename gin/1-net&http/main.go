package main

import (
	"fmt"
	"net/http"
)

// SayHello 函数是需要我们自己实现的一个函数，它有两个规定的参数
// 第一个参数用于给前端的响应（Response）写入数据，响应想要返回什么，就往这个参数里面写什么
// 第二个参数用于获取前端发送的请求（Request）
// Web 开发的本质就是一个请求对应一个响应的过程
func SayHello(w http.ResponseWriter, r *http.Request) {
	// fmt 包中的 Fprintln 函数是一个简单的可以往 w 里写东西的函数
	// 使用 Fprintln 函数将 "Hello 精弘!" 这句话以纯文本的形式写进 w 后返回给前端
	fmt.Fprintln(w, "Hello 精弘!")
}

func main() {
	// HandleFunc 函数接受两个参数，第一个是路径，即前端请求的 URL，另一个是回调函数，用于处理前端发送的请求
	// HandleFunc 函数是一个设置路由的函数，它的作用是将前端对 /hello 路径的请求映射到 SayHello 函数
	// 当前端访问 /hello 路径的时候，就去执行 SayHello 的函数，往响应写入 "Hello 精弘!" 后返回给前端
	http.HandleFunc("/hello", SayHello)

	// ListenAndServe 函数用于启动服务（Serve）并监听（Listen），接收两个参数
	// 第一个参数是 ip:port 格式的 string 参数，给 /hello 路径确定访问它的 ip 地址和端口号
	// 第二个参数指的是处理 HTTP 请求的处理器，填入nil表示使用默认的处理器
	// ":8080" 是简写，省略了 ip，默认为本机的所有 ip 如 127.0.0.1 就是其中一个，端口号指定为 8080
	// 在浏览器（前端）访问 127.0.0.1:8080/hello 就可以看到 SayHello 函数做出的响应了
	// 另外有一个 err 参数，如果端口被占用或者启动失败会返回错误，正常启动则返回 nil，即空（没有错误）
	err := http.ListenAndServe(":8080", nil)

	// 错误处理，如果错误不为 nil 则在终端打印错误
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}
