package service

import (
	"context"
	"mime/multipart"
	"mitra-kirim-be-mgmt/pkg/contants"
)

func (s *Configuration) SaveServiceImage(context context.Context, req *multipart.FileHeader) (string, error) {
	newFileName, err := s.FileSvc.UploadFile(req, "/mk-storage/assets")
	if err != nil {
		return "", err
	}

	err = s.Cache.Del(context, contants.CacheConfiguration)
	if err != nil {
		return "", err
	}

	return newFileName, nil
}
