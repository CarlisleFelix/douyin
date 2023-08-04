package initialization

import (
	"douyin/controller"
	"douyin/middleware"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	router := gin.Default()

	publicGroup := router.Group("/douyin")
	{
		//用户
		UserGroup := publicGroup.Group("/user")
		{
			UserGroup.GET("/", middleware.JwtMiddleware(), controller.User)
			UserGroup.POST("/login/", controller.UserLogin)
			UserGroup.POST("/register/", controller.UserRegister)

		}
		//视频投稿
		publishGroup := publicGroup.Group("/publish")
		{
			publishGroup.POST("/action/", middleware.JwtMiddleware(), controller.PublishAction)
			publishGroup.GET("/list/", middleware.JwtMiddleware(), controller.PublishList)
		}
		//视频浏览
		feedGroup := publishGroup.Group("/feed")
		{
			feedGroup.GET("/", controller.Feed)
		}
		//赞
		favoriteGroup := publicGroup.Group("/favorite")
		{
			favoriteGroup.POST("/action/", middleware.JwtMiddleware(), controller.FavoriteAction)
			favoriteGroup.GET("/list/", middleware.JwtMiddleware(), controller.FavoriteList)
		}
		//评论
		commentGroup := publicGroup.Group("/comment")
		{
			commentGroup.POST("/action/", middleware.JwtMiddleware(), controller.CommentAction)
			commentGroup.GET("/list/", middleware.JwtMiddleware(), controller.CommentList)
		}
		//社交
		relationGroup := publicGroup.Group("/relation")
		{
			relationGroup.POST("/action/", middleware.JwtMiddleware(), controller.RelationAction)
			relationGroup.GET("/follow/list/", middleware.JwtMiddleware(), controller.FollowList)
			relationGroup.GET("/follower/list/", middleware.JwtMiddleware(), controller.FollowerList)
			relationGroup.GET("/friend/list/", middleware.JwtMiddleware(), controller.FriendList)
		}
		//消息
		messageGroup := publicGroup.Group("/message")
		{
			messageGroup.GET("/chat/", middleware.JwtMiddleware(), controller.MessageChat)
			messageGroup.POST("/action/", middleware.JwtMiddleware(), controller.MessageAction)
		}
	}

	return router
}
