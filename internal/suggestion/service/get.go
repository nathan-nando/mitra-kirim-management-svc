package service

import (
	"context"
	"mitra-kirim-be-mgmt/internal/suggestion/model"
)

func (s Suggestion) Get(ctx context.Context) ([]model.Suggestion, error) {
	todo, err := s.repo.FindAll()

	if err != nil {
		return []model.Suggestion{}, err
	}

	return todo, nil
}
