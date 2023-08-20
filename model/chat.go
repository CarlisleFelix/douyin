package model

type Chat struct {
	Id           int64  `gorm:"column:id;type:bigint auto_increment;primaryKey" json:"id"`
	Sender_id    int64  `gorm:"column:sender_id;type:bigint;not null" json:"sender_id"`
	Receiver_id  int64  `gorm:"column:receiver_id;type:bigint;not null" json:"receiver_id"`
	Content      string `gorm:"column:content;type:varchar(255);default:null" json:"content"`
	Publish_time int64  `gorm:"column:publish_time;type:bigint;not null" json:"publish_time"`
}
