package routes

import (
	"github.com/gin-gonic/gin"

	"douyin/app/gateway/internal/http"
	"douyin/app/gateway/middleware"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	//ginRouter.Use(middleware.Cors(), middleware.ErrorMiddleware())
	//store := cookie.NewStore([]byte("something-very-secret"))
	//router.Use(sessions.Sessions("mysession", store))
	publicGroup := router.Group("/douyin")
	{
		//用户
		UserGroup := publicGroup.Group("/user")
		{
			UserGroup.GET("/", http.GetUserInfo)
			UserGroup.POST("/login/", http.UserLogin)
			UserGroup.POST("/register/", http.UserRegister)

		}
		////视频投稿
		//publishGroup := publicGroup.Group("/publish")
		//{
		//	publishGroup.POST("/action/", middleware.JwtMiddleware(), controller.PublishAction)
		//	publishGroup.GET("/list/", middleware.JwtMiddleware(), controller.PublishList)
		//}
		////视频浏览
		//feedGroup := publicGroup.Group("/feed")
		//{
		//	feedGroup.GET("/", controller.Feed)
		//}
		//赞
		publicGroup.GET("/test", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"msg": "OK!",
			})
		})
		favoriteGroup := publicGroup.Group("/favorite")
		{
			favoriteGroup.POST("/action/", http.FavoriteAction)
			favoriteGroup.GET("/list/", middleware.JwtMiddleware(), http.FavoriteList)
		}
		////评论
		//commentGroup := publicGroup.Group("/comment")
		//{
		//	commentGroup.POST("/action/", middleware.JwtMiddleware(), controller.CommentAction)
		//	commentGroup.GET("/list/", middleware.JwtMiddleware(), controller.CommentList)
		//}
		////社交
		//relationGroup := publicGroup.Group("/relation")
		//{
		//	relationGroup.POST("/action/", middleware.JwtMiddleware(), controller.RelationAction)
		//	relationGroup.GET("/follow/list/", middleware.JwtMiddleware(), controller.FollowList)
		//	relationGroup.GET("/follower/list/", middleware.JwtMiddleware(), controller.FollowerList)
		//	relationGroup.GET("/friend/list/", middleware.JwtMiddleware(), controller.FriendList)
		//}
		////消息
		//messageGroup := publicGroup.Group("/message")
		//{
		//	messageGroup.GET("/chat/", middleware.JwtMiddleware(), controller.MessageChat)
		//	messageGroup.POST("/action/", middleware.JwtMiddleware(), controller.MessageAction)
		//}
	}
	return router
}