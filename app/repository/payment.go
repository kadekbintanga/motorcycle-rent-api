package repository

import (
	"motorcycle-rent-api/app/model"

	"gorm.io/gorm"
)

type PaymentRepositoryInterface interface {
	CreatePayment(db *gorm.DB, payment model.Payment) (*model.Payment, error)
}

type PaymentRepository struct {
}

func NewPaymentRepository() PaymentRepositoryInterface {
	return &PaymentRepository{}
}

func (p *PaymentRepository) CreatePayment(db *gorm.DB, payment model.Payment) (*model.Payment, error) {
	if err := db.Create(&payment).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}
