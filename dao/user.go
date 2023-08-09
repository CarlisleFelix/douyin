package dao

import (
	"douyin/global"
	"douyin/model"
)

func GetUserById(userId int64) (model.User, error) {
	user := model.User{}
	err := global.SERVER_DB.Where("id=?", userId).First(&user).Error
	return user, err
}

func InsertUser(newUser model.User) (model.User, error) {
	err := global.SERVER_DB.Create(&newUser).Error
	return newUser, err
}
