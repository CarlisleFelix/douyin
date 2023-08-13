package global

import (
	"douyin/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	SERVER_VIPER  *viper.Viper
	SERVER_CONFIG config.Configuration
	SERVER_LOG    *zap.Logger
	SERVER_DB     *gorm.DB
)
