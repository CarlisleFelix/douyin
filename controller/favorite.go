package controller

import (
	"douyin/global"
	"douyin/response"
	"douyin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func FavoriteAction(c *gin.Context) {
	// 参数处理
	user_id, _ := c.Get("userid")
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")
	// action_type 无法解析时
	actionType, err := strconv.Atoi(action_type)
	if err != nil {
		c.JSON(http.StatusOK, response.Favorite_Action_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  global.ErrorParamFormatWrong.Error(),
			},
		})
		global.SERVER_LOG.Warn("param format wrong!", zap.String("error", err.Error()))
		return
	}
	// action_type 不为1、2时
	if actionType != 1 && actionType != 2 {
		c.JSON(http.StatusOK, response.Favorite_Action_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  global.ErrorActionType.Error(),
			},
		})
		global.SERVER_LOG.Warn("unknown action!", zap.String("error", err.Error()))
		return
	}

	// server层处理请求
	err = service.FavoriteAction(user_id.(int64), video_id, int32(actionType))

	if err != nil {
		c.JSON(http.StatusOK, response.Favorite_Action_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		global.SERVER_LOG.Warn("FavoriteAction fail!", zap.String("error", err.Error()))
		return
	}

	// 返回
	c.JSON(http.StatusOK, response.Favorite_Action_Response{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "操作成功",
		},
	})
	global.SERVER_LOG.Info("FavoriteAction success!")
}

func FavoriteList(c *gin.Context) {
	// 获取参数
	user_id := c.Query("user_id")
	userId, _ := c.Get("userid")

	// 参数校验
	// 判断token解析出的userid与user_id是否一致
	id, err := strconv.ParseInt(user_id, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, response.Favorite_List_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  global.ErrorParamFormatWrong.Error(),
			},
		})
		global.SERVER_LOG.Warn("param format wrong!", zap.String("error", err.Error()))
		return
	}
	if id != userId {
		c.JSON(http.StatusOK, response.Favorite_List_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  global.ErrorParamMismatch.Error(),
			},
		})
		//global.SERVER_LOG.Warn("parameter mismatch!", zap.String("error", err.Error()))
		return
	}

	// 处理请求
	video_list, err := service.FavoriteList(id)

	if err != nil {
		c.JSON(http.StatusOK, response.Favorite_List_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		global.SERVER_LOG.Warn("get FavoriteList fail!", zap.String("error", err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Favorite_List_Response{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "列表获取成功",
		},
		VideoList: video_list,
	})
	global.SERVER_LOG.Info("get FavoriteList success!")
}
