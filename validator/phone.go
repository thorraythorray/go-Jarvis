package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func PhoneReg(fl validator.FieldLevel) bool {
	// 定义手机号码的正则表达式
	// 此处使用简化的手机号码匹配规则，根据实际情况可进行调整
	phoneRegex := `^1[3-9]\d{9}$`

	// 使用正则表达式验证手机号码
	matched, _ := regexp.MatchString(phoneRegex, fl.Field().String())
	return matched
}
