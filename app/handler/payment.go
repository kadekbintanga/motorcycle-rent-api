package handler

import (
	"motorcycle-rent-api/app/constant"
	"motorcycle-rent-api/app/helper"
	"motorcycle-rent-api/app/resource/request"
	"motorcycle-rent-api/app/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PaymentHandler struct {
	Service   service.PaymentServiceInterface
	Validator *validator.Validate
}

func NewPaymentHandler(service service.PaymentServiceInterface, validator *validator.Validate) *PaymentHandler {
	return &PaymentHandler{
		Service:   service,
		Validator: validator,
	}
}

func (p *PaymentHandler) GetListPayment(c *gin.Context) {
	apiCallID := c.GetString(constant.RequestIDKey)
	pageString := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		helper.LogInfo(apiCallID, "Error converting page query param: "+err.Error()+". Defaulting to page 1.")
		page = 1
	}

	limitString := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		helper.LogInfo(apiCallID, "Error converting limit query param: "+err.Error()+". Defaulting to limit 10.")
		limit = 10
	}

	CustomerUUID := c.Query("customer_uuid")
	MotocycleUUID := c.Query("motorcycle_uuid")
	DateStart := c.Query("date_start")
	DateEnd := c.Query("date_end")
	Method := c.Query("method")
	Type := c.Query("type")

	param := helper.PaginationParam{
		Page:  page,
		Limit: limit,
	}

	filter := request.GetPaymentListFilter{
		CustomerUUID:   CustomerUUID,
		MotorcycleUUID: MotocycleUUID,
		DateStart:      DateStart,
		DateEnd:        DateEnd,
		Method:         Method,
		Type:           Type,
	}

	err = p.Validator.Struct(filter)
	if err != nil {
		formattedErrors := helper.ErrorValidationFormatter(err.(validator.ValidationErrors))
		helper.ResponseAPI(c, constant.Res400InvalidPayload, formattedErrors)
		return
	}

	payment, meta, res := p.Service.GetPaymentListPagination(apiCallID, param, filter)
	helper.ResponseAPI(c, res, payment, meta)
}

func (p *PaymentHandler) GetPaymentDetail(c *gin.Context) {
	apiCallID := c.GetString(constant.RequestIDKey)
	paymentUUID := c.Param("paymentUUID")

	err := p.Validator.Var(paymentUUID, "uuid")
	if err != nil {
		helper.LogInfo(apiCallID, "Invalid payment UUID: "+err.Error())
		helper.ResponseAPI(c, constant.Res400InvalidPaymentUUID)
		return
	}

	payment, res := p.Service.GetDetailPayment(apiCallID, paymentUUID)
	helper.ResponseAPI(c, res, payment)

}

func (p *PaymentHandler) GetPaymentSummary(c *gin.Context) {
	apiCallID := c.GetString(constant.RequestIDKey)

	DateStart := c.Query("date_start")
	DateEnd := c.Query("date_end")

	filter := request.GetPaymentSummaryFilter{
		DateStart: DateStart,
		DateEnd:   DateEnd,
	}

	summary, res := p.Service.GetPaymentSummary(apiCallID, filter)
	helper.ResponseAPI(c, res, summary)
}
