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

func NewRepository(db *gorm.DB) *Location {
	return &Location{Db: db}
}

func (r *Location) FindAll(limit int, offset int) ([]model.Location, error) {
	var result []model.Location

	if err := r.Db.Find(&result).Limit(limit).Offset(offset).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Location) Create(location *model.LocationRequest) (int, error) {
	loc := &model.Location{
		Name:        location.Name,
		Description: location.Description,
		IFrameLink:  location.IFrameLink,
		CreatedBy:   "SYS",
		CreatedDate: time.Now(),
	}
	if err := r.Db.Create(&loc).Error; err != nil {
		return 0, err
	}
	return loc.ID, nil
}

func (r *Location) Update(location *model.LocationRequest) (int, error) {
	loc := &model.Location{
		Name:        location.Name,
		Description: location.Description,
		IFrameLink:  location.IFrameLink,
		UpdatedBy:   "SYS",
		UpdatedDate: time.Now(),
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
