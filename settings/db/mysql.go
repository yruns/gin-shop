package db

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var Mysql *gorm.DB

func InitMysql() (err error) {
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	url := viper.GetString("mysql.host")
	database := viper.GetString("mysql.database")

	// 自定义SQL日志打印模板
	customLog := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	connectURL := username + ":" + password + "@tcp(" + url + ")/" + database
	fmt.Println(connectURL)
	Mysql, err = gorm.Open(mysql.Open(connectURL+"?charset=utf8mb4"), &gorm.Config{Logger: customLog})

	return err
}
