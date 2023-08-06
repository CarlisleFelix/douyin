package controller

import (
	"douyin/response"

	"github.com/gin-gonic/gin"
)

type Comment_Action_Response struct {
	response.Response
	response.Comment_Response `json:"comment,omitempty"`
}

type Comment_List_Response struct {
	response.Response
	CommentList []response.Comment_Response `json:"comment_list,omitempty"`
}

func CommentAction(c *gin.Context) {

	return
}

func CommentList(c *gin.Context) {

	return
}
