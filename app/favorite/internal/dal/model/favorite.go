package model

type Favorite struct {
	Favorite_id int64 `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY;not null" json:"favorite_id"`
	User_id     int64 `gorm:"column:user_id;type:int;not null" json:"user_id"`
	Video_id    int64 `gorm:"column:video_id;type:int;not null" json:"video_id"`
}
