package controller

import (
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func User(c *gin.Context) {
	// 从请求信息中获取用户名和密码
	var info map[string]string

	// 将 JSON 数据绑定到 map
	if err := c.ShouldBindJSON(&info); err != nil {
		// 处理绑定错误
		c.JSON(http.StatusBadRequest, gin.H{"请求处理错误error": err.Error()})
		return
	}

	// 处理数据，response返回结构体
	response := service.UserInformation(info)
	// 使用 c.JSON 返回响应数据
	c.JSON(int(response.Response.StatusCode), response) //！notice：包括结构体里一直都是int64，直接强制转化为int不知道会不会有问题

}

func UserLogin(c *gin.Context) {
	// 从请求信息中获取用户名和密码
	var info map[string]string

	// 将 JSON 数据绑定到 map
	if err := c.ShouldBindJSON(&info); err != nil {
		// 处理绑定错误
		c.JSON(http.StatusBadRequest, gin.H{"请求处理错误error": err.Error()})
		return
	}

	// 处理数据，response返回结构体
	response := service.Login(info)
	// 使用 c.JSON 返回响应数据
	c.JSON(int(response.Response.StatusCode), response) //！notice：包括结构体里一直都是int64，直接强制转化为int不知道会不会有问题
}

func UserRegister(c *gin.Context) {
	// 从请求信息中获取用户名和密码
	var info map[string]string

	// 将 JSON 数据绑定到 map
	if err := c.ShouldBindJSON(&info); err != nil {
		// 处理绑定错误
		c.JSON(http.StatusBadRequest, gin.H{"请求处理错误error": err.Error()})
		return
	}

	// 处理数据，response返回结构体
	response := service.Register(info)
	// 使用 c.JSON 返回响应数据
	c.JSON(int(response.Response.StatusCode), response) //！notice：包括结构体里一直都是int64，直接强制转化为int不知道会不会有问题

}
