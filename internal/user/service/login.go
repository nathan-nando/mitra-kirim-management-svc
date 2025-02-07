package service

import (
	"golang.org/x/crypto/bcrypt"
	"mitra-kirim-be-mgmt/internal/user/model"
	jwttoken "mitra-kirim-be-mgmt/pkg/jwt-token"
	"time"
)

func (s *User) Login(req *model.LoginRequest) (model.LoginResponse, error) {
	var response model.LoginResponse

	user, err := s.Repository.FindByUsername(req.Username)
	if err != nil {
		return response, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return response, err
	}

	accessTokenDuration := time.Duration(s.TokenExp) * time.Hour
	accessToken, err := jwttoken.GenerateAccessToken(user.ID, user.Email, []byte(s.JwtKey), accessTokenDuration)
	if err != nil {
		return response, err
	}

	refreshTokenDuration := time.Duration(s.RefreshTokenExp) * 24 * time.Hour
	refreshToken, err := jwttoken.GenerateRefreshToken(user.ID, user.Email, []byte(s.JwtKey), refreshTokenDuration)
	if err != nil {
		return response, err
	}

	response.Token = accessToken
	response.RefreshToken = refreshToken
	response.ExpiredAt = time.Now().Add(accessTokenDuration)

	return response, nil
}
