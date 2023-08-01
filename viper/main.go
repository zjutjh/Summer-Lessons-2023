package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

//func main() {
//	// 设置默认值
//	viper.SetDefault("fileDir", "./")
//	// 读取配置文件
//	viper.SetConfigFile("./config.yaml")   // 指定配置文件路径
//	viper.AddConfigPath("/etc/appname1/")  // 查找配置文件所在的路径
//	viper.AddConfigPath("$HOME/.appname2") // 多次调用以添加多个搜索路径
//	viper.AddConfigPath(".")               // 还可以在工作目录中查找配置
//	viper.SetConfigName("config")          // 配置文件名称(无扩展名)
//	viper.SetConfigType("yaml")            // 如果配置文件的名称中没有扩展名，则需要配置此项
//
//	err := viper.ReadInConfig() // 查找并读取配置文件
//	if err != nil {             // 处理读取配置文件的错误
//		panic(fmt.Errorf("Fatal error config file: %s \n", err))
//	}
//
//	// 实时监控配置文件的变化 WatchConfig 开始监视配置文件的更改。
//	viper.WatchConfig()
//	// OnConfigChange设置配置文件更改时调用的事件处理程序。
//	// 当配置文件变化之后调用的一个回调函数
//	viper.OnConfigChange(func(e fsnotify.Event) {
//		fmt.Println("Config file changed:", e.Name)
//	})
//
//	r := gin.Default()
//	r.GET("/version", func(c *gin.Context) {
//		// GetString以字符串的形式返回与键相关的值。
//		c.String(http.StatusOK, viper.GetString("version"))
//	})
//	r.Run()
//
//}

func main() {
	viper.SetConfigName("config1")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}

	//fmt.Println("protocols: ", viper.GetStringSlice("server.protocols"))
	//fmt.Println("ports: ", viper.GetIntSlice("server.ports"))
	//fmt.Println("timeout: ", viper.GetDuration("server.timeout"))
	//
	//fmt.Println("mysql ip: ", viper.GetString("mysql.ip"))
	//fmt.Println("mysql port: ", viper.GetInt("mysql.port"))
	//
	//if viper.IsSet("redis.port") {
	//	fmt.Println("redis.port is set")
	//} else {
	//	fmt.Println("redis.port is not set")
	//}
	//
	//fmt.Println("mysql settings: ", viper.GetStringMap("mysql"))
	//fmt.Println("redis settings: ", viper.GetStringMap("redis"))
	//fmt.Println("all settings: ", viper.AllSettings())

	viper.Set("Verbose", false)
	viper.RegisterAlias("loud", "Verbose") // 注册别名（此处loud和Verbose建立了别名）

	//viper.Set("verbose", true) // 结果与下一行相同
	viper.Set("loud", true) // 结果与前一行相同

	fmt.Println("loud", viper.GetBool("loud"))       // true
	fmt.Println("verbase", viper.GetBool("verbose")) // true
}
