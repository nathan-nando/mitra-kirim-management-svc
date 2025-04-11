package service

import (
	"context"
)

func (s *Location) Count(context context.Context) (int64, error) {
	count, err := s.Repo.CountAll()
	if err != nil {
		return 0, err
	}
	return count, nil
}
