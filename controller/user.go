package controller

import (
	"douyin/middleware"
	"douyin/response"
	"douyin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
	//参数处理

	return
}

func UserLogin(c *gin.Context) {
	//参数处理
	userName := c.Query("username")
	passWord := c.Query("password")

	//service层处理

	loginUser, err := service.UserRegisterService(userName, passWord)
	if err != nil {
		c.JSON(http.StatusOK, response.User_Login_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
	}

	//生成token
	token, err := middleware.GenerateToken(loginUser.User_id, loginUser.User_name)

	if err != nil {
		c.JSON(http.StatusOK, response.User_Register_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
	}

	c.JSON(http.StatusOK, response.User_Register_Response{
		Response: response.Response{StatusCode: 0},
		UserId:   strconv.FormatInt(loginUser.User_id, 10),
		Token:    token,
	})

	//返回

	return

	return
}

func UserRegister(c *gin.Context) {
	//参数处理
	userName := c.Query("username")
	passWord := c.Query("password")

	//service层处理

	//新建用户
	newUser, err := service.UserRegisterService(userName, passWord)

	if err != nil {
		c.JSON(http.StatusOK, response.User_Register_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	//生成token
	token, err := middleware.GenerateToken(newUser.User_id, newUser.User_name)

	if err != nil {
		c.JSON(http.StatusOK, response.User_Register_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
	}

	c.JSON(http.StatusOK, response.User_Register_Response{
		Response: response.Response{StatusCode: 0},
		UserId:   strconv.FormatInt(newUser.User_id, 10),
		Token:    token,
	})

	//返回

	return
}
