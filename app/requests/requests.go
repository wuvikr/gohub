package requests

import (
	"gohub/pkg/response"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

// ValidatorFunc 验证函数类型
type ValidatorFunc func(data interface{}, c *gin.Context) url.Values

// Validate 进行参数绑定和验证表单，调用示例：
//
//	if ok := requests.Validate(c, &requests.UserSaveRequest{}, requests.UserSave); !ok {
//	    return
//	}
func Validate(c *gin.Context, obj interface{}, handler ValidatorFunc) bool {
	// 1. 解析请求参数
	if err := c.ShouldBind(obj); err != nil {
		response.BadRequest(c, err, "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。")
		return false
	}

	// 2. 表单验证
	errs := handler(obj, c)
	if len(errs) > 0 {
		response.ValidationError(c, errs)
		return false
	}

	return true

}

// validate 创建一个验证器
func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) url.Values {
	// 配置验证规则和参数
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
	}

	return govalidator.New(opts).ValidateStruct()
}
