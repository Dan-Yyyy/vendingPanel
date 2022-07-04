package main

import (
	"flag"
	back "github.com/Dan-Yyyy/vendingPanel.git"
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/handler"
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/message"
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/repository"
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/service"
	"github.com/Dan-Yyyy/vendingPanel.git/seed"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("%s: %s", message.ReadEnvError, err.Error())
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
		log.Fatalf("%s: %s", message.DBConnectionError, err.Error())
	}

	// Отлов команды на выполнение сидеров
	handleArgs(db)

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(back.Server)
	if err := srv.Run(os.Getenv("APPLICATION_PORT"), handlers.InitRoutes()); err != nil {
		log.Fatalf("%s: %s", message.StartHTTPServerError, err.Error())
	}
}

func handleArgs(db *sqlx.DB) {
	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			seed.Execute(db, args[1:]...)
			os.Exit(0)
		}
	}
}
