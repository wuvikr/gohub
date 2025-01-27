package requests

import (
	"gohub/app/requests/validators"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ResetByPhoneRequest struct {
	Phone      string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
	Password   string `json:"password,omitempty" valid:"password"`
}

func ResetByPhone(data any, c *gin.Context) url.Values {
	rules := govalidator.MapData{
		"phone":       []string{"required", "digits:11"},
		"verify_code": []string{"required", "digits:6"},
		"password":    []string{"required", "min:6"},
	}

	messages := govalidator.MapData{
		"phone":       []string{"required:手机号为必填项", "digits:手机号长度必须为 11 位的数字"},
		"verify_code": []string{"required:验证码答案必填", "digits:验证码答案必须为 6 位的数字"},
		"password":    []string{"required:密码为必填项", "min:密码长度需大于 6"},
	}

	errs := validate(data, rules, messages)

	// 手机验证码
	_data := data.(*ResetByPhoneRequest)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)

	return errs
}

type ResetByEmailRequest struct {
	Email      string `json:"email,omitempty" valid:"email"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
	Password   string `json:"password,omitempty" valid:"password"`
}

func ResetByEmail(data any, c *gin.Context) url.Values {
	rules := govalidator.MapData{
		"email":       []string{"required", "email"},
		"verify_code": []string{"required", "digits:6"},
		"password":    []string{"required", "min:6"},
	}

	messages := govalidator.MapData{
		"email":       []string{"required:Email 为必填项", "email:Email 格式不正确"},
		"verify_code": []string{"required:验证码答案必填", "digits:验证码答案必须为 6 位的数字"},
		"password":    []string{"required:密码为必填项", "min:密码长度需大于 6"},
	}

	errs := validate(data, rules, messages)

	// 邮箱验证码
	_data := data.(*ResetByEmailRequest)
	errs = validators.ValidateVerifyCode(_data.Email, _data.VerifyCode, errs)

	return errs
}

type UserUpdatePasswordRequest struct {
	Password           string `json:"password,omitempty" valid:"password"`
	NewPassword        string `json:"new_password,omitempty" valid:"new_password"`
	NewPasswordConfirm string `json:"new_password_confirm,omitempty" valid:"new_password_confirm"`
}

func UserUpdatePassword(data any, c *gin.Context) url.Values {
	rules := govalidator.MapData{
		"password":             []string{"required", "min:6"},
		"new_password":         []string{"required", "min:6"},
		"new_password_confirm": []string{"required", "min:6"},
	}

	messages := govalidator.MapData{
		"password":             []string{"required:密码为必填项", "min:密码长度需大于 6"},
		"new_password":         []string{"required:新密码为必填项", "min:新密码长度需大于 6"},
		"new_password_confirm": []string{"required:确认密码框为必填项", "min:确认密码框长度需大于 6"},
	}

	errs := validate(data, rules, messages)

	// 比较两次输入的密码
	_data := data.(*UserUpdatePasswordRequest)
	errs = validators.ValidatePasswordConfirm(_data.NewPassword, _data.NewPasswordConfirm, errs)

	return errs
}
