package rpc

import (
	"context"
	favoritePb "douyin/idl/pb/favorite"
	"errors"
)

var SUCCESS int32 = 200

// 处理调用rpc时产生的错误
func FavoriteAction(ctx context.Context, req *favoritePb.DouyinFavoriteActionRequest) (resp *favoritePb.DouyinFavoriteActionResponse, err error) {
	r, err := FavoriteClient.FavoriteAction(ctx, req)

	if err != nil {
		return
	}

	//if *r.StatusCode != SUCCESS {
	//	err = errors.New(*r.StatusMsg)
	//	return
	//}

	return r, nil
}

func FavoriteList(ctx context.Context, req *favoritePb.DouyinFavoriteListRequest) (resp *favoritePb.DouyinFavoriteListResponse, err error) {
	r, err := FavoriteClient.FavoriteList(ctx, req)

	if err != nil {
		return
	}

	if *r.StatusCode != SUCCESS {
		err = errors.New(*r.StatusMsg)
		return
	}

	return r, nil
}
