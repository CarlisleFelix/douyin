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

	ErrorParamFormatWrong = errors.New("参数格式错误")

	ErrorDbOperation = errors.New("数据库操作出错")

	ErrorVideoDataWrong     = errors.New("视频数据错误")
	ErrorVideoDownloading   = errors.New("视频数据下载错误")
	ErrorFileOperationWrong = errors.New("文件操作错误")

	ErrorTokenincorrect = errors.New("token错误")

	ErrorVideoDuplicate = errors.New("视频重复")
)
