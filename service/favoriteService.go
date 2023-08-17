package service

import (
	"douyin/dao"
	"douyin/global"
	"douyin/model"
	"douyin/response"
	"strconv"
)

func FavoriteAction(user_id int64, video_id string, action_type int32) (StatusCode int32, StatusMsg string) {
	// 参数类型转换
	videoId, err := strconv.ParseInt(video_id, 10, 64)
	if err != nil {
		StatusCode = 1
		StatusMsg = "视频id异常"
		return
	}

	// 开始执行
	tx := global.SERVER_DB.Begin()

	// 获取点赞人信息
	var user model.User
	user, StatusCode, StatusMsg = dao.SearchUser(user_id)
	if StatusCode == 1 {
		tx.Rollback()
		return
	}

	// 获取视频信息
	var video model.Video
	video, StatusCode, StatusMsg = dao.SearchVideo(videoId)
	if StatusCode == 1 {
		tx.Rollback()
		return
	}

	// 获取作者信息
	var author model.User
	author, StatusCode, StatusMsg = dao.SearchUser(video.Author_id)
	if StatusCode == 1 {
		tx.Rollback()
		return
	}

	// 当用户对自己的视频点赞时，操作失败
	if video.Author_id == user_id {
		StatusCode = 1
		StatusMsg = "该视频为本人发布，无法进行点赞或取消点赞操作"
		tx.Rollback()
		return
	}

	// 根据user_id,video_id查找点赞信息
	favorite, result := dao.SearchFavorite(user_id, videoId)

	// 对action_type进行判断, 1-点赞， 2-取消点赞
	if action_type == 1 {
		// 如果是点赞操作
		// 先查询是否已经点赞
		if result.RowsAffected == 0 {
			// 如果未点赞，则将点赞信息存储到表中
			// 视频点赞数+1
			result = dao.UpdateVideo(video, action_type)
			if result.Error != nil {
				StatusCode = 1
				StatusMsg = "视频点赞更新异常"
				tx.Rollback()
				return
			}
			// 用户点赞数+1
			result = dao.UpdateUser(user, action_type)
			if result.Error != nil {
				StatusCode = 1
				StatusMsg = "用户点赞更新异常"
				tx.Rollback()
				return
			}
			// 作者获赞数+1
			result = dao.UpdateAuthor(author, action_type)
			if result.Error != nil {
				StatusCode = 1
				StatusMsg = "作者获赞更新异常"
				tx.Rollback()
				return
			}

			// 创建点赞记录
			StatusCode, StatusMsg = dao.CreateFavorite(user_id, videoId)
			// 根据状态码判断点赞操作是否成功
			if StatusCode == 1 {
				tx.Rollback()
			} else {
				tx.Commit()
			}
			return
		} else {
			// 如果已点赞
			StatusCode = 1
			StatusMsg = "该视频已点赞"
			tx.Rollback()
			return
		}
	} else {
		// 如果是取消点赞操作，先查询表中是否存在点赞记录，
		if result.RowsAffected == 0 {
			// 如果不存在，则返回异常
			StatusCode = 1
			StatusMsg = "未对该视频点赞，无法取消点赞"
			tx.Rollback()
			return
		} else {
			// 如果存在，则删除该记录
			// 视频点赞数-1
			result = dao.UpdateVideo(video, action_type)
			if result.Error != nil {
				StatusCode = 1
				StatusMsg = "视频点赞更新异常"
				tx.Rollback()
				return
			}
			// 用户点赞数-1
			result = dao.UpdateUser(user, action_type)
			if result.Error != nil {
				StatusCode = 1
				StatusMsg = "用户点赞更新异常"
				tx.Rollback()
				return
			}
			// 作者获赞数-1
			result = dao.UpdateAuthor(author, action_type)
			if result.Error != nil {
				StatusCode = 1
				StatusMsg = "作者获赞更新异常"
				tx.Rollback()
				return
			}

			// 删除点赞记录
			StatusCode, StatusMsg = dao.DeleteFavorite(favorite)
			// 根据状态码判断操作是否成功
			if StatusCode == 1 {
				tx.Rollback()
			} else {
				tx.Commit()
			}
			return
		}
	}
}

func FavoriteList(user_id int64) (StatusCode int32, StatusMsg string, videoList []response.Video_Response) {
	tx := global.SERVER_DB.Begin()
	favorites, result := dao.SearchFavoriteList(user_id)
	if result.Error != nil {
		StatusCode = 1
		StatusMsg = "获取用户喜欢列表异常"
		tx.Rollback()
		return StatusCode, StatusMsg, nil
	} else if result.RowsAffected == 0 {
		StatusCode = 0
		StatusMsg = "用户无喜欢视频"
		tx.Commit()
		return StatusCode, StatusMsg, nil
	}

	// 创建视频列表
	videoList = make([]response.Video_Response, len(favorites))

	//
	for i := 0; i < len(favorites); i++ {
		// 获取视频信息
		var video model.Video
		video, StatusCode, StatusMsg = dao.SearchVideo(favorites[i].Video_id)
		if StatusCode == 1 {
			tx.Rollback()
			return StatusCode, StatusMsg, nil
		}
		// 获取用户信息
		var user model.User
		user, StatusCode, StatusMsg = dao.SearchUser(video.Author_id)
		if StatusCode == 1 {
			tx.Rollback()
			return StatusCode, StatusMsg, nil
		}
		// 获取关注信息
		result = dao.SearchRelation(user_id, user.User_id)
		if result.Error != nil {
			StatusCode = 1
			StatusMsg = "获取关注信息异常"
			tx.Rollback()
			return StatusCode, StatusMsg, nil
		}

		var IsFollow bool
		// 设置关注信息
		if result.RowsAffected == 0 {
			IsFollow = false
		} else {
			IsFollow = true
		}

		videoList[i] = response.Video_Response{
			Id: video.Video_id,
			// 存储作者信息
			Author: response.User_Response{
				Id:              user.User_id,
				Name:            user.User_name,
				FollowCount:     user.Follow_count,
				FollowerCount:   user.Follower_count,
				IsFollow:        IsFollow,
				Avatar:          user.Avatar,
				BackgroundImage: user.Background_image,
				Signature:       user.Signature,
				TotalFavorited:  user.Total_favorited,
				WorkCount:       user.Work_count,
				FavoriteCount:   user.Favorite_count,
			},
			PlayUrl:       video.Play_url,
			CoverUrl:      video.Cover_url,
			FavoriteCount: video.Favorite_count,
			CommentCount:  video.Comment_count,
			IsFavorite:    true,
			Title:         video.Title,
		}
	}

	tx.Commit()
	return StatusCode, StatusMsg, videoList
}
