package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ShekleinAleksey/project-aurora/config"
	"github.com/ShekleinAleksey/project-aurora/internal/handler"
	"github.com/ShekleinAleksey/project-aurora/internal/repository"
	"github.com/ShekleinAleksey/project-aurora/internal/service"
	"github.com/ShekleinAleksey/project-aurora/pkg/postgres"
	"github.com/sirupsen/logrus"
)

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

	// 	key := "7576443951:AAGEbz7-S8AmrF1-ZS-lIEaSFPXpVCjXlqc"
	// 	bot, err := tgbotapi.NewBotAPI(key)
	// 	if err != nil {
	// 		fmt.Println("Cannot init tgbot")
	// 	}
	// 	log.Printf("Auth on account %s", bot.Self.UserName)

	// 	u := tgbotapi.NewUpdate(0)
	// 	u.Timeout = 60

	// 	updates := bot.GetUpdatesChan(u)
	// 	for update := range updates {
	// 		if update.Message != nil {
	// 			text := update.Message.Text
	// 			chatID := update.Message.Chat.ID
	// 			userID := update.Message.From.ID
	// 			log.Printf("[%s](%d) %s", update.Message.From.UserName, userID, text)
	// 			// msg := tgbotapi.NewMessage(chatID, text)

	// 			switch text {
	// 			case "/start":
	// 				msg := tgbotapi.NewMessage(chatID, "Привет! Я бот для учета расходов твоего 3D магазина.")
	// 				bot.Send(msg)

	// 			case "/help":
	// 				helpText := `Доступные команды:
	// /expense <сумма> <категория> <описание> - добавить расход
	// /expenses - показать сегодняшние расходы
	// /total - итог за сегодня`
	// 				msg := tgbotapi.NewMessage(chatID, helpText)
	// 				bot.Send(msg)

	// 			case "/info":
	// 				msg := tgbotapi.NewMessage(chatID, "Я простой Telegram бот на Go")
	// 				bot.Send(msg)

	// 			case "/calc":
	// 				msg := tgbotapi.NewMessage(chatID, "Разработано с использованием tgbotapi")
	// 				bot.Send(msg)

	// 			default:
	// 				// Если это не команда, делаем эхо
	// 				msg := tgbotapi.NewMessage(chatID, text)
	// 				bot.Send(msg)
	// 			}

	// 			// bot.Send(msg)
	// 		}
	// 	}
	router := handlers.InitRoutes()

	quit := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", router)
}
