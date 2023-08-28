package controller

import (
	"context"
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
	r, err := VideoClient.GetVideoInfo(ctx, req)

	if err != nil {
		return nil, err
	}

	return r, nil
}

// 更新视频评论
func (v *VideoSrv) UpdateCommentCount(ctx context.Context, req *pb.DouyinCommentCountRequest) (resp *pb.DouyinCommentCountResponse, err error) {
	r, err := VideoClient.UpdateCommentCount(ctx, req)

	if err != nil {
		return nil, err
	}

	return r, nil
}

// 更新视频点赞
func (v *VideoSrv) UpdateFavoriteCount(ctx context.Context, req *pb.DouyinFavoriteCountRequest) (resp *pb.DouyinFavoriteCountResponse, err error) {
	r, err := VideoClient.UpdateFavoriteCount(ctx, req)

	if err != nil {
		return nil, err
	}

	return r, nil
}

// 视频流
func (v *VideoSrv) Feed(ctx context.Context, req *pb.DouyinFeedRequest) (resp *pb.DouyinFeedResponse, err error) {
	r, err := VideoClient.Feed(ctx, req)

	if err != nil {
		return nil, err
	}

	return r, nil
}

// 视频投稿
func (v *VideoSrv) PublishAction(ctx context.Context, req *pb.DouyinPublishActionRequest) (resp *pb.DouyinPublishActionResponse, err error) {
	r, err := VideoClient.PublishAction(ctx, req)

	if err != nil {
		return nil, err
	}

	return r, nil
}

// 视频列表
func (v *VideoSrv) PublishList(ctx context.Context, req *pb.DouyinPublishListRequest) (resp *pb.DouyinPublishListResponse, err error) {
	r, err := VideoClient.PublishList(ctx, req)
	if err != nil {
		return nil, err
	}
	return r, nil
}
