package repository

import (
	"motorcycle-rent-api/app/helper"
	"motorcycle-rent-api/app/model"
	"motorcycle-rent-api/app/resource/request"

	"gorm.io/gorm"
)

type MotorcycleRepositoryInterface interface {
	CreateMotorcycle(db *gorm.DB, motorcycle model.Motorcycle) (*model.Motorcycle, error)
	GetMotorcycleByPlateNumber(db *gorm.DB, plateNumber string) (*model.Motorcycle, error)
	GetListMotorcyclesPagination(db *gorm.DB, param helper.PaginationParam, filter request.GetMotorcycleListFilter) ([]model.Motorcycle, helper.ResponseMeta, error)
	GetMotorcycleByUUID(db *gorm.DB, uuid string, withPreload bool) (*model.Motorcycle, error)
	UpdateMotorcycleMap(db *gorm.DB, motorcycle model.Motorcycle, updateData map[string]interface{}) error
	UpdateMotocycleByUUID(db *gorm.DB, motocycleUUID string, updateData model.Motorcycle) error
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

func (m *MotorcycleRepository) GetListMotorcyclesPagination(db *gorm.DB, param helper.PaginationParam, filter request.GetMotorcycleListFilter) ([]model.Motorcycle, helper.ResponseMeta, error) {
	var motorcycles []model.Motorcycle

	query := db.Model(&model.Motorcycle{})

	if param.Search != "" {
		query = query.Where("plate_number ILIKE ? OR brand ILIKE ?", "%"+param.Search+"%", "%"+param.Search+"%")
	}

	if filter.Type != "" {
		query = query.Where("type = ?", filter.Type)
	}

	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}

	if filter.Year != 0 {
		query = query.Where("year = ?", filter.Year)
	}

	query, total, err := helper.Paginate(param.Limit, param.Page, query)
	if err != nil {
		return nil, helper.ResponseMeta{}, err
	}

	if err := query.Find(&motorcycles).Order("created_at desc").Error; err != nil {
		return nil, helper.ResponseMeta{}, err
	}

	meta := helper.ResponseMeta{
		Page:  param.Page,
		Limit: param.Limit,
		Total: int(total),
	}

	return motorcycles, meta, nil
}

func (m *MotorcycleRepository) GetMotorcycleByUUID(db *gorm.DB, motorcycleUUID string, withPreload bool) (*model.Motorcycle, error) {
	var motorcycle model.Motorcycle
	query := db.Where("uuid = ?", motorcycleUUID)

	if withPreload {
		query = query
	}

	err := query.First(&motorcycle).Error
	if err != nil {
		return nil, err
	}

	return &motorcycle, nil
}

func (m *MotorcycleRepository) UpdateMotorcycleMap(db *gorm.DB, motorcycle model.Motorcycle, updateData map[string]interface{}) error {
	err := db.Model(motorcycle).Updates(updateData).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *MotorcycleRepository) UpdateMotocycleByUUID(db *gorm.DB, motocycleUUID string, updateData model.Motorcycle) error {
	err := db.Model(&model.Motorcycle{}).Where("uuid = ?", motocycleUUID).Updates(updateData).Error
	if err != nil {
		return err
	}
	return nil
}
