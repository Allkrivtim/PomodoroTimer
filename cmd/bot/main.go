package main

import (
	"log"
	"os"
	"pomodoroBot/internal/handler/commands"

	tb "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	botToken := os.Getenv("TOKEN")
	if botToken == "" {
		panic("TOKEN environment variable not set")
	}
	bot, _ := tb.NewBotAPI(botToken)
	bot.Debug = true

	log.Printf("Authorized on bot @%s", bot.Self.UserName)

	updateConfig := tb.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message.IsCommand() {
			switch update.Message.Text {
			case "/start":
				go commands.StartCommand(&update, bot)
			case "/help":
				commands.HelpCommand(&update, bot)
			}

		}
	}
}
