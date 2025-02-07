package service

import (
	"gorm.io/gorm"
	"mitra-kirim-be-mgmt/internal/user/repository"
)

type User struct {
	db              *gorm.DB
	Repository      *repository.User
	JwtKey          string
	TokenExp        int
	RefreshTokenExp int
}

func NewUser(db *gorm.DB, repo *repository.User, jwtKey string, tokenExp int, refreshTokenExp int) *User {
	return &User{
		db:              db,
		Repository:      repo,
		JwtKey:          jwtKey,
		TokenExp:        tokenExp,
		RefreshTokenExp: refreshTokenExp,
	}
}
