package service

import (
	"errors"
	"motorcycle-rent-api/app/constant"
	"motorcycle-rent-api/app/helper"
	"motorcycle-rent-api/app/repository"
	"motorcycle-rent-api/app/resource/request"
	"motorcycle-rent-api/app/resource/response"

	"gorm.io/gorm"
)

type PaymentServiceInterface interface {
	GetPaymentListPagination(apiCallID string, param helper.PaginationParam, filter request.GetPaymentListFilter) ([]response.PaymentListResponse, *helper.ResponseMeta, constant.ResponseMap)
	GetDetailPayment(apiCallID string, paymentUUID string) (*response.PaymentDetailtResponse, constant.ResponseMap)
	GetPaymentSummary(apiCallID string, filter request.GetPaymentSummaryFilter) (*response.PaymentSummary, constant.ResponseMap)
}

type PaymentService struct {
	DB                 *gorm.DB
	PaymentRespository repository.PaymentRepositoryInterface
}

func NewPaymentService(db *gorm.DB, paymentRepository repository.PaymentRepositoryInterface) PaymentServiceInterface {
	return &PaymentService{
		DB:                 db,
		PaymentRespository: paymentRepository,
	}
}

func (p *PaymentService) GetPaymentListPagination(apiCallID string, param helper.PaginationParam, filter request.GetPaymentListFilter) ([]response.PaymentListResponse, *helper.ResponseMeta, constant.ResponseMap) {
	payments, meta, err := p.PaymentRespository.GetListPaymentPagination(p.DB, param, filter)
	if err != nil {
		helper.LogError(apiCallID, "Error getting rental list: "+err.Error())
		return nil, nil, constant.Res422SomethingWentWrong
	}

	formatterPaymentList := response.PaymentListPaginationFormatter(payments)
	return formatterPaymentList, &meta, constant.Res200Get
}

func (p *PaymentService) GetDetailPayment(apiCallID string, paymentUUID string) (*response.PaymentDetailtResponse, constant.ResponseMap) {
	payment, err := p.PaymentRespository.GetPaymentByUUID(p.DB, paymentUUID, true)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.LogError(apiCallID, "Rental not found")
			return nil, constant.Res400RentalNotFound
		}
		helper.LogError(apiCallID, "Unable to get rental with error : "+err.Error())
		return nil, constant.Res422SomethingWentWrong
	}
	formatterPaymentDetail := response.PaymentDetailFormatter(*payment)
	return &formatterPaymentDetail, constant.Res200Get
}

func (p *PaymentService) GetPaymentSummary(apiCallID string, filter request.GetPaymentSummaryFilter) (*response.PaymentSummary, constant.ResponseMap) {
	summary, err := p.PaymentRespository.GetPaymentSummary(p.DB, filter)
	if err != nil {
		helper.LogError(apiCallID, "Error getting payment summary: "+err.Error())
		return nil, constant.Res422SomethingWentWrong
	}

	return summary, constant.Res200Get
}
