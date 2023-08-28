package controller

import (
	"context"
	pb "douyin/idl/pb/favorite"
	"fmt"
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

}

func (f *FavoriteSrv) FavoriteAction(ctx context.Context, req *pb.DouyinFavoriteActionRequest) (resp *pb.DouyinFavoriteActionResponse, err error) {
	fmt.Println("成功调用点赞服务")
	resp = new(pb.DouyinFavoriteActionResponse)
	var code int32 = 200
	msg := "点赞测试成功！！！"
	resp.StatusCode = &code
	resp.StatusMsg = &msg

	return
}

func (f *FavoriteSrv) FavoriteList(ctx context.Context, req *pb.DouyinFavoriteListRequest) (resp *pb.DouyinFavoriteListResponse, err error) {
	resp = new(pb.DouyinFavoriteListResponse)
	return
}
