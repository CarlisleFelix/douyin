package controller

import (
	"context"
	"douyin/app/favorite/internal/service"
	pb "douyin/idl/pb/favorite"
	"sync"
)

type FavoriteSrv struct {
	pb.UnimplementedFavoriteServiceServer
}

var FavoriteSrvIns *FavoriteSrv
var FavoriteSrvOnce sync.Once

func GetFavoriteSrv() *FavoriteSrv {
	FavoriteSrvOnce.Do(func() {
		FavoriteSrvIns = &FavoriteSrv{}
	})
	return FavoriteSrvIns
}

// 具体实现proto中定义的服务

func (f *FavoriteSrv) IsFavorite(ctx context.Context, req *pb.DouyinIsFavoriteRequest) (resp *pb.DouyinIsFavoriteResponse, err error) {
	*resp.IsFavorite = service.IsFavorite(ctx, *req.UserId, *req.VideoId)
	return resp, err
}

func (f *FavoriteSrv) FavoriteAction(ctx context.Context, req *pb.DouyinFavoriteActionRequest) (resp *pb.DouyinFavoriteActionResponse, err error) {
	err = service.FavoriteAction(ctx, *req.UserId, *req.VideoId, *req.ActionType)
	if err != nil {
		*resp.StatusCode = 1
	}
	return
}

func (f *FavoriteSrv) FavoriteList(ctx context.Context, req *pb.DouyinFavoriteListRequest) (resp *pb.DouyinFavoriteListResponse, err error) {
	resp.VideoList, err = service.FavoriteList(ctx, *req.UserId)
	*resp.StatusCode = 0
	return
}
