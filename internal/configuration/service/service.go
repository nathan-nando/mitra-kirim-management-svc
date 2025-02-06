package service

import (
	"github.com/sirupsen/logrus"
	"mitra-kirim-be-mgmt/internal/configuration/repository"
	"mitra-kirim-be-mgmt/internal/file-uploader/service"
)

type Configuration struct {
	Repo    *repository.Configuration
	FileSvc *service.FileUploader
	Logger  *logrus.Logger
}

func NewConfiguration(repo *repository.Configuration, fileService *service.FileUploader, logger *logrus.Logger) *Configuration {
	return &Configuration{repo, fileService, logger}
}
