package requests

import (
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type PaginationRequest struct {
	Sort    string `form:"sort" valid:"sort"`
	Order   string `form:"order" valid:"order"`
	PerPage int    `form:"per_page" valid:"per_page"`
}

func Pagination(data any, c *gin.Context) url.Values {
	rules := govalidator.MapData{
		"sort":     []string{"in:id,created_at,updated_at"},
		"order":    []string{"in:asc,desc"},
		"per_page": []string{"numeric_between:3,100"},
	}

	messages := govalidator.MapData{
		"sort":     []string{"in:排序字段仅支持 id, created_at, updated_at"},
		"order":    []string{"in:排序规则仅支持 asc（正序）,desc（倒序）"},
		"per_page": []string{"numeric_between:每页条数的值介于 3~100 之间"},
	}
	return validate(data, rules, messages)
}
