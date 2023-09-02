package service

import (
	"bytes"
	"context"
	"douyin/dao"
	"douyin/global"
	"douyin/model"
	"douyin/response"
	"fmt"
	"io"
	"os"
	"path/filepath"

	ffmpeg "github.com/u2takey/ffmpeg-go"
	"gorm.io/gorm"
)

const videoNum = 2

func PublishService(userId int64, title string, fileExt string, curTime int64) error {
	fileName := fmt.Sprintf("%d_%s", userId, title) //标识名字
	finalFilename := fileName + fileExt
	saveFilepath := filepath.Join("../tmp/", finalFilename) //路径+文件名
	var data io.Reader
	data, err := os.Open(saveFilepath)
	if err != nil {
		return global.ErrorFileOperationWrong
	}

	exist := UserIdExists(userId)
	if !exist {
		return global.ErrorUserNotExist
	}

	exist = VideoExists(userId, title)
	if exist {
		return global.ErrorVideoDuplicate
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
	err = dao.InsertVideo(&newVideo)
	//修改数据库表

	return err
}

func UploadVideo(fileName string, reader io.Reader) (string, error) {
	_, err := global.SERVER_COS_VIDEO.Object.Put(context.Background(), fileName, reader, nil)
	return global.SERVER_CONFIG.Cos.Video_bucket_url + "/" + fileName, err
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
	_, err = global.SERVER_COS_COVER.Object.Put(context.Background(), coverName, buf, nil)
	return global.SERVER_CONFIG.Cos.Cover_bucket_url + "/" + coverName, err
}

func PublishListService(queryUserId int64, hostUserId int64, ctx context.Context) ([]response.Video_Response, error) {

	ctx, span := global.SERVER_VIDEO_TRACER.Start(ctx, "publishlist service")
	defer span.End()

	exist := UserIdExists(queryUserId)
	if !exist {
		return nil, global.ErrorUserNotExist
	}

	queryUser, _ := dao.GetUserById(queryUserId)
	isFollow := dao.GetFollowByUserId(hostUserId, queryUserId)
	//author
	userResponse := response.User_Response{
		Id:              queryUser.User_id,
		Name:            queryUser.User_name,
		FollowCount:     queryUser.Follow_count,
		FollowerCount:   queryUser.Follower_count,
		IsFollow:        isFollow,
		Avatar:          queryUser.Avatar,
		BackgroundImage: queryUser.Background_image,
		Signature:       queryUser.Signature,
		TotalFavorited:  queryUser.Favorite_count,
		WorkCount:       queryUser.Work_count,
		FavoriteCount:   queryUser.Favorite_count,
	}

	//video
	videos, err := GetVideolistByauthor(queryUserId)
	if err != nil {
		return nil, err
	}

	//response
	var videoList []response.Video_Response
	if len(videos) == 0 {
		return nil, nil
	}
	for i := 0; i < len(videos); i++ {
		responseVideo := response.Video_Response{
			Id:            videos[i].Video_id,
			Author:        userResponse,
			PlayUrl:       videos[i].Play_url,
			CoverUrl:      videos[i].Cover_url,
			FavoriteCount: videos[i].Favorite_count,
			CommentCount:  videos[i].Comment_count,
			IsFavorite:    dao.GetifFavorite(hostUserId, videos[i].Video_id),
			Title:         videos[i].Title,
		}
		videoList = append(videoList, responseVideo)
	}
	return videoList, err
}

func UserIdExists(userId int64) bool {
	_, err := dao.GetUserById(userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		} else {
			return false
		}
	}
	return true
}

func VideoExists(userId int64, title string) bool {
	_, err := dao.GetVideoByAuthorIdandTitle(userId, title)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		} else {
			return true
		}
	}
	return true
}

func GetVideolistByauthor(userId int64) ([]model.Video, error) {
	videos, err := dao.GetVideoByAuthorId(userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return videos, nil
}

func FeedService(userId int64, latestTime int64, ctx context.Context) ([]response.Video_Response, int64, error) {

	ctx, span := global.SERVER_USER_TRACER.Start(ctx, "feed service")
	defer span.End()

	//获得视频
	videos, err := dao.GetVideoByTime(latestTime)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		} else {
			return nil, 0, err
		}
	}
	//生成response，填入作者信息
	var videoList []response.Video_Response
	var nextTime int64
	if len(videos) == 0 {
		return nil, 0, nil
	}
	for i := 0; i < len(videos); i++ {
		//有问题，得考虑作者不存在的情况..
		author, _ := dao.GetUserById(videos[i].Author_id)
		userResponse := response.User_Response{
			Id:              author.User_id,
			Name:            author.User_name,
			FollowCount:     author.Follow_count,
			FollowerCount:   author.Follower_count,
			IsFollow:        dao.GetFollowByUserId(userId, author.User_id),
			Avatar:          author.Avatar,
			BackgroundImage: author.Background_image,
			Signature:       author.Signature,
			TotalFavorited:  author.Total_favorited,
			WorkCount:       author.Work_count,
			FavoriteCount:   author.Favorite_count,
		}

		responseVideo := response.Video_Response{
			Id:            videos[i].Video_id,
			Author:        userResponse,
			PlayUrl:       videos[i].Play_url,
			CoverUrl:      videos[i].Cover_url,
			FavoriteCount: videos[i].Favorite_count,
			CommentCount:  videos[i].Comment_count,
			IsFavorite:    dao.GetifFavorite(userId, videos[i].Video_id),
			Title:         videos[i].Title,
		}

		videoList = append(videoList, responseVideo)
		if i == 0 {
			nextTime = videos[i].Publish_time
		}
	}
	return videoList, nextTime, err
}
