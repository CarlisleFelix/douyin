package controller

import (
	"douyin/global"
	"douyin/response"
	"douyin/service"
	"douyin/utils"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PublishAction(c *gin.Context) {

	curTime := utils.CurrentTimeInt()

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
				StatusMsg:  global.ErrorVideoDataWrong.Error(),
			},
		})
		global.SERVER_LOG.Warn("Publish data fail!")
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
				StatusMsg:  global.ErrorVideoDownloading.Error(),
			},
		})
		global.SERVER_LOG.Warn("Feed service fail!")
		return
	}
	//删除本地文件
	defer func() {
		err = os.Remove(saveFilepath)
		if err != nil {
			global.SERVER_LOG.Warn("File Deletion fail!")
		}
	}()

	//完成对象存储、以及数据库表活动
	err = service.PublishService(userId, title, fileExt, curTime)
	if err != nil {
		c.JSON(http.StatusOK, response.Publish_Action_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		global.SERVER_LOG.Warn("Publish service fail!")
		return
	}

	//返回成功信息
	c.JSON(http.StatusOK, response.Publish_Action_Response{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "视频发布成功",
		},
	})

	global.SERVER_LOG.Info("publish action success")

	return
}

func PublishList(c *gin.Context) {
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
				StatusMsg:  global.ErrorVideoDataWrong.Error(),
			},
		})
		global.SERVER_LOG.Warn("Id fail!")
		return
	}
	videoResponse, err := service.PublishListService(guestId, hostId)
	if err != nil {
		c.JSON(http.StatusOK, response.Publish_List_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		global.SERVER_LOG.Warn("PublishList Service fail!")
		return
	}
	c.JSON(http.StatusOK, response.Publish_List_Response{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "视频列表获取成功",
		},
		VideoList: videoResponse,
	})
	global.SERVER_LOG.Info("PublishList Success!")
	return
}
