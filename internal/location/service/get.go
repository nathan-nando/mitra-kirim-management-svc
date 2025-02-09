package service

import (
	"context"
	"encoding/json"
	"mitra-kirim-be-mgmt/internal/location/model"
	"mitra-kirim-be-mgmt/pkg/contants"
	"time"
)

func (s *Location) GetAll(context context.Context, limit int, offset int) ([]model.Location, error) {
	var results []model.Location

	dataJSON, err := s.Cache.LRange(context, contants.CacheLocations, 0, -1)
	if err == nil && len(dataJSON) > 0 {
		for _, itemJSON := range dataJSON {
			var item model.Location
			if err := json.Unmarshal([]byte(itemJSON), &item); err != nil {
				return nil, err
			}
			results = append(results, item)
		}
		s.Logger.Info("locations CACHE")
		return results, nil
	}

	dataDB, err := s.Repo.FindAll(limit, offset)
	if err != nil {
		s.Logger.Error(err)
		return []model.Location{}, err
	}
	s.Logger.Info("locations DB")

	if len(dataDB) > 0 {
		pipe := s.Cache.Pipeline()
		args := make([]interface{}, 0, len(dataDB))
		for _, item := range dataDB {
			jsonData, err := json.Marshal(item)
			if err != nil {
				return nil, err
			}
			args = append(args, jsonData)
		}

		pipe.RPush(context, contants.CacheLocations, args...)
		pipe.Expire(context, contants.CacheLocations, time.Duration(s.CacheTime)*time.Minute)
		_, err = pipe.Exec(context)
		if err != nil {
			return nil, err
		}
	}

	return dataDB, nil
}
