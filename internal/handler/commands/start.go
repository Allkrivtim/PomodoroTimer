package commands

import (
	"log"

	tb "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const startMessage = `Бот для продуктивной работы. Комманды: /help`

func StartCommand(update *tb.Update, bot *tb.BotAPI) {
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	msg := tb.NewMessage(update.Message.Chat.ID, startMessage)

	msg.ReplyToMessageID = update.Message.MessageID

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Error sending message: %v", err)
	}
}
