package requests

import (
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type TopicRequest struct {
	Title      string `valid:"title" json:"title,omitempty"`
	Body       string `valid:"body" json:"body,omitempty"`
	CategoryID string `valid:"category_id" json:"category_id,omitempty"`
}

func TopicSave(data any, c *gin.Context) url.Values {

	rules := govalidator.MapData{
		"title":       []string{"required", "min_cn:2", "max_cn:20"},
		"body":        []string{"required", "min_cn:10", "max_cn:50000"},
		"category_id": []string{"required", "exists:categories,id"},
	}
	messages := govalidator.MapData{
		"title": []string{
			"required:标题为必填项",
			"min_cn:标题长度需至少 2 个字",
			"max_cn:标题长度不能超过 20 个字",
		},
		"body": []string{
			"required:内容为必填项",
			"min_cn:内容长度需至少 10 个字",
			"max_cn:内容长度不能超过 50000 个字",
		},
		"category_id": []string{
			"required:分类 ID 为必填项",
			"exists:分类不存在",
		},
	}
	return validate(data, rules, messages)
}
