package initialization

import (
	"douyin/global"
	"fmt"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func InitializeRedis() {
	redisCfg := global.SERVER_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.SERVER_LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		fmt.Println("====4-redis====: redis init success")
		global.SERVER_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.SERVER_REDIS = client
	}
}
