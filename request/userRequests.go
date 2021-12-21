package request

import (
	"errors"
	"regexp"
)

const MinPasswordLength = 8

type SignUp struct {
	Password string `json:"password" form:"password" comment:"密码" binding:"required,min=8,max=20"`
}

type UserSignUp struct {
	Account string `json:"account" form:"account" comment:"账号" binding:"required,alphanum,min=4,max=20"`
	//Email string `json:"email" form:"email" comment:"邮箱" binding:"required"`
	//Mobile string `json:"mobile" form:"mobile" comment:"手机号" binding:"required"`
	SignUp
}

// 校验密码格式
func ValidatePassword(password string) error {

	m, err := regexp.Compile("(?=^.{8,20}$)(?=(?:.*?\\d){1})(?=.*[a-z])(?=(?:.*?[A-Z]){1})(?=(?:.*?[`·~!@#$%^&*()_+}{|:;'\",<.>/?\\=\\[\\]\\-\\\\]){1})(?!.*\\s)[0-9a-zA-Z`·~!@#$%^&*()_+}{|:;'\",<.>/?\\=\\[\\]\\-\\\\]*$")

	if err != nil {
		return err
	}

	if !m.Match([]byte(password)) {
		return errors.New("密码格式不正确")
	}

	return nil
}
