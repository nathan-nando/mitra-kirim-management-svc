package service

import (
	"github.com/sirupsen/logrus"
	"mitra-kirim-be-mgmt/internal/configuration/repository"
	"mitra-kirim-be-mgmt/internal/file-uploader/service"
	cacheService "mitra-kirim-be-mgmt/internal/gateways/cache/service"
)

type Configuration struct {
	Repo      *repository.Configuration
	FileSvc   *service.FileUploader
	Logger    *logrus.Logger
	Cache     *cacheService.Cache
	CacheTime int
}

func NewConfiguration(repo *repository.Configuration, fileService *service.FileUploader, logger *logrus.Logger, cache *cacheService.Cache, cacheTime int) *Configuration {
	return &Configuration{repo, fileService, logger, cache, cacheTime}
}
