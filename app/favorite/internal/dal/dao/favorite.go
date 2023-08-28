package dao

import (
	"context"
	"douyin/app/favorite/internal/dal/model"
	"gorm.io/gorm"
)

type FavoriteDao struct {
	*gorm.DB
}

func NewFavoriteDao(ctx context.Context) *FavoriteDao {
	return &FavoriteDao{NewDBClient(ctx)}
}

// todo: 将GetifFavorite 改名为 IsFavorite
func (dao *FavoriteDao) GetifFavorite(userId int64, videoId int64) bool {
	if userId == 0 {
		return false
	}
	favorite := model.Favorite{}
	err := _db.Where("user_id = ? AND video_id = ?", userId, videoId).First(&favorite).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		} else {
			return false
		}
	}
	return true
}

// todo: 删除searchFavorite
// SearchFavorite 根据user_id，video_id查找favorite
func (dao *FavoriteDao) SearchFavorite(user_id int64, video_id int64) (model.Favorite, error) {
	favorite := model.Favorite{}
	err := _db.Where("user_id = ?", user_id).Where("video_id = ?", video_id).Take(&favorite).Error
	return favorite, err
}

func (dao *FavoriteDao) CreateFavorite(user_id int64, video_id int64) error {
	new_favorite := model.Favorite{User_id: user_id, Video_id: video_id}
	err := _db.Create(&new_favorite).Error
	return err
}

func (dao *FavoriteDao) DeleteFavorite(favorite model.Favorite) error {
	err := _db.Delete(&favorite).Error
	return err
}

func (dao *FavoriteDao) SearchFavoriteList(user_id int64) (favorite []model.Favorite, err error) {
	err = _db.Where("user_id = ?", user_id).Find(&favorite).Error
	return
}

// todo
//// SearchUser 根据id查找用户
//func SearchUser(user_id int64) (model.User, error) {
//	user := model.User{}
//	err := global.SERVER_DB.Where("id = ?", user_id).First(&user).Error
//	return user, err
//}

// todo
//// SearchVideo 根据id查找视频
//func SearchVideo(video_id int64) (model.Video, error) {
//	video := model.Video{}
//	err := _db.Where("id = ?", video_id).First(&video).Error
//	return video, err
//}

// todo
//// UpdateVideo 更新视频点赞数
//func UpdateVideo(video model.Video, action_type int32) error {
//	if action_type == 1 {
//		video.Favorite_count += 1
//	} else if action_type == 2 {
//		video.Favorite_count -= 1
//	}
//	err := global.SERVER_DB.Save(&video).Error
//	return err
//}

// todo
//// UpdateUser 更新用户点赞数
//func  UpdateUser(user model.User, action_type int32) error {
//	if action_type == 1 {
//		user.Favorite_count += 1
//	} else if action_type == 2 {
//		user.Favorite_count -= 1
//	}
//	err := global.SERVER_DB.Save(&user).Error
//	return err
//}

// todo
//// UpdateAuthor 更新作者获赞数
//func UpdateAuthor(author model.User, action_type int32) error {
//	if action_type == 1 {
//		author.Total_favorited += 1
//	} else if action_type == 2 {
//		author.Total_favorited -= 1
//	}
//	err := global.SERVER_DB.Save(&author).Error
//	return err
//}

// todo
//func SearchRelation(host_id int64, guest_id int64) error {
//	err := global.SERVER_DB.Where("Host_id = ?", host_id).Where("Guest_id = ?", guest_id).Error
//	return err
//}
