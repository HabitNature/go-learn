package middleware

import (
	"github.com/gin-gonic/gin"
	"go-learn/common"
	"go-learn/utils"
	"net/http"
	"time"
)

func Jwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		if token == "" {
			ctx.JSONP(http.StatusUnauthorized, gin.H{"code": 400, "data": nil, "msg": "未设置token", "traceId": utils.NewUUID()})
			ctx.Abort()
			return
		}

		claims, err := common.ParseToken(token)

		if err != nil {
			ctx.JSONP(http.StatusUnauthorized, gin.H{"code": 400, "data": nil, "msg": "解析token失败", "traceId": utils.NewUUID()})
			ctx.Abort()
			return
		}

		if claims.ExpiresAt <= time.Now().Unix() {
			ctx.JSONP(http.StatusUnauthorized, gin.H{"code": 400, "data": nil, "msg": "token已过期", "traceId": utils.NewUUID()})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
