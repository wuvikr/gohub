package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type PasswordController struct {
	v1.BaseAPIController
}

// ResetByPhone 通过手机验证码重置密码
func (pc *PasswordController) ResetByPhone(c *gin.Context) {
	// 1. 验证表单
	requestData := requests.ResetByPhoneRequest{}
	if ok := requests.Validate(c, &requestData, requests.ResetByPhone); !ok {
		return
	}

	// 2. 重置密码
	userInstance := user.GetByPhone(requestData.Phone)
	if userInstance.ID == 0 {
		response.Abort404(c)
	}
	userInstance.Password = requestData.Password
	userInstance.Save()
	response.Success(c)
}

// ResetByEmail 通过 Email 验证码重置密码
func (pc *PasswordController) ResetByEmail(c *gin.Context) {
	// 1. 验证表单
	requestData := requests.ResetByEmailRequest{}
	if ok := requests.Validate(c, &requestData, requests.ResetByEmail); !ok {
		return
	}

	// 2. 重置密码
	userInstance := user.GetByEmail(requestData.Email)
	if userInstance.ID == 0 {
		response.Abort404(c)
	}

	userInstance.Password = requestData.Password
	userInstance.Save()
	response.Success(c)
}
