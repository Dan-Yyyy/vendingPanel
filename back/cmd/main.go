package main

import (
	back "github.com/Dan-Yyyy/vendingPanel.git"
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/handler"
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/messages"
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/repository"
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("%s: %s", messages.ReadEnvError, err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	})

	if err != nil {
		log.Fatalf("%s: %s", messages.DBConnectionError, err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(back.Server)
	if err := srv.Run(os.Getenv("APPLICATION_PORT"), handlers.InitRoutes()); err != nil {
		log.Fatalf("%s: %s", messages.StartHTTPServerError, err.Error())
	}
}
