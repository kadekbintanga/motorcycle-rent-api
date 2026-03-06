package repository

import (
	"motorcycle-rent-api/app/constant"
	"motorcycle-rent-api/app/helper"
	"motorcycle-rent-api/app/model"
	"motorcycle-rent-api/app/resource/request"
	"motorcycle-rent-api/app/resource/response"
	"time"

	"gorm.io/gorm"
)

type PaymentRepositoryInterface interface {
	CreatePayment(db *gorm.DB, payment model.Payment) (*model.Payment, error)
	GetListPaymentPagination(db *gorm.DB, param helper.PaginationParam, filter request.GetPaymentListFilter) ([]model.Payment, helper.ResponseMeta, error)
	GetPaymentByUUID(db *gorm.DB, paymentUUID string, withPreload bool) (*model.Payment, error)
	GetPaymentSummary(db *gorm.DB, filter request.GetPaymentSummaryFilter) (*response.PaymentSummary, error)
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

func (p *PaymentRepository) GetListPaymentPagination(db *gorm.DB, param helper.PaginationParam, filter request.GetPaymentListFilter) ([]model.Payment, helper.ResponseMeta, error) {
	var payment []model.Payment

	query := db.Model(&model.Payment{}).Joins("JOIN rentals ON rentals.uuid = payments.rental_uuid")

	if filter.CustomerUUID != "" {
		query = query.Where("rentals.customer_uuid = ?", filter.CustomerUUID)
	}

	if filter.MotorcycleUUID != "" {
		query = query.Where("rentals.motorcycle_uuid = ?", filter.MotorcycleUUID)
	}

	if filter.DateStart != "" {
		parseStartTime, err := time.Parse("2006-01-02 15:04:05", filter.DateStart+" 00:00:00")
		if err != nil {
			return nil, helper.ResponseMeta{}, err
		}
		query = query.Where("created_at >= ?", parseStartTime)
	}

	if filter.DateEnd != "" {
		parseEndTime, err := time.Parse("2006-01-02 15:04:05", filter.DateEnd+" 23:59:59")
		if err != nil {
			return nil, helper.ResponseMeta{}, err
		}
		query = query.Where("created_at <= ?", parseEndTime)
	}

	if filter.Method != "" {
		query = query.Where("method = ?", filter.Method)
	}

	if filter.Type != "" {
		query = query.Where("type = ?", filter.Type)
	}

	query, total, err := helper.Paginate(param.Limit, param.Page, query)
	if err != nil {
		return nil, helper.ResponseMeta{}, err
	}

	if err := query.Order("created_at DESC").Find(&payment).Error; err != nil {
		return nil, helper.ResponseMeta{}, err
	}

	meta := helper.ResponseMeta{
		Page:  param.Page,
		Limit: param.Limit,
		Total: int(total),
	}

	return payment, meta, nil
}

func (p *PaymentRepository) GetPaymentByUUID(db *gorm.DB, paymentUUID string, withPreload bool) (*model.Payment, error) {
	var payment model.Payment
	query := db.Where("uuid = ?", paymentUUID)

	if withPreload {
		query = query.Preload(string(constant.PaymentRental))
	}

	err := query.First(&payment).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (p *PaymentRepository) GetPaymentSummary(db *gorm.DB, filter request.GetPaymentSummaryFilter) (*response.PaymentSummary, error) {
	var result response.PaymentSummary
	query := db.Table("payments").
		Select(`
			COALESCE(SUM(amount),0) as total_amount,
			COALESCE(SUM(CASE WHEN method = 'CASH' THEN amount ELSE 0 END),0) as total_cash,
			COALESCE(SUM(CASE WHEN method = 'TRANSFER' THEN amount ELSE 0 END),0) as total_transfer,
			COALESCE(SUM(CASE WHEN method = 'QRIS' THEN amount ELSE 0 END),0) as total_qris
		`).
		Where("deleted_at IS NULL")
	if filter.DateStart != "" {
		parseStartTime, err := time.Parse("2006-01-02 15:04:05", filter.DateStart+" 00:00:00")
		if err != nil {
			return nil, err
		}
		query = query.Where("created_at >= ?", parseStartTime)
	}

	if filter.DateEnd != "" {
		parseEndTime, err := time.Parse("2006-01-02 15:04:05", filter.DateEnd+" 23:59:59")
		if err != nil {
			return nil, err
		}
		query = query.Where("created_at <= ?", parseEndTime)
	}

	err := query.Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}
