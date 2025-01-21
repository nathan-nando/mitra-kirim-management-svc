package service

import (
	"context"
	modelPub "mitra-kirim-be-mgmt/internal/gateways/publisher/model"
	"mitra-kirim-be-mgmt/internal/suggestion/model"
	"time"
)

func (s *Suggestion) Create(ctx context.Context, request *model.SuggestionCreate) (string, error) {
	suggestion, err := s.Repo.Create(request)
	if err != nil {
		return "", err
	}

	go func() {
		bgCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		s.Publisher.PublishWithRetry(
			bgCtx,
			&modelPub.PublisherEmail{
				Name:    request.Name,
				Email:   request.Email,
				Message: request.Message,
			})
	}()

	return suggestion.ID, nil
}
