package dao

import (
	"douyin/global"
	"douyin/model"
	"gorm.io/gorm"
)

// SearchUser 根据id查找用户
func SearchUser(user_id int64) (user model.User, StatusCode int32, StatusMsg string) {
	result := global.SERVER_DB.Where("id = ?", user_id).First(&user)
	if result.Error != nil {
		StatusCode = 1
		StatusMsg = "用户查询异常"
	} else if result.RowsAffected == 0 {
		StatusCode = 1
		StatusMsg = "用户不存在"
	}
	return
}

// SearchFavorite 根据user_id，video_id查找favorite
func SearchFavorite(user_id int64, video_id int64) (favorite model.Favorite, result *gorm.DB) {
	result = global.SERVER_DB.Where("user_id = ?", user_id).Where("video_id = ?", video_id).Take(&favorite)
	return
}

// SearchVideo 根据id查找视频
func SearchVideo(video_id int64) (video model.Video, StatusCode int32, StatusMsg string) {
	result := global.SERVER_DB.Where("id = ?", video_id).First(&video)
	if result.Error != nil {
		StatusCode = 1
		StatusMsg = "视频查询异常"
	} else if result.RowsAffected == 0 {
		StatusCode = 1
		StatusMsg = "视频不存在"
	}
	return
}

// UpdateVideo 更新视频点赞数
func UpdateVideo(video model.Video, action_type int32) *gorm.DB {
	if action_type == 1 {
		video.Favorite_count += 1
	} else if action_type == 2 {
		video.Favorite_count -= 1
	}
	result := global.SERVER_DB.Save(&video)
	return result
}

// UpdateUser 更新用户点赞数
func UpdateUser(user model.User, action_type int32) *gorm.DB {
	if action_type == 1 {
		user.Favorite_count += 1
	} else if action_type == 2 {
		user.Favorite_count -= 1
	}
	result := global.SERVER_DB.Save(&user)
	return result
}

// UpdateAuthor 更新作者获赞数
func UpdateAuthor(author model.User, action_type int32) *gorm.DB {
	if action_type == 1 {
		author.Total_favorited += 1
	} else if action_type == 2 {
		author.Total_favorited -= 1
	}
	result := global.SERVER_DB.Save(&author)
	return result
}

func CreateFavorite(user_id int64, video_id int64) (StatusCode int32, StatusMsg string) {
	new_favorite := model.Favorite{User_id: user_id, Video_id: video_id}
	result := global.SERVER_DB.Create(&new_favorite)
	if result.Error != nil {
		StatusCode = 1
		StatusMsg = "点赞失败"
	} else {
		StatusCode = 0
		StatusMsg = "点赞成功"
	}
	return
}

func DeleteFavorite(favorite model.Favorite) (StatusCode int32, StatusMsg string) {
	result := global.SERVER_DB.Delete(&favorite)
	if result.Error != nil {
		StatusCode = 1
		StatusMsg = "取消点赞失败"
	} else {
		StatusCode = 0
		StatusMsg = "取消点赞成功"
	}
	return
}

func SearchFavoriteList(user_id int64) (favorite []model.Favorite, result *gorm.DB) {
	result = global.SERVER_DB.Where("user_id = ?", user_id).Find(&favorite)
	return
}

func SearchRelation(host_id int64, guest_id int64) (result *gorm.DB) {
	result = global.SERVER_DB.Where("Host_id = ?", host_id).Where("Guest_id = ?", guest_id)
	return
}
