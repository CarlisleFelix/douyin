package http

import (
	"douyin/app/gateway/rpc"
	pb "douyin/idl/pb/favorite"
	"douyin/response"
	"douyin/utils/e"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FavoriteAction 点赞操作 处理参数 并调用rpc
func FavoriteAction(c *gin.Context) {
	var req pb.DouyinFavoriteActionRequest
	// 参数处理
	user_id, _ := c.Get("userid")
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")
	// action_type 无法解析时
	actionType, err := strconv.Atoi(action_type)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Favorite_Action_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  e.ErrorParamFormatWrong.Error(),
			},
		})
		// todo
		//global.SERVER_LOG.Warn("param format wrong!", zap.String("error", err.Error()))
		return
	}
	// action_type 不为1、2时
	if actionType != 1 && actionType != 2 {
		c.JSON(http.StatusInternalServerError, response.Favorite_Action_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  e.ErrorActionType.Error(),
			},
		})
		// todo
		//global.SERVER_LOG.Warn("unknown action!", zap.String("error", err.Error()))
		return
	}
	// 参数类型转换
	videoId, err := strconv.ParseInt(video_id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Favorite_Action_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  e.ErrorParamFormatWrong.Error(),
			},
		})
		// todo
		//global.SERVER_LOG.Warn("unknown video_id!", zap.String("error", err.Error()))
		return
	}

	// server层处理请求
	*req.UserId = user_id.(int64)
	req.VideoId = &videoId
	*req.ActionType = int32(actionType)
	_, err = rpc.FavoriteAction(c, &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Favorite_Action_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		// todo
		//global.SERVER_LOG.Warn("FavoriteAction fail!", zap.String("error", err.Error()))
		return
	}

	// 返回
	c.JSON(http.StatusOK, response.Favorite_Action_Response{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "操作成功",
		},
	})
	// todo
	//global.SERVER_LOG.Info("FavoriteAction success!")
}

// FavoriteList 获取点赞列表
func FavoriteList(c *gin.Context) {
	var req pb.DouyinFavoriteListRequest
	// 获取参数
	user_id := c.Query("user_id")
	userId, _ := c.Get("userid")

	// 参数校验
	// 判断token解析出的userid与user_id是否一致
	id, err := strconv.ParseInt(user_id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Favorite_List_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  e.ErrorParamFormatWrong.Error(),
			},
		})
		// todo:
		//global.SERVER_LOG.Warn("param format wrong!", zap.String("error", err.Error()))
		return
	}
	if id != userId {
		c.JSON(http.StatusInternalServerError, response.Favorite_List_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  e.ErrorParamMismatch.Error(),
			},
		})
		// todo
		//global.SERVER_LOG.Warn("parameter mismatch!", zap.String("error", err.Error()))
		return
	}

	// 处理请求
	req.UserId = &id
	videoList, err := rpc.FavoriteList(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Favorite_List_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		// todo
		//global.SERVER_LOG.Warn("get FavoriteList fail!", zap.String("error", err.Error()))
		return
	}

	videoResp := make([]response.Video_Response, len(videoList.VideoList))
	for i := 0; i < len(videoList.VideoList); i++ {
		video := videoList.VideoList[i]
		videoResp[i] = response.Video_Response{
			Id: *video.Id,
			Author: response.User_Response{
				Id:              *video.Author.Id,
				Name:            *video.Author.Name,
				FollowCount:     *video.Author.FollowCount,
				FollowerCount:   *video.Author.FollowerCount,
				IsFollow:        *video.Author.IsFollow,
				Avatar:          *video.Author.Avatar,
				BackgroundImage: *video.Author.BackgroundImage,
				Signature:       *video.Author.Signature,
				TotalFavorited:  *video.Author.TotalFavorited,
				WorkCount:       *video.Author.WorkCount,
				FavoriteCount:   *video.Author.FavoriteCount,
			},
			PlayUrl:       *video.PlayUrl,
			CoverUrl:      *video.CoverUrl,
			FavoriteCount: *video.FavoriteCount,
			CommentCount:  *video.CommentCount,
			IsFavorite:    true,
			Title:         *video.Title,
		}
	}
	c.JSON(http.StatusOK, response.Favorite_List_Response{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "列表获取成功",
		},
		VideoList: videoResp,
	})
	// todo
	//global.SERVER_LOG.Info("get FavoriteList success!")
}
