package dao

import (
	"context"
	"douyin/app/video/internal/dal/model"
	"gorm.io/gorm"
)

type VideoDao struct {
	*gorm.DB
}

func NewVideoDao(ctx context.Context) *VideoDao {
	return &VideoDao{NewDBClient(ctx)}
}

func (dao *VideoDao) InsertVideo(newVideo *model.Video) error {
	err := _db.Create(newVideo).Error
	return err
}

func (dao *VideoDao) GetVideoByAuthorId(authorId int64) ([]model.Video, error) {
	var videos []model.Video
	err := _db.Where("author_id = ?", authorId).Find(&videos).Error
	return videos, err
}

func (dao *VideoDao) GetVideoByTime(latestTime int64) ([]model.Video, error) {
	var videos []model.Video
	var err error
	// if latestTime == 0 {
	err = _db.Order("publish_time asc").Limit(5).Find(&videos).Error
	// } else {
	// 	err = global.SERVER_DB.Where("publish_time > ?", latestTime).Order("publish_time asc").Limit(5).Find(&videos).Error
	// }
	return videos, err
}

func (dao *VideoDao) GetVideoByAuthorIdandTitle(authorId int64, title string) (model.Video, error) {
	var video model.Video
	err := _db.Where("author_id = ? and title = ?", authorId, title).First(&video).Error
	return video, err
}
