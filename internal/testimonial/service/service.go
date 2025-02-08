package service

import (
	"mitra-kirim-be-mgmt/internal/file-uploader/service"
	"mitra-kirim-be-mgmt/internal/testimonial/repository"
)

type Testimonial struct {
	Repository   *repository.Testimonial
	FileUploader *service.FileUploader
}

func NewTestimonial(repo *repository.Testimonial, fileUploader *service.FileUploader) *Testimonial {
	return &Testimonial{
		Repository:   repo,
		FileUploader: fileUploader,
	}
}
