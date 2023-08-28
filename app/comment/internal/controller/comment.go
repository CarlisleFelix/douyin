package controller

import (
	"context"
	pb "douyin/idl/pb/comment"
	"sync"
)

type CommentSrv struct {
	pb.UnimplementedCommentServiceServer
}

var CommentSrvIns *CommentSrv
var CommentSrvOnce sync.Once

func GetCommentSrv() *CommentSrv {
	CommentSrvOnce.Do(func() {
		CommentSrvIns = &CommentSrv{}
	})
	return CommentSrvIns
}

// 评论操作
func (c *CommentSrv) CommentAction(ctx context.Context, req *pb.DouyinCommentActionRequest) (resp *pb.DouyinCommentActionResponse, err error) {

}

// 评论列表
func (c *CommentSrv) CommentList(ctx context.Context, req *pb.DouyinCommentListRequest) (resp *pb.DouyinCommentListResponse, err error) {

}
