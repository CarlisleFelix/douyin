package dao

import (
	"context"
	"douyin/app/user/internal/dal/model"
	"fmt"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func (dao *UserDao) GetUserById(userId int64) (model.User, error) {
	user := model.User{}
	fmt.Println("userId:", userId)
	err := _db.Where("id = ?", userId).First(&user).Error
	fmt.Println("GetUserById错误", err)
	return user, err
}

// 存在疑问
func (dao *UserDao) GetUserByName(userName string) (model.User, error) {
	user := model.User{}
	err := _db.Where("name = ?", userName).First(&user).Error
	return user, err
}

func (dao *UserDao) InsertUser(newUser *model.User) error {
	err := _db.Create(newUser).Error
	return err
}

// UpdateUserFavoriteCount 更新用户点赞数
func (dao *UserDao) UpdateUserFavoriteCount(user_id int64, action_type int32) error {
	var user model.User
	err := _db.Where("user_id = ?", user_id).First(&user).Error
	if err != nil {
		return err
	}

	if action_type == 1 {
		user.Favorite_count += 1
	} else if action_type == 2 {
		user.Favorite_count -= 1
	}

	err = _db.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateUserTotalFavorite 更新作者获赞数
func (dao *UserDao) UpdateUserTotalFavorite(user_id int64, action_type int32) error {
	var user model.User
	err := _db.Where("user_id = ?", user_id).First(&user).Error
	if err != nil {
		return err
	}

	if action_type == 1 {
		user.Total_favorited += 1
	} else if action_type == 2 {
		user.Total_favorited -= 1
	}

	err = _db.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao *UserDao) UpdateUserFollowCount(user_id int64, action_type int32) error {
	var user model.User
	err := _db.Where("user_id = ?", user_id).First(&user).Error
	if err != nil {
		return err
	}
	if action_type == 1 {
		user.Follow_count += 1
	} else if action_type == 2 {
		user.Follower_count -= 1
	}
	err = _db.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}
func (dao *UserDao) UpdateUserFollowerCount(user_id int64, action_type int32) error {
	var user model.User
	err := _db.Where("user_id = ?", user_id).First(&user).Error
	if err != nil {
		return err
	}
	if action_type == 1 {
		user.Follower_count += 1
	} else if action_type == 2 {
		user.Follower_count -= 1
	}
	err = _db.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}
