package service

import (
	"douyin/dao"
	"douyin/global"
	"douyin/model"
	"errors"

	"gorm.io/gorm"
)

// /关注////////
// IsFollowing check if HostID follows GuestID
func IsFollowing(Host_id int64, Guest_id int64) bool {
	if Host_id == Guest_id {
		return false
	}
	return dao.GetFollowByUserId(Host_id, Guest_id)
}

// FollowingList 获取关注表
func FollowingList(Host_id int64) ([]model.User, error) {
	//2.查HostID的关注表
	userList, err := dao.FollowingList(Host_id)
	if err != nil {
		return userList, err
	}
	return userList, nil
}

// FollowAction 关注操作
func FollowAction(Host_id int64, Guest_id int64, actionType int64) error {
	//创建关注操作
	if actionType == 1 {
		//判断关注是否存在
		if dao.GetFollowByUserId(Host_id, Guest_id) {
			//关注存在
			return global.ErrorRelationExist
		} else {
			//关注不存在,创建关注(启用事务Transaction)
			err1 := global.SERVER_DB.Transaction(func(db *gorm.DB) error {
				err := dao.CreateFollowing(Host_id, Guest_id)
				if err != nil {
					return err
				}
				// err = dao.CreateFollower(Guest_id, Host_id)
				// if err != nil {
				// 	return err
				// }
				//增加host_id的关注数
				err = dao.IncreaseFollowCount(Host_id)
				if err != nil {
					return err
				}
				//增加guest_id的粉丝数
				err = dao.IncreaseFollowerCount(Guest_id)
				if err != nil {
					return err
				}
				return nil
			})
			if err1 != nil {
				return err1
			}
		}
	}
	if actionType == 2 {
		//判断关注是否存在
		if dao.GetFollowByUserId(Host_id, Guest_id) {
			//关注存在,删除关注(启用事务Transaction)
			if err1 := global.SERVER_DB.Transaction(func(db *gorm.DB) error {
				err := dao.DeleteFollowing(Host_id, Guest_id)
				if err != nil {
					return err
				}
				// err = dao.DeleteFollower(Host_id, Guest_id)
				// if err != nil {
				// 	return err
				// }
				//减少host_id的关注数
				err = dao.DecreaseFollowCount(Host_id)
				if err != nil {
					return err
				}
				//减少guest_id的粉丝数
				err = dao.DecreaseFollowerCount(Guest_id)
				if err != nil {
					return err
				}
				return nil
			}); err1 != nil {
				return err1
			}

		} else {
			//关注不存在
			return global.ErrorRelationNull
		}
	}
	return nil
}

// ///粉丝///////////

// IsFollower 判断HostID是否有GuestID这个粉丝
func IsFollower(Host_id int64, Guest_id int64) bool {
	//2.查询粉丝表中粉丝是否存在
	err := dao.IsFollower(Host_id, Guest_id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// follower not found
		return false
	}
	// follower found
	return true
}

// FollowerList  获取粉丝表
func FollowerList(Host_id int64) ([]model.User, error) {
	//2.查HostID的关注表
	userList, err := dao.FollowerList(Host_id)
	if err != nil {
		return userList, err
	}
	return userList, nil
}

// ///////好友////////////
// FriendList 获取朋友列表（互相关注）
func FriendList(Id int64) ([]model.User, error) {
	var friendList []model.User
	// 查询 Id 的关注列表
	// 检查 关注列表中的用户是否也关注 Id
	followList, err := FollowingList(Id)
	if err != nil {
		return nil, err
	} else {
		for _, user := range followList {
			if IsFollowing(user.User_id, Id) {
				friendList = append(friendList, user)
			}
		}
		return friendList, nil
	}
}
