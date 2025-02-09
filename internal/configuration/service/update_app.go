package service

import (
	"context"
	"mitra-kirim-be-mgmt/internal/configuration/model"
	"mitra-kirim-be-mgmt/pkg/contants"
	"mitra-kirim-be-mgmt/pkg/converter"
)

func (s *Configuration) UpdateLogoApp(context context.Context, req *model.UpdateAppLogoRequest) ([]converter.KeyValue, error) {
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

	err = s.Cache.Del(context, contants.CacheConfiguration)
	if err != nil {
		return nil, err
	}

	return keyVal, nil
}

func (s *Configuration) UpdateApp(context context.Context, req model.UpdateAppRequest) ([]converter.KeyValue, error) {
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
