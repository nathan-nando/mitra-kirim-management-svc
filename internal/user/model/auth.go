package model

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	UserID    string `json:"id"`
	Email     string `json:"email"`
	TokenType string `json:"tokenType"`
	jwt.RegisteredClaims
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type LoginResponse struct {
	ExpiredAt    time.Duration `json:"expiredAt"`
	Token        string        `json:"token"`
	RefreshToken string        `json:"refreshToken"`
}
