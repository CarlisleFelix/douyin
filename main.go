package main

import (
	"douyin/initialization"
	"douyin/middleware/rabbitmq"
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
	rabbitmq.InitRabbitMQ()
	// 初始化Follow的相关消息队列，并开启消费。
	rabbitmq.InitFollowRabbitMQ()
	//初始化Comment的消息队列，并开启消费
	rabbitmq.InitCommentRabbitMQ()
	// TODO：5.开启服务器
	initialization.RunServer()
}
