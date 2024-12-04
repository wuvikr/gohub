package requests

import (
	"gohub/pkg/captcha"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type VerifyCodePhoneRequest struct {
	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`

	Phone string `json:"phone,omitempty" valid:"phone"`
}

func VerifyCodePhone(data interface{}, c *gin.Context) url.Values {
	// 自定义验证规则
	rules := govalidator.MapData{
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
		"phone":          []string{"required", "digits:11"},
	}

	// 自定义验证出错的提示语
	messages := govalidator.MapData{
		"captcha_id":     []string{"required:验证码 ID 为必填项"},
		"captcha_answer": []string{"required:验证码答案必填", "digits:验证码答案必须为 6 位的数字"},
		"phone":          []string{"required:手机号为必填项", "digits:手机号长度必须为 11 位的数字"},
	}

	errs := validate(data, rules, messages)

	// 图片验证验证码
	_data := data.(*VerifyCodePhoneRequest)
	if ok := captcha.NewCaptcha().VerifyCaptcha(_data.CaptchaID, _data.CaptchaAnswer); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "验证码答案错误")
	}

	return errs
}
