package service

import (
	"douyin/global"
	"douyin/model"
	"douyin/utils"

	"gorm.io/gorm"
)

const (
	MaxUsernameLength = 32 //用户名最大长度
	MaxPasswordLength = 32 //密码最大长度
	MinPasswordLength = 6  //密码最小长度
)

func UserRegisterService(userName string, passWord string) (model.User, error) {
	//准备参数
	hashedPwd, err := utils.Hash(passWord)
	newUser := model.User{
		User_name: userName,
		Password:  hashedPwd,
	}

	//检查参数
	err = CheckUserParam(userName, passWord)
	if err != nil {
		return newUser, err
	}

	//检查参数
	exist := UserNameExists(userName)
	if exist == true {
		return newUser, global.ErrorUserExist
	}

	//创建用户
	if err = global.SERVER_DB.Create(&newUser).Error; err != nil {
		global.SERVER_LOG.Info("error of database while create user!")
		return newUser, err
	}

	//返回
	return newUser, err
}

func CheckUserParam(userName string, passWord string) error {
	if userName == "" {
		return global.ErrorUserNameNull
	}
	if len(userName) > MaxUsernameLength {
		return global.ErrorUserNameExtend
	}
	if len(passWord) > MaxPasswordLength || len(passWord) < MinPasswordLength {
		return global.ErrorPasswordLength
	}
	return nil
}

func UserNameExists(userName string) bool {
	user := model.User{}
	err := global.SERVER_DB.Where("name = ?", userName).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		} else {
			return true
		}
	}
	return true
}

func UserLoginService(userName string, passWord string) (model.User, error) {
	loginUser := model.User{}

	//检查参数
	err := CheckUserParam(userName, passWord)
	if err != nil {
		return loginUser, err
	}

	//检查用户是否存在
	exist := UserNameExists(userName)
	if exist == false {
		return loginUser, global.ErrorUserNotExist
	}

	//检查密码是否一致
	passwordMatch := CheckUserPassword(userName, passWord, &loginUser)
	if !passwordMatch {
		return loginUser, global.ErrorPasswordFalse
	}

	//返回
	return loginUser, err
}

func CheckUserPassword(userName string, passWord string, loginUser *model.User) bool {
	global.SERVER_DB.Where("name=?", userName).First(loginUser)
	if !utils.Compare(loginUser.Password, passWord) {
		return false
	}
	return true
}
