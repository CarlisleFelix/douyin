package global

import (
	"douyin/config"
	"sync"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"github.com/tencentyun/cos-go-sdk-v5"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	SERVER_VIPER      *viper.Viper
	SERVER_CONFIG     config.Configuration
	SERVER_LOG        *zap.Logger
	SERVER_DB         *gorm.DB
	SERVER_COS_VIDEO  *cos.Client
	SERVER_COS_COVER  *cos.Client
	SERVER_COS_AVATAR *cos.Client
	SERVER_REDIS      *redis.Client
	GlobalConfig      config.Configuration
	Wg                sync.WaitGroup
)
