package http

import (
	"douyin/app/gateway/middleware"
	"douyin/app/gateway/rpc"
	pb "douyin/idl/pb/video"
	"douyin/response"
	"douyin/utils/e"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func Feed(ctx *gin.Context) {
	var req pb.DouyinFeedRequest
	//参数处理
	//通过token获取用户id
	token := ctx.Query("token")
	var userId int64
	if token == "" {
		userId = 0
	} else {
		tokenStruck, ok := middleware.CheckToken(token)
		//如果token无效
		if !ok {
			ctx.JSON(http.StatusOK, response.Feed_Response{
				Response: response.Response{
					StatusCode: 1,
					StatusMsg:  "token incorrect",
				},
			})
			// todo
			//global.SERVER_LOG.Warn("Token fail!")
			return
		}
		userId = tokenStruck.UserId
	}

	//获取最近时间
	strLatesttime := ctx.Query("latest_time")
	var latestTime int64

	latestTime, err := strconv.ParseInt(strLatesttime, 10, 64)
	if err != nil {
		latestTime = 0
	}

	//fmt.Println("userId:%v", userId)
	//fmt.Println("latesttime:%v", latestTime)

	//获取视频
	req.UserId = &userId
	req.LatestTime = &latestTime
	resp, err := rpc.Feed(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, response.Feed_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  "视频流获取失败",
			},
		})
		// todo
		//global.SERVER_LOG.Warn("Feed service fail!")
		return
	}

	//fmt.Println("%v", videoResponse)

	//返回
	videoResp := make([]response.Video_Response, len(resp.VideoList))
	for i := 0; i < len(resp.VideoList); i++ {
		video := resp.VideoList[i]
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
			IsFavorite:    *video.IsFavorite,
			Title:         *video.Title,
		}
	}
	ctx.JSON(http.StatusOK, response.Feed_Response{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "视频流获取成功",
		},
		VideoList: videoResp,
		NextTime:  *resp.NextTime,
	})
	// todo
	//global.SERVER_LOG.Info("Feed Success!")
	return
}

func PublishAction(c *gin.Context) {
	var req pb.DouyinPublishActionRequest

	//获得用户id
	getUserId, _ := c.Get("userid")
	var userId int64
	if v, ok := getUserId.(int64); ok {
		userId = v
	}

	//获得视频相关信息
	title := c.PostForm("title")
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, response.Publish_Action_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  e.ErrorVideoDataWrong.Error(),
			},
		})
		// todo
		//global.SERVER_LOG.Warn("Publish data fail!")
		return
	}

	//获得文件名并存储到本地
	fileName := fmt.Sprintf("%d_%s", userId, title) //标识名字
	//fmt.Println("filename:%s", fileName)
	fileExt := filepath.Ext(data.Filename)
	//fmt.Println("fileExt:%s", fileExt)
	finalFilename := fileName + fileExt
	//fmt.Println("finalFilename:%s", finalFilename)
	saveFilepath := filepath.Join("../tmp/", finalFilename) //路径+文件名
	//fmt.Println("saveFilepath:%s", saveFilepath)
	if err := c.SaveUploadedFile(data, saveFilepath); err != nil {
		c.JSON(http.StatusOK, response.Publish_Action_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  e.ErrorVideoDownloading.Error(),
			},
		})
		// todo
		//global.SERVER_LOG.Warn("Feed service fail!")
		return
	}
	//删除本地文件
	defer func() {
		err = os.Remove(saveFilepath)
		if err != nil {
			// todo
			//global.SERVER_LOG.Warn("File Deletion fail!")
		}
	}()

	//完成对象存储、以及数据库表活动
	req.UserId = &userId
	req.Title = &title
	req.FileExt = &fileExt
	_, err = rpc.PublishAction(c, &req)
	if err != nil {
		c.JSON(http.StatusOK, response.Publish_Action_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		// todo
		//global.SERVER_LOG.Warn("Publish service fail!")
		return
	}

	//返回成功信息
	c.JSON(http.StatusOK, response.Publish_Action_Response{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "视频发布成功",
		},
	})

	// todo
	//global.SERVER_LOG.Info("publish action success")

	return
}

func PublishList(c *gin.Context) {
	var req pb.DouyinPublishListRequest
	getHostId, _ := c.Get("userid")
	var hostId int64
	if v, ok := getHostId.(int64); ok {
		hostId = v
	}
	//2.查询要查看用户的id的所有视频，返回页面
	guestId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, response.Publish_List_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  e.ErrorVideoDataWrong.Error(),
			},
		})
		// todo
		//global.SERVER_LOG.Warn("Id fail!")
		return
	}
	req.QueryId = &guestId
	req.HostId = &hostId
	resp, err := rpc.PublishList(c, &req)
	if err != nil {
		c.JSON(http.StatusOK, response.Publish_List_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		// todo
		//global.SERVER_LOG.Warn("PublishList Service fail!")
		return
	}

	videoResp := make([]response.Video_Response, len(resp.VideoList))
	for i := 0; i < len(resp.VideoList); i++ {
		video := resp.VideoList[i]
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
			IsFavorite:    *video.IsFavorite,
			Title:         *video.Title,
		}
	}
	c.JSON(http.StatusOK, response.Publish_List_Response{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "视频列表获取成功",
		},
		VideoList: videoResp,
	})
	// todo
	//global.SERVER_LOG.Info("PublishList Success!")
	return
}
