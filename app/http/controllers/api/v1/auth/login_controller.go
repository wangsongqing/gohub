package auth

import (
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/requests"
	"gohub/pkg/auth"
	"gohub/pkg/jwt"
	"gohub/pkg/response"
)

// LoginController 用户控制器
type LoginController struct {
	v1.BaseAPIController
}

// LoginByPhone 手机登录
func (lc *LoginController) LoginByPhone(c *gin.Context) {

	// 1. 验证表单
	request := requests.LoginByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPhone); !ok {
		return
	}

	// 2. 尝试登录
	user, err := auth.LoginByPhone(request.Phone)
	if err != nil {
		// 失败，显示错误提示
		response.Error(c, err, "账号不存在或密码错误")
	} else {
		// 登录成功
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)

		response.JSON(c, gin.H{
			"token": token,
		})
	}
}

// LoginByEmail 邮箱登录
func (lc *LoginController) LoginByEmail(c *gin.Context) {
	request := requests.LoginByEmailRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByEmail); !ok {
		return
	}

	user, err := auth.LoginByEmail(request.Email)
	if err != nil {
		response.Error(c, err, "邮箱账号不存在或密码错误")
	} else {
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)

		response.JSON(c, gin.H{
			"token": token,
		})
	}
}

// LoginByPassword 账号密码登录
func (lc *LoginController) LoginByPassword(c *gin.Context) {
	request := requests.LoginPasswordRequest{}
	if ok := requests.Validate(c, &request, requests.LoginPassword); !ok {
		return
	}

	userInfo, err := auth.LoginByAccount(request.Account)
	if userInfo.ID > 0 {
		token := jwt.NewJWT().IssueToken(userInfo.GetStringID(), userInfo.Name)
		response.JSON(c, gin.H{
			"token": token,
		})
	} else {
		response.Error(c, err, "邮箱账号不存在或密码错误")
	}
}