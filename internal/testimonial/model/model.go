package model

import "time"

type Testimonial struct {
	ID          string    `gorm:"column:id;primaryKey"`
	Name        string    `gorm:"column:name"`
	Img         string    `gorm:"column:img"`
	Description string    `gorm:"column:description"`
	Sort        int       `gorm:"column:sort"`
	CreatedDate time.Time `gorm:"column:createdDate"`
	CreatedBy   string    `gorm:"column:createdBy"`
	UpdatedBy   string    `gorm:"column:updatedBy"`
	UpdatedDate time.Time `gorm:"column:updatedDate"`
}

func (a *Testimonial) TableName() string {
	return "testimonial"
}
