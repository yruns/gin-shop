package main

import (
	"fmt"
	"gin-shop/controller"
	"gin-shop/middlewares"
	"gin-shop/settings"
	"gin-shop/settings/mysql"
	"gin-shop/settings/redis"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

func main() {

	// 1.加载配置
	if err := settings.InitSettings(); err != nil {
		// 配置加载失败
		log.Fatal(err.Error())
		return
	}

	// 2.初始化Mysql
	if err := mysql.InitMysql(); err != nil {
		log.Fatal(err.Error())
		return
	}

	// 3.初始化Redis
	if err := redis.InitRedis(); err != nil {
		log.Fatal(err.Error())
		return
	}

	app := gin.Default()
	// 设置全局前置中间件
	app.Use(Cors(), middlewares.AuthMiddleware())
	// 注册路由
	registerRouter(app)

	// 启动
	if err := app.Run(fmt.Sprintf(":%d", viper.GetInt("app.port"))); err != nil {
		log.Fatal(err.Error())
	}
}

// 注册路由
func registerRouter(router *gin.Engine) {
	new(controller.UserController).Register(router)
}

// Cors 设置跨域访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Allow-Credentials", "true")

		}
		if method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
