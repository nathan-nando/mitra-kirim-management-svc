package model

import "time"

type Configuration struct {
	ID          int       `gorm:"column:id;primaryKey" json:"id"`
	Key         string    `gorm:"column:key" json:"key"`
	Type        string    `gorm:"column:type" json:"type"`
	Value       string    `gorm:"column:value" json:"value"`
	FormType    string    `gorm:"column:form_type" json:"formType"`
	CreatedDate time.Time `gorm:"column:created_date" json:"createdDate"`
	CreatedBy   string    `gorm:"column:created_by" json:"createdBy"`
	UpdatedDate time.Time `gorm:"column:updated_date" json:"updatedDate"`
	UpdatedBy   string    `gorm:"column:updated_by" json:"updatedBy"`
}

func (a *Configuration) TableName() string {
	return "configuration"
}
