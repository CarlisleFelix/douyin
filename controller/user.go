package controller

import (
	"douyin/global"
	"douyin/middleware"
	"douyin/response"
	"douyin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {

	//参数处理
	queryUserId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	hostUserId, err := strconv.ParseInt(c.Query("userid"), 10, 64)

	if err != nil {
		c.JSON(http.StatusOK, response.User_Interface_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  global.ErrorParamFormatWrong.Error(),
			},
		})
		return
	}

	//service层处理
	userResponse, err := service.UserService(queryUserId, hostUserId)

	if err != nil {
		c.JSON(http.StatusOK, response.User_Interface_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	//返回
	c.JSON(http.StatusOK, response.User_Interface_Response{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "查询成功",
		},
		User_Response: userResponse,
	})

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
