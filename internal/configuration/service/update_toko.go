package service

import (
	"context"
	"mitra-kirim-be-mgmt/internal/configuration/model"
	"mitra-kirim-be-mgmt/pkg/contants"
	"mitra-kirim-be-mgmt/pkg/converter"
)

func (s *Configuration) UpdateToko(context context.Context, req model.UpdateTokoRequest) ([]converter.KeyValue, error) {
	keyVal := converter.ConvertKeyValue(req)

	err := s.Repo.UpdateByKey(keyVal)
	if err != nil {
		return nil, err
	}

	err = s.Cache.Del(context, contants.CacheConfiguration)
	if err != nil {
		return nil, err
	}

	return keyVal, nil
}
