package controller

import (
	"douyin/dao"
	"douyin/model"
	"douyin/response"
	"douyin/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FriendUser struct {
	response.User_Response
	LatestMessage string `json:"chat"`     // 和该好友的最新聊天消息
	MessageType   int64  `json:"msg_type"` // message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息

}

// RelationAction 关注/取消关注操作
func RelationAction(c *gin.Context) {
	//1.取数据
	//1.1 从token中获取用户id
	// strToken := c.Query("token")
	// tokenStruct, _ := middleware.CheckToken(strToken)
	// hostID := tokenStruct.UserId
	_userId, _ := c.Get("userid")
	hostID, _ := _userId.(int64)
	//1.2 获取待关注的用户id
	getToUserID, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	guestID := int64(getToUserID)
	//1.3 获取关注操作（关注1，取消关注2）
	getActionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)
	actionType := int64(getActionType)

	//2.自己关注/取消关注自己不合法
	if hostID == guestID {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 405,
			StatusMsg:  "不能关注自己",
		})
		c.Abort()
		return
	}

	//3.关注/取关
	err := service.FollowAction(hostID, guestID, actionType)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 0,
			StatusMsg:  "关注/取消关注成功！",
		})
	}
}

// FollowList 获取用户关注列表
func FollowList(c *gin.Context) {

	//1.数据预处理
	//1.1获取用户本人id
	// strToken := c.Query("token")
	// tokenStruct, _ := middleware.CheckToken(strToken)
	// hostID := tokenStruct.UserId
	_userId, _ := c.Get("userid")
	userid, _ := _userId.(int64)
	//1.2获取其他用户id
	getUserId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	user_id := int64(getUserId)

	//2.判断查询类型，从数据库取用户列表
	var err error
	var userList []model.User
	var FuserID int64
	if user_id == 0 {
		//查本人
		FuserID = userid
	} else {
		//查对方
		FuserID = user_id
	}
	userList, err = service.FollowingList(FuserID)
	//构造返回的数据
	var ReturnFollowerList = make([]response.User_Response, len(userList))
	for i, m := range userList {
		ReturnFollowerList[i].Id = m.User_id
		ReturnFollowerList[i].Name = m.User_name
		ReturnFollowerList[i].FollowCount = m.Follow_count
		ReturnFollowerList[i].FollowerCount = m.Follower_count
		ReturnFollowerList[i].Avatar = m.Avatar
		ReturnFollowerList[i].BackgroundImage = m.Background_image
		ReturnFollowerList[i].IsFollow = service.IsFollowing(userid, m.User_id)
		ReturnFollowerList[i].TotalFavorited = m.Total_favorited
		ReturnFollowerList[i].WorkCount = m.Work_count
		ReturnFollowerList[i].FavoriteCount = m.Favorite_count
	}
	fmt.Printf("找到关注表: %+v\n", ReturnFollowerList)

	//3.响应返回
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Relation_Follow_List_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  "查找关注列表失败！",
			},
			UserList: nil,
		})
	} else {
		c.JSON(http.StatusOK, response.Relation_Follow_List_Response{
			Response: response.Response{
				StatusCode: 0,
				StatusMsg:  "已找到关注列表！",
			},
			UserList: ReturnFollowerList,
		})
	}
}

// FollowerList 获取用户粉丝列表
func FollowerList(c *gin.Context) {

	//1.数据预处理
	//1.1获取用户本人id
	// strToken := c.Query("token")
	// tokenStruct, _ := middleware.CheckToken(strToken)
	// hostID := tokenStruct.UserId
	_userId, _ := c.Get("userid")
	userid, _ := _userId.(int64)
	//1.2获取其他用户id
	getUserId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	user_id := int64(getUserId)
	var FuserID int64
	//2.判断查询类型
	var err error
	var userList []model.User
	if user_id == 0 {
		//查本人的粉丝表
		FuserID = userid
	} else {
		//查对方的粉丝表
		FuserID = user_id
	}
	userList, err = service.FollowerList(FuserID)
	//3.判断查询类型，从数据库取用户列表
	var ReturnFollowerList = make([]response.User_Response, len(userList))
	for i, m := range userList {
		ReturnFollowerList[i].Id = m.User_id
		ReturnFollowerList[i].Name = m.User_name
		ReturnFollowerList[i].FollowCount = m.Follow_count
		ReturnFollowerList[i].FollowerCount = m.Follower_count
		ReturnFollowerList[i].Avatar = m.Avatar
		ReturnFollowerList[i].BackgroundImage = m.Background_image
		ReturnFollowerList[i].IsFollow = service.IsFollowing(userid, m.User_id)
		ReturnFollowerList[i].TotalFavorited = m.Total_favorited
		ReturnFollowerList[i].WorkCount = m.Work_count
		ReturnFollowerList[i].FavoriteCount = m.Favorite_count
	}
	fmt.Printf("找到粉丝表: %+v\n", ReturnFollowerList)

	//3.处理
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Relation_Follower_List_Response{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  "查找粉丝列表失败！",
			},
			UserList: nil,
		})
	} else {
		c.JSON(http.StatusOK, response.Relation_Follower_List_Response{
			Response: response.Response{
				StatusCode: 0,
				StatusMsg:  "已找到粉丝列表！",
			},
			UserList: ReturnFollowerList,
		})
	}
}

type FriendListResponse struct {
	response.Response
	FriendList []FriendUser `json:"user_list"` // 用户列表
}

// / FriendList 好友列表
// 获取朋友列表，并且会带着和该用户的最新的一条消息
func FriendList(c *gin.Context) {
	// 取 token
	// token := c.Query("token")
	// tokenStruct, _ := middleware.CheckToken(token)

	// //获取用户本人id
	// hostId := tokenStruct.UserId
	_userId, _ := c.Get("userid")
	userid, _ := _userId.(int64)
	//1.2获取其他用户id
	getUserId, _ := strconv.ParseInt(c.Query("user_id "), 10, 64)
	user_id := int64(getUserId)
	//2.判断查询类型
	var err error
	var tmpFriendList []model.User
	var FuserID int64
	if user_id == 0 {
		//查本人的好友表
		FuserID = userid
	} else {
		//查对方的好友表
		FuserID = user_id
	}
	tmpFriendList, err = service.FriendList(FuserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, FriendListResponse{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  "查询好友列表失败",
			},
			FriendList: nil,
		})
		return
	} else {
		// 对返回列表二次加工
		var returnFriendList []FriendUser
		for _, u := range tmpFriendList {
			var msg string
			var msgType int64
			latestMsg1, msgType1, err1 := dao.GetLatestChat(FuserID, u.User_id)
			latestMsg2, msgType2, err2 := dao.GetLatestChat(u.User_id, FuserID)
			if err1 != nil && err2 != nil {
				msg = ""
				msgType = -1
			} else if err1 != nil {
				msg = latestMsg2.Content
				msgType = msgType2
			} else if err2 != nil {
				msg = latestMsg1.Content
				msgType = msgType1
			} else {
				if (latestMsg1.Publish_time) < (latestMsg2.Publish_time) {
					msg = latestMsg1.Content
					msgType = msgType1
				} else {
					msg = latestMsg2.Content
					msgType = msgType2
				}
			}
			curFriend := FriendUser{
				User_Response: response.User_Response{
					Id:              u.User_id,
					Name:            u.User_name,
					FollowCount:     u.Follow_count,
					FollowerCount:   u.Follower_count,
					IsFollow:        service.IsFollowing(userid, u.User_id),
					FavoriteCount:   u.Favorite_count,
					Signature:       u.Signature,
					WorkCount:       u.Work_count,
					Avatar:          u.Avatar,
					TotalFavorited:  u.Total_favorited,
					BackgroundImage: u.Background_image,
				},
				LatestMessage: msg,
				MessageType:   msgType, // 0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
			}
			returnFriendList = append(returnFriendList, curFriend)
		}
		c.JSON(http.StatusOK, FriendListResponse{
			Response: response.Response{
				StatusCode: 0,
				StatusMsg:  "查询好友列表成功",
			},
			FriendList: returnFriendList,
		})
		return
	}
}
