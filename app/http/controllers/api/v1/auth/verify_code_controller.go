package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/pkg/captcha"
	"gohub/pkg/logger"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

// VerifyCodeController 验证码控制器
type VerifyCodeController struct {
	v1.BaseAPIController
}

// ShowCaptcha 生成验证码
func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	// 生成验证码
	id, b64s, _, err := captcha.NewCaptcha().GenerateCaptcha()
	logger.LogIf(err)

	response.JSON(c, gin.H{
		"captcha_id":  id,
		"captcha_img": b64s,
	})
}
