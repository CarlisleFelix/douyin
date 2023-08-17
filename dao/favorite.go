package dao

import (
	"douyin/global"
	"douyin/model"

	"gorm.io/gorm"
)

func GetifFavorite(userId int64, videoId int64) bool {
	if userId == 0 {
		return false
	}
	favorite := model.Favorite{}
	err := global.SERVER_DB.Where("user_id = ? AND video_id = ?", userId, videoId).First(&favorite).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		} else {
			return false
		}
	}
	return true
}
