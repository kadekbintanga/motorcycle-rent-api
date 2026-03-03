package repository

import (
	"motorcycle-rent-api/app/model"

	"gorm.io/gorm"
)

type MotorcycleRepositoryInterface interface {
	CreateMotorcycle(db *gorm.DB, motorcycle model.Motorcycle) (*model.Motorcycle, error)
	GetMotorcycleByPlateNumber(db *gorm.DB, plateNumber string) (*model.Motorcycle, error)
}

type MotorcycleRepository struct{}

func NewMotorcycleRepository() MotorcycleRepositoryInterface {
	return &MotorcycleRepository{}
}

func (m *MotorcycleRepository) CreateMotorcycle(db *gorm.DB, motorcycle model.Motorcycle) (*model.Motorcycle, error) {
	if err := db.Create(&motorcycle).Error; err != nil {
		return nil, err
	}
	return &motorcycle, nil
}

func (m *MotorcycleRepository) GetMotorcycleByPlateNumber(db *gorm.DB, plateNumber string) (*model.Motorcycle, error) {
	var motorcycle model.Motorcycle
	if err := db.Where("plate_number = ?", plateNumber).First(&motorcycle).Error; err != nil {
		return nil, err
	}
	return &motorcycle, nil
}
