package controller

import (
	"context"
	"douyin/app/gateway/middleware"
	"douyin/app/user/internal/dal/model"
	"douyin/app/user/internal/service"
	pb "douyin/idl/pb/user"
	"sync"
)

type UserSrv struct {
	pb.UnimplementedUserServiceServer
}

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

func (u *UserSrv) UserLogin(ctx context.Context, req *pb.DouyinUserLoginRequest) (resp *pb.DouyinUserLoginResponse, err error) {
	var loginUser model.User
	loginUser, err = service.UserLoginService(ctx, *req.Username, *req.Password)
	if err != nil {
		return resp, err
	}

	//生成token
	token, err := middleware.GenerateToken(loginUser.User_id, loginUser.User_name)
	resp.Token = &token
	resp.UserId = &loginUser.User_id
	return
}

func (u *UserSrv) UserRegister(ctx context.Context, req *pb.DouyinUserRegisterRequest) (resp *pb.DouyinUserRegisterResponse, err error) {
	var newUser model.User
	newUser, err = service.UserRegisterService(ctx, *req.Username, *req.Password)
	resp.UserName = &newUser.User_name
	resp.UserId = &newUser.User_id
	if err != nil {
		return nil, err
	}
	return
}

func (u *UserSrv) GetUserInfo(ctx context.Context, req *pb.DouyinUserRequest) (resp *pb.DouyinUserResponse, err error) {
	resp.User, err = service.GetUserInfo(ctx, *req.HostId)
	if err != nil {
		return nil, err
	}
	return
}

func (u *UserSrv) UpdateTotalFavorite(ctx context.Context, req *pb.DouyinTotalFavoriteRequest) (resp *pb.DouyinTotalFavoriteResponse, err error) {
	err = service.UpdateTotalFavorite(ctx, *req.UserId, *req.ActionType)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (u *UserSrv) UpdateFavoriteCount(ctx context.Context, req *pb.DouyinFavoriteCountRequest) (resp *pb.DouyinFavoriteCountResponse, err error) {
	err = service.UpdateFavoriteCount(ctx, *req.UserId, *req.ActionType)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (u *UserSrv) UpdateFollowCount(ctx context.Context, req *pb.DouyinFollowCountRequest) (resp *pb.DouyinFollowCountResponse, err error) {
	err = service.UpdateFollowCount(ctx, *req.UserId, *req.ActionType)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (u *UserSrv) UpdateFollowerCount(ctx context.Context, req *pb.DouyinFollowerCountRequest) (resp *pb.DouyinFollowerCountResponse, err error) {
	err = service.UpdateFollowerCount(ctx, *req.UserId, *req.ActionType)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
