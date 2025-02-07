package service

import (
	"golang.org/x/crypto/bcrypt"
	"mitra-kirim-be-mgmt/internal/user/model"
)

func (s *User) Register(request *model.RegisterRequest) (model.RegisterResponse, error) {
	var response model.RegisterResponse
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return response, err
	}
	request.Password = string(hashedPassword)
	user, err := s.Repository.CreateUser(request)
	if err != nil {
		return response, err
	}
	response.Username = user.Username

	return response, nil
}
