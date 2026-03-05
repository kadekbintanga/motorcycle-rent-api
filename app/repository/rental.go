package repository

import (
	"motorcycle-rent-api/app/constant"
	"motorcycle-rent-api/app/helper"
	"motorcycle-rent-api/app/model"
	"motorcycle-rent-api/app/resource/request"
	"time"

	"gorm.io/gorm"
)

type RentalRepositoryInterface interface {
	CreateRental(db *gorm.DB, rental model.Rental) (*model.Rental, error)
	GetCustomerOngoingRental(db *gorm.DB, customerUUID string) (*model.Rental, error)
	GetRentalByUUID(db *gorm.DB, rentalUUID string, withPreload bool) (*model.Rental, error)
	UpdateRentalMap(db *gorm.DB, rental model.Rental, updateData map[string]interface{}) error
	GetListRentalPagination(db *gorm.DB, param helper.PaginationParam, filter request.GetRentalListFilter) ([]model.Rental, helper.ResponseMeta, error)
}

type RentalRepository struct {
}

func NewRentalRepository() RentalRepositoryInterface {
	return &RentalRepository{}
}

func (r *RentalRepository) CreateRental(db *gorm.DB, rental model.Rental) (*model.Rental, error) {
	if err := db.Create(&rental).Error; err != nil {
		return nil, err
	}
	return &rental, nil
}

func (r *RentalRepository) GetCustomerOngoingRental(db *gorm.DB, customerUUID string) (*model.Rental, error) {
	var rental model.Rental
	err := db.Where("customer_uuid = ? AND status = ?", customerUUID, constant.RentalStatusOngoing).First(&rental).Error
	if err != nil {
		return nil, err
	}
	return &rental, nil
}

func (r *RentalRepository) GetRentalByUUID(db *gorm.DB, rentalUUID string, withPreload bool) (*model.Rental, error) {
	var rental model.Rental
	query := db.Where("uuid = ?", rentalUUID)

	if withPreload {
		query = query.Preload(string(constant.RentalPayment))
	}

	err := query.First(&rental).Error
	if err != nil {
		return nil, err
	}
	return &rental, nil
}

func (r *RentalRepository) UpdateRentalMap(db *gorm.DB, rental model.Rental, updateData map[string]interface{}) error {
	err := db.Model(rental).Updates(updateData).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *RentalRepository) GetListRentalPagination(db *gorm.DB, param helper.PaginationParam, filter request.GetRentalListFilter) ([]model.Rental, helper.ResponseMeta, error) {
	var rental []model.Rental

	query := db.Model(&model.Rental{}).Debug()

	if param.Search != "" {
		query = query.Where("customer_name_captured ILIKE ? OR customer_id_number_captured ILIKE ? OR customer_sim_number_captured ILIKE ? OR customer_phone_captured ILIKE ? OR motorcycle_plate_number_captured ILIKE ?", "%"+param.Search+"%", "%"+param.Search+"%", "%"+param.Search+"%", "%"+param.Search+"%", "%"+param.Search+"%")
	}

	if filter.CustomerUUID != "" {
		query = query.Where("customer_uuid = ?", filter.CustomerUUID)
	}

	if filter.MotorcycleUUID != "" {
		query = query.Where("motorcycle_uuid = ?", filter.MotorcycleUUID)
	}

	if filter.RentDateStart != "" {
		parseStartTime, err := time.Parse("2006-01-02", filter.RentDateStart)
		if err != nil {
			return nil, helper.ResponseMeta{}, err
		}
		query = query.Where("rent_date >= ?", parseStartTime)
	}

	if filter.RentDateEnd != "" {
		parseEndTime, err := time.Parse("2006-01-02", filter.RentDateEnd)
		if err != nil {
			return nil, helper.ResponseMeta{}, err
		}
		query = query.Where("rent_date <= ?", parseEndTime)
	}

	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}

	query, total, err := helper.Paginate(param.Limit, param.Page, query)
	if err != nil {
		return nil, helper.ResponseMeta{}, err
	}

	if err := query.Order("created_at DESC").Find(&rental).Error; err != nil {
		return nil, helper.ResponseMeta{}, err
	}

	meta := helper.ResponseMeta{
		Page:  param.Page,
		Limit: param.Limit,
		Total: int(total),
	}

	return rental, meta, nil

}
