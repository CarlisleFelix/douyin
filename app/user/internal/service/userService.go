package service

import (
	"context"
	"douyin/app/user/internal/dal/dao"
	"douyin/app/user/internal/dal/model"
	"douyin/app/video/utils"
	pb "douyin/idl/pb/user"
	"douyin/response"
	"douyin/utils/e"
	"gorm.io/gorm"
)

const (
	MaxUsernameLength = 32 //用户名最大长度
	MaxPasswordLength = 32 //密码最大长度
	MinPasswordLength = 6  //密码最小长度
)

func UserRegisterService(ctx context.Context, userName string, passWord string) (model.User, error) {
	//准备参数
	hashedPwd, err := utils.Hash(passWord)
	newUser := model.User{
		User_name: userName,
		Password:  hashedPwd,
	}

	//检查参数
	err = CheckUserParam(ctx, userName, passWord)
	if err != nil {
		return newUser, err
	}

	//检查参数
	exist := UserNameExists(ctx, userName)
	if exist == true {
		return newUser, e.ErrorUserExist
	}

	//创建用户

	if err = dao.NewUserDao(ctx).InsertUser(&newUser); err != nil {
		return newUser, err
	}

	//返回
	return newUser, err
}

func CheckUserParam(ctx context.Context, userName string, passWord string) error {
	if userName == "" {
		return e.ErrorUserNameNull
	}
	if len(userName) > MaxUsernameLength {
		return e.ErrorUserNameExtend
	}
	if len(passWord) > MaxPasswordLength || len(passWord) < MinPasswordLength {
		return e.ErrorPasswordLength
	}
	return nil
}

func UserNameExists(ctx context.Context, userName string) bool {
	_, err := dao.NewUserDao(ctx).GetUserByName(userName)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		} else {
			return true
		}
	}
	return true
}

func UserLoginService(ctx context.Context, userName string, passWord string) (model.User, error) {
	loginUser := model.User{
		User_name: userName,
		Password:  passWord,
	}

	//检查参数
	err := CheckUserParam(ctx, userName, passWord)
	if err != nil {
		return loginUser, err
	}

	//检查用户是否存在
	exist := UserNameExists(ctx, userName)
	if exist == false {
		return loginUser, e.ErrorUserNotExist
	}

	//检查密码是否一致
	passwordMatch := CheckUserPassword(ctx, userName, passWord, &loginUser)
	if !passwordMatch {
		return loginUser, e.ErrorPasswordFalse
	}

	//返回
	return loginUser, err
}

func CheckUserPassword(ctx context.Context, userName string, passWord string, loginUser *model.User) bool {
	tmpUser, err := dao.NewUserDao(ctx).GetUserByName(loginUser.User_name)
	*loginUser = tmpUser
	if err != nil || !utils.Compare(loginUser.Password, passWord) {
		return false
	}
	return true
}

func UserService(ctx context.Context, queryUserId int64, hostUserId int64) (response.User_Response, error) {
	userResponse := response.User_Response{}
	queryUser, err := dao.NewUserDao(ctx).GetUserById(queryUserId)
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

func GetUserInfo(ctx context.Context, user_id int64) (*pb.User, error) {
	var user model.User
	user, err := dao.NewUserDao(ctx).GetUserById(user_id)
	userResp := pb.User{
		Id:              &user.User_id,
		Name:            &user.User_name,
		FollowCount:     &user.Follow_count,
		FollowerCount:   &user.Follower_count,
		Avatar:          &user.Avatar,
		BackgroundImage: &user.Background_image,
		Signature:       &user.Signature,
		TotalFavorited:  &user.Total_favorited,
		WorkCount:       &user.Work_count,
		FavoriteCount:   &user.Favorite_count,
	}
	if err != nil {
		return &userResp, err
	}
	return &userResp, nil
}

func UpdateTotalFavorite(ctx context.Context, user_id int64, action_type int32) error {
	err := dao.NewUserDao(ctx).UpdateUserFavoriteCount(user_id, action_type)
	return err
}

func UpdateFavoriteCount(ctx context.Context, user_id int64, action_type int32) error {
	err := dao.NewUserDao(ctx).UpdateUserTotalFavorite(user_id, action_type)
	return err
}

func UpdateFollowCount(ctx context.Context, user_id int64, action_type int32) error {
	err := dao.NewUserDao(ctx).UpdateUserFollowerCount(user_id, action_type)
	return err
}

func UpdateFollowerCount(ctx context.Context, user_id int64, action_type int32) error {
	err := dao.NewUserDao(ctx).UpdateUserFollowerCount(user_id, action_type)
	return err
}
