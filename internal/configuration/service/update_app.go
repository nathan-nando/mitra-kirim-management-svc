package service

import (
	"mitra-kirim-be-mgmt/internal/configuration/model"
	"mitra-kirim-be-mgmt/pkg/converter"
)

func (s *Configuration) UpdateLogoApp(req *model.UpdateAppLogoRequest) ([]converter.KeyValue, error) {
	var result []converter.KeyValue

	newFileName, err := s.FileSvc.UploadFile(req.AppLogo, "/mk-storage/assets")
	if err != nil {
		return result, err
	}

	keyVal := converter.ConvertKeyValue(model.UpdateAppLogoFileName{
		AppLogo: newFileName,
	})

	err = s.Repo.UpdateByKey(keyVal)
	if err != nil {
		return nil, err
	}

	return keyVal, nil
}

func (s *Configuration) UpdateApp(req model.UpdateAppRequest) ([]converter.KeyValue, error) {
	keyVal := converter.ConvertKeyValue(req)

	err := s.Repo.UpdateByKey(keyVal)
	if err != nil {
		return nil, err
	}

	return keyVal, nil
}
