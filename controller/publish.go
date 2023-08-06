package controller

import (
	"douyin/response"

	"github.com/gin-gonic/gin"
)

type Publish_Action_Response struct {
	response.Response
}

type Publish_List_Response struct {
	response.Response
	VideoList []response.Video_Response `json:"video_list,omitempty"`
}

func PublishAction(c *gin.Context) {

	return
}

func PublishList(c *gin.Context) {

	return
}
