package http

import (
	"douyin/app/gateway/middleware"
	"douyin/app/gateway/rpc"
	pb "douyin/idl/pb/user"
	"douyin/response"
	"douyin/utils/e"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func User(c *gin.Context) {
	var req pb.DouyinUserRequest
	//参数处理
	queryUserId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	rawId, _ := c.Get("userid")
	hostUserId, _ := rawId.(int64)

	if err != nil {
		c.JSON(http.StatusOK, response.User_Interface_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  e.ErrorParamFormatWrong.Error(),
			},
		})
		// todo
		//global.SERVER_LOG.Warn("param format wrong!", zap.String("error", err.Error()))
		return
	}

	//service层处理
	req.QueryId = &queryUserId
	req.HostId = &hostUserId
	userResp, err := rpc.GetUserInfo(c, &req)

	if err != nil {
		c.JSON(http.StatusOK, response.User_Interface_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		// todo
		//global.SERVER_LOG.Warn("UserService fail!", zap.String("error", err.Error()))
		return
	}

	//返回
	c.JSON(http.StatusOK, response.User_Interface_Response{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "查询成功",
		},
		User_Response: response.User_Response{
			Id:              *userResp.User.Id,
			Name:            *userResp.User.Name,
			FollowCount:     *userResp.User.FollowCount,
			FollowerCount:   *userResp.User.FollowerCount,
			IsFollow:        true,
			Avatar:          *userResp.User.Avatar,
			BackgroundImage: *userResp.User.BackgroundImage,
			Signature:       *userResp.User.Signature,
			TotalFavorited:  *userResp.User.TotalFavorited,
			WorkCount:       *userResp.User.WorkCount,
			FavoriteCount:   *userResp.User.FavoriteCount,
		},
	})
	// todo
	//global.SERVER_LOG.Info("User Success!")

	return
}

func UserLogin(c *gin.Context) {
	var req pb.DouyinUserLoginRequest
	//参数处理
	userName := c.Query("username")
	passWord := c.Query("password")

	//service层处理
	req.Username = &userName
	req.Password = &passWord
	loginUser, err := rpc.UserLogin(c, &req)
	if err != nil {
		c.JSON(http.StatusOK, response.User_Login_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		// todo
		//global.SERVER_LOG.Warn("UserLoginService fail!", zap.String("error", err.Error()))
		return
	}

	if err != nil {
		c.JSON(http.StatusOK, response.User_Register_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		// todo
		//global.SERVER_LOG.Warn("GenerateToken fail! Inconsistency has occured!", zap.String("error", err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.User_Register_Response{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "登录成功",
		},
		UserId: strconv.FormatInt(*loginUser.UserId, 10),
		Token:  *loginUser.Token,
	})
	// todo
	//global.SERVER_LOG.Info("UserLogin success!")
	//返回

	return
}

//问题在于如果插入数据库成功了，生成token失败了的话就会不一致，所以应该删除掉,还是先不考虑了

func UserRegister(c *gin.Context) {
	var req pb.DouyinUserRegisterRequest
	//参数处理
	userName := c.Query("username")
	passWord := c.Query("password")

	//service层处理

	//新建用户
	req.Username = &userName
	req.Password = &passWord
	newUser, err := rpc.UserRegister(c, &req)

	if err != nil {
		c.JSON(http.StatusOK, response.User_Register_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		// todo
		//global.SERVER_LOG.Warn("UserRegisterService fail!", zap.String("error", err.Error()))
		return
	}

	//生成token
	token, err := middleware.GenerateToken(*newUser.UserId, *newUser.UserName)

	if err != nil {
		c.JSON(http.StatusOK, response.User_Register_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		// todo
		//global.SERVER_LOG.Warn("GenerateToken fail! Inconsistency has occurred!", zap.String("error", err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.User_Register_Response{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "注册成功！",
		},
		UserId: strconv.FormatInt(*newUser.UserId, 10),
		Token:  token,
	})
	// todo
	//global.SERVER_LOG.Info("UserRegister success!")
	//返回

	return
}
