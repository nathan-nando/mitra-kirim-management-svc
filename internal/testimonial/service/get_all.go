package service

import (
	"mitra-kirim-be-mgmt/internal/testimonial/model"
)

func (s *Testimonial) GetAll(limit int, offset int) ([]model.Testimonial, error) {
	response, err := s.Repository.GetAll(limit, offset)
	if err != nil {
		return response, err
	}

	return response, nil
}
