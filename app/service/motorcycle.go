package service

import (
	"errors"
	"motorcycle-rent-api/app/constant"
	"motorcycle-rent-api/app/helper"
	"motorcycle-rent-api/app/model"
	"motorcycle-rent-api/app/repository"
	"motorcycle-rent-api/app/resource/request"
	"motorcycle-rent-api/app/resource/response"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MotorcycleServiceInterface interface {
	CreateMotorcycle(apiCallID string, payload request.CreateMotorcycleRequest) constant.ResponseMap
	GetListMotorcyclesPagination(apiCallID string, param helper.PaginationParam, filter request.GetMotorcycleListFilter) ([]response.MotorcycleListpaginationResponse, *helper.ResponseMeta, constant.ResponseMap)
	UpdateMotorcycleDetail(apiCallID string, motorcycleUUID string, payload request.UpdateMotorcycleRequest) constant.ResponseMap
	UpdateMotorcycleStatus(apiCallID string, motorcycleUUID string, payload request.UpdateMotorcycleStatusRequest) constant.ResponseMap
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
			PlateNumber: strings.ToUpper(payload.PlateNumber),
			Brand:       payload.Brand,
			Type:        constant.MotorcycleType(payload.Type),
			Year:        payload.Year,
			Status:      constant.MotorcycleStatus(payload.Status),
			PricePerDay: payload.PricePerDay,
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
	return constant.Res200Save
}

func (m *MotorcycleService) GetListMotorcyclesPagination(apiCallID string, param helper.PaginationParam, filter request.GetMotorcycleListFilter) ([]response.MotorcycleListpaginationResponse, *helper.ResponseMeta, constant.ResponseMap) {
	motorcycles, meta, err := m.MotorcycleRepository.GetListMotorcyclesPagination(m.DB, param, filter)
	if err != nil {
		helper.LogError(apiCallID, "Error getting motorcycle list: "+err.Error())
		return nil, nil, constant.Res422SomethingWentWrong
	}

	formattedMotorcycleList := response.MotorcycleListpaginationResponseFormatter(motorcycles)
	return formattedMotorcycleList, &meta, constant.Res200Get
}

func (m *MotorcycleService) UpdateMotorcycleDetail(apiCallID string, motorcycleUUID string, payload request.UpdateMotorcycleRequest) constant.ResponseMap {
	err := m.DB.Clauses(clause.Locking{Strength: "UPDATE"}).Transaction(func(tx *gorm.DB) error {
		motorcycle, err := m.MotorcycleRepository.GetMotorcycleByUUID(tx, motorcycleUUID, false)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				helper.LogError(apiCallID, "Motorcycle not found: "+motorcycleUUID)
				return errors.New("motorcycle not found")
			}
			helper.LogError(apiCallID, "Error getting motorcycle by UUID: "+err.Error())
			return errors.New("error getting motorcycle by uuid")
		}

		if motorcycle.PlateNumber != payload.PlateNumber {
			checkPlateNumber, err := m.MotorcycleRepository.GetMotorcycleByPlateNumber(tx, payload.PlateNumber)
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				helper.LogError(apiCallID, "Error checking plate number: "+err.Error())
				return errors.New("error get motorcycle by plate number")
			}

			if checkPlateNumber != nil {
				helper.LogError(apiCallID, "Plate number already exists: "+payload.PlateNumber)
				return errors.New("plate number already exists")
			}
		}

		updateMotorcycle := map[string]interface{}{
			"plate_number":  strings.ToUpper(payload.PlateNumber),
			"brand":         payload.Brand,
			"type":          constant.MotorcycleType(payload.Type),
			"year":          payload.Year,
			"status":        constant.MotorcycleStatus(payload.Status),
			"price_per_day": payload.PricePerDay,
		}

		err = m.MotorcycleRepository.UpdateMotorcycleMap(tx, *motorcycle, updateMotorcycle)
		return nil
	})

	if err != nil {
		switch err.Error() {
		case "motorcycle not found":
			return constant.Res404MotorcycleNotFound
		case "plate number already exists":
			return constant.Res400PlateNumberExists
		default:
			return constant.Res422SomethingWentWrong
		}
	}
	return constant.Res200Update
}

func (m *MotorcycleService) UpdateMotorcycleStatus(apiCallID string, motorcycleUUID string, payload request.UpdateMotorcycleStatusRequest) constant.ResponseMap {
	err := m.DB.Clauses(clause.Locking{Strength: "UPDATE"}).Transaction(func(tx *gorm.DB) error {
		motorcycle, err := m.MotorcycleRepository.GetMotorcycleByUUID(tx, motorcycleUUID, false)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				helper.LogError(apiCallID, "Motorcycle not found: "+motorcycleUUID)
				return errors.New("motorcycle not found")
			}
			helper.LogError(apiCallID, "Error getting motorcycle by UUID: "+err.Error())
			return errors.New("error getting motorcycle by uuid")
		}

		updateMotorcycle := map[string]interface{}{
			"status": constant.MotorcycleStatus(payload.Status),
		}

		err = m.MotorcycleRepository.UpdateMotorcycleMap(tx, *motorcycle, updateMotorcycle)
		return nil
	})

	if err != nil {
		switch err.Error() {
		case "motorcycle not found":
			return constant.Res404MotorcycleNotFound
		default:
			return constant.Res422SomethingWentWrong
		}
	}
	return constant.Res200Update
}
