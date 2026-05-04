package commands

import (
	"context"
	"log"
	"pomodoroBot/internal/handler/updates"
	"pomodoroBot/internal/utils"

	tb "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/redis/go-redis/v9"
)

func NewtimerCommand(update *tb.Update, bot *tb.BotAPI, rdb redis.UniversalClient) error {
	ctx := context.Background()
	errmsg := tb.NewMessage(update.Message.Chat.ID, "Что-то сломалось. Попробуй еще раз, только позже.")

	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	check_timer, err := utils.CheckTimers(ctx, update.Message.From.UserName, rdb)
	if check_timer != "" || err != nil {
		if err != nil {
			return err
		}
		msg := tb.NewMessage(update.Message.Chat.ID, "У тебя уже есть таймер, или у нас что-то сломалось.")
		_, err := bot.Send(msg)
		if err != nil {
			bot.Send(errmsg)
			return err
		}
	} else {
		_, err := utils.CreateTimer(ctx, update.Message.From.UserName, rdb)
		if err != nil {
			bot.Send(errmsg)
			return err
		}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		for key := range updates.GetRedisUpdate(ctx, rdb) {
			if update.Message.From.UserName == key {
				msg := tb.NewMessage(update.Message.Chat.ID, "Время вышло!")
				bot.Send(msg)
			}
		}
		msg := tb.NewMessage(update.Message.Chat.ID, "Таймер поставлен. Работаем!")
		bot.Send(msg)
	}
	return nil
}
