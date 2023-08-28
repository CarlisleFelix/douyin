package controller

import (
	"context"
	"douyin/app/message/internal/dal/model"
	"douyin/app/message/internal/service"
	pb "douyin/idl/pb/message"
	"douyin/response"
	"net/http"
	"strconv"
	"sync"
)

type MessageSrv struct {
	pb.UnimplementedMessageServiceServer
}

var MessageSrvIns *MessageSrv
var FavoriteSrvOnce sync.Once

func GetMessageSrv() *MessageSrv {
	FavoriteSrvOnce.Do(func() {
		MessageSrvIns = &MessageSrv{}
	})
	return MessageSrvIns
}

// 聊天记录
func (m *MessageSrv) MessageChat(ctx context.Context, req *pb.DouyinMessageChatRequest) (resp *pb.DouyinMessageChatResponse, err error) {
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

	err = service.ActionService(userId, toUserId, content)
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

// 消息操作
func (m *MessageSrv) MessageAction(ctx context.Context, req *pb.DouyinMessageActionRequest) (resp *pb.DouyinMessageActionResponse, err error) {
	_userId, _ := c.Get("userid")
	userId, _ := _userId.(int64)
	toUserId, err := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	lastTime, _ := strconv.ParseInt(c.Query("pre_msg_time"), 10, 64)
	lastTime = lastTime / 1000

	chatList, err := service.ChatService(userId, toUserId, lastTime)
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
