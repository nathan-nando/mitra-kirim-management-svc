package model

import "time"

type Location struct {
	ID          int       `gorm:"column:id;primaryKey" json:"id"`
	Name        string    `gorm:"column:name" json:"nama"`
	Description string    `gorm:"column:description" json:"deskripsi"`
	Email       string    `gorm:"column:email" json:"email"`
	Phone       string    `gorm:"column:phone" json:"whatsapp"`
	IFrameLink  string    `gorm:"column:iframe_link" json:"iframeLink"`
	CreatedDate time.Time `gorm:"column:created_date" json:"createdDate"`
	CreatedBy   string    `gorm:"column:created_by" json:"createdBy"`
	UpdatedDate time.Time `gorm:"column:updated_date" json:"updatedDate"`
	UpdatedBy   string    `gorm:"column:updated_by" json:"updatedBy"`
}

func (a *Location) TableName() string {
	return "location"
}
