package repository

import (
	"gorm.io/gorm"
	"mitra-kirim-be-mgmt/internal/suggestion/model"
)

type Suggestion struct {
	Db *gorm.DB
}

func NewSuggestion(db *gorm.DB) *Suggestion {
	return &Suggestion{db}
}

func (r *Suggestion) FindAll() ([]model.Suggestion, error) {
	var result []model.Suggestion
	if err := r.Db.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
