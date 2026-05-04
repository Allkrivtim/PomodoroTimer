package utils

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

func CheckTimers(ctx context.Context, usrid string, rdb redis.UniversalClient) (string, error) {
	if usrid == "" {
		return "", errors.New("no user id")
	}
	res, err := rdb.Get(ctx, usrid).Result()
	if errors.Is(err, redis.Nil) {
		return "", nil // таймера нет — это не ошибка
	}
	if err != nil {
		return "", err
	}
	return res, nil
}

func CreateTimer(ctx context.Context, usrid string, rdb redis.UniversalClient) (string, error) {
	if usrid == "" {
		return "", errors.New("No user id")
	}
	res, err := rdb.Set(ctx, usrid, time.Now(), 25*time.Minute).Result()
	if err != nil {
		return "ERROR", err
	}
	return res, nil
}

func DeleteTimer(ctx context.Context, usrid string, rdb redis.UniversalClient) (int64, error) {
	if usrid == "" {
		return 0, errors.New("No user id")
	}
	res, err := rdb.Del(ctx, usrid).Result()
	if err != nil {
		return 0, err
	}
	return res, nil
}
