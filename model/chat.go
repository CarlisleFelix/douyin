package model

type Chat struct {
	Chat_id      int64 `gorm:"column:id;type:bigint auto_increment;primaryKey" json:"chat_id"`
	Sender_id    int64 `gorm:"column:sender_id;type:bigint;not null" json:"sender_id"`
	Receiver_id  int64 `gorm:"column:receiver_id;type:bigint;not null" json:"receiver_id"`
	Publish_time int64 `gorm:"column:publish_time;type:varchar(255);not null" json:"publish_time"`
}
