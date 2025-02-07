package model

import (
	"time"
)

type User struct {
	ID          string     `gorm:"column:id;primaryKey" json:"id"`
	Name        string     `gorm:"column:name"`
	Username    string     `gorm:"column:username"`
	Title       string     `gorm:"column:title"`
	Email       string     `gorm:"column:email"`
	Password    string     `gorm:"column:password"`
	Phone       string     `gorm:"column:phone"`
	Gender      string     `gorm:"column:gender"`
	Img         string     `gorm:"column:img"`
	Status      int        `gorm:"column:status"`
	CreatedDate time.Time  `gorm:"column:created_date"`
	CreatedBy   string     `gorm:"column:created_by"`
	UpdatedDate *time.Time `gorm:"column:updated_date"`
	UpdatedBy   string     `gorm:"column:updated_by"`
}

func (a *User) TableName() string {
	return "users"
}
