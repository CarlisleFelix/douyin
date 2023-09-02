package service

import (
	"context"
	"douyin/dao"
	"douyin/global"
	"douyin/model"
	"douyin/response"
	"strconv"

	"gorm.io/gorm"
)

func FavoriteAction(user_id int64, video_id string, action_type int32) error {

	// 参数类型转换
	videoId, err := strconv.ParseInt(video_id, 10, 64)
	if err != nil {
		return global.ErrorParamFormatWrong
	}

	// 获取点赞人信息
	var user model.User
	user, err = dao.SearchUser(user_id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return global.ErrorUserNotExist
		}
	}

	// 获取视频信息
	var video model.Video
	video, err = dao.SearchVideo(videoId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return global.ErrorVideoNotExist
		}
	}

	// 获取作者信息
	var author model.User
	author, err = dao.SearchUser(video.Author_id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return global.ErrorUserNotExist
		}

	}

	// 根据user_id,video_id查找点赞信息
	favorite, err := dao.SearchFavorite(user_id, videoId)

	// 开始执行
	tx := global.SERVER_DB.Begin()

	// 对action_type进行判断, 1-点赞， 2-取消点赞
	if action_type == 1 {
		// 如果是点赞操作
		// 先查询是否已经点赞
		if err == gorm.ErrRecordNotFound {
			// 如果未点赞，则将点赞信息存储到表中
			// 视频点赞数+1
			err = dao.UpdateVideo(video, action_type)
			if err != nil {
				tx.Rollback()
				return err
			}
			// 用户点赞数+1
			err = dao.UpdateUser(user, action_type)
			if err != nil {
				tx.Rollback()
				return err
			}
			// 作者获赞数+1
			err = dao.UpdateAuthor(author, action_type)
			if err != nil {
				tx.Rollback()
				return err
			}

			// 创建点赞记录
			err = dao.CreateFavorite(user_id, videoId)
			// 根据状态码判断点赞操作是否成功
			if err != nil {
				tx.Rollback()
				return err
			} else {
				tx.Commit()
				return nil
			}
		} else {
			// 如果已点赞
			tx.Rollback()
			return global.ErrorFavoriteExist
		}
	} else {
		// 如果是取消点赞操作，先查询表中是否存在点赞记录，
		if err == gorm.ErrRecordNotFound {
			// 如果不存在，则返回异常
			tx.Rollback()
			return global.ErrorFavoriteNotExist
		} else {
			// 如果存在，则删除该记录
			// 视频点赞数-1
			err = dao.UpdateVideo(video, action_type)
			if err != nil {
				tx.Rollback()
				return err
			}
			// 用户点赞数-1
			err = dao.UpdateUser(user, action_type)
			if err != nil {
				tx.Rollback()
				return err
			}
			// 作者获赞数-1
			err = dao.UpdateAuthor(author, action_type)
			if err != nil {
				tx.Rollback()
				return err
			}

			// 删除点赞记录
			err = dao.DeleteFavorite(favorite)
			// 根据状态码判断操作是否成功
			if err != nil {
				tx.Rollback()
				return err
			} else {
				tx.Commit()
				return nil
			}
		}
	}
}

func FavoriteList(user_id int64, ctx context.Context) (videoList []response.Video_Response, err error) {

	ctx, span := global.SERVER_FAVORITE_TRACER.Start(ctx, "favoritelist service")
	defer span.End()

	favorites, err := dao.SearchFavoriteList(user_id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 用户无喜欢视频
			return nil, nil
		} else {
			return nil, err
		}
	}

	// 创建视频列表
	videoList = make([]response.Video_Response, len(favorites))

	for i := 0; i < len(favorites); i++ {
		// 获取视频信息
		var video model.Video
		video, err = dao.SearchVideo(favorites[i].Video_id)
		if err != nil {
			return nil, err
		}
		// 获取用户信息
		var user model.User
		user, err = dao.SearchUser(video.Author_id)
		if err != nil {
			return nil, err
		}
		// 获取关注信息
		err = dao.SearchRelation(user_id, user.User_id)
		if err != nil {
			return nil, err
		}

		var IsFollow bool
		// 设置关注信息
		if err == gorm.ErrRecordNotFound {
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

	return videoList, nil
}
