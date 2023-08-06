package controller

import (
	"douyin/response"

	"github.com/gin-gonic/gin"
)

type Favorite_Action_Response struct {
	response.Response
}

type Favorite_List_Response struct {
	response.Response
	VideoList []response.Video_Response `json:"video_list,omitempty"`
}

func FavoriteAction(c *gin.Context) {

	return
}

func FavoriteList(c *gin.Context) {

	return
}
