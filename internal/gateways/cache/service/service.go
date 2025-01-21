package service

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type Cache struct {
	client *redis.Client
}

func NewCache(client *redis.Client) *Cache {
	return &Cache{client: client}
}

func (r *Cache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *Cache) Get(ctx context.Context, key string, value interface{}) error {
	return r.client.Get(ctx, key).Scan(value)
}

func (r *Cache) Del(ctx context.Context, keys ...string) error {
	return r.client.Del(ctx, keys...).Err()
}
