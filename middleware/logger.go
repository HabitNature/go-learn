package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-learn/common/loggger"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		end := time.Now()
		m := ctx.Request.Method
		url := ctx.Request.URL
		ip := ctx.ClientIP()
		code := ctx.Writer.Status()
		loggger.GLogger.WithFields(logrus.Fields{
			"host":      ip,
			"method":    m,
			"path":      url.Path,
			"code":      code,
			"consuming": end.Sub(start).Milliseconds(),
		}).Info("")
	}
}
