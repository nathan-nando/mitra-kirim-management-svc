package model

import "time"

type Testimonial struct {
	ID          int       `gorm:"column:id;primaryKey" json:"id"`
	Name        string    `gorm:"column:name" json:"nama"`
	Img         string    `gorm:"column:img" json:"img"`
	Description string    `gorm:"column:description" json:"deskripsi"`
	IsCarousel  int       `gorm:"column:is_carousel" json:"slide"`
	CreatedDate time.Time `gorm:"column:created_date" json:"createdDate"`
	CreatedBy   string    `gorm:"column:created_by" json:"createdBy"`
	UpdatedBy   string    `gorm:"column:updated_by" json:"updatedBy"`
	UpdatedDate time.Time `gorm:"column:updated_date" json:"updatedDate"`
}

func (a *Testimonial) TableName() string {
	return "testimonial"
}
