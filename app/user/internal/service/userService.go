package service

import (
	"context"
	"douyin/app/user/internal/dal/dao"
	"douyin/app/user/internal/dal/model"
	pb "douyin/idl/pb/user"
	"douyin/utils/e"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

const (
	MaxUsernameLength = 32 //用户名最大长度
	MaxPasswordLength = 32 //密码最大长度
	MinPasswordLength = 6  //密码最小长度
)

// 加密密码
func Hash(passWord string) (pwdHash string, err error) {
	pwd := []byte(passWord)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return
	}
	pwdHash = string(hash)
	return
}

// 比较加密的密码和输入的密码
func Compare(hashedPwd string, passWord string) bool {
	byteHash := []byte(hashedPwd)
	bytePwd := []byte(passWord)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	if err != nil {
		return false
	}
	return true
}

//func UserRegisterService(userName string, passWord string) (model.User, error) {
//	//准备参数
//	hashedPwd, err := Hash(passWord)
//	newUser := model.User{
//		User_name: userName,
//		Password:  hashedPwd,
//	}
//
//	//检查参数
//	err = CheckUserParam(userName, passWord)
//	if err != nil {
//		return newUser, err
//	}
//
//	//检查参数
//	exist := UserNameExists(userName)
//	if exist == true {
//		return newUser, e.ErrorUserExist
//	}
//
//	//创建用户
//
//	if err = dao.InsertUser(&newUser); err != nil {
//		return newUser, err
//	}
//
//	//返回
//	return newUser, err
//}

func CheckUserParam(userName string, passWord string) error {
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

func UserNameExists(userName string) bool {
	//_, err := dao.GetUserByName(userName)
	//if err != nil {
	//	if err == gorm.ErrRecordNotFound {
	//		return false
	//	} else {
	//		return true
	//	}
	//}
	return true
}

func UserLoginService(ctx context.Context, userName string, passWord string) (model.User, error) {
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
		return loginUser, e.ErrorUserExist
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
	if err != nil || !Compare(loginUser.Password, passWord) {
		return false
	}
	return true
}

func UserService(ctx context.Context, queryUserId int64, hostUserId int64) (resp *pb.User, err error) {

	queryUser, err := dao.NewUserDao(ctx).GetUserById(queryUserId)

	// todo:待处理
	//isFollow := dao.GetFollowByUserId(hostUserId, queryUserId)
	isFollow := false
	fmt.Println("isFollow:", isFollow)
	if err != nil {
		return
	}
	resp = &pb.User{
		Id:              &queryUser.User_id,
		Name:            &queryUser.User_name,
		FollowCount:     &queryUser.Follow_count,
		FollowerCount:   &queryUser.Follower_count,
		IsFollow:        &isFollow,
		Avatar:          &queryUser.Avatar,
		BackgroundImage: &queryUser.Background_image,
		Signature:       &queryUser.Signature,
		TotalFavorited:  &queryUser.Favorite_count,
		WorkCount:       &queryUser.Work_count,
		FavoriteCount:   &queryUser.Favorite_count,
	}
	return
}
