package service

import (
	"context"
	"douyin/dao"
	"douyin/global"
	"douyin/middleware"
	"douyin/model"
	"douyin/response"
	"douyin/utils"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func CreateComment(userID int64, videoID int64, commentText string, ctx context.Context) (response.Comment_Action_Response, error) {

	ctx, span := global.SERVER_COMMENT_TRACER.Start(ctx, "addcomment service")
	defer span.End()

	// 获取评论时间
	currentTime := time.Now().Unix()
	// 1-发布评论
	// 1.1 创建comment结构体
	comment := model.Comment{
		User_id:      userID,
		Video_id:     videoID,
		Comment:      commentText,
		Publish_time: currentTime,
	}
	// 1.2 将comment增添到数据库中
	tx := dao.BeginTransaction()
	err := dao.CreateComment(&comment)
	if err != nil {
		// 如果发生错误，将数据库回滚到未添加评论的初始状态
		defer dao.RollbackTransaction(tx)
		fmt.Printf("添加评论异常：%s", err)
		global.SERVER_LOG.Error("Failed to create comment:", zap.String("error", err.Error()))

		commentActionResponse := response.Comment_Action_Response{
			Response: response.Response{StatusCode: http.StatusInternalServerError, StatusMsg: "添加评论异常"},
		}
		return commentActionResponse, err

	}
	// 更新视频表评论总数+1
	err = dao.UpdateVideoCommentCount(videoID, 1)
	if err != nil {
		// 如果发生错误，将数据库回滚到未添加评论的初始状态
		defer dao.RollbackTransaction(tx)
		fmt.Printf("更新评论总数异常：%s", err)
		global.SERVER_LOG.Error("Failed to count comment:", zap.String("error", err.Error()))

		commentActionResponse := response.Comment_Action_Response{
			Response: response.Response{StatusCode: http.StatusInternalServerError, StatusMsg: "更新视频评论数异常"},
		}
		return commentActionResponse, err

	}

	// 1.3 创建Comment_Response响应结构体
	createDate := utils.IntTime2CommentTime(currentTime)
	commenter, err := dao.GetUserById(userID)
	if err != nil {
		// 如果发生错误，将数据库回滚到未添加评论的初始状态
		defer dao.RollbackTransaction(tx)
		fmt.Printf("获取用户异常：%s", err)
		global.SERVER_LOG.Error("Failed to fetch user:", zap.String("error", err.Error()))

		commentActionResponse := response.Comment_Action_Response{
			Response: response.Response{StatusCode: http.StatusInternalServerError, StatusMsg: "添加评论异常"},
		}
		return commentActionResponse, err
	}
	// 1.4返回响应
	commentActionResponse := response.Comment_Action_Response{
		Response: response.Response{StatusCode: 0, StatusMsg: "OK"},
		Comment_Response: response.Comment_Response{
			Id: 0,
			Commenter: response.User_Response{
				Id:              commenter.User_id,
				Name:            commenter.User_name,
				FollowCount:     commenter.Follow_count,
				FollowerCount:   commenter.Follower_count,
				IsFollow:        false, // 待确定自己与自己的关注状态
				Avatar:          commenter.Avatar,
				BackgroundImage: commenter.Background_image,
				Signature:       commenter.Signature,
				TotalFavorited:  commenter.Total_favorited,
				WorkCount:       commenter.Work_count,
				FavoriteCount:   commenter.Favorite_count,
			},
			Content:    commentText,
			CreateDate: createDate,
		},
	}
	return commentActionResponse, err
}

func DeleteComment(userID int64, videoID int64, commentID int64, ctx context.Context) (response.Comment_Action_Response, error) {
	ctx, span := global.SERVER_COMMENT_TRACER.Start(ctx, "deletecomment service")
	defer span.End()

	var commentActionResponse response.Comment_Action_Response
	// 2-删除评论
	// 2.1 根据commentID在数据库中找到待删除的评论

	// 2.2 判断是否有权限删除
	// 		2.2.1 通过commentID找到commenterID
	comment, err := dao.GetCommentById(commentID)
	if err != nil {
		fmt.Printf("获取评论异常：%s", err)
		global.SERVER_LOG.Error("Failed to fetch comment:", zap.String("error", err.Error()))

		commentActionResponse = response.Comment_Action_Response{
			Response: response.Response{StatusCode: http.StatusInternalServerError, StatusMsg: "获取评论异常"},
		}
		return commentActionResponse, err
	}
	commenterID := comment.User_id
	// 2.3 若有权限，则删除id为commentID评论;若无权限，则拒绝删除
	if commenterID == userID {
		tx := dao.BeginTransaction()
		err = dao.DeleteCommentById(commentID)
		if err != nil {
			// 如果发生错误，将数据库回滚到未删除评论的初始状态
			defer dao.RollbackTransaction(tx)
			fmt.Printf("删除评论异常：%s", err)
			global.SERVER_LOG.Error("Failed to delete comment:", zap.String("error", err.Error()))

			commentActionResponse = response.Comment_Action_Response{
				Response: response.Response{StatusCode: http.StatusInternalServerError, StatusMsg: "删除评论异常"},
			}

			return commentActionResponse, err
		}
		// 更新视频表评论总数-1
		err = dao.UpdateVideoCommentCount(videoID, -1)
		if err != nil {
			// 如果发生错误，将数据库回滚到未删除评论的初始状态
			defer dao.RollbackTransaction(tx)
			fmt.Printf("更新视频评论数异常：%s", err)
			global.SERVER_LOG.Error("Failed to count comment:", zap.String("error", err.Error()))

			commentActionResponse = response.Comment_Action_Response{
				Response: response.Response{StatusCode: http.StatusInternalServerError, StatusMsg: "更新视频评论数异常"},
			}

			return commentActionResponse, err
		}

		commentActionResponse = response.Comment_Action_Response{
			Response:         response.Response{StatusCode: 0, StatusMsg: "删除成功"},
			Comment_Response: response.Comment_Response{},
		}

	} else {
		commentActionResponse = response.Comment_Action_Response{
			Response:         response.Response{StatusCode: 0, StatusMsg: "无删除权限"},
			Comment_Response: response.Comment_Response{},
		}
	}
	return commentActionResponse, err
}

func GetCommentList(videoID int64, userID int64, ctx context.Context) ([]response.Comment_Response, error) {

	ctx, span := global.SERVER_COMMENT_TRACER.Start(ctx, "getcommentlist service")
	defer span.End()

	// 从数据库中获取id为video_id的全部评论
	comments, err := dao.GetCommentByIdListById(videoID)
	if err != nil {
		err = global.ErrorVideoDataWrong
		return nil, err
	}

	// 将获取到的评论添加到commentList列表中
	// 将model.comment解析为response.Comment_Response格式
	// 将获取到的评论添加到commentList列表中
	var commentList []response.Comment_Response
	// 将model.comment解析为response.Comment_Response格式
	for _, comment := range comments {

		// 根据评论者id构建user_response
		commenter, err := dao.GetUserById(comment.User_id)
		if err != nil {
			// 处理获取用户信息错误
			// 在日志中记录错误信息
			global.SERVER_LOG.Error("Failed to fetch user:", zap.String("error", err.Error()))
			continue // 继续处理下一个评论
		}

		isFollow, err := middleware.GetUserRelationState(userID, commenter.User_id)
		//没有该记录,查询后设置
		if err == global.ErrorCacheMiss {
			isFollow = dao.GetFollowByUserId(userID, commenter.User_id)
			go middleware.SetUserRelation(userID, commenter.User_id, isFollow)
			//redis操作出错 从数据库中查询
		} else if err != nil {
			isFollow = dao.GetFollowByUserId(userID, commenter.User_id)
			global.SERVER_LOG.Warn("redis operation fail!")
		}
		//正常从redis中取出数据

		// 构建Comment_Response中嵌套的User_Response字段
		userResponse := response.User_Response{
			Id:              commenter.User_id,
			Avatar:          commenter.Avatar,
			BackgroundImage: commenter.Background_image,
			FavoriteCount:   commenter.Favorite_count,
			FollowCount:     commenter.Follow_count,
			IsFollow:        isFollow,
			FollowerCount:   commenter.Follower_count,
			Name:            commenter.User_name,
			Signature:       commenter.Signature,
			TotalFavorited:  commenter.Total_favorited,
			WorkCount:       commenter.Work_count,
		}
		//查询该用户是否被关注

		//userResponse.IsFollow = dao.GetFollowByUserId(userID, commenter.User_id)

		commentResponse := response.Comment_Response{
			Id:         comment.Comment_id,
			Content:    comment.Comment,
			CreateDate: utils.IntTime2StrTime(comment.Publish_time),
			Commenter:  userResponse,
		}
		commentList = append(commentList, commentResponse)
	}
	return commentList, nil
}
