package service

import (
	"context"
	"mitra-kirim-be-mgmt/internal/configuration/model"
	"mitra-kirim-be-mgmt/pkg/contants"
	"mitra-kirim-be-mgmt/pkg/converter"
)

func (s *Configuration) UpdateHero(context context.Context, req *model.UpdateHeroLogoRequest) ([]converter.KeyValue, error) {
	var result []converter.KeyValue

	newFileName, err := s.FileSvc.UploadFile(req.HeroImg, "/mk-storage/assets")
	if err != nil {
		return result, err
	}

	keyVal := converter.ConvertKeyValue(model.UpdateHeroKeyVal{
		HeroImg:  newFileName,
		HeroDesc: req.HeroDesc,
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
