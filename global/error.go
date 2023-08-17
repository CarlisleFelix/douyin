package global

import "errors"

var (
	ErrorUserNameNull   = errors.New("用户名为空")
	ErrorUserNameExtend = errors.New("用户名长度不符合规范")

	ErrorPasswordNull   = errors.New("密码为空")
	ErrorPasswordLength = errors.New("密码长度不符合规范")
	ErrorPasswordFalse  = "密码错误"

	ErrorUserExist    = ("用户已存在")
	ErrorUserNotExist = "用户不存在" //notice：这里改了

	ErrorParamFormatWrong = errors.New("参数格式错误")

	ErrorTokenCreatedWrong = "token生成失败" //！notice：增加
	ErrorDatabase          = "无法录入数据库"   //！notice：增加
	ErrorTokenExpired      = "token验证失败"
)
