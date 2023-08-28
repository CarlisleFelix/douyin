package controller

import (
	"douyin/global"
	"douyin/response"
	"douyin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CommentAction(c *gin.Context) {

	ctx, span := global.SERVER_COMMENT_TRACER.Start(c.Request.Context(), "commentaction controller")
	defer span.End()

	// 获取请求参数
	videoID, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		// 处理videoID解析错误
		c.JSON(http.StatusBadRequest, response.Comment_Action_Response{
			Response: response.Response{StatusCode: http.StatusBadRequest, StatusMsg: "无效的video_id"},
		})
		return
	}
	actionType := c.Query("action_type")
	commentText := c.Query("comment_text")

	value, _ := c.Get("userid")
	userID, ok := value.(int64)
	if !ok {
		// 处理userid类型断言失败的情况
		c.JSON(http.StatusBadRequest, response.Comment_Action_Response{
			Response: response.Response{StatusCode: http.StatusBadRequest, StatusMsg: "无效的userid"},
		})
		return
	}

	// 判断操作类型
	if actionType == "1" {
		commentActionResponse, _ := service.CreateComment(userID, videoID, commentText, ctx)
		c.JSON(http.StatusOK, commentActionResponse)

	} else if actionType == "2" {
		commentID, err := strconv.ParseInt(c.Query("comment_id"), 10, 64)
		if err != nil {
			// 处理commentID解析错误
			c.JSON(http.StatusBadRequest, response.Comment_Action_Response{
				Response: response.Response{StatusCode: http.StatusBadRequest, StatusMsg: "无效的comment_id"},
			})
			return
		}
		commentActionResponse, _ := service.DeleteComment(userID, videoID, commentID, ctx)
		c.JSON(http.StatusOK, commentActionResponse)
	}

	return
}

func CommentList(c *gin.Context) {

	ctx, span := global.SERVER_COMMENT_TRACER.Start(c.Request.Context(), "commentlist controller")
	defer span.End()

	videoID, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		// 处理videoID解析错误
		c.JSON(http.StatusBadRequest, response.Comment_Action_Response{
			Response: response.Response{StatusCode: http.StatusBadRequest, StatusMsg: "无效的video_id"},
		})
		return
	}
	value, _ := c.Get("userid")
	userID, ok := value.(int64)
	if !ok {
		// 处理userid类型断言失败的情况
		c.JSON(http.StatusBadRequest, response.Comment_Action_Response{
			Response: response.Response{StatusCode: http.StatusBadRequest, StatusMsg: "无效的userid"},
		})
		return
	}

	// 将获取到的评论添加到commentList列表中
	commentList, _ := service.GetCommentList(videoID, userID, ctx)
	// 返回response
	c.JSON(http.StatusOK,
		response.Comment_List_Response{
			Response:    response.Response{StatusCode: 0, StatusMsg: "OK"},
			CommentList: commentList,
		})
	return
}
