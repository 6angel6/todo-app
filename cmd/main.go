package main

import (
	"TODOapp/pkg/handler"
	"TODOapp/pkg/repository"
	"TODOapp/pkg/service"
	todo "TODOapp/server"
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	"os"
	"os/signal"
	"syscall"
)

//	@title			Todo App API
//	@version		1.0
//	@description	API server for TodoList Application

//	@host		localhost:8000
//	@BasePath	/

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := initConfig(); err != nil {
		logrus.Fatal(err)
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatal(err)
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatal("Error to init fb", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	go func() {
		if err := srv.Run(viper.GetString("server.port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error occurred while running HTTP server: %s", err.Error())
		}
	}()

	logrus.Print("TodoApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("TodoApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Error occurred on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("Error occurred on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
