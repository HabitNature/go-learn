package controller

import (
	"github.com/gin-gonic/gin"
	"go-learn/request"
	"go-learn/response"
	"go-learn/service"
)

type UserController struct {
}

// @Summary 注册用户
// @Description 注册用户
// @Tags 用户
// @Accept json
// @Produce json
// @Param name formData string true "用户名：4-20位,数字，字母组合"
// @Param password formData string true "密码; 8-20位，数字，字母，特殊字符的组合"
// @Router /user/signup [Post]
func (c UserController) AccountSignUp(ctx *gin.Context) {
	req := request.UserSignUp{}

	if err := request.Validate(ctx, &req); err != nil {
		response.Error(ctx, err)
		return
	}

	request.ValidatePassword(req.Password)

	u, err := service.SignUp(req)

	if err != nil {
		response.Error(ctx, err)
		return
	}

	ur := response.UserResponseFrom(u)

	response.Ok(ctx, ur)
}

// @Summary 登录
// @Description 登录
// @Tags 用户
// @Accept json
// @Produce json
// @Param name formData string true "账号：4-20位,数字，字母组合"
// @Param password formData string true "密码; 8-20位，数字，字母，特殊字符的组合"
// @Router /user/signin [Post]
func (c UserController) AccountSignIn(ctx *gin.Context) {
	req := request.UserSignUp{}

	if err := request.Validate(ctx, &req); err != nil {
		response.Error(ctx, err)
		return
	}

	user, token, err := service.SignIn(req)

	if err != nil {
		response.Error(ctx, err)
		return
	}

	res := response.SignInResponseFrom(user, token)
	response.Ok(ctx, res)
}
