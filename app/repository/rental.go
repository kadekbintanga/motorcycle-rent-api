package repository

import (
	"motorcycle-rent-api/app/constant"
	"motorcycle-rent-api/app/model"

	"gorm.io/gorm"
)

type RentalRepositoryInterface interface {
	CreateRental(db *gorm.DB, rental model.Rental) (*model.Rental, error)
	GetCustomerOngoingRental(db *gorm.DB, customerUUID string) (*model.Rental, error)
	GetRentalByUUID(db *gorm.DB, rentalUUID string, withPreload bool) (*model.Rental, error)
	UpdateRentalMap(db *gorm.DB, rental model.Rental, updateData map[string]interface{}) error
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
