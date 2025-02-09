package service

import (
	"gorm.io/gorm"
	fileUploaderService "mitra-kirim-be-mgmt/internal/file-uploader/service"
	"mitra-kirim-be-mgmt/internal/user/repository"
)

type User struct {
	db              *gorm.DB
	Repository      *repository.User
	FileUpSvc       *fileUploaderService.FileUploader
	JwtKey          string
	TokenExp        int
	RefreshTokenExp int
}

func NewUser(db *gorm.DB, fileSvc *fileUploaderService.FileUploader, repo *repository.User, jwtKey string, tokenExp int, refreshTokenExp int) *User {
	return &User{
		db:              db,
		FileUpSvc:       fileSvc,
		Repository:      repo,
		JwtKey:          jwtKey,
		TokenExp:        tokenExp,
		RefreshTokenExp: refreshTokenExp,
	}
}
