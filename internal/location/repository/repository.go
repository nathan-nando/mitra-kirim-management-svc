package repository

import (
	"fmt"
	"gorm.io/gorm"
	"mitra-kirim-be-mgmt/internal/location/model"
	"time"
)

type Location struct {
	Db *gorm.DB
}

func NewLocation(db *gorm.DB) *Location {
	return &Location{Db: db}
}

func (r *Location) FindAll(limit int, offset int) ([]model.Location, error) {
	var result []model.Location

	if err := r.Db.Order("created_date desc").Find(&result).Limit(limit).Offset(offset).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Location) CountAll() (int64, error) {
	var count int64

	err := r.Db.Model(&model.Location{}).Count(&count).Error

	return count, err
}

func (r *Location) Create(location *model.LocationRequest, username string) (int, error) {
	loc := &model.Location{
		Name:        location.Name,
		Description: location.Description,
		IFrameLink:  location.IFrameLink,
		Email:       location.Email,
		Address:     location.Address,
		Phone:       location.Whatsapp,
		CreatedBy:   username,
		CreatedDate: time.Now(),
	}
	if err := r.Db.Create(&loc).Error; err != nil {
		return 0, err
	}
	return loc.ID, nil
}

func (r *Location) Update(location *model.LocationRequest, username string) (int, error) {
	now := time.Now()
	loc := &model.Location{
		Name:        location.Name,
		Description: location.Description,
		IFrameLink:  location.IFrameLink,
		Email:       location.Email,
		Address:     location.Address,
		Phone:       location.Whatsapp,
		UpdatedBy:   username,
		UpdatedDate: &now,
	}
	if err := r.Db.Model(&loc).Where(location.Id).Updates(loc).Error; err != nil {
		fmt.Println("ERR", err)
		return 0, err
	}
	return loc.ID, nil
}

func (r *Location) Delete(id int) (int, error) {
	if err := r.Db.Where(id).Delete(&model.Location{}).Error; err != nil {
		return 0, err
	}
	return id, nil
}
