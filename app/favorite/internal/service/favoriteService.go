package service

import (
	"context"
	"douyin/app/favorite/internal/dal/dao"
	"douyin/app/gateway/rpc"
	pb "douyin/idl/pb/favorite"
	"douyin/idl/pb/relation"
	"douyin/idl/pb/user"
	"douyin/idl/pb/video"
	"douyin/utils/e"
	"gorm.io/gorm"
)

func IsFavorite(ctx context.Context, user_id int64, video_id int64) bool {
	isFavorite := dao.NewFavoriteDao(ctx).IsFavorite(user_id, video_id)
	return isFavorite
}

func FavoriteAction(ctx context.Context, user_id int64, video_id int64, action_type int32) error {
	// 获取点赞人信息
	_, err := getUserInfo(ctx, user_id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return e.ErrorUserNotExist
		}
	}
	// 获取视频信息
	videoResp, err := getVideoInfo(ctx, video_id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return e.ErrorVideoNotExist
		}
	}
	// 获取作者信息
	_, err = getUserInfo(ctx, *videoResp.Video.Author.Id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return e.ErrorUserNotExist
		}
	}
	// 根据user_id,video_id查找点赞信息
	isFavorite := dao.NewFavoriteDao(ctx).IsFavorite(user_id, video_id)

	// 开始执行
	tx := dao.NewFavoriteDao(ctx).Begin()

	// 对action_type进行判断, 1-点赞， 2-取消点赞
	if action_type == 1 {
		// 如果是点赞操作
		// 先查询是否已经点赞
		if isFavorite == false {
			// 如果未点赞，则将点赞信息存储到表中
			// 视频点赞数+1
			err := updateVideoFavoriteCount(ctx, video_id, action_type)
			if err != nil {
				tx.Rollback()
				return err
			}
			// 用户点赞数+1
			err = updateUserFavoriteCount(ctx, user_id, action_type)
			if err != nil {
				tx.Rollback()
				return err
			}
			// 作者获赞数+1
			err = updateUserTotalFavorite(ctx, user_id, action_type)
			if err != nil {
				tx.Rollback()
				return err
			}

			// 创建点赞记录
			err = dao.NewFavoriteDao(ctx).CreateFavorite(user_id, video_id)
			// 判断点赞操作是否成功
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
			return e.ErrorFavoriteExist
		}
	} else {
		// 如果是取消点赞操作，先查询表中是否存在点赞记录，
		if err == gorm.ErrRecordNotFound {
			// 如果不存在，则返回异常
			tx.Rollback()
			return e.ErrorFavoriteNotExist
		} else {
			// 如果存在，则删除该记录
			// 视频点赞数-1
			err = updateVideoFavoriteCount(ctx, video_id, action_type)
			if err != nil {
				tx.Rollback()
				return err
			}
			// 用户点赞数-1
			err = updateUserFavoriteCount(ctx, user_id, action_type)
			if err != nil {
				tx.Rollback()
				return err
			}
			// 作者获赞数-1
			err = updateUserTotalFavorite(ctx, user_id, action_type)
			if err != nil {
				tx.Rollback()
				return err
			}

			// 删除点赞记录
			err = dao.NewFavoriteDao(ctx).DeleteFavorite(user_id, video_id)
			// 判断操作是否成功
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

func FavoriteList(ctx context.Context, user_id int64) (videoList []*pb.Video, err error) {
	favorites, err := dao.NewFavoriteDao(ctx).SearchFavoriteList(user_id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 用户无喜欢视频
			return nil, nil
		} else {
			return nil, err
		}
	}

	// 创建视频列表
	videoList = make([]*pb.Video, len(favorites))
	for i := 0; i < len(favorites); i++ {
		// 获取视频信息
		video, err := getVideoInfo(ctx, favorites[i].Video_id)
		if err != nil {
			continue
		}
		// 获取用户信息
		author, err := getUserInfo(ctx, favorites[i].User_id)
		if err != nil {
			continue
		}
		// 获取关注信息
		var isRelationReq relation.DouyinIsFollowRequest
		isRelationReq.HostId = &user_id
		isRelationReq.GuestId = author.User.Id
		isRelationResp, err := rpc.IsFollow(ctx, &isRelationReq)
		// 设置关注信息
		if err != nil {
			continue
		}

		videoList[i] = &pb.Video{
			Id: video.Video.Id,
			// 存储作者信息
			Author: &pb.User{
				Id:              author.User.Id,
				Name:            author.User.Name,
				FollowCount:     author.User.FollowCount,
				FollowerCount:   author.User.FollowerCount,
				IsFollow:        isRelationResp.IsFollow,
				Avatar:          author.User.Avatar,
				BackgroundImage: author.User.BackgroundImage,
				Signature:       author.User.Signature,
				TotalFavorited:  author.User.TotalFavorited,
				WorkCount:       author.User.WorkCount,
				FavoriteCount:   author.User.FavoriteCount,
			},
			PlayUrl:       video.Video.PlayUrl,
			CoverUrl:      video.Video.CoverUrl,
			FavoriteCount: video.Video.FavoriteCount,
			CommentCount:  video.Video.CommentCount,
			IsFavorite:    nil,
			Title:         video.Video.Title,
		}
	}

	return videoList, nil
}

func getVideoInfo(ctx context.Context, video_id int64) (video.DouyinVideoResponse, error) {
	var videoReq video.DouyinVideoRequest
	videoReq.VideoId = &video_id
	videoResp, err := rpc.GetVideoInfo(ctx, &videoReq)
	return *videoResp, err
}

func getUserInfo(ctx context.Context, user_id int64) (user.DouyinUserResponse, error) {
	var userReq user.DouyinUserRequest
	userReq.HostId = &user_id
	userResp, err := rpc.GetUserInfo(ctx, &userReq)
	return *userResp, err
}

func updateVideoFavoriteCount(ctx context.Context, video_id int64, action_type int32) error {
	var videoFavReq video.DouyinFavoriteCountRequest
	videoFavReq.VideoId = &video_id
	videoFavReq.ActionType = &action_type
	_, err := rpc.UpdateFavoriteCount(ctx, &videoFavReq)
	return err
}

func updateUserFavoriteCount(ctx context.Context, user_id int64, action_type int32) error {
	var userFavReq user.DouyinFavoriteCountRequest
	userFavReq.UserId = &user_id
	userFavReq.ActionType = &action_type
	_, err := rpc.UpdateUserFavoriteCount(ctx, &userFavReq)
	return err
}

func updateUserTotalFavorite(ctx context.Context, user_id int64, action_type int32) error {
	var authorTotReq user.DouyinFavoriteCountRequest
	authorTotReq.UserId = &user_id
	authorTotReq.ActionType = &action_type
	_, err := rpc.UpdateUserFavoriteCount(ctx, &authorTotReq)
	return err
}
