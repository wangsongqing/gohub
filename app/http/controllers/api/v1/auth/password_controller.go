package auth

import (
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"gohub/pkg/hash"
	"gohub/pkg/response"
)

type PasswordController struct {
	v1.BaseAPIController
}

// PhoneByPassword 手机号修改密码
func (pc *PasswordController) PhoneByPassword(c *gin.Context) {
	// 参数验证
	request := requests.PasswordFindPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.PasswordFindPhone); !ok {
		return
	}

	userModel := user.GetByPhone(request.Phone)
	userModel.Password = hash.BcryptHash(request.Password)
	userModel.Save()
	response.Success(c)
}

// EmailByPassword 邮箱修改密码
func (pc *PasswordController) EmailByPassword(c *gin.Context) {
	// 参数验证
	request := requests.PasswordFindEmailRequest{}
	if ok := requests.Validate(c, &request, requests.PasswordFindEmail); !ok {
		return
	}

	userModel := user.GetByEmail(request.Email)
	userModel.Password = hash.BcryptHash(request.Password)
	userModel.Save()
	response.Success(c)
}
