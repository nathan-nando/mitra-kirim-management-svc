package service

import (
	"github.com/sirupsen/logrus"
	"mitra-kirim-be-mgmt/internal/gateways/publisher/service"
	"mitra-kirim-be-mgmt/internal/suggestion/repository"
)

type Suggestion struct {
	Repo      *repository.Suggestion
	Logger    *logrus.Logger
	Publisher *service.Publisher
}

func NewSuggestion(repo *repository.Suggestion, logger *logrus.Logger, publisher *service.Publisher) *Suggestion {
	return &Suggestion{repo, logger, publisher}
}
