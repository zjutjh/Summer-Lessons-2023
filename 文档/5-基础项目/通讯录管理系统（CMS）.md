# 项目规范说明
## 风格规范
- 所有的文件/文件夹名采用小驼峰命名（首单词全小写，后面的单词首字母大写）
- 代码采用tab制表符缩进而不是空格（Go 的默认风格）tab默认四空格

## 接口规范
+ 采用 RESTful 风格 API
+ 接受和返回参数采用蛇形小写命名

## 接口响应数据规范
### 响应数据结构
```json
[
    "code":int,
    "msg":string,
    "data":any
]
```
响应数据说明：
- code 为自定义业务状态码
- msg 为详细响应信息
- data 为响应数据

## 项目目录结构
```go
CMS
├── app 
│   ├── controllers // 控制器，主要包括了处理函数
│   ├── midwares // 中间件
│   ├── models // 模型
│   ├── services // 业务逻辑，与数据库进行交互
│   └── utils // 一些封装好的工具
├── config // 配置相关（数据库、路由、viper）
│   ├── config
│   ├── database
│   └── router
```
+ Controller 用于接收和响应
	+ 在 Controller 层定义的是请求和响应，是外部请求的入口和响应的出口，在这里，不应该出现业务逻辑的代码
+ Service 用于处理业务逻辑，会调用 DAO 层的 API
	+ Service 层是业务逻辑层，在该层进行复杂的业务逻辑处理，对在多个Dao层查到的数据进行组装、处理，然后将结果返回给 Controller
+ DAO 是对应一张或多张表的，是持久层 API，是针对单张或多张表的增删改查
	+ DAO 层全称数据访问层，用于数据库的增删改查，它表达的是对 SQL 语句的封装，是组成 Service 层的一部分
+ 作用
	+ 解耦、提高代码整洁性，提高了代码的可维护性

# 通讯录管理系统
## 基本要求：
能实现简单的用户登录注册，对联系人进行添加、删除、修改、查询操作

## 开始
```sh
$ go mod init CMS                    // 初始化项目
$ go get -u github.com/gin-gonic/gin // 安装 gin
$ go get -u gorm.io/gorm             // 安装 gorm
$ go get -u gorm.io/driver/mysql     // 安装 MySQL 驱动
$ go get -u github.com/spf13/viper   // 安装 viper
```

## 模型定义
+ 定义一个 User （用户）类，包括 ID、用户名、性别、手机号（作为账号）、专业、密码等
```go
type User struct {
    ID       uint   `json:"id"`
    Username string `json:"username"`
    Sex      string `json:"sex"`
    PhoneNum string `json:"phone_num"`
    Major    string `json:"major"`
    Password string `json:"-"`
}
```
+ 定义一个 Contact（联系人）类，包括 ID、对应用户 ID、学号、姓名、性别、手机号、专业、是否加入黑名单等
```go
type Contact struct {
    ID        uint   `json:"id"`
    OwnerID   uint   `json:"owner_id"`
    StudentID string `json:"student_id"`
    Name      string `json:"name"`
    Sex       string `json:"sex"`
    PhoneNum  string `json:"phone_num"`
    Major     string `json:"major"`
    Blacklist bool   `json:"blacklist"`
}
```

## 接口文档
+ 1.Web 端
	+ [通讯录管理系统](https://apifox.com/apidoc/shared-ee5ecd53-4ee6-4e24-8ab7-be0c398aa9ed)
	+ 本地测试需要 Chrome 或 Edge 安装插件
+ 2.apifox 导入
	+ CMS/doc/通讯录.openapi.json