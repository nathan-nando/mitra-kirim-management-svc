package repository

import (
	"gorm.io/gorm"
	"mitra-kirim-be-mgmt/internal/testimonial/model"
	"time"
)

type Testimonial struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Testimonial {
	return &Testimonial{Db: db}
}

func (a *Testimonial) GetAll(offset int, limit int) ([]model.Testimonial, error) {
	var testimonials []model.Testimonial
	err := a.Db.Find(&testimonials).Offset(offset).Limit(limit).Error
	return testimonials, err
}

func (a *Testimonial) GetSlide(offset int, limit int) ([]model.Testimonial, error) {
	var testimonials []model.Testimonial
	err := a.Db.Where("is_carousel", 1).Find(&testimonials).Offset(offset).Limit(limit).Error
	return testimonials, err
}

func (a *Testimonial) Create(req *model.TestimonialCreate, username string) (model.Testimonial, error) {
	testimonial := model.Testimonial{
		Name:        req.Name,
		Img:         req.Img,
		Description: req.Description,
		IsCarousel:  0,
		CreatedBy:   username,
	}
	err := a.Db.Create(&testimonial).Error
	return testimonial, err
}

func (a *Testimonial) UpdateSlide(req *model.TestimonialUpdateSlide, userID string) (bool, error) {
	now := time.Now()
	testimonial := model.Testimonial{
		IsCarousel:  req.Slide,
		UpdatedDate: &now,
		UpdatedBy:   userID,
	}
	err := a.Db.Model(&model.Testimonial{}).
		Select("is_carousel").
		Where("id", req.Id).Updates(&testimonial).Error
	if err != nil {
		return false, err
	}
	return true, err
}

func (a *Testimonial) Delete(id int) (bool, error) {
	err := a.Db.Where(id).Delete(&model.Testimonial{})
	if err != nil {
		return false, nil
	}
	return true, nil
}
