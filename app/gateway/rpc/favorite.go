package rpc

import (
	"context"
	pb "douyin/idl/pb/favorite"
	"errors"
)

var SUCCESS int32 = 200

func IsFavorite(ctx context.Context, req *pb.DouyinIsFavoriteRequest) (resp *pb.DouyinIsFavoriteResponse, err error) {
	r, err := FavoriteClient.IsFavorite(ctx, req)

	if err != nil {
		return nil, err
	}

	//if *r.StatusCode != SUCCESS {
	//	err = errors.New(*r.StatusMsg)
	//	return
	//}

	return r, nil
}

// 处理调用rpc时产生的错误
func FavoriteAction(ctx context.Context, req *pb.DouyinFavoriteActionRequest) (resp *pb.DouyinFavoriteActionResponse, err error) {
	r, err := FavoriteClient.FavoriteAction(ctx, req)

	if err != nil {
		return nil, err
	}

	//if *r.StatusCode != SUCCESS {
	//	err = errors.New(*r.StatusMsg)
	//	return
	//}

	return r, nil
}

func FavoriteList(ctx context.Context, req *pb.DouyinFavoriteListRequest) (resp *pb.DouyinFavoriteListResponse, err error) {
	r, err := FavoriteClient.FavoriteList(ctx, req)

	if err != nil {
		return nil, err
	}

	if *r.StatusCode != SUCCESS {
		err = errors.New(*r.StatusMsg)
		return
	}

	return r, nil
}
