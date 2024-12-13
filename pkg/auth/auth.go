package auth

import (
	"errors"
	"gohub/app/models/user"
)

// Attempt 尝试登录
func Attempt(email string, password string) (user.User, error) {
	_user := user.GetByMulti(email)
	if _user.ID == 0 {
		return user.User{}, errors.New("用户不存在")
	}

	if !_user.ComparePassword(password) {
		return user.User{}, errors.New("密码不正确")
	}

	return _user, nil
}

// LoginByPhone 通过手机号登录
func LoginByPhone(phone string) (user.User, error) {
	_user := user.GetByPhone(phone)
	if _user.ID == 0 {
		return user.User{}, errors.New("用户不存在")
	}
	return _user, nil
}
