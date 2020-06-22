package app

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"

	"go-backend/pkg/config"
	"go-backend/pkg/log"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(ctx context.Context, cfg *config.Config) (*RedisClient, error) {
	logger := log.GetLogger(ctx)
	logger.Info("creating new Redis client")
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("redis:%s", cfg.REDIS_PORT),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &RedisClient{client: rdb}, nil
}
