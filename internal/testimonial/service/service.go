package service

import (
	"github.com/sirupsen/logrus"
	"mitra-kirim-be-mgmt/internal/file-uploader/service"
	cacheSvc "mitra-kirim-be-mgmt/internal/gateways/cache/service"
	"mitra-kirim-be-mgmt/internal/testimonial/repository"
)

type Testimonial struct {
	Repository   *repository.Testimonial
	FileUploader *service.FileUploader
	Logger       *logrus.Logger
	CacheSvc     *cacheSvc.Cache
	CacheTime    int
}

func NewTestimonial(repo *repository.Testimonial, fileUploader *service.FileUploader, logger *logrus.Logger, cache *cacheSvc.Cache, cacheTime int) *Testimonial {
	return &Testimonial{
		Repository:   repo,
		FileUploader: fileUploader,
		Logger:       logger,
		CacheSvc:     cache,
		CacheTime:    cacheTime,
	}
}
