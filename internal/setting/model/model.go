package model

import "time"

type Configuration struct {
	ID          string    `gorm:"column:id;primaryKey"`
	key         string    `gorm:"column:key"`
	Type        string    `gorm:"column:type"`
	Value       string    `gorm:"column:value"`
	Description string    `gorm:"column:description"`
	Sort        int       `gorm:"column:sort"`
	CreatedDate time.Time `gorm:"column:createdDate"`
	CreatedBy   string    `gorm:"column:createdBy"`
	UpdatedBy   string    `gorm:"column:updatedBy"`
	UpdatedDate time.Time `gorm:"column:updatedDate"`
}

func (a *Configuration) TableName() string {
	return "configuration"
}
