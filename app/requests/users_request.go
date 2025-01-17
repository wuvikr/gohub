package requests

import (
	"gohub/app/requests/validators"
	"gohub/pkg/auth"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type UserUpdateProfileRequest struct {
	Name         string `valid:"name" json:"name"`
	City         string `valid:"city" json:"city"`
	Introduction string `valid:"introduction" json:"introduction"`
}

func UserUpdateProfile(data any, c *gin.Context) url.Values {

	uid := auth.CurrentUID(c)
	rules := govalidator.MapData{
		"name":         []string{"required", "alpha_num", "between:2,20", "not_exists:users,name," + uid},
		"city":         []string{"min_cn:2", "max_cn:10"},
		"introduction": []string{"min_cn:3", "max_cn:240"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:用户名为必填项",
			"alpha_num:用户名格式错误，只允许数字和字母",
			"between:用户名长度需在 2~20 之间",
			"not_exists:用户名已存在",
		},
		"city": []string{
			"min_cn:城市长度需至少 2 个字",
			"max_cn:城市长度不能超过 10 个字",
		},
		"introduction": []string{
			"min_cn:简介长度需至少 3 个字",
			"max_cn:简介长度不能超过 240 个字",
		},
	}
	return validate(data, rules, messages)
}

type UserUpdateEmailRequest struct {
	Email      string `valid:"email" json:"email,omitempty"`
	VerifyCode string `valid:"verify_code" json:"verify_code,omitempty"`
}

func UserUpdateEmail(data any, c *gin.Context) url.Values {
	currentUser := auth.CurrentUser(c)
	rules := govalidator.MapData{
		"email": []string{
			"required",
			"email",
			"min:4",
			"max:30",
			"not_exists:users,email," + currentUser.GetStringID(),
			"not_in:" + currentUser.Email,
		},
		"verify_code": []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"email:Email 格式不正确",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"not_exists:Email 已被占用",
			"not_in:新的Email 与当前 Email 相同",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
	}

	errs := validate(data, rules, messages)

	// 邮箱验证码
	_data := data.(*UserUpdateEmailRequest)
	errs = validators.ValidateVerifyCode(_data.Email, _data.VerifyCode, errs)
	return errs
}

type UserUpdatePhoneRequest struct {
	Phone      string `valid:"phone" json:"phone,omitempty"`
	VerifyCode string `valid:"verify_code" json:"verify_code,omitempty"`
}

func UserUpdatePhone(data any, c *gin.Context) url.Values {
	currentUser := auth.CurrentUser(c)

	rules := govalidator.MapData{
		"phone": []string{
			"required",
			"digits:11",
			"not_exists:users,phone," + currentUser.GetStringID(),
			"not_in:" + currentUser.Phone,
		},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项",
			"digits:手机号长度必须为 11 位的数字",
			"not_exists:手机号已被占用",
			"not_in:新的手机号与当前手机号相同",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
	}

	errs := validate(data, rules, messages)

	// 手机验证码
	_data := data.(*UserUpdatePhoneRequest)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)

	return errs
}
