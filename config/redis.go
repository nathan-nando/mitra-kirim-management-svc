package config

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type RedisConfig struct {
	Client *redis.Client
}

func NewRedisPublisher(cfg *Config, log *logrus.Logger) *RedisConfig {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("localhost:%s", cfg.RedisPort),
		Password: "",
		DB:       0,
	})

	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	return &RedisConfig{
		Client: rdb,
	}
}

func (r *RedisConfig) Publish(ctx context.Context, channel, message string) error {
	return r.Client.Publish(ctx, channel, message).Err()
}
