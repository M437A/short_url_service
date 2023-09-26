package config

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var Redis *redis.Client
var Ctx context.Context

func RunRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.host"),
		Password: "",
		DB:       0,
	})
	Ctx = context.Background()
	Redis = client
	return client
}
