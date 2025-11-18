package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ShekleinAleksey/project-aurora/config"
	"github.com/ShekleinAleksey/project-aurora/pkg/postgres"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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

	key := "7576443951:AAGEbz7-S8AmrF1-ZS-lIEaSFPXpVCjXlqc"
	bot, err := tgbotapi.NewBotAPI(key)
	if err != nil {
		fmt.Println("Cannot init tgbot")
	}
	log.Printf("Auth on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil {
			text := update.Message.Text
			chatID := update.Message.Chat.ID
			userID := update.Message.From.ID
			log.Printf("[%s](%d) %s", update.Message.From.UserName, userID, text)
			// msg := tgbotapi.NewMessage(chatID, text)

			switch text {
			case "/start":
				msg := tgbotapi.NewMessage(chatID, "Привет! Я бот для учета расходов твоего 3D магазина.")
				bot.Send(msg)

			case "/help":
				helpText := `Доступные команды:
/expense <сумма> <категория> <описание> - добавить расход
/expenses - показать сегодняшние расходы
/total - итог за сегодня`
				msg := tgbotapi.NewMessage(chatID, helpText)
				bot.Send(msg)

			case "/info":
				msg := tgbotapi.NewMessage(chatID, "Я простой Telegram бот на Go")
				bot.Send(msg)

			case "/calc":
				msg := tgbotapi.NewMessage(chatID, "Разработано с использованием tgbotapi")
				bot.Send(msg)

			default:
				// Если это не команда, делаем эхо
				msg := tgbotapi.NewMessage(chatID, text)
				bot.Send(msg)
			}

			// bot.Send(msg)
		}
	}
	http.ListenAndServe(":8080", nil)
}
