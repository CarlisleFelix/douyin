package controller

import (
	"douyin/response"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FavoriteAction(c *gin.Context) {
	// 解析token生成user_id, 获取video_id, action_type
	user_id, _ := c.Get("userid")
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")

	// 初始化返回值
	var StatusCode int32
	var StatusMsg string

	// 参数校验
	// action_type 不为1或2时
	actionType, err := strconv.Atoi(action_type)
	if err != nil {
		StatusCode = 1
		StatusMsg = "未知操作"
		c.JSON(http.StatusOK, response.Favorite_Action_Response{
			Response: response.Response{StatusCode: StatusCode, StatusMsg: StatusMsg},
		})
	} else if actionType != 1 && actionType != 2 {
		StatusCode = 1
		StatusMsg = "未知操作"
		c.JSON(http.StatusOK, response.Favorite_Action_Response{
			Response: response.Response{StatusCode: StatusCode, StatusMsg: StatusMsg},
		})
	}

	// 处理请求
	StatusCode, StatusMsg = service.FavoriteAction(user_id.(int64), video_id, int32(actionType))

	// 返回
	c.JSON(http.StatusOK, response.Favorite_Action_Response{
		Response: response.Response{StatusCode: StatusCode, StatusMsg: StatusMsg},
	})
}

func FavoriteList(c *gin.Context) {
	// 获取参数
	user_id := c.Query("user_id")
	userId, _ := c.Get("userid")

	// 参数校验
	// 判断token解析出的userid与user_id是否一致
	id, _ := strconv.ParseInt(user_id, 10, 64)
	if id != userId {
		c.JSON(http.StatusOK, response.Favorite_List_Response{
			Response:  response.Response{StatusCode: 1, StatusMsg: "token与user_id不一致"},
			VideoList: nil,
		})
		return
	}

	// 处理请求
	statusCode, statusMsg, video_list := service.FavoriteList(id)

	// 返回
	c.JSON(http.StatusOK, response.Favorite_List_Response{
		Response:  response.Response{StatusCode: statusCode, StatusMsg: statusMsg},
		VideoList: video_list,
	})
}
