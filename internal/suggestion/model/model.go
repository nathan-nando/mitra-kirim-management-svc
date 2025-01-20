package model

import "time"

type Suggestion struct {
	ID          int        `gorm:"column:id;primaryKey" json:"id"`
	Name        string     `gorm:"column:name" json:"name" json:"name"`
	Email       string     `gorm:"column:email" json:"email"`
	Message     string     `gorm:"column:message" json:"message"`
	HasReplied  int8       `gorm:"column:has_replied" json:"hasReplied"`
	CreatedDate time.Time  `gorm:"column:created_date" json:"createdDate"`
	CreatedBy   string     `gorm:"column:created_by" json:"createdBy"`
	UpdatedDate *time.Time `gorm:"column:updated_date" json:"updatedDate"`
	UpdatedBy   string     `gorm:"column:updated_by" json:"updatedBy"`
}

func (a *Suggestion) TableName() string {
	return "suggestion"
}
