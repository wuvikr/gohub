package auth

import (
	"errors"
	"gohub/app/models/user"
	"gohub/pkg/logger"

	"github.com/gin-gonic/gin"
)

// Attempt 尝试登录
func Attempt(email string, password string) (user.User, error) {
	UserInstance := user.GetByMulti(email)
	if UserInstance.ID == 0 {
		return user.User{}, errors.New("用户不存在")
	}

	if !UserInstance.ComparePassword(password) {
		return user.User{}, errors.New("密码不正确")
	}

	return UserInstance, nil
}

// LoginByPhone 通过手机号登录
func LoginByPhone(phone string) (user.User, error) {
	UserInstance := user.GetByPhone(phone)
	if UserInstance.ID == 0 {
		return user.User{}, errors.New("用户不存在")
	}
	return UserInstance, nil
}

func CurrentUser(c *gin.Context) user.User {
	UserInstance, ok := c.MustGet("current_user").(user.User)
	if !ok {
		logger.LogIf(errors.New("无法获取当前用户"))
		return user.User{}
	}
	return UserInstance
}
