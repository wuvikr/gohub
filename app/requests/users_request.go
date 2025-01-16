package requests

import (
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
