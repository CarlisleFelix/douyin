package controller

import (
	"douyin/global"
	"douyin/model"
	"douyin/response"
	"douyin/service"
	"douyin/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 发送消息
func MessageAction(c *gin.Context) {

	ctx, span := global.SERVER_MESSAGE_TRACER.Start(c.Request.Context(), "messageaction controller")
	defer span.End()

	toUserId, err := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	actionType, err := strconv.ParseInt(c.Query("action_type"), 10, 64)
	_userId, _ := c.Get("userid")
	userId, _ := _userId.(int64)
	content := c.Query("content")
	if actionType != 1 {
		c.JSON(http.StatusOK, response.Message_Action_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  "操作类型不为0",
			},
		})

		return
	}

	err = service.ActionService(userId, toUserId, content, ctx)
	if err != nil {
		c.JSON(http.StatusOK, response.Message_Action_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, response.Message_Action_Response{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "发送成功",
		},
	})
	return
}

// 接收消息
func MessageChat(c *gin.Context) {

	ctx, span := global.SERVER_MESSAGE_TRACER.Start(c.Request.Context(), "messagechat controller")
	defer span.End()

	_userId, _ := c.Get("userid")
	userId, _ := _userId.(int64)
	toUserId, err := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	lastTime, _ := strconv.ParseInt(c.Query("pre_msg_time"), 10, 64)
	lastTime = lastTime / 1000

	chatList, err := service.ChatService(userId, toUserId, lastTime, ctx)
	if err != nil {
		c.JSON(http.StatusOK, response.Message_Chat_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, response.Message_Chat_Response{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "接收成功",
		},
		MessageList: ChatList2MessageResponseList(chatList),
	})
	return
}

func ChatList2MessageResponseList(chatList []model.Chat) []response.Message_Response {
	responseList := []response.Message_Response{}
	for i := 0; i < len(chatList); i++ {
		// fmt.Println(chatList[i])
		responseList = append(responseList, response.Message_Response{
			Id:         chatList[i].Id,
			ToUserId:   chatList[i].Receiver_id,
			FromUserID: chatList[i].Sender_id,
			Content:    chatList[i].Content,
			CreateTime: utils.IntTime2ChatTime(chatList[i].Publish_time),
		})
	}
	return responseList
}
