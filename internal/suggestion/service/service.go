package service

import (
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"mitra-kirim-be-mgmt/internal/suggestion/repository"
)

type Suggestion struct {
	repo      *repository.Suggestion
	publisher *redis.Client
	logger    *logrus.Logger
}

func NewService(repo *repository.Suggestion, publisher *redis.Client, logger *logrus.Logger) *Suggestion {
	return &Suggestion{repo, publisher, logger}
}
