package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/requests"
	"gohub/pkg/auth"
	"gohub/pkg/jwt"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	v1.BaseAPIController
}

func (lc *LoginController) LoginByPhone(c *gin.Context) {

	// 1. 验证表单
	requestData := requests.LoginByPhoneRequest{}
	if ok := requests.Validate(c, &requestData, requests.LoginByPhone); !ok {
		return
	}

	// 2. 用户登录
	user, err := auth.LoginByPhone(requestData.Phone)
	if err != nil {
		response.Error(c, err, "账户不存在，请先注册。")
	} else {
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)
		response.JSON(c, gin.H{
			"data":  user,
			"token": token,
		})
	}
}

func (lc *LoginController) LoginByPassword(c *gin.Context) {

	// 1. 验证表单
	requestData := requests.LoginByPasswordRequest{}
	if ok := requests.Validate(c, &requestData, requests.LoginByPassword); !ok {
		return
	}

	// 2. 用户登录
	user, err := auth.Attempt(requestData.LoginID, requestData.Password)
	if err != nil {
		response.Unauthorized(c, "账户不存在或者密码错误。")
	} else {
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)
		response.JSON(c, gin.H{
			"token": token,
		})
	}
}
