package service

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"mime/multipart"
	"mitra-kirim-be-mgmt/internal/file-uploader/repository"
)

type FileUploader struct {
	Repository *repository.FileUploader
	Logger     *logrus.Logger
}

func NewFileUploader(fileUploader *repository.FileUploader, logger *logrus.Logger) *FileUploader {
	return &FileUploader{fileUploader, logger}
}

func (s *FileUploader) UploadFile(fileHeader *multipart.FileHeader, dir string) (string, error) {
	src, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("error opening file: %w", err)
	}
	defer src.Close()

	fileName, filePath, err := s.Repository.SaveFile(fileHeader, src, dir)
	s.Logger.Infof("File saved to %s", filePath)

	if err != nil {
		return "", fmt.Errorf("error saving file via repository: %w", err)
	}

	return fileName, nil
}
