package model

type Relation struct {
	Relation_id int64 `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY;not null" json:"relation_id"`
	Follow_id   int64 `gorm:"column:follow_id;type:int;not null" json:"follow_id"`
	Follower_id int64 `gorm:"column:follower_id;type:int;not null" json:"follower_id"`
}
