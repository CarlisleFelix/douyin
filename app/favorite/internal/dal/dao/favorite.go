package dao

import (
	"context"
	"douyin/app/favorite/internal/dal/model"
	"gorm.io/gorm"
)

type FavoriteDao struct {
	*gorm.DB
}

func NewFavoriteDao(ctx context.Context) *FavoriteDao {
	return &FavoriteDao{NewDBClient(ctx)}
}

func (dao *FavoriteDao) IsFavorite(userId int64, videoId int64) bool {
	if userId == 0 {
		return false
	}
	favorite := model.Favorite{}
	err := _db.Where("user_id = ? AND video_id = ?", userId, videoId).First(&favorite).Error
	if err != nil {
		return false
	}
	return true
}

func (dao *FavoriteDao) CreateFavorite(user_id int64, video_id int64) error {
	new_favorite := model.Favorite{User_id: user_id, Video_id: video_id}
	err := _db.Create(&new_favorite).Error
	return err
}

func (dao *FavoriteDao) DeleteFavorite(user_id int64, video_id int64) error {
	favorite := model.Favorite{User_id: user_id, Video_id: video_id}
	err := _db.Delete(&favorite).Error
	return err
}

func (dao *FavoriteDao) SearchFavoriteList(user_id int64) (favorite []model.Favorite, err error) {
	err = _db.Where("user_id = ?", user_id).Find(&favorite).Error
	return
}

func (dao *FavoriteDao) BeginTransaction() *gorm.DB {
	return _db.Begin()
}
