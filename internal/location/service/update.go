package service

import (
	"mitra-kirim-be-mgmt/internal/location/model"
)

func (s *Location) Update(request *model.LocationRequest, username string) (int, error) {
	id, err := s.Repo.Update(request, username)
	if err != nil {
		return id, err
	}
	return id, nil
}
