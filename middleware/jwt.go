package middleware

import (
	"douyin/response"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 私匙
var Private_Key = []byte("douyin sucks")

// token中payload内容
type MyClaims struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

// 根据用户id和用户名生成token
func GenerateToken(userId int64, userName string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add((24 * time.Hour))
	claims := MyClaims{
		UserId:   userId,
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  nowTime.Unix(),
			Issuer:    "awesome guys",
			Subject:   "userToken",
		},
	}
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return tokenStruct.SignedString(Private_Key)
}

// 根据token进行校验
func CheckToken(token string) (*MyClaims, bool) {
	tokenObj, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Private_Key, nil
	})
	if key, _ := tokenObj.Claims.(*MyClaims); tokenObj.Valid {
		return key, true
	} else {
		return nil, false
	}
}

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//先查看url里面有无token传参
		tokenStr := c.Query("token")
		//再查看body里面有无token传参
		if tokenStr == "" {
			tokenStr = c.PostForm("token")
		}
		//无token
		if tokenStr == "" {
			c.JSON(http.StatusOK, response.Response{StatusCode: 401, StatusMsg: "用户不存在"})
			c.Abort() //阻止执行
			return
		}
		//验证token
		tokenStruck, ok := CheckToken(tokenStr)
		//如果token无效
		if !ok {
			c.JSON(http.StatusOK, response.Response{
				StatusCode: 403,
				StatusMsg:  "token incorrect",
			})
			c.Abort() //阻止执行
			return
		}
		//如果token超时
		if time.Now().Unix() > tokenStruck.ExpiresAt {
			c.JSON(http.StatusOK, response.Response{
				StatusCode: 402,
				StatusMsg:  "token过期",
			})
			c.Abort() //阻止执行
			return
		}
		//fmt.Println(tokenStruck.UserName)
		//fmt.Println(tokenStruck.UserId)
		c.Set("username", tokenStruck.UserName)
		c.Set("userid", tokenStruck.UserId)
		c.Next()
	}
}
