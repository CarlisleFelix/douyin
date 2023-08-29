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

// GetVideoById
func (dao *VideoDao) GetVideoById(video_id int64) (model.Video, error) {
	video := model.Video{}
	err := _db.Where("id = ?", video_id).First(&video).Error
	return video, err
}

// UpdateVideoFavoriteCount 更新视频点赞数
func (dao *VideoDao) UpdateVideoFavoriteCount(video_id int64, action_type int32) error {
	var video model.Video
	err := _db.Where("video_id = ?", video_id).First(&video).Error
	if err != nil {
		return err
	}

	if action_type == 1 {
		video.Favorite_count += 1
	} else if action_type == 2 {
		video.Favorite_count -= 1
	}

	err = _db.Save(&video).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao *VideoDao) UpdateVideoCommentCount(videoID int64, action_type int32) error {
	// 查询视频数据
	var video model.Video
	err := _db.First(&video, videoID).Error
	if err != nil {
		return err
	}

	// 更新评论总数字段
	if action_type == 1 {
		video.Comment_count += 1
	} else if action_type == 2 {
		video.Comment_count -= 1
	}

	// 保存更新后的视频数据
	err = _db.Save(&video).Error
	if err != nil {
		return err
	}
	return nil
}
