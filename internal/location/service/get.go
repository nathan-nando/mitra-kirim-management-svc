package service

import "mitra-kirim-be-mgmt/internal/location/model"

func (s *Location) GetAll(limit int, offset int) ([]model.Location, error) {
	result, err := s.Repo.FindAll(limit, offset)
	if err != nil {
		s.Log.Error(err)
		return []model.Location{}, err
	}
	return result, nil
}
