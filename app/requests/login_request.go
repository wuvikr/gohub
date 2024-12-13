package requests

import (
	"gohub/app/requests/validators"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type LoginByPhoneRequest struct {
	Phone      string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
}

// LoginByPhone 手机登录验证表单
func LoginByPhone(data interface{}, c *gin.Context) url.Values {

	rules := govalidator.MapData{
		"phone":       []string{"required", "digits:11"},
		"verify_code": []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"phone":       []string{"required:手机号为必填项", "digits:手机号长度必须为 11 位的数字"},
		"verify_code": []string{"required:验证码答案必填", "digits:验证码答案必须为 6 位的数字"},
	}

	errs := validate(data, rules, messages)

	// 手机验证码
	_data := data.(*LoginByPhoneRequest)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)

	return errs
}

type LoginByPasswordRequest struct {
	CaptchaID     string `json:"captcha_id" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer" valid:"captcha_answer"`
	LoginID       string `json:"login_id" valid:"login_id"`
	Password      string `json:"password" valid:"password"`
}

func LoginByPassword(data interface{}, c *gin.Context) url.Values {
	rules := govalidator.MapData{
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
		"login_id":       []string{"required", "min:3"},
		"password":       []string{"required", "min:6"},
	}

	messages := govalidator.MapData{
		"captcha_id":     []string{"required:验证码 ID 为必填项"},
		"captcha_answer": []string{"required:验证码答案必填", "digits:验证码答案必须为 6 位的数字"},
		"login_id":       []string{"required:登录 ID 为必填项", "min:登录 ID 长度需大于 3"},
		"password":       []string{"required:密码为必填项", "min:密码长度需大于 6"},
	}

	errs := validate(data, rules, messages)

	// 图片验证验证码
	_data := data.(*LoginByPasswordRequest)
	errs = validators.ValidateCaptcha(_data.CaptchaID, _data.CaptchaAnswer, errs)

	return errs
}
