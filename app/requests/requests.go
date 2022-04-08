package requests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

// ValidatorFunc 验证函数类型
type ValidatorFunc func(interface{}, *gin.Context) map[string][]string


func validate (data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {

	// 配置初始化
	opts := govalidator.Options{
		Data: data,
		Rules: rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages: messages,
	}

	// 开始验证
	return govalidator.New(opts).ValidateStruct()
}

func Validate(c *gin.Context, obj interface{}, handle ValidatorFunc) bool {
	if err := c.ShouldBindJSON(&obj); err != nil {
		// 解析失败，返回 422 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。",
			"error": err.Error(),
		})
		// 打印错误信息
		fmt.Println(err.Error())
		// 出错了，中断请求
		return false
	}

	// 表单验证
	errs := handle(obj,c)

	// 判断验证是否通过
	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "请求验证不通过，具体请查看 errors",
			"errors":  errs,
		})
		return false
	}

	return true
}