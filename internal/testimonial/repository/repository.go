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

func (a *Testimonial) Create(req *model.TestimonialCreate, userID string) (model.Testimonial, error) {
	testimonial := model.Testimonial{
		Name:        req.Name,
		Img:         req.Img,
		Description: req.Description,
		IsCarousel:  0,
		CreatedDate: time.Now(),
		CreatedBy:   userID,
	}
	err := a.Db.Create(&testimonial).Error
	return testimonial, err
}

func (a *Testimonial) Delete(id string) (string, error) {
	err := a.Db.Where(id).Delete(&model.Testimonial{})
	if err != nil {
		return "", nil
	}
	return id, nil
}
