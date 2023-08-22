package dao

import (
	"douyin/global"
	"douyin/model"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	//"gorm.io/gorm"//和上一个global必须一起导入是吧
)

// GetUserByUsername 根据用户名从数据库中查询用户信息。
// 参数：
//
//	db: *gorm.DB 数据库连接对象，用于执行查询操作。
//	username: string 要查询的用户名。
//
// 返回值：
//
//	*model.User: 查询到的用户信息，如果不存在则为 nil。
//	error: 查询过程中的错误，如果查询成功则为 nil。
func GetUserByUsername(username string) (*model.User, error) {
	var user model.User

	// 使用给定的数据库连接查询指定用户名的用户信息
	err := global.SERVER_DB.Where("name = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 如果未找到记录，返回自定义错误
			return nil, fmt.Errorf("user not found")
		}
		// 如果发生其他错误，返回错误
		return nil, err
	}

	// 返回查询到的用户信息和 nil 错误
	return &user, nil
}

// 保存user到数据库
func CreateUser(db *gorm.DB, user *model.User) error {
	// 在这里执行插入操作
	err := db.Create(user).Error
	if err != nil {
		fmt.Println("Error while creating user:", err)
	} //notice：加了这两句0822
	return err
}
