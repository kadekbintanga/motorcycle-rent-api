package repository

import (
	"motorcycle-rent-api/app/constant"
	"motorcycle-rent-api/app/model"

	"gorm.io/gorm"
)

type ConfigRepositoryInterface interface {
	GetAppStatus(db *gorm.DB) (*model.Config, error)
}

type ConfigRepository struct {
}

func NewConfigRepository() ConfigRepositoryInterface {
	return &ConfigRepository{}
}

func (*ConfigRepository) GetAppStatus(db *gorm.DB) (*model.Config, error) {
	var config model.Config

	err := db.Where("key = ?", constant.AppStatusKey).First(&config).Error
	if err != nil {
		return nil, err
	}

	return &config, nil
}
