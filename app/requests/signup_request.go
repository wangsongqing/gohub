package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
	Name string `json:"name,omitempty" valid:"name"`
}

type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

func ValidateSignupPhoneExist(data interface{}, c *gin.Context) map[string][]string {
	// 自定义验证规则ValidateSignupEmailExist
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}

	message := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
	}

	opts := govalidator.Options{
		Data : data,
		Rules: rules,
		TagIdentifier:"valid",
		Messages: message,
	}

	return govalidator.New(opts).ValidateStruct()
}

func ValidateSignupEmailExist(data interface{}, c *gin.Context) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}

	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"email": []string{
			"required: Email为必填项",
			"min: 长度需大于 4",
			"max: 长度小于 30",
			"email: Email格式不正确",
		},
	}

	// 配置初始化
	opts := govalidator.Options{
		Data: data,
		Rules: rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages: messages,
	}

	return govalidator.New(opts).ValidateStruct()
}