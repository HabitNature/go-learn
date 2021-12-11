package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-learn/controller/v1"
	"go-learn/utils"
	"net/http"
)

func InitGinRoutes() {
	g := gin.New()

	g.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "data": "", "msg": "路径不存在", "traceId": utils.NewUUID()})
	})

	user := g.Group("/v1/user")
	userController := controller.UserController{}
	user.POST("/create", userController.Create)

	addr := fmt.Sprintf("localhost:%s", viper.GetString("serverPort"))
	err := g.Run(addr)

	if err != nil {
		panic(err)
	}
}
