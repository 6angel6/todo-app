package main

import (
	"TODOapp/pkg/handler"
	"TODOapp/pkg/repository"
	"TODOapp/pkg/service"
	todo "TODOapp/server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatal(err)
	}
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
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
		log.Fatal("Error to init fb", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(todo.Server)
	if err := server.Run(viper.GetString("server.port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error starting server: %v", err)
	}

}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
