// Package user 存放用户模块的模型
package user

import (
	"gohub/app/models"
	"gohub/pkg/database"
	"gohub/pkg/hash"
)

// User 用户模型
type User struct {
	models.BaseModel

	Name         string `json:"name,omitempty"`
	City         string `json:"city,omitempty"`
	Introduction string `json:"introduction,omitempty"`
	Avatar       string `json:"avatar,omitempty"`
	Email        string `json:"email,omitempty"`
	Phone        string `json:"-"`
	Password     string `json:"-"`

	models.CommonTimestampsField
}

func (u *User) Create() {
	database.DB.Create(&u)
}

func (u *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, u.Password)
}

func (u *User) Save() (rowsAffected int64) {
	result := database.DB.Save(&u)
	return result.RowsAffected
}
