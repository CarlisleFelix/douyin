package rpc

import (
	"context"
	pb "douyin/idl/pb/user"
	"fmt"
)

// 处理调用rpc时产生的错误,以及微服务与微服务之间的调用

func UserRegister(ctx context.Context, req *pb.DouyinUserRegisterRequest) (resp *pb.DouyinUserRegisterResponse, err error) {
	r, err := UserClient.UserRegister(ctx, req)

	if err != nil {
		return
	}

	return r, nil
}

func UserLogin(ctx context.Context, req *pb.DouyinUserLoginRequest) (resp *pb.DouyinUserLoginResponse, err error) {
	r, err := UserClient.UserLogin(ctx, req)

	if err != nil {
		return
	}

	return r, nil
}

func GetUserInfo(ctx context.Context, req *pb.DouyinUserRequest) (resp *pb.DouyinUserResponse, err error) {
	r, err := UserClient.GetUserInfo(ctx, req)

	if err != nil {
		fmt.Println("调用rpc错误")
		return
	}
	fmt.Println("成功调用rpc")

	return r, nil
}

// todo:待修改传入参数
func UpdateTotalFavorite(ctx context.Context, req *pb.DouyinTotalFavoriteRequest) (resp *pb.DouyinTotalFavoriteResponse, err error) {
	r, err := UserClient.UpdateTotalFavorite(ctx, req)

	if err != nil {
		return
	}

	return r, nil
}

func UpdateFollowCount(ctx context.Context, req *pb.DouyinFollowCountRequest) (resp *pb.DouyinFollowCountResponse, err error) {
	r, err := UserClient.UpdateFollowCount(ctx, req)

	if err != nil {
		return
	}

	return r, nil
}

func UpdateFollowerCount(ctx context.Context, req *pb.DouyinFollowerCountRequest) (resp *pb.DouyinFollowerCountResponse, err error) {
	r, err := UserClient.UpdateFollowerCount(ctx, req)

	if err != nil {
		return
	}

	return r, nil
}
