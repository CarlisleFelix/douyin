package controller

import (
	"context"
	"douyin/app/user/internal/service"
	pb "douyin/idl/pb/user"
	"douyin/utils/e"
	"fmt"
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
	return
}

func (u *UserSrv) UserRegister(ctx context.Context, req *pb.DouyinUserRegisterRequest) (resp *pb.DouyinUserRegisterResponse, err error) {
	return
}

func (u *UserSrv) GetUserInfo(ctx context.Context, req *pb.DouyinUserRequest) (resp *pb.DouyinUserResponse, err error) {
	fmt.Println("开始调用8002……")
	//参数处理
	queryUserId := req.UserId
	hostUserId := req.HostUserId

	if err != nil {
		msg := err.Error()
		resp = &pb.DouyinUserResponse{
			StatusCode: &e.FAILED_CODE,
			StatusMsg:  &msg,
		}
		//global.SERVER_LOG.Warn("param format wrong!", zap.String("error", err.Error()))
		return
	}

	//service层处理
	userResponse, err := service.UserService(ctx, *queryUserId, *hostUserId)

	if err != nil {
		msg := err.Error()
		resp = &pb.DouyinUserResponse{
			StatusCode: &e.FAILED_CODE,
			StatusMsg:  &msg,
		}
		fmt.Println("servive层调用错误")

		//global.SERVER_LOG.Warn("UserService fail!", zap.String("error", err.Error()))
		return
	}

	msg := "查询成功"
	//返回
	resp = &pb.DouyinUserResponse{
		StatusCode: &e.SUCCESS_CODE,
		StatusMsg:  &msg,
		User:       userResponse,
	}

	//global.SERVER_LOG.Info("User Success!")
	fmt.Println(msg)
	return

}

func (u *UserSrv) UpdateTotalFavorite(ctx context.Context, req *pb.DouyinTotalFavoriteRequest) (resp *pb.DouyinTotalFavoriteResponse, err error) {
	return
}

func (u *UserSrv) UpdateFollowCount(ctx context.Context, req *pb.DouyinFollowCountRequest) (resp *pb.DouyinFollowCountResponse, err error) {
	return
}

func (u *UserSrv) UpdateFollowerCount(ctx context.Context, req *pb.DouyinFollowerCountRequest) (resp *pb.DouyinFollowerCountResponse, err error) {
	return
}
