package service

import (
	"mitra-kirim-be-mgmt/internal/configuration/model"
	"mitra-kirim-be-mgmt/pkg/converter"
)

func (s *Configuration) UpdateToko(req model.UpdateTokoRequest) ([]converter.KeyValue, error) {
	keyVal := converter.ConvertKeyValue(req)

	err := s.Repo.UpdateByKey(keyVal)
	if err != nil {
		return nil, err
	}

	return keyVal, nil
}
