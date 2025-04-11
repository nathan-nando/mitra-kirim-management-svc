package service

import "context"

func (s *Testimonial) Count(ctx context.Context) (int64, error) {
	count, err := s.Repository.CountAll()
	if err != nil {
		return count, err
	}

	return count, nil
}
