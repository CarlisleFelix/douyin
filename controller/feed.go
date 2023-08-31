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

func Feed(c *gin.Context) {

	//参数处理

	//通过token获取用户id
	token := c.Query("token")
	var userId int64
	if token == "" {
		userId = 0
	} else {
		tokenStruck, ok := middleware.CheckToken(token)
		//如果token无效
		if !ok {
			c.JSON(http.StatusOK, response.Feed_Response{
				Response: response.Response{
					StatusCode: 1,
					StatusMsg:  "token incorrect",
				},
			})
			global.SERVER_LOG.Warn("Token fail!")
			return
		}
		userId = tokenStruck.UserId
	}

	//获取最近时间
	strLatesttime := c.Query("latest_time")
	var latestTime int64

	latestTime, err := strconv.ParseInt(strLatesttime, 10, 64)
	if err != nil {
		latestTime = 0
	}

	//fmt.Println("userId:%v", userId)
	//fmt.Println("latesttime:%v", latestTime)

	//获取视频
	videoResponse, nextTime, err := service.FeedService(userId, latestTime)
	if err != nil {
		c.JSON(http.StatusOK, response.Feed_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  "视频流获取失败",
			},
		})
		global.SERVER_LOG.Warn("Feed service fail!")
		return
	}

	//fmt.Println("%v", videoResponse)

	//返回
	c.JSON(http.StatusOK, response.Feed_Response{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "视频流获取成功",
		},
		VideoList: videoResponse,
		NextTime:  nextTime,
	})
	global.SERVER_LOG.Info("Feed Success!")
	return
}
