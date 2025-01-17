package v1

import (
	"gohub/app/models/user"
	"gohub/app/requests"
	"gohub/pkg/auth"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	BaseAPIController
}

// CurrentUser 当前登录用户信息
func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Data(c, userModel)
}

// Index 用户列表
func (ctrl *UsersController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := user.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (ctrl *UsersController) UpdateProfile(c *gin.Context) {
	request := requests.UserUpdateProfileRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdateProfile); !ok {
		return
	}

	currentUser := auth.CurrentUser(c)
	currentUser.Name = request.Name
	currentUser.City = request.City
	currentUser.Introduction = request.Introduction
	rowsAffected := currentUser.Save()
	if rowsAffected > 0 {
		response.Data(c, currentUser)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

// UpdateEmail 更新用户邮箱
func (ctrl *UsersController) UpdateEmail(c *gin.Context) {
	// 1. 表单验证
	request := requests.UserUpdateEmailRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdateEmail); !ok {
		return
	}

	// 2. 更新用户结构体信息
	currentUser := auth.CurrentUser(c)
	currentUser.Email = request.Email

	// 3. 保存用户信息到数据库
	rowsAffected := currentUser.Save()
	if rowsAffected > 0 {
		response.Success(c)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

// UpdatePhone 更新用户手机号
func (ctrl *UsersController) UpdatePhone(c *gin.Context) {
	// 1. 表单验证
	request := requests.UserUpdatePhoneRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdatePhone); !ok {
		return
	}

	// 2. 更新用户结构体信息
	currentUser := auth.CurrentUser(c)
	currentUser.Phone = request.Phone

	// 3. 保存用户信息到数据库
	rowsAffected := currentUser.Save()
	if rowsAffected > 0 {
		response.Success(c)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

// UpdatePassword 更新用户密码
func (ctrl *UsersController) UpdatePassword(c *gin.Context) {
	// 1. 表单验证
	request := requests.UserUpdatePasswordRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdatePassword); !ok {
		return
	}

	currentUser := auth.CurrentUser(c)

	// 2. 验证原先密码是否正确
	_, err := auth.Attempt(currentUser.Name, request.Password)
	if err != nil {
		response.Unauthorized(c, "原密码不正确")
		return
	}

	// 3. 更新用户结构体信息
	currentUser.Password = request.NewPassword
	rowsAffected := currentUser.Save()
	if rowsAffected > 0 {
		response.Success(c)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}
