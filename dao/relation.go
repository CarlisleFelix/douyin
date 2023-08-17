package dao

import (
	"douyin/global"
	"douyin/model"

	"gorm.io/gorm"
)

// 检查user1是否follow了user2
func GetFollowByUserId(userId1 int64, userId2 int64) bool {
	relationship := model.Relation{}
	if userId1 == userId2 || userId1 == 0 {
		return false
	}
	if err := global.SERVER_DB.Model(&model.Relation{}).Where("host_id=? And guest_id=?", userId1, userId2).First(&relationship).Error; err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}
