package models

import "time"

// ID 自增主键
type ID struct {
	ID uint `json:"user_id" gorm:"primaryKey" column:"id"`
}

// Timestamp 创建与更新时间
type Timestamp struct {
	CreateTime time.Time `json:"create_time" column:"create_time"`
	UpdateTime time.Time `json:"update_time" column:"update_time"`
}

// PseudoDeletion 伪删除
type PseudoDeletion struct {
	DeleteTime time.Time `json:"delete_time" column:"delete_time"`
}
