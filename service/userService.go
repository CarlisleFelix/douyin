package service

import (
	"douyin/dao"
	"douyin/global"
	"douyin/middleware"
	"douyin/model"
	"douyin/response"
	"douyin/utils"

	"gorm.io/gorm"
)

const (
	MaxUsernameLength = 32 //用户名最大长度
	MaxPasswordLength = 32 //密码最大长度
	MinPasswordLength = 6  //密码最小长度
)

func UserRegisterService(userName string, passWord string) (model.User, error) {
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

func UserLoginService(userName string, passWord string) (model.User, error) {
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

func UserService(queryUserId int64, hostUserId int64) (response.User_Response, error) {
	userResponse := response.User_Response{}
	queryUser, err := dao.GetUserById(queryUserId)
	//先从redis里面取出
	isFollow, err := middleware.GetUserRelationState(hostUserId, queryUserId)
	//没有该记录,查询后设置
	if err == global.ErrorCacheMiss {
		isFollow = dao.GetFollowByUserId(hostUserId, queryUserId)
		//fmt.Printf("isfollow:%+v\n", isFollow)
		go middleware.SetUserRelation(hostUserId, queryUserId, isFollow)
		//redis操作出错 从数据库中查询
	} else if err != nil {
		isFollow = dao.GetFollowByUserId(hostUserId, queryUserId)
		global.SERVER_LOG.Warn("redis operation fail!")
	}

	//isFollow, err = middleware.GetUserRelationState(hostUserId, queryUserId)
	//fmt.Printf("isfollow:%+v\n", isFollow)
	//fmt.Printf("err:%+v", err.Error())

	//正常从redis中取出数据
	userResponse = response.User_Response{
		Id:              queryUser.User_id,
		Name:            queryUser.User_name,
		FollowCount:     queryUser.Follow_count,
		FollowerCount:   queryUser.Follower_count,
		IsFollow:        isFollow,
		Avatar:          queryUser.Avatar,
		BackgroundImage: queryUser.Background_image,
		Signature:       queryUser.Signature,
		TotalFavorited:  queryUser.Favorite_count,
		WorkCount:       queryUser.Work_count,
		FavoriteCount:   queryUser.Favorite_count,
	}
	return userResponse, nil
}
