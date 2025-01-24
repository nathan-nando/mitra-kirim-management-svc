package service

import (
	"mitra-kirim-be-mgmt/internal/location/model"
)

func (s *Location) Create(request *model.LocationRequest) (int, error) {
	id, err := s.Repo.Create(request)
	if err != nil {
		return id, err
	}
	return id, nil
}
