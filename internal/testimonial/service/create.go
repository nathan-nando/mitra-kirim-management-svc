package service

import (
	"mime/multipart"
	"mitra-kirim-be-mgmt/internal/testimonial/model"
)

func (s *Testimonial) Create(file *multipart.FileHeader, data *model.TestimonialCreate, userID string) (model.TestimonialResponse, error) {
	var response model.TestimonialResponse

	newFileName, err := s.FileUploader.UploadFile(file, "/mk-storage/testimonials/")
	if err != nil {
		return response, err
	}

	data.Img = newFileName

	res, err := s.Repository.Create(data, userID)
	if err != nil {
		return response, err
	}

	response.Filename = newFileName
	response.Id = res.ID

	return response, nil
}
