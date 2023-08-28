package http

import (
	"douyin/app/gateway/rpc"
	pb "douyin/idl/pb/message"
	"douyin/response"
	"douyin/utils/ctl"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 发送消息
func MessageAction(ctx *gin.Context) {
	var req pb.DouyinMessageActionRequest
	toUserId, err := strconv.ParseInt(ctx.Query("to_user_id"), 10, 64)
	actionType, err := strconv.ParseInt(ctx.Query("action_type"), 10, 64)
	_userId, _ := ctx.Get("userid")
	userId, _ := _userId.(int64)
	content := ctx.Query("content")
	if actionType != 1 {
		ctx.JSON(http.StatusOK, response.Message_Action_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  "操作类型不为0",
			},
		})

		return
	}

	req.UserId = &userId
	req.ToUserId = &toUserId
	req.Content = &content
	_, err = rpc.MessageAction(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, response.Message_Action_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Message_Action_Response{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "发送成功",
		},
	})
	return
}

// 接收消息
func MessageChat(ctx *gin.Context) {
	var req pb.DouyinMessageChatRequest
	_userId, _ := ctx.Get("userid")
	userId, _ := _userId.(int64)
	toUserId, err := strconv.ParseInt(ctx.Query("to_user_id"), 10, 64)
	lastTime, _ := strconv.ParseInt(ctx.Query("pre_msg_time"), 10, 64)
	lastTime = lastTime / 1000

	req.UserId = &userId
	req.ToUserId = &toUserId
	req.PreMsgTime = &lastTime
	chatList, err := rpc.MessageChat(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, response.Message_Chat_Response{
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
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, chatList))
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
