package dao

import (
	"douyin/global"
	"douyin/model"
)

// GetCommentByIdListById 根据video_id返回视频评论列表
func GetCommentByIdListById(videoID int64) ([]model.Comment, error) {
	var comments []model.Comment
	err := global.SERVER_DB.Where("video_id = ?", videoID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// GetCommentById 通过commentID 返回comment结构体
func GetCommentById(commentID int64) (model.Comment, error) {
	var comment model.Comment
	err := global.SERVER_DB.Where("id = ?", commentID).First(&comment).Error
	if err != nil {
		return model.Comment{}, err
	}
	return comment, nil
}

// CreateComment 创建评论
func CreateComment(comment *model.Comment) error {
	err := global.SERVER_DB.Create(comment).Error
	return err
}

// DeleteCommentById 根据id删除评论
func DeleteCommentById(commentID int64) error {
	err := global.SERVER_DB.Where("id = ?", commentID).Delete(model.Comment{}).Error
	//使用mq进行数据库中评论的删除-评论状态更新
	//评论id传入消息队列
	// rabbitmq.RmqCommentDel.Publish(strconv.FormatInt(commentID, 10))
	return err
}

// UpdateVideoCommentCount 根据视频ID更新视频表的评论总数字段
func UpdateVideoCommentCount(videoID int64, operand int64) error {
	// 查询视频数据
	var video model.Video
	err := global.SERVER_DB.First(&video, videoID).Error
	if err != nil {
		return err
	}

	// 更新评论总数字段
	video.Comment_count += operand

	// 保存更新后的视频数据
	err = global.SERVER_DB.Save(&video).Error
	if err != nil {
		return err
	}

	return nil
}
