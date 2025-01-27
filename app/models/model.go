package models

import (
	"time"

	"github.com/spf13/cast"
)

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primary_key;autoIncrement;" json:"id,omitempty"`
}

type CommonTimestampsField struct {
	CreatedAT time.Time `gorm:"column:created_at;index;autoCreateTime" json:"created_at,omitempty"`
	UpdatedAT time.Time `gorm:"column:updated_at;index;autoUpdateTime" json:"updated_at,omitempty"`
}

func (a BaseModel) GetStringID() string {
	return cast.ToString(a.ID)
}
