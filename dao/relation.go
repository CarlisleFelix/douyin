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

// /////////////////////////关注/////////////////////////////
var relations = "relations"

// IncreaseFollowCount 增加UserID的关注数（UserID 的 follow_count+1）
func IncreaseFollowCount(UserID int64) error {
	if err := global.SERVER_DB.Model(&model.User{}).
		Where("id=?", UserID).
		Update("follow_count", gorm.Expr("follow_count+?", 1)).Error; err != nil {
		return err
	}
	return nil
}

// DecreaseFollowCount 减少UserID的关注数（UserID 的 follow_count-1）
func DecreaseFollowCount(UserID int64) error {
	if err := global.SERVER_DB.Model(&model.User{}).
		Where("id=?", UserID).
		Update("follow_count", gorm.Expr("follow_count-?", 1)).Error; err != nil {
		return err
	}
	return nil
}

// CreateFollowing 创建关注
func CreateFollowing(Host_id int64, Guest_id int64) error {

	//1.Following数据模型准备
	newFollowing := model.Relation{
		Host_id:  Host_id,
		Guest_id: Guest_id,
	}

	//2.新建following
	if err := global.SERVER_DB.Model(&model.Relation{}).Create(&newFollowing).Error; err != nil {
		return err
	}
	return nil
}

// DeleteFollowing 删除关注
func DeleteFollowing(Host_id int64, Guest_id int64) error {
	//1.Following数据模型准备
	deleteFollowing := model.Relation{
		Host_id:  Host_id,
		Guest_id: Guest_id,
	}

	//2.删除following
	if err := global.SERVER_DB.Model(&model.Relation{}).Where("host_id=? AND guest_id=?", Host_id, Guest_id).Delete(&deleteFollowing).Error; err != nil {
		return err
	}

	return nil
}

// FollowingList 获取关注表
func FollowingList(UserID int64) ([]model.User, error) {
	//1.userList数据模型准备
	var userList []model.User
	//2.查HostID的关注表
	err := global.SERVER_DB.Model(&model.User{}).
		Joins("left join "+relations+" on "+users+".id = "+relations+".guest_id").
		Where(relations+".host_id=?", UserID).
		Scan(&userList).Error
	return userList, err
}

// ///////////粉丝/////////////////////
// 粉丝表
var followers = "relations"

// 用户表
var users = "users"

func IsFollower(Host_id int64, Guest_id int64) (err error) {
	var relationExist = &model.Relation{}
	//判断关注是否存在
	err = global.SERVER_DB.Model(&model.Relation{}).
		Where("host_id=? AND guest_id=?", Host_id, Guest_id).
		First(&relationExist).Error
	return
}

// FollowerList  获取粉丝表
func FollowerList(UserID int64) ([]model.User, error) {
	//1.userList数据模型准备
	var userList []model.User
	//2.查UserID的粉丝表
	err := global.SERVER_DB.Model(&model.User{}).
		Joins("left join "+followers+" on "+users+".id = "+followers+".host_id").
		Where(followers+".guest_id=? ", UserID).
		Scan(&userList).Error
	return userList, err
}

// IncreaseFollowerCount 增加UserID的粉丝数（UserID 的 follower_count+1）
func IncreaseFollowerCount(UserID int64) error {
	if err := global.SERVER_DB.Model(&model.User{}).
		Where("id=?", UserID).
		Update("follower_count", gorm.Expr("follower_count+?", 1)).Error; err != nil {
		return err
	}
	return nil
}

// DecreaseFollowerCount 减少UserID的粉丝数（Host_id 的 follow_count-1）
func DecreaseFollowerCount(UserID int64) error {
	if err := global.SERVER_DB.Model(&model.User{}).
		Where("id=?", UserID).
		Update("follower_count", gorm.Expr("follower_count-?", 1)).Error; err != nil {
		return err
	}
	return nil
}

// 获取好友表，即互相都是对方粉丝
func FriendList(userID int64) ([]model.User, error) {
	var userList []model.User
	//查粉丝表
	err := global.SERVER_DB.Model(&model.User{}).
		Where("users.ID IN (SELECT a.host_id FROM followers a JOIN followers b ON a.host_id  = b.guest_id AND a.guest_id = b.host_id  AND a.guest_id = ? )", userID).
		Scan(&userList).Error
	return userList, err
}
