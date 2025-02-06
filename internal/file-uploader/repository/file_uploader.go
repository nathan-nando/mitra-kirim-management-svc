package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

type FileUploader struct {
	UploadDir string
	logger    *logrus.Logger
}

func NewFileUploader(uploadDir string, logger *logrus.Logger) *FileUploader {
	return &FileUploader{UploadDir: uploadDir, logger: logger}
}

func (r *FileUploader) SaveFile(fileHeader *multipart.FileHeader, src multipart.File, dir string) (string, string, error) {

	targetDir := r.UploadDir + dir
	newUUID := uuid.New().String()

	ext := filepath.Ext(fileHeader.Filename)
	name := strings.TrimSuffix(fileHeader.Filename, ext)

	newFileName := fmt.Sprintf("%s-%s%s", name, newUUID, ext)

	dstPath := filepath.Join(targetDir, newFileName)

	if err := os.MkdirAll(targetDir, os.ModePerm); err != nil {
		return "", "", fmt.Errorf("failed to create storage directory: %w", err)
	}

	dst, err := os.Create(dstPath)
	if err != nil {
		return "", "", fmt.Errorf("error creating destination file: %w", err)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return "", "", fmt.Errorf("error saving file: %w", err)
	}

	return newFileName, dstPath, nil
}
