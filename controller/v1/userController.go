package controller

import (
	"github.com/gin-gonic/gin"
	"go-learn/model"
	"go-learn/request"
	"go-learn/utils"
	"net/http"
)

type UserController struct {
}

func (c UserController) Create(ctx *gin.Context) {
	req := request.UserCreate{}

	err := ctx.ShouldBind(&req)

	if err != nil {
		ctx.JSONP(http.StatusOK, gin.H{"code": 400, "data": "", "msg": err.Error(), "traceId": utils.NewUUID()})
		return
	}

	user := model.User{}
	user.Name = req.Name

	err = model.DbInstance.Create(&user).Error

	if err != nil {
		ctx.JSONP(http.StatusOK, gin.H{"code": 400, "data": "", "msg": err.Error(), "traceId": utils.NewUUID()})
	} else {
		ctx.JSONP(http.StatusOK, gin.H{"code": 200, "data": user, "msg": "", "traceId": utils.NewUUID()})
	}
}
