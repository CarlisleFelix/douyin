package controller

import (
	"context"
	pb "douyin/idl/pb/relation"
	"sync"
)

type RelationSrv struct {
	pb.UnimplementedRelationServiceServer
}

var RelationSrvIns *RelationSrv
var RelationSrvOnce sync.Once

func GetRelationSrv() *RelationSrv {
	RelationSrvOnce.Do(func() {
		RelationSrvIns = &RelationSrv{}
	})
	return RelationSrvIns
}

// 关注信息
func (r *RelationSrv) IsFollow(ctx context.Context, req *pb.DouyinIsFollowRequest) (resp *pb.DouyinIsFollowResponse, err error) {
	r, err := RelationClient.IsFollow(ctx, req)

	if err != nil {
		return
	}

	return r, nil
}

// 关系操作
func (r *RelationSrv) RelationAction(ctx context.Context, req *pb.DouyinRelationActionRequest) (resp *pb.DouyinRelationActionResponse, err error) {
	r, err := RelationClient.RelationAction(ctx, req)

	if err != nil {
		return
	}

	return r, nil
}

// 用户关注列表
func (r *RelationSrv) RelationFollowList(ctx context.Context, req *pb.DouyinRelationFollowListRequest) (resp *pb.DouyinRelationFollowListResponse, err error) {
	r, err := RelationClient.RelationFollowList(ctx, req)

	if err != nil {
		return
	}

	return r, nil
}

// 用户粉丝列表
func (r *RelationSrv) RelationFollowerList(ctx context.Context, req *pb.DouyinRelationFollowerListRequest) (resp *pb.DouyinRelationFollowerListResponse, err error) {
	r, err := RelationClient.RelationFollowerList(ctx, req)

	if err != nil {
		return
	}

	return r, nil
}

// 用户好友列表
func (r *RelationSrv) RelationFriendList(ctx context.Context, req *pb.DouyinRelationFriendListRequest) (resp *pb.DouyinRelationFriendListResponse, err error) {
	r, err := RelationClient.RelationFriendList(ctx, req)

	if err != nil {
		return
	}

	return r, nil
}
