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
