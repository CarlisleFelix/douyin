package rpc

import (
	"context"
	pb "douyin/idl/pb/comment"
)

// 评论操作
func CommentAction(ctx context.Context, req *pb.DouyinCommentActionRequest) (resp *pb.DouyinCommentActionResponse, err error) {
	r, err := CommentClient.CommentAction(ctx, req)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// 评论列表
func CommentList(ctx context.Context, req *pb.DouyinCommentListRequest) (resp *pb.DouyinCommentListResponse, err error) {
	r, err := CommentClient.CommentList(ctx, req)
	if err != nil {
		return nil, err
	}
	return r, nil
}
