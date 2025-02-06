package repository

import (
	"fmt"
	"gorm.io/gorm"
	"mitra-kirim-be-mgmt/internal/configuration/model"
	"mitra-kirim-be-mgmt/pkg/converter"
	"time"
)

type Configuration struct {
	Db *gorm.DB
}

func NewConfiguration(db *gorm.DB) *Configuration {
	return &Configuration{Db: db}
}

func (r *Configuration) FindByTypes(types []string) ([]model.Configuration, error) {
	var result []model.Configuration
	if err := r.Db.Where("type IN ?", types).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Configuration) FindByID(id string) (model.Configuration, error) {
	var result model.Configuration
	if err := r.Db.First(&result).Where("ID = ?", id).Error; err != nil {
		return model.Configuration{}, err
	}
	return result, nil
}

func (r *Configuration) Create(cfg *model.ConfigurationCreate) (model.Configuration, error) {
	var result model.Configuration
	if err := r.Db.Create(
		&model.Configuration{
			Key:       cfg.Key,
			Type:      cfg.Type,
			Value:     cfg.Value,
			CreatedBy: "SYS",
		},
	).Error; err != nil {
		return model.Configuration{}, err
	}
	return result, nil
}

func (r *Configuration) UpdateByKey(keyVal []converter.KeyValue) error {

	for _, v := range keyVal {

		result := r.Db.Model(&model.Configuration{}).
			Where("key = ?", v.Key).
			Updates(&model.Configuration{
				Value:       converter.ConvertToString(v.Value),
				UpdatedBy:   "SYS",
				UpdatedDate: time.Now(),
			})
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			fmt.Println("no rows affected")
			return nil
		}
	}

	return nil
}
