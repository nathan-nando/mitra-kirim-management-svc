package service

import (
	"mitra-kirim-be-mgmt/internal/location/model"
)

func (s *Location) Update(request *model.LocationRequest) (int, error) {
	id, err := s.Repo.Update(request)
	if err != nil {
		return id, err
	}
	return id, nil
}
