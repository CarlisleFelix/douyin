package controller

import (
	"douyin/global"
	"douyin/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func User(c *gin.Context) {
	// 从请求信息中获取用户名和密码
	var info struct {
		User_id int64  `json:"user_id"`
		Token   string `json:"token"`
	}

	// 从查询参数中获取 id 和 token
	info.User_id, _ = strconv.ParseInt(c.Query("id"), 10, 64)
	info.Token = c.Query("token")
	fmt.Println("Request Params:", info.User_id, info.Token) //test

	// 处理数据，response返回结构体
	response := service.UserInformation(info.User_id, info.Token)

	// 使用 c.JSON 返回响应数据
	c.JSON(http.StatusOK, response) //！notice：包括结构体里一直都是int64，直接强制转化为int不知道会不会有问题
	global.SERVER_LOG.Info("UserLogin success!")

}

func UserLogin(c *gin.Context) {
	// 从请求信息中获取用户名和密码
	var info struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// 将 JSON 数据绑定到 map
	if err := c.ShouldBindJSON(&info); err != nil {
		// 处理绑定错误
		c.JSON(http.StatusBadRequest, gin.H{"请求处理错误error": err.Error()})
		return
	}
	fmt.Println("Request Params:", info.Username, info.Password) //test

	// 处理数据，response返回结构体
	response := service.Login(info.Username, info.Password)

	// 使用 c.JSON 返回响应数据
	c.JSON(http.StatusOK, response) //！notice：包括结构体里一直都是int64，直接强制转化为int不知道会不会有问题
	global.SERVER_LOG.Info("UserLogin success!")
}

func UserRegister(c *gin.Context) {

	var info struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// 放在body中必须用bind
	if err := c.ShouldBindJSON(&info); err != nil {
		// 处理绑定错误
		c.JSON(http.StatusBadRequest, gin.H{"请求处理错误error": err.Error()})
		return
	}
	fmt.Println("Request Params:", info.Username, info.Password)

	// 处理数据，response返回结构体
	response := service.Register(info.Username, info.Password)

	// 使用 c.JSON 返回响应数据
	c.JSON(http.StatusOK, response) //！notice：包括结构体里一直都是int64，直接强制转化为int不知道会不会有问题
	global.SERVER_LOG.Info("UserRegister success!")

	return
}

// 手机号验证登录
func UserSmsRegister(c *gin.Context) {

	var info struct {
		Username  string `json:"username"`
		Telephone string `json:"telephone"`
	}

	// 放在body中必须用bind
	if err := c.ShouldBindJSON(&info); err != nil {
		// 处理绑定错误
		c.JSON(http.StatusBadRequest, gin.H{"请求处理错误error": err.Error()})
		return
	}
	fmt.Println("Request Params:", info.Username, info.Telephone)

	// 处理数据，response返回结构体
	response := service.Register(info.Username, info.Telephone)

	// 使用 c.JSON 返回响应数据
	c.JSON(http.StatusOK, response) //！notice：包括结构体里一直都是int64，直接强制转化为int不知道会不会有问题
	global.SERVER_LOG.Info("UserRegister success!")

	return
}
