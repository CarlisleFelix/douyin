package main

import (
	"douyin/global"
	"douyin/initialization"
	"douyin/model"
	"time"
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

	// TODO：5.开启服务器
	// core.RunServer()

	nowtime := time.Now().Format("01-02")
	comment := model.Comment{
		User_id:      12,
		Video_id:     14,
		Comment:      "hhhhhhh",
		Publish_time: nowtime,
	}
	result := global.SERVER_DB.Create(&comment)
	if result.Error != nil {
		global.SERVER_LOG.Info("comment1 insertion failed")
	}

	nowtime = time.Now().Format("02-03")
	comment = model.Comment{
		User_id:      12,
		Video_id:     14,
		Comment:      "hhhhhhh",
		Publish_time: nowtime,
	}
	result = global.SERVER_DB.Create(&comment)
	if result.Error != nil {
		global.SERVER_LOG.Info("comment2 insertion failed")
	}

	nowtime = time.Now().Format("04-05")
	comment = model.Comment{
		User_id:      12,
		Video_id:     14,
		Comment:      "hhhhhhh",
		Publish_time: nowtime,
	}
	result = global.SERVER_DB.Create(&comment)
	if result.Error != nil {
		global.SERVER_LOG.Info("comment3 insertion failed")
	}

}
