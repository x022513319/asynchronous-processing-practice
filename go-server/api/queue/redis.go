package queue

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// 定義 client 可以使用的功能
type Client interface {
	Enqueue(ctx context.Context, payload []byte) error
}

type RedisClient struct {
	rdb       *redis.Client
	queueName string
}

func NewRedisClient(addr string) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	return &RedisClient{
		rdb:       rdb,
		queueName: "tasks",
	}
}

func (c *RedisClient) Enqueue(ctx context.Context, payload []byte) error {
	return c.rdb.LPush(ctx, c.queueName, payload).Err()
}
