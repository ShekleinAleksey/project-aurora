package main

import (
	"log"
	"net/http"

	"github.com/ShekleinAleksey/project-aurora/config"
	"github.com/ShekleinAleksey/project-aurora/internal/handler"
	"github.com/ShekleinAleksey/project-aurora/internal/repository"
	"github.com/ShekleinAleksey/project-aurora/internal/service"
	"github.com/ShekleinAleksey/project-aurora/pkg/postgres"
	"github.com/sirupsen/logrus"
)

// @title Aurora
// @version 1.0
// @description Система для контроля материалов и отслеживания статусов заказов
// @host localhost:8080
// @BasePath /api/v1
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error get env variables %v", err)
	}
	log.Println("cfg", cfg)
	db, err := postgres.NewDB(cfg)
	if err != nil {
		log.Fatalf("Error opening database %v", err)
	}
	defer db.Close()

	logrus.Info("Initializing repository...")
	repo := repository.NewRepository(db)
	logrus.Info("Initializing service...")
	service := service.NewService(repo)
	logrus.Info("Initializing handler...")
	handlers := handler.NewHandler(service)

	router := handlers.InitRoutes()

	// signalChan := make(chan os.Signal)
	// signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// <-signalChan

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", router)
}
