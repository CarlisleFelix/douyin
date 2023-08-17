package controller

import (
	"douyin/global"
	"douyin/middleware"
	"douyin/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func User(c *gin.Context) {
}

func UserLogin(c *gin.Context) {
	username := c.Query("username")
	//password := c.Query("password")

	// 查询密码是否与用户名匹配
	// 查询用户id
	var user model.User
	global.SERVER_DB.Where("name = ?", username).First(&user)
	id := user.User_id
	// 通过id与用户名生成token
	token, _ := middleware.GenerateToken(id, username)
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "测试",
		"user_id":     id,
		"token":       token,
	})
	return
}

func UserRegister(c *gin.Context) {

}
