package main

import (
	"TODOapp/pkg/handler"
	"TODOapp/pkg/repository"
	"TODOapp/pkg/service"
	todo "TODOapp/server"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(todo.Server)
	if err := server.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
