package dao

import (
	"context"
	"douyin/app/comment/internal/dal/model"
	"gorm.io/gorm"
)

type CommentDao struct {
	*gorm.DB
}

func NewCommentDao(ctx context.Context) *CommentDao {
	return &CommentDao{NewDBClient(ctx)}
}

// GetCommentByIdListById 根据video_id返回视频评论列表
func (dao *CommentDao) GetCommentByIdListById(videoID int64) ([]model.Comment, error) {
	var comments []model.Comment
	err := _db.Where("video_id = ?", videoID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// GetCommentById 通过commentID 返回comment结构体
func (dao *CommentDao) GetCommentById(commentID int64) (model.Comment, error) {
	var comment model.Comment
	err := _db.Where("id = ?", commentID).First(&comment).Error
	if err != nil {
		return model.Comment{}, err
	}
	return comment, nil
}

// CreateComment 创建评论
func (dao *CommentDao) CreateComment(comment *model.Comment) error {
	err := _db.Create(comment).Error
	return err
}

// DeleteCommentById 根据id删除评论
func (dao *CommentDao) DeleteCommentById(commentID int64) error {
	err := _db.Where("id = ?", commentID).Delete(model.Comment{}).Error
	return err
}

//// todo: IsFollow 查询该用户是否被关注
//func IsFollow(hostID int64, guestID int64) (bool, error) {
//	var relation model.Relation
//	// 判断用户id是否相等
//	if hostID == guestID {
//		return false, nil
//	}
//	// 查询是否存在关系记录
//	err := global.SERVER_DB.Where("host_id = ? AND guest_id = ?", hostID, guestID).First(&relation).Error
//	// 如果查询不到被关注记录，则将userResponse.IsFollow置为false
//	if errors.Is(err, gorm.ErrRecordNotFound) {
//		return false, nil
//	} else if err != nil {
//		return false, err
//	}
//	return true, nil
//}

// BeginTransaction 开始事务
func BeginTransaction() *gorm.DB {
	return _db.Begin()
}

// RollbackTransaction 回滚事务
func RollbackTransaction(tx *gorm.DB) {
	tx.Rollback()
}
