package settings

import (
	"github.com/spf13/viper"
	"log"
)

func InitSettings() (err error) {

	// 加载项目配置
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		// 读取失败
		log.Fatal(err.Error())
	}

	return
}
