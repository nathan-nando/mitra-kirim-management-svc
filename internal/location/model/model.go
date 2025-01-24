package model

import "time"

type Location struct {
	ID          int        `gorm:"column:id;primaryKey"`
	Name        string     `gorm:"column:name"`
	Latitude    float64    `gorm:"column:latitude"`
	Longitude   float64    `gorm:"column:longitude"`
	Sort        int        `gorm:"column:sort"`
	CreatedDate time.Time  `gorm:"column:created_date" json:"createdDate"`
	CreatedBy   string     `gorm:"column:created_by" json:"createdBy"`
	UpdatedDate *time.Time `gorm:"column:updated_date" json:"updatedDate"`
	UpdatedBy   string     `gorm:"column:updated_by" json:"updatedBy"`
}

func (a *Location) TableName() string {
	return "location"
}
