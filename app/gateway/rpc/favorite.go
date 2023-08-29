package rpc

import (
	"context"
	pb "douyin/idl/pb/favorite"
)

var SUCCESS int32 = 200

func IsFavorite(ctx context.Context, req *pb.DouyinIsFavoriteRequest) (resp *pb.DouyinIsFavoriteResponse, err error) {
	resp, err = FavoriteClient.IsFavorite(ctx, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func FavoriteAction(ctx context.Context, req *pb.DouyinFavoriteActionRequest) (resp *pb.DouyinFavoriteActionResponse, err error) {
	resp, err = FavoriteClient.FavoriteAction(ctx, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func FavoriteList(ctx context.Context, req *pb.DouyinFavoriteListRequest) (resp *pb.DouyinFavoriteListResponse, err error) {
	resp, err = FavoriteClient.FavoriteList(ctx, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
