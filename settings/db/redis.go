package db

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var Redis *redis.Client

func InitRedis() (err error) {
	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.database"),
	})

	_, err = Redis.Ping(context.Background()).Result()
	return err
}
