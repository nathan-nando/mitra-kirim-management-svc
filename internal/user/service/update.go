package service

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"mitra-kirim-be-mgmt/internal/user/model"
)

func (s *User) Update(req *model.UserUpdate, userID string) (string, error) {

	user, err := s.Repository.FindByUserID(userID)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	req.Password = string(hashedPassword)

	if user.ID != userID {
		return "", errors.New("User not matched")
	}

	_, err = s.Repository.UpdateUser(req, userID, user.Username)
	if err != nil {
		return "", err
	}

	return userID, nil
}
