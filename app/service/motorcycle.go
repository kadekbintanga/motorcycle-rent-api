package service

import (
	"errors"
	"motorcycle-rent-api/app/constant"
	"motorcycle-rent-api/app/helper"
	"motorcycle-rent-api/app/model"
	"motorcycle-rent-api/app/repository"
	"motorcycle-rent-api/app/resource/request"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MotorcycleServiceInterface interface {
	CreateMotorcycle(apiCallID string, payload request.CreateMotorcycleRequest) constant.ResponseMap
}

type MotorcycleService struct {
	DB                   *gorm.DB
	MotorcycleRepository repository.MotorcycleRepositoryInterface
}

func NewMotorcycleService(db *gorm.DB, motorcycleRepository repository.MotorcycleRepositoryInterface) MotorcycleServiceInterface {
	return &MotorcycleService{
		DB:                   db,
		MotorcycleRepository: motorcycleRepository,
	}
}

func (m *MotorcycleService) CreateMotorcycle(apiCallID string, payload request.CreateMotorcycleRequest) constant.ResponseMap {
	err := m.DB.Clauses(clause.Locking{Strength: "UPDATE"}).Transaction(func(tx *gorm.DB) error {
		checkPlateNumber, err := m.MotorcycleRepository.GetMotorcycleByPlateNumber(tx, payload.PlateNumber)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			helper.LogError(apiCallID, "Error checking plate number: "+err.Error())
			return errors.New("error get motorcycle by plate number")
		}

		if checkPlateNumber != nil {
			helper.LogError(apiCallID, "Plate number already exists: "+payload.PlateNumber)
			return errors.New("plate number already exists")
		}

		createMotorcycle := model.Motorcycle{
			PlateNumber: payload.PlateNumber,
			Brand:       payload.Brand,
			Type:        constant.MotorcycleType(payload.Type),
			Year:        payload.Year,
			Status:      constant.MotorcycleStatus(payload.Status),
		}

		_, err = m.MotorcycleRepository.CreateMotorcycle(tx, createMotorcycle)
		if err != nil {
			helper.LogError(apiCallID, "Error creating motorcycle: "+err.Error())
			return errors.New("error creating motorcycle")
		}

		return nil
	})
	if err != nil {
		switch err.Error() {
		case "plate number already exists":
			return constant.Res400PlateNumberExists
		default:
			return constant.Res422SomethingWentWrong
		}
	}
	return constant.Res200Success
}
