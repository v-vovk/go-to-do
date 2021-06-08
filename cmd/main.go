package main

import (
	"github.com/spf13/viper"
	todo "github.com/wolfykey/go-to-do"
	"github.com/wolfykey/go-to-do/pkg/handler"
	"github.com/wolfykey/go-to-do/pkg/repository"
	"github.com/wolfykey/go-to-do/pkg/service"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewSRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server on port: %s", err.Error())
	}
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
