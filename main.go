package main

import (
	"douyin/initialization"
)

//const AppMode = "debug"

func main() {
	//gin.SetMode(AppMode)

	// TODO：1.配置初始化
	initialization.InitializeViper()

	// TODO：2.日志初始化
	initialization.InitializeZap()

	// TODO: 3.数据库初始化，数据表自动创建
	initialization.InitializeGorm()

	// TODO: 4.路由等其他初始化，中间件初始化
	initialization.InitializeCos()

	initialization.InitializeRedis()
	// 初始化rabbitMQ。
	initialization.InitRabbitMQ()

	initialization.InitializeJaeger()

	initialization.InitializeContext()

	// TODO：5.开启服务器
	initialization.RunServer()
}
