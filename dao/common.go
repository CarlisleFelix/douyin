package dao

import (
	"douyin/global"
	"douyin/model"
	"errors"
	"gorm.io/gorm"
)

// IsFollow 查询该用户是否被关注
func IsFollow(hostID int64, guestID int64) (bool, error) {
	var relation model.Relation
	// 判断用户id是否相等
	if hostID == guestID {
		return false, nil
	}
	// 查询是否存在关系记录
	err := global.SERVER_DB.Where("host_id = ? AND guest_id = ?", hostID, guestID).First(&relation).Error
	// 如果查询不到被关注记录，则将userResponse.IsFollow置为false
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

// BeginTransaction 开始事务
func BeginTransaction() *gorm.DB {
	return global.SERVER_DB.Begin()
}

// RollbackTransaction 回滚事务
func RollbackTransaction(tx *gorm.DB) {
	tx.Rollback()
}
