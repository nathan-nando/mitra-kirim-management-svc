package service

import (
	"mitra-kirim-be-mgmt/internal/configuration/model"
	"mitra-kirim-be-mgmt/pkg/converter"
)

func (s *Configuration) UpdateSocial(req model.UpdateSocialRequest) ([]converter.KeyValue, error) {
	keyVal := converter.ConvertKeyValue(req)

	err := s.Repo.UpdateByKey(keyVal)
	if err != nil {
		return nil, err
	}

	return keyVal, nil
}
