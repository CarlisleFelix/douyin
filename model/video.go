package model

type Video struct {
	Video_id       int64  `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY;not null" json:"video_id"`
	Author_id      int64  `gorm:"column:author_id;type:int;not null" json:"author_id"`
	Play_url       string `gorm:"column:play_url;type:varchar(255);not null" json:"play_url"`
	Cover_url      string `gorm:"column:name;type:varchar(255);not null" json:"cover_url"`
	Favorite_count int64  `gorm:"column:favorite_count;type:int;default:0" json:"favorite_count"`
	Comment_count  int64  `gorm:"column:comment_count;type:int;default:0" json:"comment_count"`
	Title          string `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Publish_time   string `gorm:"column:publish_time;type:varchar(255);not null" json:"publish_time"`
}
