package service

import (
	"context"
	"encoding/json"
	"mitra-kirim-be-mgmt/internal/suggestion/model"
)

func (s *Suggestion) Create(ctx context.Context, request *model.SuggestionCreate) (string, error) {
	_, err := s.repo.Create(request)
	if err != nil {
		return "", err
	}

	jsonMessage, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	err = s.publisher.Publish(context.Background(), "EMAIL_NOTIFICATION", jsonMessage).Err()
	if err != nil {
		return "", err
	}

	return "OK", nil
}
