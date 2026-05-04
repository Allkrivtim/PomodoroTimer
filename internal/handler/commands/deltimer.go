package commands

import (
	"context"
	"log"
	"pomodoroBot/internal/utils"

	tb "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/redis/go-redis/v9"
)

func DeleteCommand(update *tb.Update, bot *tb.BotAPI, rdb redis.UniversalClient) error {
	ctx := context.Background()
	errmsg := tb.NewMessage(update.Message.Chat.ID, "Что-то сломалось. Попробуй еще раз, только позже.")

	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	check_timer, err := utils.CheckTimers(ctx, update.Message.From.UserName, rdb)
	if check_timer == "" || err != nil {
		if err != nil {
			return err
		}
		msg := tb.NewMessage(update.Message.Chat.ID, "У тебя нету активного таймера или у нас что-то сломалось.")
		_, err := bot.Send(msg)
		if err != nil {
			bot.Send(errmsg)
			return err
		}
	} else {
		_, err := utils.DeleteTimer(ctx, update.Message.From.UserName, rdb)
		if err != nil {
			bot.Send(errmsg)
			return err
		}
		msg := tb.NewMessage(update.Message.Chat.ID, "Таймер удален!")
		bot.Send(msg)
	}
	return nil
}
