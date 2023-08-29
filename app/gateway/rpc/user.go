package rpc

import (
	"context"
	pb "douyin/idl/pb/user"
)

// 处理调用rpc时产生的错误,以及微服务与微服务之间的调用

func UserRegister(ctx context.Context, req *pb.DouyinUserRegisterRequest) (resp *pb.DouyinUserRegisterResponse, err error) {
	resp, err = UserClient.UserRegister(ctx, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func UserLogin(ctx context.Context, req *pb.DouyinUserLoginRequest) (resp *pb.DouyinUserLoginResponse, err error) {
	resp, err = UserClient.UserLogin(ctx, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func GetUserInfo(ctx context.Context, req *pb.DouyinUserRequest) (resp *pb.DouyinUserResponse, err error) {
	resp, err = UserClient.GetUserInfo(ctx, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// todo:待修改传入参数
func UpdateTotalFavorite(ctx context.Context, req *pb.DouyinTotalFavoriteRequest) (resp *pb.DouyinTotalFavoriteResponse, err error) {
	resp, err = UserClient.UpdateTotalFavorite(ctx, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func UpdateUserFavoriteCount(ctx context.Context, req *pb.DouyinFavoriteCountRequest) (resp *pb.DouyinFavoriteCountResponse, err error) {
	resp, err = UserClient.UpdateUserFavoriteCount(ctx, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func UpdateFollowCount(ctx context.Context, req *pb.DouyinFollowCountRequest) (resp *pb.DouyinFollowCountResponse, err error) {
	resp, err = UserClient.UpdateFollowCount(ctx, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func UpdateFollowerCount(ctx context.Context, req *pb.DouyinFollowerCountRequest) (resp *pb.DouyinFollowerCountResponse, err error) {
	resp, err = UserClient.UpdateFollowerCount(ctx, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
