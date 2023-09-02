package dao

import (
	"douyin/global"
	"douyin/model"
)

func InsertVideo(newVideo *model.Video) error {
	err := global.SERVER_DB.Create(newVideo).Error
	return err
}

func GetVideoByAuthorId(authorId int64) ([]model.Video, error) {
	var videos []model.Video
	err := global.SERVER_DB.Where("author_id = ?", authorId).Find(&videos).Error
	return videos, err
}

func GetVideoByTime(latestTime int64) ([]model.Video, error) {
	var videos []model.Video
	var err error
	// if latestTime == 0 {
	err = global.SERVER_DB.Order("publish_time desc").Limit(5).Find(&videos).Error
	// } else {
	// 	err = global.SERVER_DB.Where("publish_time > ?", latestTime).Order("publish_time asc").Limit(5).Find(&videos).Error
	// }
	return videos, err
}

func GetVideoByAuthorIdandTitle(authorId int64, title string) (model.Video, error) {
	var video model.Video
	err := global.SERVER_DB.Where("author_id = ? and title = ?", authorId, title).First(&video).Error
	return video, err
}
