package service

import (
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type Publisher struct {
	client   *redis.Client
	Logger   *logrus.Logger
	MaxRetry int
}

func NewPublisher(client *redis.Client, logger *logrus.Logger, maxRetry int) *Publisher {
	return &Publisher{client: client, Logger: logger, MaxRetry: maxRetry}
}
