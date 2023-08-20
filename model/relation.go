package model

type Relation struct {
	Relation_id int64 `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY;not null" json:"relation_id"`
	Host_id     int64 `gorm:"column:host_id;type:int;not null" json:"follow_id"`
	Guest_id    int64 `gorm:"column:guest_id;type:int;not null" json:"follower_id"`
}
