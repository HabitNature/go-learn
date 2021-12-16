package controller

import (
	"github.com/gin-gonic/gin"
	"go-learn/model"
	"go-learn/request"
	"go-learn/response"
)

type UserController struct {
}

// @Summary 创建用户
// @Description 创建用户
// @Tags 用户
// @Accept json
// @Produce json
// @Param name formData string true "用户名"
// @Router /user/create [Post]
func (c UserController) Create(ctx *gin.Context) {
	req := request.UserCreate{}

	err := request.Validate(ctx, &req)

	if err != nil {
		response.Error(ctx, err)
		return
	}

	user := model.User{}
	user.Name = req.Name

	err = model.DbInstance.Create(&user).Error

	if err != nil {
		response.Error(ctx, err)
	} else {
		response.Ok(ctx, user)
	}
}
