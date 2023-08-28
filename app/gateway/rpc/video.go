package rpc

import (
	"context"
	pb "douyin/idl/pb/video"
)

// 查找视频
func GetVideoInfo(ctx context.Context, req *pb.DouyinVideoRequest) (resp *pb.DouyinVideoResponse, err error) {
	r, err := VideoClient.GetVideoInfo(ctx, req)

	if err != nil {
		return nil, err
	}

	return r, nil
}

// 更新视频评论
func UpdateCommentCount(ctx context.Context, req *pb.DouyinCommentCountRequest) (resp *pb.DouyinCommentCountResponse, err error) {
	r, err := VideoClient.UpdateCommentCount(ctx, req)

	if err != nil {
		return nil, err
	}

	return r, nil
}

// 更新视频点赞
func UpdateFavoriteCount(ctx context.Context, req *pb.DouyinFavoriteCountRequest) (resp *pb.DouyinFavoriteCountResponse, err error) {
	r, err := VideoClient.UpdateFavoriteCount(ctx, req)

	if err != nil {
		return nil, err
	}

	return r, nil
}

// 视频流
func Feed(ctx context.Context, req *pb.DouyinFeedRequest) (resp *pb.DouyinFeedResponse, err error) {
	r, err := VideoClient.Feed(ctx, req)

	if err != nil {
		return nil, err
	}

	return r, nil
}

// 视频投稿
func PublishAction(ctx context.Context, req *pb.DouyinPublishActionRequest) (resp *pb.DouyinPublishActionResponse, err error) {
	r, err := VideoClient.PublishAction(ctx, req)

	if err != nil {
		return nil, err
	}

	return r, nil
}

// 视频列表
func PublishList(ctx context.Context, req *pb.DouyinPublishListRequest) (resp *pb.DouyinPublishListResponse, err error) {
	r, err := VideoClient.PublishList(ctx, req)
	if err != nil {
		return nil, err
	}
	return r, nil
}
