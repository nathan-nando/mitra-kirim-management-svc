package service

import (
	"context"
)

func (s *Suggestion) Count(ctx context.Context) (int64, error) {
	result, err := s.Repo.CountAll()

	if err != nil {
		return 0, err
	}

	return result, nil
}
