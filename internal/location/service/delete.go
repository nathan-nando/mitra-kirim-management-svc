package service

func (s *Location) Delete(id int) (int, error) {
	id, err := s.Repo.Delete(id)
	if err != nil {
		return id, err
	}
	return id, nil
}
