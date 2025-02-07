package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"mitra-kirim-be-mgmt/internal/user/model"
	"time"
)

type User struct {
	Db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{Db: db}
}

func (r *User) FindByUsername(username string) (model.User, error) {
	var user model.User
	if err := r.Db.Where("username", username).Take(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *User) CreateUser(req *model.RegisterRequest) (model.User, error) {

	user := model.User{
		ID:          uuid.NewString(),
		Name:        req.Name,
		Username:    req.Username,
		Title:       req.Title,
		Email:       req.Email,
		Password:    req.Password,
		Phone:       req.Phone,
		Gender:      req.Gender,
		Img:         req.Img,
		Status:      1,
		CreatedDate: time.Now(),
		CreatedBy:   "SYS",
	}

	if err := r.Db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
