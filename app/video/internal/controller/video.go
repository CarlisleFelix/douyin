package controller

import (
	"context"
	"douyin/app/video/internal/service"
	"douyin/app/video/utils"
	pb "douyin/idl/pb/video"
	"sync"
)

type VideoSrv struct {
	pb.UnimplementedVideoServiceServer
}

var VideoSrvIns *VideoSrv
var VideoSrvOnce sync.Once

func GetVideoSrv() *VideoSrv {
	VideoSrvOnce.Do(func() {
		VideoSrvIns = &VideoSrv{}
	})
	return VideoSrvIns
}

// 查找视频
func (v *VideoSrv) GetVideoInfo(ctx context.Context, req *pb.DouyinVideoRequest) (resp *pb.DouyinVideoResponse, err error) {
	video, err := service.GetVideoInfo(ctx, *req.VideoId)
	if err != nil {
		return nil, err
	}
	*resp.StatusCode = 0
	resp.Video = &video
	return resp, nil
}

// 更新视频评论
func (v *VideoSrv) UpdateCommentCount(ctx context.Context, req *pb.DouyinCommentCountRequest) (resp *pb.DouyinCommentCountResponse, err error) {
	err = service.UpdateCommentCount(ctx, *req.VideoId, *req.ActionType)
	if err != nil {
		return nil, err
	}
	*resp.StatusCode = 0
	return resp, nil
}

// 更新视频点赞
func (v *VideoSrv) UpdateFavoriteCount(ctx context.Context, req *pb.DouyinFavoriteCountRequest) (resp *pb.DouyinFavoriteCountResponse, err error) {
	err = service.UpdateFavoriteCount(ctx, *req.VideoId, *req.ActionType)
	if err != nil {
		return nil, err
	}
	*resp.StatusCode = 0
	return resp, nil
}

// 视频流
func (v *VideoSrv) Feed(ctx context.Context, req *pb.DouyinFeedRequest) (resp *pb.DouyinFeedResponse, err error) {
	//获取视频
	videoResponse, nextTime, err := service.FeedService(ctx, *req.UserId, *req.LatestTime)
	resp.NextTime = &nextTime
	resp.VideoList = videoResponse
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// 视频投稿
func (v *VideoSrv) PublishAction(ctx context.Context, req *pb.DouyinPublishActionRequest) (resp *pb.DouyinPublishActionResponse, err error) {
	curTime := utils.CurrentTimeInt()
	err = service.PublishService(ctx, *req.UserId, *req.Title, *req.FileExt, curTime)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// 视频列表
func (v *VideoSrv) PublishList(ctx context.Context, req *pb.DouyinPublishListRequest) (resp *pb.DouyinPublishListResponse, err error) {
	video, err := service.PublishListService(ctx, *req.QueryId, *req.HostId)
	resp.VideoList = video
	if err != nil {
		return nil, err
	}
	return resp, nil
}
