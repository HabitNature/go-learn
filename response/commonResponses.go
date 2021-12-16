package response

import (
	"github.com/gin-gonic/gin"
	"go-learn/utils"
	"net/http"
)

func OkResult(ctx *gin.Context, code int, data interface{}, msg string) {
	ctx.JSONP(http.StatusOK, gin.H{"code": code, "data": data, "msg": msg, "uuid": utils.NewUUID()})
}

func OkDataMsg(ctx *gin.Context, data interface{}, msg string) {
	ctx.JSONP(http.StatusOK, gin.H{"code": CodeOk, "data": data, "msg": msg, "uuid": utils.NewUUID()})
}

func Ok(ctx *gin.Context, data interface{}) {
	ctx.JSONP(http.StatusOK, gin.H{"code": CodeOk, "data": data, "msg": nil, "uuid": utils.NewUUID()})
}

func ErrorResult(ctx *gin.Context, code int, msg string) {
	ctx.JSONP(http.StatusOK, gin.H{"code": code, "data": nil, "msg": msg, "uuid": utils.NewUUID()})
}

func ErrorMsg(ctx *gin.Context, msg string) {
	ctx.JSONP(http.StatusOK, gin.H{"code": CodeError, "data": nil, "msg": msg, "uuid": utils.NewUUID()})
}

func Error(ctx *gin.Context, err error) {
	ctx.JSONP(http.StatusOK, gin.H{"code": CodeError, "data": nil, "msg": err.Error(), "uuid": utils.NewUUID()})
}
