package main

import (
	"log"
	"os"
	"pomodoroBot/internal/database"
	"pomodoroBot/internal/handler/commands"

	tb "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	botToken := os.Getenv("TOKEN")
	if botToken == "" {
		panic("TOKEN environment variable not set")
	}
	bot, err := tb.NewBotAPI(botToken)
	if err != nil {
		panic(err)
	}
	bot.Debug = false
	rdb, err := database.InitRedis(os.Getenv("REDIS_URL"), "", 0)
	if err != nil {
		panic("Redis initialization failure")
	}

	log.Printf("Authorized on bot @%s", bot.Self.UserName)

	updateConfig := tb.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil || !update.Message.IsCommand() {
			continue
		}
		switch update.Message.Command() {
		case "start":
			go commands.StartCommand(&update, bot)
		case "help":
			go commands.HelpCommand(&update, bot)
		case "newtimer":
			go commands.NewtimerCommand(&update, bot, rdb)
		case "deltimer":
			go commands.DeleteCommand(&update, bot, rdb)
		}
	}
}
