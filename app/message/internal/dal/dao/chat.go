package dao

import (
	"context"
	"douyin/app/message/internal/dal/model"
	"douyin/app/message/internal/global"
	"fmt"
	"gorm.io/gorm"
)

type ChatDao struct {
	*gorm.DB
}

func NewChatDao(ctx context.Context) *ChatDao {
	return &ChatDao{NewDBClient(ctx)}
}

func (dao *ChatDao) InsertChat(newChat model.Chat) error {
	err := _db.Create(&newChat).Error
	if err != nil {
		global.SERVER_LOG.Info("error of database while create message!")
	}
	return err
}

func (dao *ChatDao) GetChatList(userId int64, toUserId int64, lastTime int64) ([]model.Chat, error) {
	chatList := []model.Chat{}
	fmt.Println("lastTime:", lastTime)
	//err := global.SERVER_DB.Where("(sender_id=? and receiver_id=?) or (receiver_id=? and sender_id=?) and publish_time > ? and publish_time != ?", userId, toUserId, userId, toUserId, lastTime, lastTime).Order("publish_time asc").Find(&chatList).Error
	err := _db.Where("(sender_id=? and receiver_id=?) or (receiver_id=? and sender_id=?)", userId, toUserId, userId, toUserId).Order("publish_time asc").Find(&chatList).Error
	finalchatList := []model.Chat{}
	for i := 0; i < len(chatList); i++ {
		if chatList[i].Publish_time > lastTime {
			finalchatList = chatList[i:]
			break
		}
	}
	return finalchatList, err
}

// GetLatestChat 获取 from 与 to 之间 最近的一条消息内容(from 发给 to)
func (dao *ChatDao) GetLatestChat(fromId int64, toId int64) (model.Chat, int64, error) {
	var fromTo []model.Chat // from 发给 to
	err := _db.Table("chats").Where("sender_id = ? AND receiver_id = ?", fromId, toId).Order("publish_time desc").Limit(1).Find(&fromTo).Error
	if err != nil || len(fromTo) == 0 {
		return model.Chat{}, -1, err
	} else {
		return fromTo[0], 1, nil
	}
}

// QueryNewestMessageByUserIdAndToUserID 通过两者的用户Id查询最新最新的两者之间的聊天记录 0-接受 1-发送 有点问题
func (dao *ChatDao) QueryNewestMessageByUserIdAndToUserID(Receiver_id int64, Sender_id int64) (string, int8, error) {
	message := model.Chat{}
	result := _db.Debug().Where("sender_id = ? AND receiver_id = ?", Receiver_id, Sender_id).Or("receiver_id = ? AND sender_id = ?", Receiver_id, Sender_id).Order("createTime desc").Limit(1).Find(&message)
	if result.Error != nil {
		return "", -1, result.Error
	}
	if Sender_id == Receiver_id {
		return message.Content, 1, nil
	} else {
		return message.Content, 0, nil
	}
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
