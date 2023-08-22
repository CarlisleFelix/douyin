package model

type User struct {
	User_id          int64  `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY;not null" json:"user_id"`
	User_name        string `gorm:"column:name;type:varchar(32);not null" json:"user_name"`
	Password         string `gorm:"column:password;type:varchar(255);not null" json:"password"` //长度改为255
	Follow_count     int64  `gorm:"column:follow_count;type:int;default:0" json:"follow_count"`
	Follower_count   int64  `gorm:"column:follower_count;type:int;default:0" json:"follower_count"`
	Avatar           string `gorm:"column:avatar;type:varchar(255);default:null" json:"avatar"`
	Background_image string `gorm:"column:background_image;type:varchar(255);default:null" json:"background_image"`
	Signature        string `gorm:"column:signature;type:varchar(255);default:null" json:"signature"`
	Total_favorited  int64  `gorm:"column:total_favorited;type:int;default:0" json:"total_favorited"`
	Work_count       int64  `gorm:"column:work_count;type:int;default:0" json:"work_count"`
	Favorite_count   int64  `gorm:"column:favorite_count;type:int;default:0" json:"favorite_count"`
}
