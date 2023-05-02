package redis

import (
	"context"
	"fmt"
	Redis "github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var DB *Redis.Client

func InitRedis() (err error) {
	redis := Redis.NewClient(&Redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.database"),
	})

	_, err = redis.Ping(context.Background()).Result()
	return err
}
