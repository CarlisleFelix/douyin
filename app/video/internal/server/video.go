package server

import (
	"douyin/config"

	"github.com/tencentyun/cos-go-sdk-v5"
)

var (
	SERVER_CONFIG     config.Config
	SERVER_COS_VIDEO  *cos.Client
	SERVER_COS_COVER  *cos.Client
	SERVER_COS_AVATAR *cos.Client
)
