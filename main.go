package main

import (
	"fmt"
	"go-learn/common"
	"go-learn/common/loggger"
	"go-learn/config"
	"go-learn/model"
	"go-learn/route"
)

// 参见 : https://ieevee.com/tech/2018/04/19/go-swag.html

// @title go-learn
// @version 1.0
// @description go-learn apis
// @termsOfService http://www.habit.com

// @contact.name habit
// @contact.url http://www.habit.com
// @contact.email go-lean@habit.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8888
// @BasePath /v1
func main() {
	fmt.Println("go learn start")

	defer func() {
		err := recover()
		if err != nil {
			loggger.GLogger.Panicf("程序异常退出 : %v", err)
			loggger.GLogger.Exit(-1)
			panic(err)
		}
	}()

	inits()
}

func inits() {
	common.InitPaths()
	config.InitConfig()
	loggger.InitLogger()
	loggger.GLogger.Info("初始化日志完毕")
	model.InitDB()
	loggger.GLogger.Info("初始化数据库完毕")
	route.InitGinRoutes()
}
