package service

import (
	"mitra-kirim-be-mgmt/internal/suggestion/repository"
)

type Suggestion struct {
	repo *repository.Suggestion
}

func NewService(repo *repository.Suggestion) *Suggestion {
	return &Suggestion{repo}
}
