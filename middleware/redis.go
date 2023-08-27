package middleware

import (
	"context"
	"douyin/global"
	"douyin/model"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var favorite = "favorite"
var relation = "relation"
var ctx = context.Background()

// 更新视频点赞缓存 如果有favarite:userid:videoid就是有缓存 没有就是没有缓存 更新出错err
func SetVideoFavoriteState(userId int64, videoId int64, state bool) error {
	//准备键值
	key := fmt.Sprintf("%s:%d:%d", favorite, userId, videoId)
	value, err := json.Marshal(state)
	if err != nil {
		return global.ErrorRedisOperationFail
	}
	//设置键值
	err = global.SERVER_REDIS.Set(ctx, key, value, 30*time.Second).Err()
	if err != nil {
		return global.ErrorRedisOperationFail
	}
	return nil
}

// 获得视频点赞缓存 cachemiss没有记录 operationfail操作失败
func GetVideoFavoriteState(userId int64, videoId int64) (bool, error) {
	key := fmt.Sprintf("%s:%d:%d", favorite, userId, videoId)
	val, err := global.SERVER_REDIS.Get(ctx, key).Result()
	//若没有这个记录
	if err == redis.Nil {
		return false, global.ErrorCacheMiss
		//如果出错了
	} else if err != nil {
		return false, global.ErrorRedisOperationFail
	}
	var state bool
	err = json.Unmarshal([]byte(val), &state)
	if err != nil {
		return false, global.ErrorRedisOperationFail
	}
	return state, nil
}

// 更新关系缓存，如果有relation:userid1:userid2就是有缓存 没有就是没有缓存 更新出错err
func SetUserRelation(userId1 int64, userId2 int64, state bool) error {
	// 准备键值
	key := fmt.Sprintf("%s:%d:%d", relation, userId1, userId2)
	value, err := json.Marshal(state)
	if err != nil {
		//fmt.Printf("marshal fail!\n")
		return global.ErrorRedisOperationFail
	}
	// 设置键值
	err = global.SERVER_REDIS.Set(ctx, key, value, 30*time.Second).Err()
	if err != nil {
		//fmt.Printf("redis set fail!\n")
		return global.ErrorRedisOperationFail
	}
	//fmt.Printf("redis set success!\n")
	return nil
}

// cachemiss没有记录 operationfail操作失败
func GetUserRelationState(userId1 int64, userId2 int64) (bool, error) {
	key := fmt.Sprintf("%s:%d:%d", relation, userId1, userId2)
	val, err := global.SERVER_REDIS.Get(ctx, key).Result()
	//若没有这个记录
	if err == redis.Nil {
		//fmt.Printf("redis no record!\n")
		return false, global.ErrorCacheMiss
		//如果出错了
	} else if err != nil {
		//fmt.Printf("redis problem!\n")
		return false, global.ErrorRedisOperationFail
	}
	//fmt.Printf("redis hit!\n")
	var state bool
	err = json.Unmarshal([]byte(val), &state)
	if err != nil {
		return false, global.ErrorRedisOperationFail
	}
	return state, nil
}

//temporarily postponed

// 更新用户
func SetUser(userId int64, user model.User) error {
	//准备键值
	key := fmt.Sprintf("%s:%d", "userinfo", userId)
	value, err := json.Marshal(user)
	if err != nil {
		return global.ErrorRedisOperationFail
	}
	//设置键值
	err = global.SERVER_REDIS.Set(ctx, key, value, 30*time.Second).Err()
	if err != nil {
		return global.ErrorRedisOperationFail
	}
	return nil
}

func GetUser(userId int64) (model.User, error) {
	key := fmt.Sprintf("%s:%d", "user", userId)
	val, err := global.SERVER_REDIS.Get(ctx, key).Result()
	//若没有这个记录
	if err == redis.Nil {
		return model.User{}, global.ErrorCacheMiss
		//如果出错了
	} else if err != nil {
		return model.User{}, global.ErrorRedisOperationFail
	}
	var user model.User
	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		return model.User{}, global.ErrorRedisOperationFail
	}
	return user, nil
}

func SetVideo(videoId int64, video model.Video) error {
	// 准备键值
	key := fmt.Sprintf("%s:%d", "video", videoId)
	value, err := json.Marshal(video)
	if err != nil {
		return global.ErrorRedisOperationFail
	}
	// 设置键值
	err = global.SERVER_REDIS.Set(ctx, key, value, 30*time.Second).Err()
	if err != nil {
		return global.ErrorRedisOperationFail
	}
	return nil
}

func GetVideo(videoId int64) (model.Video, error) {
	key := fmt.Sprintf("%s:%d", "video", videoId)
	val, err := global.SERVER_REDIS.Get(ctx, key).Result()
	//若没有这个记录
	if err == redis.Nil {
		return model.Video{}, global.ErrorCacheMiss
		//如果出错了
	} else if err != nil {
		return model.Video{}, global.ErrorRedisOperationFail
	}
	var video model.Video
	err = json.Unmarshal([]byte(val), &video)
	if err != nil {
		return model.Video{}, global.ErrorRedisOperationFail
	}
	return video, nil
}
