// 没有“成功”的mesage提示
// global增加了success文件
// id都被我换成int64类型了，是随机数吗，是随机字符串吗
package service

import (
	"douyin/dao"
	"douyin/global"
	"douyin/middleware"
	"douyin/model"
	"douyin/response"
	//"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("wenjiandakai") // 把秘钥放到txt文件，读取文件内容。避免秘钥明文显示

// Login 函数用于处理用户登录逻辑
func Login(info map[string]string) response.User_Login_Response {

	// 解析数据 获取 map 中的用户名和密码
	username := info["username"]
	password := info["password"]

	// 定义一个 User 变量来存储查询结果
	user, err := dao.GetUserByUsername(username) //只要传递用户的名，查询是否存在
	if err != nil {
		return response.User_Login_Response{
			Response: response.Response{
				StatusCode: 401,                      // 设置响应状态为错误
				StatusMsg:  global.ErrorUserNotExist, // 设置错误消息
			},
			UserId: -1, // 设置为空表示没有有效的 UserId
			Token:  "", // 设置为空表示没有有效的 Token
		}
	}

	// 用户名存在，获取用户信息。其中的password被使用

	// 使用 bcrypt.CompareHashAndPassword 检查密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		// 有错
		return response.User_Login_Response{
			Response: response.Response{
				StatusCode: 402,
				StatusMsg:  global.ErrorPasswordFalse,
			},
			UserId: user.User_id, // 设置 UserId
			Token:  "",           // 设置为空表示没有有效的 Token
		}
	}

	// 密码正确，生成token，成功登录

	// 使用 GenerateToken 函数生成 token
	token, err := middleware.GenerateToken(user.User_id, user.User_name)
	if err != nil {
		return response.User_Login_Response{
			Response: response.Response{
				StatusCode: 403,
				StatusMsg:  global.ErrorTokenCreatedWrong,
			},
			UserId: user.User_id,
			Token:  "",
		}
	}

	return response.User_Login_Response{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  global.SucLogin,
		},
		UserId: user.User_id,
		Token:  token,
	}
}

// Register 函数用于处理用户注册逻辑
func Register(info map[string]string) response.User_Register_Response {
	// 从请求信息中获取各项注册信息
	username := info["username"]
	password := info["password"]

	// checkPassword := info["checkPassword"]	// 不知道能不能加这个
	// // 检查两次输入的密码是否一致
	// if password != checkPassword {
	// 	return map[string]interface{}{"msg": "两次密码输入不同,请重新注册！"}, 405
	// }

	// 检查用户名是否已被注册
	user, err := dao.GetUserByUsername(username) //只要传递用户的名，查询是否存在
	if user != nil {
		return response.User_Register_Response{
			Response: response.Response{
				StatusCode: 404,                   // 设置响应状态为错误
				StatusMsg:  global.ErrorUserExist, // 被注册
			},
			UserId: -1, // 设置为空表示没有有效的 UserId
			Token:  "", // 设置为空表示没有有效的 Token
		}
	}

	// 用户名可注册，获取用户password存储.现在user是nil

	// 使用 bcrypt 对密码进行哈希处理
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// 使用 GenerateToken 函数生成 token
	token, err := middleware.GenerateToken(user.User_id, user.User_name)
	if err != nil {
		return response.User_Register_Response{
			Response: response.Response{
				StatusCode: 403,
				StatusMsg:  global.ErrorTokenCreatedWrong,
			},
			UserId: user.User_id,
			Token:  "",
		}
	}

	// 创建新用户对象并保存到数据库
	user = &model.User{ //Q：不知道为什么要加&，不加不行
		User_name: username,
		Password:  string(hashedPassword), //！notice：userid应该会自增吧，不会的话要设置
	}
	// 调用 CreateUser 函数保存用户信息到数据库
	err2 := dao.CreateUser(global.SERVER_DB, user)
	if err2 != nil {
		return response.User_Register_Response{
			Response: response.Response{
				StatusCode: 405,                  // 设置响应状态为错误
				StatusMsg:  global.ErrorDatabase, // 设置错误消息
			},
			UserId: -1,    // 设置为空表示没有有效的 UserId
			Token:  token, // 设置为空表示没有有效的 Token
		}
	}

	return response.User_Register_Response{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  global.SucRegister,
		},
		UserId: user.User_id,
		Token:  token,
	}
}

// UserInformation 处理查询用户信息的请求
func UserInformation(info map[string]string) response.User_Interface_Response {

	// 取出token
	token := info["token"]

	// 检查 token 是否有效
	account, flag := middleware.CheckToken(token)

	// 无效
	if !flag {
		return response.User_Interface_Response{
			Response: response.Response{
				StatusCode: 406,                      // 设置响应状态为错误
				StatusMsg:  global.ErrorTokenExpired, // 设置错误消息
			},
		}
	}
	// 有效

	// 解码成功，可以从 claims 中获取用户信息
	//userid := account.UserId
	username := account.UserName

	// 用户是否存在
	user, err := dao.GetUserByUsername(username) //只要传递用户的名，查询是否存在
	if err != nil {                              //用户不存在
		return response.User_Interface_Response{
			Response: response.Response{
				StatusCode: 401,                      // 设置响应状态为错误
				StatusMsg:  global.ErrorUserNotExist, // 不存在
			},
		}
	}

	//user用户存在
	return response.User_Interface_Response{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  global.SucInterface,
		},
		User_Response: ConvertToUserResponse(*user),
	}

}

// 根据user信息构造userresponse结构体
func ConvertToUserResponse(user model.User) response.User_Response {
	userResponse := response.User_Response{
		Id:              user.User_id,
		Name:            user.User_name,
		FollowCount:     user.Follow_count,
		FollowerCount:   user.Follower_count,
		Avatar:          user.Avatar,
		BackgroundImage: user.Background_image,
		Signature:       user.Signature,
		TotalFavorited:  user.Total_favorited,
		WorkCount:       user.Work_count,
		FavoriteCount:   user.Favorite_count,
	}
	return userResponse
}
