package service

import (
	"mitra-kirim-be-mgmt/internal/user/model"
)

func (s *User) GetInformation(userID string) (model.UserInformation, error) {
	userData, err := s.Repository.FindByUserID(userID)
	if err != nil {
		return model.UserInformation{}, err
	}

	user := model.UserInformation{
		Name:     userData.Name,
		Username: userData.Username,
		Title:    userData.Title,
		Email:    userData.Email,
		Phone:    userData.Phone,
		Img:      userData.Img,
		Status:   userData.Status,
	}

	return user, nil
}
