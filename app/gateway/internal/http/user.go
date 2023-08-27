package http

import (
	"douyin/app/gateway/rpc"
	pb "douyin/idl/pb/user"
	"douyin/utils/ctl"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UserLogin(ctx *gin.Context) {
	var req pb.DouyinUserLoginRequest
	// todo:参数绑定
	username := ctx.Query("user_name")
	password := ctx.Query("password")

	//if err := ctx.Bind(&req); err != nil {
	//	ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "绑定参数错误"))
	//	return
	//}
	req.Username = &username
	req.Password = &password

	// 调用rpc
	r, err := rpc.UserLogin(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "FavoriteAction RPC服务调用错误"))
		return
	}

	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, r))
}

func UserRegister(ctx *gin.Context) {
	var req pb.DouyinUserRegisterRequest
	// todo:参数绑定
	username := ctx.Query("user_name")
	password := ctx.Query("password")

	//if err := ctx.Bind(&req); err != nil {
	//	ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "绑定参数错误"))
	//	return
	//}
	req.Username = &username
	req.Password = &password

	// 调用rpc
	r, err := rpc.UserRegister(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "FavoriteAction RPC服务调用错误"))
		return
	}

	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, r))
}

func GetUserInfo(ctx *gin.Context) {
	var req pb.DouyinUserRequest
	// 参数绑定
	hostId := ctx.Query("user_id")
	// todo:从token中解析出user_id
	var userId int64 = 3

	// 将参数值转换为int64类型
	hostId64, err := strconv.ParseInt(hostId, 10, 64)
	if err != nil {
		// 转换失败，处理错误情况
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter"})
		return
	}
	req.UserId = &userId
	req.HostUserId = &hostId64

	r, err := rpc.GetUserInfo(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "GetUserInfo RPC服务调用错误"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, r))

}
