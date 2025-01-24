package model

import (
	"time"
)

type User struct {
	ID          int        `gorm:"column:id;primaryKey" json:"id"`
	Name        string     `gorm:"column:name"`
	Title       string     `gorm:"column:title"`
	Email       string     `gorm:"column:email"`
	Password    string     `gorm:"column:password"`
	Phone       string     `gorm:"column:phone"`
	Address     string     `gorm:"column:address"`
	Gender      string     `gorm:"column:gender"`
	Img         string     `gorm:"column:img"`
	Status      int        `gorm:"column:status"`
	CreatedDate time.Time  `gorm:"column:createdDate"`
	CreatedBy   string     `gorm:"column:createdBy"`
	UpdatedDate *time.Time `gorm:"column:updatedDate"`
	UpdatedBy   string     `gorm:"column:updatedBy"`
}

func (a *User) TableName() string {
	return "users"
}
