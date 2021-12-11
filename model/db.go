package model

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DbInstance *gorm.DB

func InitDB() {
	mysqlCfg := viper.GetStringMapString("mysql")
	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlCfg["user"], mysqlCfg["password"], mysqlCfg["host"], mysqlCfg["port"], mysqlCfg["database"])

	var err error
	DbInstance, err = gorm.Open(mysql.Open(dbSource), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = DbInstance.AutoMigrate(&User{})

	if err != nil {
		panic(err)
	}
}
