package service

import (
	"motorcycle-rent-api/app/constant"
	"motorcycle-rent-api/app/helper"
	"motorcycle-rent-api/app/model"
	"motorcycle-rent-api/app/repository"

	"gorm.io/gorm"
)

type HealthServiceInterface interface {
	Health(apiCallID string) (*model.Config, constant.ResponseMap)
}
type HealthService struct {
	DB               *gorm.DB
	ConfigRepository repository.ConfigRepositoryInterface
}

func NewHealthService(db *gorm.DB, configRepository repository.ConfigRepositoryInterface) HealthServiceInterface {
	return &HealthService{
		DB:               db,
		ConfigRepository: configRepository,
	}
}

func (h *HealthService) Health(apiCallID string) (*model.Config, constant.ResponseMap) {
	data, err := h.ConfigRepository.GetAppStatus(h.DB)
	if err != nil {
		helper.LogError(apiCallID, "Error Health Check : "+err.Error())
		return nil, constant.Res400Failed
	}

	return data, constant.Res200Success
}
