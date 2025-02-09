package repository

import (
	"fmt"
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

func (r *User) FindByUserID(id string) (model.User, error) {
	var user model.User
	if err := r.Db.Where("id", id).Take(&user).Error; err != nil {
		return user, err
	}
	return user, nil
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

func (r *User) UpdateUser(update *model.UserUpdate, userID string, username string) (string, error) {
	now := time.Now()
	user := &model.User{
		Name:        update.Name,
		Username:    update.Username,
		Password:    update.Password,
		Title:       update.Title,
		Email:       update.Email,
		Phone:       update.Phone,
		UpdatedBy:   username,
		UpdatedDate: &now,
	}
	if err := r.Db.Model(&user).Where("id", userID).Updates(user).Error; err != nil {
		fmt.Println("ERR", err)
		return "", err
	}
	return user.ID, nil
}

func (r *User) UpdateUserImg(newImgName, username string) (string, error) {
	now := time.Now()
	user := &model.User{
		Img:         newImgName,
		UpdatedBy:   username,
		UpdatedDate: &now,
	}
	if err := r.Db.Model(&user).Where("username", username).Updates(user).Error; err != nil {
		fmt.Println("ERR", err)
		return "", err
	}
	return user.ID, nil
}
