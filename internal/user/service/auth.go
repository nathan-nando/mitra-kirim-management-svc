package service

import (
	"github.com/golang-jwt/jwt/v5"
	"mitra-kirim-be-mgmt/internal/user/model"
	"time"
)

var (
	JwtKey          = []byte("n4thsecret") // Replace with a secure secret key
	accessTokenExp  = 1 * time.Minute      // Access token expires in 15 minutes
	refreshTokenExp = 3 * 24 * time.Hour   // Refresh token expires in 7 days
)

func (s *User) GenerateAccessToken(userID, email string) (string, error) {
	return s.generateToken(userID, email, "access", accessTokenExp)
}

func (s *User) GenerateRefreshToken(userID, email string) (string, error) {
	return s.generateToken(userID, email, "refresh", refreshTokenExp)
}

func (s *User) generateToken(userID, email, tokenType string, expiration time.Duration) (string, error) {
	claims := &model.Claims{
		UserID:    userID,
		Email:     email,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "your-app-name",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}
