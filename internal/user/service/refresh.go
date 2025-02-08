package service

import (
	"github.com/golang-jwt/jwt/v5"
	"mitra-kirim-be-mgmt/internal/user/model"
	jwt_token "mitra-kirim-be-mgmt/pkg/jwt-token"
	"time"
)

func (s *User) Refresh(req *model.RefreshTokenRequest) (string, error) {
	claims := &model.Claims{}
	token, err := jwt.ParseWithClaims(req.RefreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.JwtKey), nil
	})

	if !token.Valid {
		return "", err
	}

	if err != nil || !token.Valid || claims.TokenType != "refresh" {
		return "", err
	}

	accessTokenDuration := time.Duration(s.TokenExp) * time.Hour
	accessToken, err := jwt_token.GenerateAccessToken(claims.UserID, claims.Username, []byte(s.JwtKey), accessTokenDuration)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
