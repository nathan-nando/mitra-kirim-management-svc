package service

import (
	"context"
	"mitra-kirim-be-mgmt/internal/location/model"
	"mitra-kirim-be-mgmt/pkg/contants"
)

func (s *Location) Create(context context.Context, request *model.LocationRequest, username string) (int, error) {
	id, err := s.Repo.Create(request, username)
	if err != nil {
		return id, err
	}

	err = s.Cache.Del(context, contants.CacheLocations)
	if err != nil {
		return 0, err
	}

	return id, nil
}
