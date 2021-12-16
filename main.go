package main

import (
	"fmt"
	"go-learn/common"
	"go-learn/common/loggger"
	"go-learn/config"
	"go-learn/model"
	"go-learn/route"
)

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
