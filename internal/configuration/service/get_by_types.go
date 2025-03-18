package service

import (
	"context"
	"encoding/json"
	"mitra-kirim-be-mgmt/internal/configuration/model"
	"mitra-kirim-be-mgmt/pkg/contants"
	"time"
)

func (s *Configuration) GetByTypes(context context.Context, types []string) ([]model.Configuration, error) {
	//var configs []model.Configuration
	//
	//dataJSON, err := s.Cache.LRange(context, contants.CacheConfiguration, 0, -1)
	//if err == nil && len(dataJSON) > 0 {
	//	for _, itemJSON := range dataJSON {
	//		var item model.Configuration
	//		if err = json.Unmarshal([]byte(itemJSON), &item); err != nil {
	//			return nil, err
	//		}
	//		configs = append(configs, item)
	//	}
	//	s.Logger.Info("config CACHE")
	//	return configs, nil
	//}

	dataDB, err := s.Repo.FindByTypes(types)
	if err != nil {
		return dataDB, err
	}
	s.Logger.Info("config DB")

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

		pipe.RPush(context, contants.CacheConfiguration, args...)
		pipe.Expire(context, contants.CacheConfiguration, time.Duration(s.CacheTime)*time.Minute)
		_, err = pipe.Exec(context)
		if err != nil {
			return nil, err
		}
	}

	return dataDB, nil
}
