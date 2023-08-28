package service

import (
	"context"
	"douyin/app/message/internal/dal/dao"
	"douyin/app/message/internal/dal/model"
	"time"
)

func ActionService(ctx context.Context, userId int64, toUserId int64, content string) error {
	newChat := model.Chat{
		Sender_id:    userId,
		Receiver_id:  toUserId,
		Content:      content,
		Publish_time: time.Now().Unix(),
	}
	err := dao.NewChatDao(ctx).InsertChat(newChat)
	return err
}

func ChatService(ctx context.Context, userId int64, toUserId int64, lastTime int64) ([]model.Chat, error) {
	return dao.NewChatDao(ctx).GetChatList(userId, toUserId, lastTime)
}
