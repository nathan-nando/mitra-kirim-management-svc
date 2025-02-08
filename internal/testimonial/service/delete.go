package service

func (s *Testimonial) Delete(id int) (bool, error) {
	ok, err := s.Repository.Delete(id)
	if err != nil {
		return ok, err
	}

	return ok, nil
}
