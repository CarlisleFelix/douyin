package utils

import (
	"errors"
	"math/rand"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
)

//@function: PathExists
//@description: 文件目录是否存在
//@param: path string
//@return: bool, error

func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("file with same name")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 新建文件夹
func Mkdir(path string) error {
	if ok, _ := PathExists(path); ok {
		return nil
	}
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// 生成16位随机字符串
func RandomString() string {
	var letters = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
	result := make([]byte, 16)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// 生成人可读的时间字符串,精确到秒
func CurrentTimeStr() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

// 生成时间int64，精确到秒
func CurrentTimeInt() int64 {
	return time.Now().Unix()
}

// 字符串时间到int64时间
func StrTime2IntTime(strTime string) int64 {
	template := "2006-01-02 15:04:05"
	middleForm, err := time.ParseInLocation(template, strTime, time.Local)
	if err != nil {
		return -1
	}
	return middleForm.Unix()
}

// int64时间到字符串时间
func IntTime2StrTime(intTime int64) string {
	template := "2006-01-02 15:04:05"
	return time.Unix(intTime, 0).Format(template)
}

func IntTime2CommentTime(intTime int64) string {
	return ""
}

func StrTime2CommentTime(strTime int64) string {
	return ""
}

func GetCommentTime() string {
	return ""
}

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
