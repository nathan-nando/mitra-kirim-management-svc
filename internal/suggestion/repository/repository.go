package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"mitra-kirim-be-mgmt/internal/suggestion/model"
	"time"
)

type Suggestion struct {
	Db *gorm.DB
}

func NewSuggestion(db *gorm.DB) *Suggestion {
	return &Suggestion{Db: db}
}

func (r *Suggestion) FindAll() ([]model.Suggestion, error) {
	var result []model.Suggestion
	if err := r.Db.Order("created_date desc").Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Suggestion) CountAll() (int64, error) {
	var result int64
	if err := r.Db.Model(&model.Suggestion{}).Count(&result).Error; err != nil {
		return 0, err
	}
	return result, nil
}

func (r *Suggestion) FindByID(id string) (model.Suggestion, error) {
	var result model.Suggestion
	if err := r.Db.First(&result).Where("ID = ?", id).Error; err != nil {
		return model.Suggestion{}, err
	}
	return result, nil
}

func (r *Suggestion) Create(suggestion *model.SuggestionCreate) (model.Suggestion, error) {
	var result model.Suggestion
	if err := r.Db.Create(
		&model.Suggestion{
			ID:          uuid.New().String(),
			Name:        suggestion.Name,
			Email:       suggestion.Email,
			Message:     suggestion.Message,
			HasReplied:  0,
			CreatedDate: time.Now(),
			CreatedBy:   "SYSTEM",
		},
	).Error; err != nil {
		return model.Suggestion{}, err
	}
	return result, nil
}
