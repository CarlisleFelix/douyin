package global

import (
	"context"
	"douyin/config"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"github.com/tencentyun/cos-go-sdk-v5"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	SERVER_VIPER           *viper.Viper
	SERVER_CONFIG          config.Configuration
	SERVER_LOG             *zap.Logger
	SERVER_DB              *gorm.DB
	SERVER_COS_VIDEO       *cos.Client
	SERVER_COS_COVER       *cos.Client
	SERVER_COS_AVATAR      *cos.Client
	SERVER_REDIS           *redis.Client
	SERVER_TRACE_PROVIDER  *tracesdk.TracerProvider
	SERVER_USER_TRACER     trace.Tracer
	SERVER_VIDEO_TRACER    trace.Tracer
	SERVER_RELATION_TRACER trace.Tracer
	SERVER_COMMENT_TRACER  trace.Tracer
	SERVER_FAVORITE_TRACER trace.Tracer
	SERVER_MESSAGE_TRACER  trace.Tracer
	SERVER_CONTEXT         *context.Context
)
