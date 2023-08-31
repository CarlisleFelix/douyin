package controller

import (
	"douyin/global"
	"douyin/middleware"
	"douyin/response"
	"douyin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func User(c *gin.Context) {

	//参数处理
	queryUserId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	rawId, _ := c.Get("userid")
	hostUserId, _ := rawId.(int64)

	if err != nil {
		c.JSON(http.StatusOK, response.User_Interface_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  global.ErrorParamFormatWrong.Error(),
			},
		})
		global.SERVER_LOG.Warn("param format wrong!", zap.String("error", err.Error()))
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
		global.SERVER_LOG.Warn("UserService fail!", zap.String("error", err.Error()))
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

	global.SERVER_LOG.Info("User Success!")

	return
}

func UserLogin(c *gin.Context) {
	//参数处理
	userName := c.Query("username")
	passWord := c.Query("password")

	//service层处理

	loginUser, err := service.UserLoginService(userName, passWord)
	if err != nil {
		c.JSON(http.StatusOK, response.User_Login_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		global.SERVER_LOG.Warn("UserLoginService fail!", zap.String("error", err.Error()))
		return
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
		global.SERVER_LOG.Warn("GenerateToken fail! Inconsistency has occured!", zap.String("error", err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.User_Register_Response{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "登录成功",
		},
		UserId: strconv.FormatInt(loginUser.User_id, 10),
		Token:  token,
	})

	global.SERVER_LOG.Info("UserLogin success!")
	//返回

	return
}

//问题在于如果插入数据库成功了，生成token失败了的话就会不一致，所以应该删除掉,还是先不考虑了

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
		global.SERVER_LOG.Warn("UserRegisterService fail!", zap.String("error", err.Error()))
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
		global.SERVER_LOG.Warn("GenerateToken fail! Inconsistency has occurred!", zap.String("error", err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.User_Register_Response{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "注册成功！",
		},
		UserId: strconv.FormatInt(newUser.User_id, 10),
		Token:  token,
	})

	global.SERVER_LOG.Info("UserRegister success!")
	//返回

	return
}
