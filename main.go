package main

import (
	"douyin/initialization"
	"fmt"
)

//const AppMode = "debug"

func main() {
	fmt.Println("Hello, World!")
	//gin.SetMode(AppMode)

	// TODO：1.配置初始化
	initialization.InitializeViper()

	// TODO：2.日志初始化
	initialization.InitializeZap()

	// TODO: 3.数据库初始化，数据表自动创建
	initialization.InitializeGorm()

	// TODO: 4.路由等其他初始化，中间件初始化

	// TODO：5.开启服务器
	initialization.RunServer()
}
