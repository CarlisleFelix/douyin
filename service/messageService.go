package service

import (
	"context"
	"douyin/dao"
	"douyin/global"
	"douyin/model"
	"time"
)

func ActionService(userId int64, toUserId int64, content string, ctx context.Context) error {

	ctx, span := global.SERVER_MESSAGE_TRACER.Start(ctx, "messageaction service")
	defer span.End()

	newChat := model.Chat{
		Sender_id:    userId,
		Receiver_id:  toUserId,
		Content:      content,
		Publish_time: time.Now().Unix(),
	}
	err := dao.InsertChat(newChat)
	return err
}

func ChatService(userId int64, toUserId int64, lastTime int64, ctx context.Context) ([]model.Chat, error) {
	ctx, span := global.SERVER_MESSAGE_TRACER.Start(ctx, "messageaction service")
	defer span.End()
	return dao.GetChatList(userId, toUserId, lastTime)
}
