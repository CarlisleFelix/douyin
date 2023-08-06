package controller

import (
	"douyin/response"
	"douyin/service"

	"github.com/gin-gonic/gin"
)

type User_Register_Response struct {
	response.Response
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

type User_Login_Response struct {
	response.Response
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

type User_Interface_Response struct {
	response.Response
	response.User_Response
}

func User(c *gin.Context) {

	return
}

func UserLogin(c *gin.Context) {

	return
}

func UserRegister(c *gin.Context) {

	//参数处理
	userName := c.Query("username")
	passWord := c.Query("password")

	//service层处理

	service.UserRegisterService(userName, passWord)

	//返回

	return
}
