package dao

import (
	"douyin/global"
	"douyin/model"
)

func InsertChat(newChat model.Chat) error {
	err := global.SERVER_DB.Create(&newChat).Error
	if err != nil {
		global.SERVER_LOG.Info("error of database while create chat!")
	}
	return err
}

func GetChatList(userId int64, toUserId int64) ([]model.Chat, error) {
	chatList := []model.Chat{}
	err := global.SERVER_DB.Where("sender_id=? and receiver_id=?", userId, toUserId).Find(&chatList).Error
	return chatList, err
}

// func GetLatestChat(userId int64, toUserId int64) (model.Chat, int, error) {
// 	latestRecvChat := model.Chat{}
// 	latestSendChat := model.Chat{}
// 	err := global.SERVER_DB.Where("id=? and to_user_id=?", userId, toUserId).Last(&latestRecvChat).Error
// 	err = global.SERVER_DB.Where("id=? and to_user_id=?", toUserId, userId).Last(&latestSendChat).Error
// 	if latestRecvChat.Publish_time > latestSendChat.Publish_time {
// 		return latestRecvChat, 0, err
// 	}
// 	return latestSendChat, 1, err
// }

// GetLatestChat 获取 from 与 to 之间 最近的一条消息内容(from 发给 to)
func GetLatestChat(fromId int64, toId int64) (model.Chat, int64, error) {
	var fromTo []model.Chat // from 发给 to
	err := global.SERVER_DB.Table("chats").Where("sender_id = ? AND receiver_id = ?", fromId, toId).Order("publish_time desc").Limit(1).Find(&fromTo).Error
	if err != nil || len(fromTo) == 0 {
		return model.Chat{}, -1, err
	} else {
		return fromTo[0], 1, nil
	}
}

// QueryNewestMessageByUserIdAndToUserID 通过两者的用户Id查询最新最新的两者之间的聊天记录 0-接受 1-发送 有点问题
func QueryNewestMessageByUserIdAndToUserID(Receiver_id int64, Sender_id int64) (string, int8, error) {
	message := model.Chat{}
	result := global.SERVER_DB.Debug().Where("sender_id = ? AND receiver_id = ?", Receiver_id, Sender_id).Or("receiver_id = ? AND sender_id = ?", Receiver_id, Sender_id).Order("createTime desc").Limit(1).Find(&message)
	if result.Error != nil {
		return "", -1, result.Error
	}
	if Sender_id == Receiver_id {
		return message.Content, 1, nil
	} else {
		return message.Content, 0, nil
	}
}
