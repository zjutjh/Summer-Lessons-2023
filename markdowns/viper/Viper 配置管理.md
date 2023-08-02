# Viper

Viper是适用于Go应用程序的完整配置解决方案。它被设计用于在应用程序中工作，并且可以处理**所有类型的配置需求和格式**。
## 安装
```sh
$ go get github.com/spf13/viper
```
## 什么是Viper
Viper是一个功能强大的配置管理库，特别适用于Go应用程序。它可以帮助Go开发者轻松加载、解析、更新和保存配置信息，从而使应用程序配置管理更加方便和灵活。

以下是Go Viper库的一些主要功能和用途：
+ 加载配置文件：Viper可以从不同的配置文件格式（如JSON、YAML、TOML、HCL等）中读取配置信息，并将其映射到Go数据结构中。
+ 默认值：可以设置默认值，以确保配置文件或其他来源中没有指定的配置项具有合理的默认值。
+ 监控配置变化：Viper还支持监控配置文件的更改，并在配置文件发生变化时重新加载配置，从而允许应用程序在运行时动态更新配置。
+ 支持多种数据类型：Viper可以处理各种基本数据类型（如字符串、整数、浮点数等）以及复杂数据类型（如数组、切片、结构体等）的配置项。
+ 环境变量：Viper允许使用环境变量覆盖配置文件中的值，这对于在不同环境中使用不同配置非常有用。
+ 命令行标志：Viper支持从命令行参数中读取配置值，这使得在启动应用程序时通过命令行传递配置选项变得非常简便。

通过使用Viper，开发者可以在Go应用程序中更轻松地管理和维护配置信息，从而实现更高度的灵活性和可配置性，特别是在不同环境（如开发、测试、生产）中使用不同的配置。这对于构建现代化的、可伸缩的应用程序是非常有益的。

Viper会按照下面的优先级：
+ 显示调用Set设置值
+ 命令行参数（flag）
+ 环境变量
+ 配置文件
+ key/value存储
+ 默认值
  
重要： 目前Viper配置的键（Key）是大小写不敏感的。目前正在讨论是否将这一选项设为可选。
# Viper 实操
## 建立默认值
一个好的配置系统将支持默认值。

Examples:
```go
viper.SetDefault("ContentDir", "content")

viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})
```

如果在配置文件中没有设置ContentDir这个参数的值，那么默认使用"content"作为内容目录。

如果在配置文件中没有设置Taxonomies这个参数的值，那么默认使用一个包含键值对的映射。其中，键是字符串类型，可以是"tag"或"category"，而值是对应的目录名，可以是"tags"或"categories"。

通过设置这些默认值，代码确保在读取配置文件时，如果某些参数没有在配置文件中明确指定，那么就会使用这里设置的默认值。这样可以确保代码在处理缺失配置的情况下能够正常运行，而不会因为找不到配置参数而出现错误。
## 读取配置文件

Viper需要最少知道在哪里查找配置文件的配置。Viper支持`JSON`、`TOML`、`YAML`、`HCL`、`envfile`和`Java properties`格式的配置文件。Viper可以搜索多个路径，但目前单个Viper实例只支持单个配置文件。Viper不默认任何配置搜索路径，将默认决策留给应用程序。

下面是一个如何使用Viper搜索和读取配置文件的示例。不需要任何特定的路径，但是至少应该提供一个配置文件预期出现的路径。

```go
viper.SetConfigFile("./config.yaml") // 指定配置文件路径
viper.SetConfigName("config") // 配置文件名称(无扩展名)
viper.SetConfigType("yaml") // 如果配置文件的名称中没有扩展名，则需要配置此项
viper.AddConfigPath("/etc/appname/")   // 查找配置文件所在的路径
viper.AddConfigPath("$HOME/.appname")  // 多次调用以添加多个搜索路径
viper.AddConfigPath(".")               // 还可以在工作目录中查找配置
err := viper.ReadInConfig() // 查找并读取配置文件
if err != nil { // 处理读取配置文件的错误
    panic(fmt.Errorf("Fatal error config file: %s \n", err))
}
```

在加载配置文件出错时，你可以像下面这样处理找不到配置文件的特定情况：

```go
if err := viper.ReadInConfig(); err != nil {
    if _, ok := err.(viper.ConfigFileNotFoundError); ok {
        // 配置文件未找到错误；如果需要可以忽略
    } else {
        // 配置文件被找到，但产生了另外的错误
    }
}
// 配置文件找到并成功解析
```

注意:自1.6起,你也可以有不带扩展名的文件，并以编程方式指定其格式。对于位于用户$HOME目录中的配置文件没有任何扩展名，如.zshrc。
## 写入配置文件

从配置文件中读取配置文件是有用的，但是有时你想要存储在运行时所做的所有修改。为此，可以使用下面一组命令，每个命令都有自己的用途:

+ WriteConfig - 将当前的viper配置写入预定义的路径并覆盖（如果存在的话）。如果没有预定义的路径，则报错。
+ SafeWriteConfig - 将当前的viper配置写入预定义的路径。如果没有预定义的路径，则报错。如果存在，将不会覆盖当前的配置文件。
+ WriteConfigAs - 将当前的viper配置写入给定的文件路径。将覆盖给定的文件(如果它存在的话)。
+ SafeWriteConfigAs - 将当前的viper配置写入给定的文件路径。不会覆盖给定的文件(如果它存在的话)。

根据经验，标记为safe的所有方法都不会覆盖任何文件，而是直接创建（如果不存在），而默认行为是创建或截断。

Examples:
```go
viper.WriteConfig() // 将当前配置写入“viper.AddConfigPath()”和“viper.SetConfigName”设置的预定义路径
viper.SafeWriteConfig()
viper.WriteConfigAs("/path/to/my/.config")
viper.SafeWriteConfigAs("/path/to/my/.config") // 因为该配置文件写入过，所以会报错
viper.SafeWriteConfigAs("/path/to/my/.other_config")
```

## 监控并重新读取配置文件

Viper支持在运行时实时读取配置文件的功能。

需要重新启动服务器以使配置生效的日子已经一去不复返了，viper驱动的应用程序可以在运行时读取配置文件的更新，而不会错过任何消息。

只需告诉viper实例watchConfig。可选地，你可以为Viper提供一个回调函数，以便在每次发生更改时运行。

确保在调用WatchConfig()之前添加了所有的配置路径。

```go
viper.WatchConfig()
viper.OnConfigChange(func(e fsnotify.Event) {
  // 配置文件发生变更之后会调用的回调函数
    fmt.Println("Config file changed:", e.Name)
})
```

对于这段代码中的回调函数，参数 e 是 fsnotify.Event 类型的变量。fsnotify.Event 是 Viper 库中用于表示文件系统事件的结构体类型。当配置文件被修改时，Viper 会创建一个 fsnotify.Event 对象，并将其作为参数传递给回调函数。

fsnotify.Event 结构体中包含了与文件系统事件相关的信息，包括文件名称、事件类型、事件时间等。在这段代码中，我们可以通过 e.Name 来访问文件名，获取发生变化的配置文件的名称。
## 读取键

viper 提供了多种形式的读取方法。在刚刚的例子中，我们看到了GetString方法的用法。

但viper最基础的读取键的方式是Get方法，其返回一个interface{}的值，使用有所不便。GetType系列方法可以返回指定类型的值。其中，Type 可以为Bool/Float64/Int/String/Time/Duration/IntSlice/StringSlice。但是请注意，如果指定的键不存在或类型不正确，GetType方法返回对应类型的零值。

如果要判断某个键是否存在，使用`IsSet`方法。另外，`GetStringMap`和`GetStringMapString`直接以 map 返回某个键下面所有的键值对，前者返回`map[string]interface{}`，后者返回`map[string]string`。`AllSettings`以`map[string]interface{}`返回所有设置。
## 设置键

覆盖设置：这些可能来自命令行标志，也可能来自你自己的应用程序逻辑。

```go
viper.Set("Verbose", false)
```

注册和使用别名：别名允许多个键引用单个值
```go
viper.RegisterAlias("loud", "Verbose")  // 注册别名（此处loud和Verbose建立了别名）

viper.Set("verbose", true) // 结果与下一行相同
viper.Set("loud", true)   // 结果与前一行相同

viper.GetBool("loud") // true
viper.GetBool("verbose") // true
```

## 思考问题

当你使用如下方式读取配置时，viper会从`./conf`目录下查找任何以`config`为文件名的配置文件，如果同时存在`./conf/config.json`和`./conf/config.yaml`两个配置文件的话，viper会从哪个配置文件加载配置呢？

```go
viper.SetConfigName("config")
viper.AddConfigPath("./conf")
```

在上面两个语句下搭配使用`viper.SetConfigType("yaml")`指定配置文件类型可以实现预期的效果吗？