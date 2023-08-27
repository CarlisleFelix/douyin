package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"douyin/app/gateway/rpc"
	pb "douyin/idl/pb/favorite"
	"douyin/utils/ctl"
)

// FavoriteAction 点赞操作 处理参数 并调用rpc
func FavoriteAction(ctx *gin.Context) {
	var req pb.DouyinFavoriteActionRequest
	// todo:参数绑定
	token := ctx.Query("token")
	video_id := ctx.Query("video_id")
	// 将参数值转换为int64类型
	video_id_64, err := strconv.ParseInt(video_id, 10, 64)
	if err != nil {
		// 转换失败，处理错误情况
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter"})
		return
	}
	action_type := ctx.Query("action_type")
	// 将参数值转换为int32类型
	action_type_32, err := strconv.ParseInt(action_type, 10, 64)
	if err != nil {
		// 转换失败，处理错误情况
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter"})
		return
	}
	a := int32(action_type_32)

	//if err := ctx.Bind(&req); err != nil {
	//	ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "绑定参数错误"))
	//	return
	//}
	req.Token = &token
	req.VideoId = &video_id_64
	req.ActionType = &a

	// 调用rpc
	r, err := rpc.FavoriteAction(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "FavoriteAction RPC服务调用错误"))
		return
	}

	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, r))
}

// FavoriteList 获取点赞列表
func FavoriteList(ctx *gin.Context) {
	var req pb.DouyinFavoriteListRequest
	r, err := rpc.FavoriteList(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "UserRegister RPC服务调用错误"))
		return
	}

	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, r))
}
