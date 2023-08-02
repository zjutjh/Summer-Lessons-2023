# ORM
## 什么是 ORM
+ ORM 是 Object Relational Mapping 的缩写，译为 ”对象关系映射“，他解决了对象和关系型数据库之间的数据交互问题
+ 将*程序中的对象/实例*与*关系型数据库*映射起来
+ 三个映射关系
	+ 数据表对应结构体
	+ 数据行对应结构体实例
	+ 字段对应结构体字段
+ 以 Go 为例：
+ ![](https://image.bluebird.icu/img/84878fff-e7ef-4ab6-b6ae-3b9a52a5ff3a.webp)
## 为什么使用 ORM
+ 提高开发效率
```go
type UserInfo struct {
    ID   int
    Name string
    Age  int
    Sex  string
}

func main() {
    user := UserInfo{1, "XiMo", 3, "male"}
    // 将 user 存入数据库
    // SQL 语句
    // insert into userinfo values(1,"XiMo",3,"male");

	// orm 语句
	// orm.Create(&user)
}
```

+ 缺点
	+ 自动生成 SQL 语句，会牺牲一定的性能
	+ 对于复杂的数据库操作，ORM 通常难以处理，即使能够处理，也不如直接手写原生 SQL 语句灵活
	+ 弱化 SQL 能力

# Gorm
+ [GORM](https://gorm.io/zh_CN/) 是 Go 语言目前比较热门的数据库 ORM 操作库，对开发者也比较友好，使用非常简单

## 连接
### 安装
在使用 gorm 之前我们需要下载 gorm 以及对应数据库 mysql 的驱动
```go
$ go get -u gorm.io/gorm  
$ go get -u gorm.io/driver/mysql
```

### 简单连接
```go
// 定义一个全局变量
var DB *gorm.DB
  
func Init() {
    user := "root"
    pass := "123456"
    host := "127.0.0.1"
    port := "3306"
    DBname := "gorm_learn"
  
    // dsn data-source-name 告知数据库所在的位置以及数据库相关的属性
    // user:pass@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        user, pass, host, port, DBname)

    // 连接数据库，获得 DB 类型实例，用于后面对数据库进行的操作
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database, error = " + err.Error())
    }
  
    // 赋值给全局变量
    DB = db
}

func main() {
	// 做一层简单的封装，后续操作可以直接利用全局变量 DB 进行
    Init()
}
```