// Package user 存放用户模块的模型
package user

import (
	"gohub/app/models"
	"gohub/pkg/database"
)

// User 用户模型
type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}

func (u *User) Create() {
	database.DB.Create(&u)
}
