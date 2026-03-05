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

type CustomerHandler struct {
	Service   service.CustomerServiceInterface
	Validator *validator.Validate
}

func NewCustomerHandler(service service.CustomerServiceInterface, validator *validator.Validate) *CustomerHandler {
	return &CustomerHandler{
		Service:   service,
		Validator: validator,
	}
}

func (ch *CustomerHandler) CreateCustomer(c *gin.Context) {
	apiCallID := c.GetString(constant.RequestIDKey)
	var payload request.CreateCustomerRequest

	if err := c.ShouldBind(&payload); err != nil {
		helper.ResponseAPI(c, constant.Res400InvalidPayload)
		return
	}

	err := ch.Validator.Struct(payload)
	if err != nil {
		formattedErrors := helper.ErrorValidationFormatter(err.(validator.ValidationErrors))
		helper.ResponseAPI(c, constant.Res400InvalidPayload, formattedErrors)
		return
	}

	res := ch.Service.CreateCustomer(apiCallID, payload)
	helper.ResponseAPI(c, res)
}

func (ch *CustomerHandler) UpdateCustomerDetail(c *gin.Context) {
	apiCallID := c.GetString(constant.RequestIDKey)
	customerUUID := c.Param("customerUUID")
	err := ch.Validator.Var(customerUUID, "uuid")
	if err != nil {
		helper.LogInfo(apiCallID, "Invalid customer UUID: "+err.Error())
		helper.ResponseAPI(c, constant.Res400InvalidCustomerUUID)
		return
	}

	var payload request.UpdateCustomerDetailRequest
	if err := c.ShouldBind(&payload); err != nil {
		helper.ResponseAPI(c, constant.Res400InvalidPayload)
		return
	}

	err = ch.Validator.Struct(payload)
	if err != nil {
		formattedErrors := helper.ErrorValidationFormatter(err.(validator.ValidationErrors))
		helper.ResponseAPI(c, constant.Res400InvalidPayload, formattedErrors)
		return
	}

	res := ch.Service.UpdateCustomerDetail(apiCallID, customerUUID, payload)
	helper.ResponseAPI(c, res)
}

func (ch *CustomerHandler) UpdateCustomerStatus(c *gin.Context) {
	apiCallID := c.GetString(constant.RequestIDKey)
	customerUUID := c.Param("customerUUID")
	err := ch.Validator.Var(customerUUID, "uuid")
	if err != nil {
		helper.LogInfo(apiCallID, "Invalid customer UUID: "+err.Error())
		helper.ResponseAPI(c, constant.Res400InvalidCustomerUUID)
		return
	}

	var payload request.UpdateCustomerStatusRequest
	if err := c.ShouldBind(&payload); err != nil {
		helper.ResponseAPI(c, constant.Res400InvalidPayload)
		return
	}

	err = ch.Validator.Struct(payload)
	if err != nil {
		formattedErrors := helper.ErrorValidationFormatter(err.(validator.ValidationErrors))
		helper.ResponseAPI(c, constant.Res400InvalidPayload, formattedErrors)
		return
	}

	res := ch.Service.UpdateCustomerStatus(apiCallID, customerUUID, payload)
	helper.ResponseAPI(c, res)
}

func (ch *CustomerHandler) GetListCustomers(c *gin.Context) {
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

	search := c.Query("search")
	status := c.Query("status")

	param := helper.PaginationParam{
		Page:   page,
		Limit:  limit,
		Search: search,
	}

	filter := request.GetCustomerListFilter{
		Status: status,
	}

	err = ch.Validator.Struct(filter)
	if err != nil {
		formattedErrors := helper.ErrorValidationFormatter(err.(validator.ValidationErrors))
		helper.ResponseAPI(c, constant.Res400InvalidPayload, formattedErrors)
		return
	}

	customer, meta, res := ch.Service.GetListCustomerPagination(apiCallID, param, filter)
	helper.ResponseAPI(c, res, customer, meta)
}

func (ch *CustomerHandler) GetCustomerDetail(c *gin.Context) {
	apiCallID := c.GetString(constant.RequestIDKey)
	customerUUID := c.Param("customerUUID")
	err := ch.Validator.Var(customerUUID, "uuid")
	if err != nil {
		helper.LogInfo(apiCallID, "Invalid customer UUID: "+err.Error())
		helper.ResponseAPI(c, constant.Res400InvalidCustomerUUID)
		return
	}

	customer, res := ch.Service.GetDetailCustomer(apiCallID, customerUUID)
	helper.ResponseAPI(c, res, customer)
}

func (ch *CustomerHandler) GetCustomerSummary(c *gin.Context) {
	apiCallID := c.GetString(constant.RequestIDKey)
	summary, res := ch.Service.GetCustomerSummary(apiCallID)
	helper.ResponseAPI(c, res, summary)
}
