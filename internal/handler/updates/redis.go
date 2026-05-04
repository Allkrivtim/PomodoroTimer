package updates

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func GetRedisUpdate(ctx context.Context, rdb redis.UniversalClient) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)
		pubsub := rdb.Subscribe(ctx, "__keyevent@0__:expired")
		defer pubsub.Close()

		for msg := range pubsub.Channel() {
			select {
			case ch <- msg.Payload:
			case <-ctx.Done():
				return
			}
		}
	}()

	return ch
}
