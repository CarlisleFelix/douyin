package model

type Comment struct {
	Comment_id   int64  `gorm:"column:id;type:bigint AUTO_INCREMENT;not null;primary_key" json:"comment_id"`
	User_id      int64  `gorm:"column:user_id;type:bigint;not null" json:"user_id"`
	Video_id     int64  `gorm:"column:video_id;type:bigint;not null" json:"video_id"`
	Comment      string `gorm:"column:comment;type:varchar(255);not null" json:"comment"`
	Publish_time string `gorm:"column:publish_time;type:varchar(255);not null" json:"publish_time"`
}
