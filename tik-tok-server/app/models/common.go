package models

import "time"

// ID 自增主键
type ID struct {
	ID uint `json:"id" gorm:"primaryKey"`
}

// Timestamp 创建与更新时间
type Timestamp struct {
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

// PseudoDeletion 伪删除
type PseudoDeletion struct {
	DeleteTime time.Time `json:"delete_time"`
}
