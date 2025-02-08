package service

import "mitra-kirim-be-mgmt/internal/testimonial/model"

func (s *Testimonial) UpdateSlide(req *model.TestimonialUpdateSlide, userID string) (bool, error) {
	ok, err := s.Repository.UpdateSlide(req, userID)
	if err != nil {
		return ok, err
	}

	return ok, nil
}
