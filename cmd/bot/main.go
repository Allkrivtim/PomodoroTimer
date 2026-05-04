package main

import (
	"os"

	tb "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	dotenv "github.com/joho/godotenv"
)

func main() {
	err := dotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	bot, _ := tb.NewBotAPI(os.Getenv("TOKEN"))

	bot.Debug = true
}
