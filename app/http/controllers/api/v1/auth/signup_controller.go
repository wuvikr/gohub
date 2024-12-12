package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist 判断手机号是否已经注册
func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	// 请求对象
	request := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(c, &request, requests.ValidateSignupPhoneExist); !ok {
		return
	}

	// 检查数据库并返回响应
	response.JSON(c, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

func (sc *SignupController) IsEmailExist(c *gin.Context) {
	// 初始化请求对象
	request := requests.SignupEmailExistRequest{}
	if ok := requests.Validate(c, &request, requests.ValidateSignupEmailExist); !ok {
		return
	}

	response.JSON(c, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}

func (sc *SignupController) SignupUsingPhone(c *gin.Context) {
	// 1. 验证表单
	request := requests.SignupUsingPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.SignupUsingPhone); !ok {
		return
	}

	// 2. 表单验证通过，创建数据
	_user := user.User{
		Name:     request.Name,
		Phone:    request.Phone,
		Password: request.Password,
	}
	_user.Create()

	if _user.ID > 0 {
		response.CreatedJSON(c, gin.H{
			"data": _user,
		})
	} else {
		response.Abort500(c, "创建用户失败, 请稍后再试")
	}

}

func (sc *SignupController) SignupUsingEmail(c *gin.Context) {
	// 1. 验证表单
	request := requests.SignupUsingEmailRequest{}
	if ok := requests.Validate(c, &request, requests.SignupUsingEmail); !ok {
		return
	}

	// 2. 表单验证通过，创建数据
	_user := user.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	_user.Create()

	if _user.ID > 0 {
		response.CreatedJSON(c, gin.H{
			"data": _user,
		})
	} else {
		response.Abort500(c, "创建用户失败, 请稍后再试")
	}
}
