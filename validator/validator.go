package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func newValidator() *validator.Validate {
	v := validator.New()
	v.RegisterValidation("phone", PhoneReg)
	return v
}

func ValidateWithSturct(s interface{}) string {
	v := newValidator()
	tagList := []string{}
	if errs := v.Struct(s); errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			fmt.Println(err.Error())
			tagList = append(tagList, strings.ToLower(err.Field()))
		}
	}
	var errString string = ""
	if len(tagList) != 0 {
		errString = strings.Join(tagList, ",") + "等参数格式有误"
	}
	return errString
}
