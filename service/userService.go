package service

import (
	"context"
	"douyin/dao"
	"douyin/global"
	"douyin/model"
	"douyin/response"
	"douyin/utils"
	"fmt"

	"gorm.io/gorm"
)

const (
	MaxUsernameLength = 32 //用户名最大长度
	MaxPasswordLength = 32 //密码最大长度
	MinPasswordLength = 6  //密码最小长度
)

func UserRegisterService(userName string, passWord string, ctx context.Context) (model.User, error) {

	ctx, span := global.SERVER_USER_TRACER.Start(ctx, "userregister service")
	defer span.End()

	//准备参数
	hashedPwd, err := utils.Hash(passWord)
	newUser := model.User{
		User_name: userName,
		Password:  hashedPwd,
	}

	//检查参数
	err = CheckUserParam(userName, passWord)
	if err != nil {
		return newUser, err
	}

	//检查参数
	exist := UserNameExists(userName)
	if exist == true {
		return newUser, global.ErrorUserExist
	}

	//创建用户

	if err = dao.InsertUser(&newUser); err != nil {
		return newUser, err
	}

	//返回
	return newUser, err
}

func CheckUserParam(userName string, passWord string) error {
	if userName == "" {
		return global.ErrorUserNameNull
	}
	if len(userName) > MaxUsernameLength {
		return global.ErrorUserNameExtend
	}
	if len(passWord) > MaxPasswordLength || len(passWord) < MinPasswordLength {
		return global.ErrorPasswordLength
	}
	return nil
}

func UserNameExists(userName string) bool {
	_, err := dao.GetUserByName(userName)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		} else {
			return true
		}
	}
	return true
}

func UserLoginService(userName string, passWord string, ctx context.Context) (model.User, error) {

	ctx, span := global.SERVER_USER_TRACER.Start(ctx, "userlogin service")
	defer span.End()

	loginUser := model.User{
		User_name: userName,
		Password:  passWord,
	}

	//检查参数
	err := CheckUserParam(userName, passWord)
	if err != nil {
		return loginUser, err
	}

	//检查用户是否存在
	exist := UserNameExists(userName)
	if exist == false {
		return loginUser, global.ErrorUserNotExist
	}

	//检查密码是否一致
	passwordMatch := CheckUserPassword(userName, passWord, &loginUser)
	if !passwordMatch {
		return loginUser, global.ErrorPasswordFalse
	}

	//返回
	return loginUser, err
}

func CheckUserPassword(userName string, passWord string, loginUser *model.User) bool {
	tmpUser, err := dao.GetUserByName(loginUser.User_name)
	*loginUser = tmpUser
	if err != nil || !utils.Compare(loginUser.Password, passWord) {
		return false
	}
	return true
}

func UserService(queryUserId int64, hostUserId int64, ctx context.Context) (response.User_Response, error) {

	ctx, span := global.SERVER_USER_TRACER.Start(ctx, "user service")
	defer span.End()

	userResponse := response.User_Response{}
	queryUser, err := dao.GetUserById(queryUserId)
	isFollow := dao.GetFollowByUserId(hostUserId, queryUserId)
	fmt.Println("isFolllow:", isFollow)
	if err != nil {
		return userResponse, err
	}
	userResponse = response.User_Response{
		Id:              queryUser.User_id,
		Name:            queryUser.User_name,
		FollowCount:     queryUser.Follow_count,
		FollowerCount:   queryUser.Follower_count,
		IsFollow:        true,
		Avatar:          queryUser.Avatar,
		BackgroundImage: queryUser.Background_image,
		Signature:       queryUser.Signature,
		TotalFavorited:  queryUser.Favorite_count,
		WorkCount:       queryUser.Work_count,
		FavoriteCount:   queryUser.Favorite_count,
	}
	return userResponse, err
}
