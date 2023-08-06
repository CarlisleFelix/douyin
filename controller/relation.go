package controller

import (
	"douyin/response"

	"github.com/gin-gonic/gin"
)

type Relation_Action_Response struct {
	response.Response
}

type Relation_Follow_List_Response struct {
	response.Response
	UserList []response.User_Response `json:"user_list,omitempty"`
}

type Relation_Follower_List_Response struct {
	response.Response
	UserList []response.User_Response `json:"user_list,omitempty"`
}

type Relation_Friend_List_Response struct {
	response.Response
	UserList []response.Friend_Response `json:"user_list,omitempty"`
}

func RelationAction(c *gin.Context) {

	return
}

func FollowList(c *gin.Context) {

	return
}

func FollowerList(c *gin.Context) {

	return
}

func FriendList(c *gin.Context) {

	return
}
