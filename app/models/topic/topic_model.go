package topic

import (
	"gohub/app/models"
	"gohub/app/models/category"
	"gohub/app/models/user"
	"gohub/pkg/database"
)

type Topic struct {
	models.BaseModel

	Title      string `json:"title,omitempty"`
	Body       string `json:"body,omitempty" gorm:"type:text"`
	UserID     string `json:"user_id,omitempty"`
	CategoryID string `json:"category_id,omitempty"`

	// 关联用户
	User user.User `json:"user"`

	// 关联分类
	Category category.Category `json:"category"`

	models.CommonTimestampsField
}

func (topic *Topic) Create() {
	database.DB.Create(&topic)
}

func (topic *Topic) Save() (rowsAffected int64) {
	result := database.DB.Save(&topic)
	return result.RowsAffected
}

func (topic *Topic) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&topic)
	return result.RowsAffected
}
