package requests

import (
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

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

	// 配置初始化
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
	}

	return govalidator.New(opts).ValidateStruct()
}

type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

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

	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
	}

	return govalidator.New(opts).ValidateStruct()
}
