# VSCode 及 Golang 插件的安装
![Untitled](https://bu.dusays.com/2023/07/30/64c6703ef0233.png)

在搜索引擎中搜索 vscode，进入官方下载网站 [https://code.visualstudio.com/](https://code.visualstudio.com/)

![Untitled](https://bu.dusays.com/2023/07/30/64c67041c793c.png)

直接单击最大的下载按钮即可进行下载。

如果下载速度太慢，可以尝试将下载链接中的域名( [az764295.vo.msecnd.net](https://az764295.vo.msecnd.net/stable/2ccd690cbff1569e4a83d7c43d45101f817401dc/VSCodeUserSetup-x64-1.80.2.exe) )替换为 [vscode.cdn.azure.cn](http://vscode.cdn.azure.cn)

双击下载的安装包，按照常规的安装流程进行安装
  
![Untitled](https://bu.dusays.com/2023/07/30/64c670442da57.png)


![Untitled](https://bu.dusays.com/2023/07/30/64c67045e0c8e.png)


![Untitled](https://bu.dusays.com/2023/07/30/64c6705048825.png)


在安装成功后弹出的主界面上，单击左侧栏第五个按钮打开插件管理器

![Untitled](https://bu.dusays.com/2023/07/30/64c67058948c8.png)

在搜索栏中搜索 go，安装第一个插件

![Untitled](https://bu.dusays.com/2023/07/30/64c6705d3cc14.png)

# Golang 的安装与配置

下载链接：[https://golang.google.cn](https://golang.google.cn/)

正常下载安装

![Untitled](https://bu.dusays.com/2023/07/30/64c67062c4bf1.png)

![Untitled](https://bu.dusays.com/2023/07/30/64c670659c2a3.png)

![Untitled](https://bu.dusays.com/2023/07/30/64c6706a4b5fc.png)

![Untitled](https://bu.dusays.com/2023/07/30/64c67071b0030.png)

  
安装成功后，可以通过 `Win+R` 快捷键输入 cmd 打开终端，输入 go 并敲下回车，查看 Path 环境变量是否有被自动配置好


![Untitled](https://bu.dusays.com/2023/07/30/64c6707818591.png)

![Untitled](https://bu.dusays.com/2023/07/30/64c6707e2fa07.png)

并顺势复制以下命令配置境内的镜像源

```bash
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

配置完后重新打开 VSCode 按住 `Ctrl+Shift+P` 输入 `Go:Install/Update Tools`

![](https://image.bluebird.icu/img/b6891e50-b040-4c2f-8856-03c8cc8c216c.webp)

点击后全选

![](https://image.bluebird.icu/img/73b025ef-5557-43ae-962a-0efb59b9ff08.webp)

点击确定后等待安装完毕即可

![](https://image.bluebird.icu/img/1be00716-23a1-431f-a948-1987049a040d.webp)

# 编写第一个 Go 程序
+ 打开文件夹
+ 新建文件
+ 将文件重命名为 main.go

![](https://image.bluebird.icu/img/5972d392-1ceb-4e39-a58c-a0560ecb4f78.webp)

+ 打开终端（终端->新建终端 或者 输入 Ctrl+Shift+\` 快捷键）
+ 输入 go mod init [name] 来初始化 Go 应用，name 名字可以自己取

![](https://image.bluebird.icu/img/f1dc474f-a249-44e0-9b65-b871218c0f6f.webp)

然后在 main 文件中粘贴下面代码

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}
```

在终端输入 `go run main.go` 即可运行显示结果

![](https://image.bluebird.icu/img/585b8c0a-0495-4566-aa9b-ca6ad2722a25.webp)
# 安装 MySQL

推荐直接从下面两个镜像站下载 MySQL 的离线安装包

[https://mirrors.huaweicloud.com/mysql/Downloads/MySQLInstaller/](https://mirrors.huaweicloud.com/mysql/Downloads/MySQLInstaller/)

[https://mirrors.aliyun.com/mysql/MySQLInstaller/](https://mirrors.aliyun.com/mysql/MySQLInstaller/)
  

![Untitled](https://bu.dusays.com/2023/07/30/64c6708259c55.png)

请务必确保自己下载的是 `mysql-installer-community` ，不要下载 web 在线安装版本。

![Untitled](https://bu.dusays.com/2023/07/30/64c67086504f1.png)

刚点开安装包时可能会在这个界面卡很久，耐心等待即可

![Untitled](https://bu.dusays.com/2023/07/30/64c6708a6481d.png)

随后可能会弹出更新提示，直接选择 No 忽略即可

随后在安装程序的主页面，选择 Server only，剩下的安装步骤中除了 Root Password 以外全部下一步即可。

![Untitled](https://bu.dusays.com/2023/07/30/64c6708ecd42f.png)

随后我们需要将 `C:\Program Files\MySQL\MySQL Server 8.0\bin` 加入系统变量中

![Untitled](https://bu.dusays.com/2023/07/30/64c67095165e7.png)

右键我的电脑，选择属性

![Untitled](https://bu.dusays.com/2023/07/30/64c6709b2389d.png)

单击高级系统设置

![Untitled](https://bu.dusays.com/2023/07/30/64c6709de81cb.png)

选择环境变量

![Untitled](https://bu.dusays.com/2023/07/30/64c670a50890c.png)

双击当前用户的 Path 变量

![Untitled](https://bu.dusays.com/2023/07/30/64c670aaacc47.png)

单击新建，输入 `C:\Program Files\MySQL\MySQL Server 8.0\bin`

确定后即完成了 mysql 的 Path 变量配置