package controller

import (
	"douyin/response"

	"github.com/gin-gonic/gin"
)

type Message_Chat_Response struct {
	response.Comment_Response
	MessageList []response.Message_Response `json:"message_list,omitempty"`
}

type Message_Action_Response struct {
	response.Comment_Response
}

func MessageChat(c *gin.Context) {

	return
}

func MessageAction(c *gin.Context) {

	return
}
