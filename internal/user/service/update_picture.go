package service

import (
	"mime/multipart"
)

func (s *User) UpdatePicture(img *multipart.FileHeader, username string) (bool, error) {
	newFileName, err := s.FileUpSvc.UploadFile(img, "/mk-storage/users/")
	if err != nil {
		return false, err
	}

	_, err = s.Repository.UpdateUserImg(newFileName, username)
	if err != nil {
		return false, err
	}

	return true, nil
}
