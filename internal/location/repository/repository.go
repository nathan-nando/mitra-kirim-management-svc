package repository

import "gorm.io/gorm"

type Location struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Location {
	return &Location{Db: db}
}
