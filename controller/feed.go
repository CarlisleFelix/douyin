package controller

import (
	"github.com/gin-gonic/gin"
)

func Feed(c *gin.Context) {

	// //参数处理

	// //通过token获取用户id
	// token := c.Query("token")
	// var userId int64

	// if token == "" {
	// 	userId = 0
	// } else {

	// 	userId, err := strconv.ParseInt(c.Query("userid"), 10, 64)
	// 	if err != nil {
	// 		userId = 0
	// 	}
	// }

	// //获取最近时间
	// strLatesttime := c.Query("latest_time")
	// var latestTime int64

	// if strLatesttime == "" {
	// 	latestTime = 0
	// } else {

	// 	latestTime, err := strconv.ParseInt(strLatesttime, 10, 64)
	// 	if err != nil {
	// 		latestTime = 0
	// 	}
	// }

	// //准备数据，service层进行处理

	// //获取视频
	// videoList := service.GetVideolist(userId, latestTime)

}
