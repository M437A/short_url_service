package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	"net/http"
	"short_url/config"
	"short_url/iternal/exceptions"
	"short_url/iternal/kafka"
	"short_url/iternal/routers"
	"short_url/iternal/services"
)

func main() {
	config.YamlConfig()

	dbpool := config.GetDBPool()
	defer dbpool.Close()

	go kafka.KafkaRun()

	analytics := config.RunAnalytic()
	defer analytics.Shutdown()

	redis := config.RunRedis()
	defer redis.Close()

	router := chi.NewRouter()
	routers.CreateRouters(router)

	go services.CheckQueueOpenUrl()

	err := http.ListenAndServe(viper.GetString("port"), router)
	exceptions.CheckMainException(err)
}
