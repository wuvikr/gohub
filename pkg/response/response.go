package response

import (
	"gohub/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// JSON 返回 200 和 JSON数据
func JSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

// Success 返回 200 预设『操作成功！』的 JSON 数据
// 执行某些没有返回值的操作时使用，例如删除，修改密码等
func Success(c *gin.Context) {
	JSON(c, gin.H{
		"success": true,
		"message": "操作成功",
	})
}

// Data 返回 200 和带 data 的 JSON
// 执行成功后，用于返回已更新的数据，例如更新用户信息，更新话题
func Data(c *gin.Context, data interface{}) {
	JSON(c, gin.H{
		"success": true,
		"data":    data,
	})
}

// Created 返回 201 和带 data 的 JSON
// 执行更新操作成功后，用于返回已创建的数据，例如创建话题
func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    data,
	})
}

// CreatedJSON 响应 201 和 JSON 数据
func CreatedJSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, data)
}

// Abort404 响应 404
func Abort404(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"message": defaultMessage("请求资源不存在，请确认 url 和请求方法是否正确。", msg...),
	})
}

// Abort403 响应 403
func Abort403(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"message": defaultMessage("权限不足，请确保有操作权限。", msg...),
	})
}

// Abort500 响应 500
func Abort500(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"message": defaultMessage("服务器内部错误，请稍后再试。", msg...),
	})
}

// BadRequest 响应 400, 传参错误
func BadRequest(c *gin.Context, err error, msg ...string) {
	logger.LogIf(err)
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": defaultMessage("请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。", msg...),
		"error":   err.Error(),
	})
}

// Error 响应 404 或者 422 错误
func Error(c *gin.Context, err error, msg ...string) {
	logger.LogIf(err)

	if err == gorm.ErrRecordNotFound {
		Abort404(c)
		return
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"message": defaultMessage("服务器内部错误，请稍后再试。", msg...),
		"error":   err.Error(),
	})
}

// ValidationError 处理表单验证错误
func ValidationError(c *gin.Context, errors map[string][]string) {
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"message": "请求验证不通过，请修正后重试。",
		"errors":  errors,
	})
}

// Unauthorized 响应 401
// 未登录, jwt 过期等错误时调用
func Unauthorized(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"message": defaultMessage("未授权，请检查您的登录状态。", msg...),
	})
}

// defaultMessage 内部辅助函数，用于设置默认消息
func defaultMessage(defaultMsg string, msg ...string) string {
	if len(msg) > 0 {
		return msg[0]
	}
	return defaultMsg
}
