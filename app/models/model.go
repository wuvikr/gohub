package models

import "time"

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primary_key;autoIncrement;" json:"id,omitempty"`
}

type CommonTimestampsField struct {
	CreatedAT time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
	UpdatedAT time.Time `gorm:"column:updated_at;index;" json:"updated_at,omitempty"`
}
