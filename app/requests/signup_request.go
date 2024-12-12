package requests

import (
	"gohub/app/requests/validators"
	"gohub/pkg/logger"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

// ValidateSignupPhoneExist 验证手机号是否存在
func ValidateSignupPhoneExist(data interface{}, c *gin.Context) url.Values {
	// 自定义验证规则
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}

	// 自定义验证出错的提示语
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
	}

	return validate(data, rules, messages)
}

type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

// ValidateSignupEmailExist 验证邮箱是否存在
func ValidateSignupEmailExist(data interface{}, c *gin.Context) url.Values {
	// 自定义验证规则
	rules := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}

	// 自定义验证出错的提示语
	messages := govalidator.MapData{
		"email": []string{
			"required:邮箱为必填项，参数名称 email",
			"min:邮箱长度需大于 4",
			"max:邮箱长度需小于 30",
			"email: Email 格式不正确, 请输入正确的邮箱格式",
		},
	}

	return validate(data, rules, messages)
}

// SignupUsingPhoneRequest 手机注册请求
type SignupUsingPhoneRequest struct {
	Phone           string `json:"phone,omitempty" valid:"phone"`
	VerifyCode      string `json:"verify_code,omitempty" valid:"verify_code"`
	Name            string `json:"name" valid:"name"`
	Password        string `json:"password,omitempty" valid:"password"`
	PasswordConfirm string `json:"password_confirm,omitempty"`
}

func SignupUsingPhone(data interface{}, c *gin.Context) url.Values {

	// 自定义验证规则
	rules := govalidator.MapData{
		"phone":            []string{"required", "digits:11", "not_exists:users,phone"},
		"verify_code":      []string{"required", "digits:6"},
		"name":             []string{"required", "alpha_num", "between:3,20", "not_exists:users,name"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
	}

	// 自定义验证出错的提示语
	messages := govalidator.MapData{
		"phone":            []string{"required:手机号为必填项", "digits:手机号长度必须为 11 位的数字"},
		"verify_code":      []string{"required:验证码答案必填", "digits:验证码答案必须为 6 位的数字"},
		"name":             []string{"required:用户名为必填项", "alpha_num:用户名格式错误，只允许数字和字母", "between:用户名长度需在 3~20 之间"},
		"password":         []string{"required:密码为必填项", "min:密码长度需大于 6"},
		"password_confirm": []string{"required:确认密码框为必填项", "equal:两次输入的密码不相同"},
	}

	errs := validate(data, rules, messages)

	_data := data.(*SignupUsingPhoneRequest)
	logger.DebugJSON("SignupUsingPhoneRequest", "_data", _data)
	// 图片验证验证码
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)

	// 比较两次输入的密码
	errs = validators.ValidatePasswordConfirm(_data.Password, _data.PasswordConfirm, errs)

	return errs
}

// SignupUsingEmailRequest 通过邮箱注册的请求信息
type SignupUsingEmailRequest struct {
	Email           string `json:"email,omitempty" valid:"email"`
	VerifyCode      string `json:"verify_code,omitempty" valid:"verify_code"`
	Name            string `json:"name" valid:"name"`
	Password        string `json:"password,omitempty" valid:"password"`
	PasswordConfirm string `json:"password_confirm,omitempty" valid:"password_confirm"`
}

func SignupUsingEmail(data interface{}, c *gin.Context) url.Values {

	// 自定义验证规则
	rules := govalidator.MapData{
		"email":            []string{"required", "min:4", "max:30", "email", "not_exists:users,email"},
		"verify_code":      []string{"required", "digits:6"},
		"name":             []string{"required", "alpha_num", "between:3,20", "not_exists:users,name"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
	}

	// 自定义验证出错的提示语
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
			"not_exists:Email 已被占用",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
		"name": []string{
			"required:用户名为必填项",
			"alpha_num:用户名格式错误，只允许数字和英文",
			"between:用户名长度需在 3~20 之间",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"password_confirm": []string{
			"required:确认密码框为必填项",
		},
	}

	errs := validate(data, rules, messages)

	// 图片验证验证码
	_data := data.(*SignupUsingEmailRequest)
	errs = validators.ValidateVerifyCode(_data.Email, _data.VerifyCode, errs)

	// 比较两次输入的密码
	errs = validators.ValidatePasswordConfirm(_data.Password, _data.PasswordConfirm, errs)

	return errs
}
