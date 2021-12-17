package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go-learn/controller/v1"
	"go-learn/middleware"
	"go-learn/utils"
	"net/http"
)

func InitGinRoutes() {
	g := gin.New()

	g.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "data": "", "msg": "路径不存在", "traceId": utils.NewUUID()})
	})

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 设置参数验证中间件
	g.Use(middleware.ValidateMiddleware())

	// 设置gin日志中间件
	g.Use(middleware.Logger())

	// 设置jwt token中间件
	//g.Use(middleware.Jwt())

	userController := controller.UserController{}
	user := g.Group("/v1/user")
	{
		user.POST("/create", userController.Create)
	}

	addr := fmt.Sprintf("localhost:%s", viper.GetString("server_port"))
	err := g.Run(addr)

	if err != nil {
		panic(err)
	}
}
