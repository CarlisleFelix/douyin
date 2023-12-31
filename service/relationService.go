package service

import (
	"douyin/dao"
	"douyin/global"
	"douyin/middleware/rabbitmq"
	"douyin/model"
	"errors"
	"log"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

// /关注////////
// IsFollowing check if HostID follows GuestID
func IsFollowing(UserID int64, ToUserID int64) bool {
	if UserID == ToUserID {
		return false
	}
	return dao.GetFollowByUserId(UserID, ToUserID)
}

// FollowingList 获取关注表
func FollowingList(UserID int64) ([]model.User, error) {
	//2.查HostID的关注表
	userList, err := dao.FollowingList(UserID)
	if err != nil {
		return userList, err
	}
	return userList, nil
}

// FollowAction 关注操作
func FollowAction(UserID int64, ToUserID int64, actionType int64) error {
	//创建关注操作
	if actionType == 1 {
		//判断关注是否存在
		if dao.GetFollowByUserId(UserID, ToUserID) {
			//关注存在
			return global.ErrorRelationExist
		} else {
			//关注不存在,创建关注(启用事务Transaction)
			err1 := global.SERVER_DB.Transaction(func(db *gorm.DB) error {
				// 加信息打入消息队列。
				sb := strings.Builder{}
				sb.WriteString(strconv.Itoa(int(UserID)))
				sb.WriteString(" ")
				sb.WriteString(strconv.Itoa(int(ToUserID)))
				rabbitmq.RmqFollowAdd.Publish(sb.String())
				// 记录日志
				log.Println("消息打入成功。")
				err := dao.CreateFollowing(UserID, ToUserID)
				if err != nil {
					return err
				}
				//增加UserID的关注数
				err = dao.IncreaseFollowCount(UserID)
				if err != nil {
					return err
				}
				//增加ToUserID的粉丝数
				err = dao.IncreaseFollowerCount(ToUserID)
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
		if dao.GetFollowByUserId(UserID, ToUserID) {
			//关注存在,删除关注(启用事务Transaction)
			if err1 := global.SERVER_DB.Transaction(func(db *gorm.DB) error {
				// 加信息打入消息队列。
				sb := strings.Builder{}
				sb.WriteString(strconv.Itoa(int(UserID)))
				sb.WriteString(" ")
				sb.WriteString(strconv.Itoa(int(ToUserID)))
				rabbitmq.RmqFollowDel.Publish(sb.String())
				// 记录日志
				log.Println("消息打入成功。")
				err := dao.DeleteFollowing(UserID, ToUserID)
				if err != nil {
					return err
				}
				//减少UserID的关注数
				err = dao.DecreaseFollowCount(UserID)
				if err != nil {
					return err
				}
				//减少ToUserID的粉丝数
				err = dao.DecreaseFollowerCount(ToUserID)
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
func IsFollower(UserID int64, ToUserID int64) bool {
	//2.查询粉丝表中粉丝是否存在
	err := dao.IsFollower(UserID, ToUserID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// follower not found
		return false
	}
	// follower found
	return true
}

// FollowerList  获取粉丝表
func FollowerList(UserID int64) ([]model.User, error) {
	//2.查UserID的关注表
	userList, err := dao.FollowerList(UserID)
	if err != nil {
		return userList, err
	}
	return userList, nil
}

// ///////好友////////////
// FriendList 获取朋友列表（互相关注）
func FriendList(UserID int64) ([]model.User, error) {
	var friendList []model.User
	// 查询 UserID 的关注列表
	// 检查 关注列表中的用户是否也关注 UserID
	followList, err := FollowingList(UserID)
	if err != nil {
		return nil, err
	} else {
		for _, user := range followList {
			if IsFollowing(user.User_id, UserID) {
				friendList = append(friendList, user)
			}
		}
		return friendList, nil
	}
}
