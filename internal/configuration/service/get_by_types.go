package service

import "mitra-kirim-be-mgmt/internal/configuration/model"

func (s *Configuration) GetByTypes(types []string) ([]model.Configuration, error) {
	data, err := s.Repo.FindByTypes(types)
	if err != nil {
		return data, err
	}
	return data, nil
}
