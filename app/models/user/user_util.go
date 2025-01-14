package user

import (
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"

	"github.com/gin-gonic/gin"
)

// IsEmailExist 判断邮箱是否存在
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// IsPhoneExist 判断手机号是否存在
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

// GetByMulti 通过 手机号/Email/用户名 来获取用户
func GetByMulti(loginID string) (user User) {
	database.DB.
		Where("phone = ?", loginID).
		Or("email = ?", loginID).
		Or("name = ?", loginID).
		First(&user)
	return user
}

// GetByPhone 通过手机号获取用户
func GetByPhone(phone string) (user User) {
	database.DB.Where("phone = ?", phone).First(&user)
	return user
}

// GetByEmail 通过 Email 获取用户
func GetByEmail(email string) (user User) {
	database.DB.Where("email = ?", email).First(&user)
	return user
}

// Get 通过 id 获取用户
func Get(idstr string) User {
	var user User
	database.DB.Where("id = ?", idstr).First(&user)
	return user
}

func All() (users []User) {
	database.DB.Find(&users)
	return
}

func Paginate(c *gin.Context, perPage int) (users []User, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(User{}),
		&users,
		app.V1URL(database.TableName(&User{})),
		perPage,
	)
	return
}
