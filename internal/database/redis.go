package database

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func InitRedis(addr string, pass string, db int) (redis.UniversalClient, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       db,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		rdb.Close()
		return nil, err
	}
	return rdb, nil
}
