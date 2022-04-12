// Package auth 处理用户身份认证相关逻辑
package auth

import (
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"gohub/pkg/helpers"
	"gohub/pkg/response"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist 检测手机号是否被注册
func (sc *SignupController) IsPhoneExist(c *gin.Context) {

	// 初始化请求对象
	request := requests.SignupPhoneExistRequest{}

	ok := requests.Validate(c, &request, requests.ValidateSignupPhoneExist)
	if !ok {
		return
	}

	//  检查数据库并返回响应
	response.JSON(c, gin.H{
		"name":  request.Name,
		"phone": request.Phone,
		"exist": user.IsPhoneExist(request.Phone),
	})
}

func (sc *SignupController) IsEmailExist(c *gin.Context) {
	request := requests.SignupEmailExistRequest{}
	if ok := requests.Validate(c, &request, requests.ValidateSignupEmailExist); !ok {
		return
	}

	response.JSON(c, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}

// 邮件注册
func (sc *SignupController) SignupUsingEmail(c *gin.Context) {
	request := requests.SignupUsingEmail{}
	if ok := requests.Validate(c, &request, requests.ValidateSignupEmail); !ok {
		return
	}

	userModel := user.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: helpers.GetStringMd5(request.Password),
	}
	userModel.Create()
	if userModel.ID > 0 {
		response.JSON(c, gin.H{
			"ok": true,
		})
	} else {
		response.JSON(c, gin.H{
			"ok": false,
		})
	}

}

//SignupUsingPhone 使用手机和验证码进行注册
func (sc *SignupController) SignupUsingPhone(c *gin.Context) {
	request := requests.SignupUsingPhoneRequest{}

	// 1. 验证表单
	if ok := requests.Validate(c, &request, requests.SignupUsingPhone); !ok {
		return
	}

	// 2. 验证成功，创建数据
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
		response.Abort500(c, "创建用户失败，请稍后尝试~")
	}
}
