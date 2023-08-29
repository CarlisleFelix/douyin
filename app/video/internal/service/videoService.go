package service

import (
	"bytes"
	"context"
	"douyin/app/gateway/rpc"
	"douyin/app/video/internal/dal/dao"
	"douyin/app/video/internal/dal/model"
	"douyin/app/video/internal/server"
	"douyin/idl/pb/favorite"
	"douyin/idl/pb/relation"
	"douyin/idl/pb/user"
	pb "douyin/idl/pb/video"
	"douyin/utils/e"
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"gorm.io/gorm"
	"io"
	"os"
	"path/filepath"
)

const videoNum = 2

// todo 未加入isFollow
func GetVideoInfo(ctx context.Context, video_id int64) (pb.Video, error) {
	var videoResp pb.Video
	video, err := dao.NewVideoDao(ctx).GetVideoById(video_id)
	if err != nil {
		return videoResp, err
	}

	var isFavoriteReq favorite.DouyinIsFavoriteRequest
	isFavoriteReq.UserId = &video.Author_id
	isFavoriteReq.VideoId = &video.Video_id
	isFavoriteResp, err := rpc.IsFavorite(ctx, &isFavoriteReq)
	if err != nil {
		return videoResp, err
	}

	var getUserInfoReq user.DouyinUserRequest
	getUserInfoReq.HostId = &video.Author_id
	getUserInfoResp, err := rpc.GetUserInfo(ctx, &getUserInfoReq)
	if err != nil {
		return videoResp, err
	}

	videoResp = pb.Video{
		Id: &video.Video_id,
		Author: &pb.User{
			Id:              getUserInfoResp.User.Id,
			Name:            getUserInfoResp.User.Name,
			FollowCount:     getUserInfoResp.User.FollowCount,
			FollowerCount:   getUserInfoResp.User.FollowerCount,
			IsFollow:        nil,
			Avatar:          getUserInfoResp.User.Avatar,
			BackgroundImage: getUserInfoResp.User.BackgroundImage,
			Signature:       getUserInfoResp.User.Signature,
			TotalFavorited:  getUserInfoResp.User.TotalFavorited,
			WorkCount:       getUserInfoResp.User.WorkCount,
			FavoriteCount:   getUserInfoResp.User.FavoriteCount,
		},
		PlayUrl:       &video.Play_url,
		CoverUrl:      &video.Cover_url,
		FavoriteCount: &video.Favorite_count,
		CommentCount:  &video.Comment_count,
		IsFavorite:    isFavoriteResp.IsFavorite,
		Title:         &video.Title,
	}
	return videoResp, err

}

func UpdateFavoriteCount(ctx context.Context, video_id int64, action_type int32) error {
	return dao.NewVideoDao(ctx).UpdateVideoFavoriteCount(video_id, action_type)
}

func UpdateCommentCount(ctx context.Context, video_id int64, action_type int32) error {
	return dao.NewVideoDao(ctx).UpdateVideoCommentCount(video_id, action_type)
}

func PublishService(ctx context.Context, userId int64, title string, fileExt string, curTime int64) error {
	fileName := fmt.Sprintf("%d_%s", userId, title) //标识名字
	finalFilename := fileName + fileExt
	saveFilepath := filepath.Join("../tmp/", finalFilename) //路径+文件名
	var data io.Reader
	data, err := os.Open(saveFilepath)
	if err != nil {
		return e.ErrorFileOperationWrong
	}

	exist := UserIdExists(ctx, userId)
	if !exist {
		return e.ErrorUserNotExist
	}

	exist = VideoExists(ctx, userId, title)
	if exist {
		return e.ErrorVideoDuplicate
	}

	//上传视频
	videoUrl, err := UploadVideo(finalFilename, data)
	if err != nil {
		return err
	}
	//提取封面并上传
	coverUrl, err := ExtractCoverandUpload(finalFilename, saveFilepath, 3)
	if err != nil {
		return err
	}

	newVideo := model.Video{
		Author_id:      userId,
		Play_url:       videoUrl,
		Cover_url:      coverUrl,
		Favorite_count: 0,
		Comment_count:  0,
		Title:          title,
		Publish_time:   curTime,
	}
	err = dao.NewVideoDao(ctx).InsertVideo(&newVideo)
	//修改数据库表

	return err
}

func UploadVideo(fileName string, reader io.Reader) (string, error) {
	_, err := server.SERVER_COS_VIDEO.Object.Put(context.Background(), fileName, reader, nil)
	return server.SERVER_CONFIG.Cos.Video_bucket_url + "/" + fileName, err
}

func ExtractCoverandUpload(finalFilename string, saveFilepath string, frameNum int) (string, error) {
	//提取封面
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(saveFilepath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		//fmt.Println("err in extraction")
		return "", err
	}

	fileExt := filepath.Ext(finalFilename)
	pureName := finalFilename[0 : len(finalFilename)-len(fileExt)]
	coverName := pureName + ".jpeg"
	_, err = server.SERVER_COS_COVER.Object.Put(context.Background(), coverName, buf, nil)
	return server.SERVER_CONFIG.Cos.Cover_bucket_url + "/" + coverName, err
}

func PublishListService(ctx context.Context, queryUserId int64, hostUserId int64) ([]*pb.Video, error) {
	exist := UserIdExists(ctx, queryUserId)
	if !exist {
		return nil, e.ErrorUserNotExist
	}

	var getUserReq user.DouyinUserRequest
	getUserReq.HostId = &queryUserId
	getUserResp, _ := rpc.GetUserInfo(ctx, &getUserReq)

	var isFollowReq relation.DouyinIsFollowRequest
	isFollowReq.HostId = &hostUserId
	isFollowReq.GuestId = &queryUserId
	isFollowResp, _ := rpc.IsFollow(ctx, &isFollowReq)
	//author
	userResponse := pb.User{
		Id:              getUserResp.User.Id,
		Name:            getUserResp.User.Name,
		FollowCount:     getUserResp.User.FollowCount,
		FollowerCount:   getUserResp.User.FollowerCount,
		IsFollow:        isFollowResp.IsFollow,
		Avatar:          getUserResp.User.Avatar,
		BackgroundImage: getUserResp.User.BackgroundImage,
		Signature:       getUserResp.User.Signature,
		TotalFavorited:  getUserResp.User.TotalFavorited,
		WorkCount:       getUserResp.User.WorkCount,
		FavoriteCount:   getUserResp.User.FavoriteCount,
	}

	//video
	videos, err := GetVideolistByauthor(ctx, queryUserId)
	if err != nil {
		return nil, err
	}

	//response
	var videoList []*pb.Video
	if len(videos) == 0 {
		return nil, nil
	}
	for i := 0; i < len(videos); i++ {
		var isFavoriteReq favorite.DouyinIsFavoriteRequest
		isFavoriteReq.UserId = &hostUserId
		isFavoriteReq.VideoId = &videos[i].Video_id
		isFavoriteResp, err := rpc.IsFavorite(ctx, &isFavoriteReq)

		if err != nil {
			return videoList, err
		}
		responseVideo := pb.Video{
			Id:            &videos[i].Video_id,
			Author:        &userResponse,
			PlayUrl:       &videos[i].Play_url,
			CoverUrl:      &videos[i].Cover_url,
			FavoriteCount: &videos[i].Favorite_count,
			CommentCount:  &videos[i].Comment_count,
			IsFavorite:    isFavoriteResp.IsFavorite,
			Title:         &videos[i].Title,
		}
		videoList = append(videoList, &responseVideo)
	}
	return videoList, err
}

func UserIdExists(ctx context.Context, userId int64) bool {
	var getUserReq user.DouyinUserRequest
	getUserReq.HostId = &userId
	_, err := rpc.GetUserInfo(ctx, &getUserReq)
	if err != nil {
		return false
	}
	return true
}

func VideoExists(ctx context.Context, userId int64, title string) bool {
	_, err := dao.NewVideoDao(ctx).GetVideoByAuthorIdandTitle(userId, title)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		} else {
			return true
		}
	}
	return true
}

func GetVideolistByauthor(ctx context.Context, userId int64) ([]model.Video, error) {
	videos, err := dao.NewVideoDao(ctx).GetVideoByAuthorId(userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return videos, nil
}

func FeedService(ctx context.Context, userId int64, latestTime int64) ([]*pb.Video, int64, error) {
	//获得视频
	videos, err := dao.NewVideoDao(ctx).GetVideoByTime(latestTime)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		} else {
			return nil, 0, err
		}
	}
	//生成response，填入作者信息
	var videoList []*pb.Video
	var nextTime int64
	if len(videos) == 0 {
		return nil, 0, nil
	}
	for i := 0; i < len(videos); i++ {
		var getUserReq user.DouyinUserRequest
		getUserReq.HostId = &videos[i].Author_id
		getUserResp, err := rpc.GetUserInfo(ctx, &getUserReq)

		var isFollowReq relation.DouyinIsFollowRequest
		isFollowReq.HostId = &userId
		isFollowReq.GuestId = getUserResp.User.Id
		isFollowResp, err := rpc.IsFollow(ctx, &isFollowReq)
		if err != nil {
			*isFollowResp.IsFollow = false
		}

		var isFavoriteReq favorite.DouyinIsFavoriteRequest
		isFavoriteReq.UserId = &userId
		isFavoriteReq.VideoId = &videos[i].Video_id
		isFavoriteResp, err := rpc.IsFavorite(ctx, &isFavoriteReq)
		if err != nil {
			*isFavoriteResp.IsFavorite = false
		}

		responseVideo := pb.Video{
			Id: &videos[i].Video_id,
			//有问题，得考虑作者不存在的情况..
			Author: &pb.User{
				Id:              getUserResp.User.Id,
				Name:            getUserResp.User.Name,
				FollowCount:     getUserResp.User.FollowCount,
				FollowerCount:   getUserResp.User.FollowerCount,
				IsFollow:        isFollowResp.IsFollow,
				Avatar:          getUserResp.User.Avatar,
				BackgroundImage: getUserResp.User.BackgroundImage,
				Signature:       getUserResp.User.Signature,
				TotalFavorited:  getUserResp.User.TotalFavorited,
				WorkCount:       getUserResp.User.WorkCount,
				FavoriteCount:   getUserResp.User.FavoriteCount,
			},
			PlayUrl:       &videos[i].Play_url,
			CoverUrl:      &videos[i].Cover_url,
			FavoriteCount: &videos[i].Favorite_count,
			CommentCount:  &videos[i].Comment_count,
			IsFavorite:    isFavoriteResp.IsFavorite,
			Title:         &videos[i].Title,
		}

		videoList = append(videoList, &responseVideo)
		if i == 0 {
			nextTime = videos[i].Publish_time
		}
	}
	return videoList, nextTime, err
}
