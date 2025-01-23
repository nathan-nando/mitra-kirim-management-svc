package repository

import (
	"gorm.io/gorm"
	"mitra-kirim-be-mgmt/internal/user/model"
)

type User struct {
	Db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{Db: db}
}

func (r *User) FindById(id int) (model.User, error) {
	var user model.User
	if err := r.Db.Where(id).Take(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
