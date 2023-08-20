package service

import (
	"douyin/dao"
	"douyin/model"
	"time"
)

func ActionService(userId int64, toUserId int64, content string) error {
	newChat := model.Chat{
		Sender_id:    userId,
		Receiver_id:  toUserId,
		Content:      content,
		Publish_time: time.Now().Unix(),
	}
	err := dao.InsertChat(newChat)
	return err
}

func ChatService(userId int64, toUserId int64) ([]model.Chat, error) {
	return dao.GetChatList(userId, toUserId)
}
