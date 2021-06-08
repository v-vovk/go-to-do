package main

import (
	todo "github.com/wolfykey/go-to-do"
	"github.com/wolfykey/go-to-do/pkg/handler"
	"github.com/wolfykey/go-to-do/pkg/repository"
	"github.com/wolfykey/go-to-do/pkg/service"
	"log"
)

func main() {
	repos := repository.NewSRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server on port: %s", err.Error())
	}
}
