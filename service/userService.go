package service

import "douyin/global"

const (
	MaxUsernameLength = 32 //用户名最大长度
	MaxPasswordLength = 32 //密码最大长度
	MinPasswordLength = 6  //密码最小长度
)

func UserRegisterService(userName string, passWord string) {
	err := CheckUserParam(userName, passWord)
	if err != nil {
		return
	}
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
