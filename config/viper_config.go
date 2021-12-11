package config

import (
	"github.com/spf13/viper"
	"go-learn/common"
)

func InitConfig() {
	viper.SetConfigName("config_dev")
	viper.SetConfigType("yml")
	viper.AddConfigPath(common.WORK_DIR + "/config")

	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
}
