// Package routes 注册路由
package routes

import (
	"gohub/app/http/controllers/api/v1/auth"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			// 判断手机是否已注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)
			authGroup.POST("/signup/email", suc.SignupUsingEmail)
			authGroup.POST("/signup/using-phone", suc.SignupUsingPhone)

			vcc := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)

			user := new(auth.UserController)
			authGroup.POST("/getUser", user.GetUser)

			authGroup.POST("/verify-codes/phone", vcc.SendUsingPhone)

			authGroup.POST("/verify-codes/email", vcc.SendUsingEmail)

			lgc := new(auth.LoginController)
			// 使用手机号，短信验证码进行登录
			authGroup.POST("/login/using-phone", lgc.LoginByPhone)
			// 使用邮箱登录
			authGroup.POST("/login/using-email", lgc.LoginByEmail)
			// 账号密码登录
			authGroup.POST("/login/account", lgc.LoginByPassword)
			// 刷新 token
			authGroup.POST("/login/refresh-token", lgc.RefreshToken)
		}
	}
}
