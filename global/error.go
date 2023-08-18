package global

import "errors"

var (
	ErrorUserNameNull   = errors.New("用户名为空")
	ErrorUserNameExtend = errors.New("用户名长度不符合规范")

	ErrorPasswordNull   = errors.New("密码为空")
	ErrorPasswordLength = errors.New("密码长度不符合规范")
	ErrorPasswordFalse  = errors.New("密码错误")

	ErrorUserExist    = errors.New("用户已存在")
	ErrorUserNotExist = errors.New("用户不存在")

	ErrorParamMismatch    = errors.New("参数不匹配")
	ErrorParamFormatWrong = errors.New("参数格式错误")

	ErrorVideoNotExist = errors.New("视频不存在")

	ErrorActionType = errors.New("未知操作")

	ErrorFavoriteExist    = errors.New("视频已点赞")
	ErrorFavoriteNotExist = errors.New("视频未点赞")
)
