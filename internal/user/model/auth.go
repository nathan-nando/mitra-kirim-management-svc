package model

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	UserID    string `json:"id"`
	Username  string `json:"username"`
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
	ExpiredAt    time.Time `json:"expiredAt"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refreshToken"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Title    string `json:"title"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Gender   string `json:"gender"`
	Img      string `json:"img"`
}

type RegisterResponse struct {
	Username string `json:"username"`
}
