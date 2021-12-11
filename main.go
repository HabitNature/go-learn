package main

import (
	"fmt"
	"go-learn/common"
	"go-learn/config"
	"go-learn/model"
	"go-learn/route"
)

func main() {
	fmt.Println("go learn start")

	defer func() {
		err := recover()

		if err != nil {

		}
	}()

	inits()
}

func inits() {
	common.InitPaths()
	config.InitConfig()
	model.InitDB()
	route.InitGinRoutes()
}
