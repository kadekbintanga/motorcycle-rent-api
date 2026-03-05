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

type RentalHandler struct {
	Service   service.RentalServiceInterface
	Validator *validator.Validate
}

func NewRentalHandler(service service.RentalServiceInterface, validator *validator.Validate) *RentalHandler {
	return &RentalHandler{
		Service:   service,
		Validator: validator,
	}
}

func (r *RentalHandler) CreateRental(c *gin.Context) {
	apiCallID := c.GetString(constant.RequestIDKey)
	var payload request.CreateRentalRequest

	if err := c.ShouldBind(&payload); err != nil {
		helper.ResponseAPI(c, constant.Res400InvalidPayload)
		return
	}

	err := r.Validator.Struct(payload)
	if err != nil {
		formattedErrors := helper.ErrorValidationFormatter(err.(validator.ValidationErrors))
		helper.ResponseAPI(c, constant.Res400InvalidPayload, formattedErrors)
		return
	}

	rental, res := r.Service.CreateRental(apiCallID, payload)
	helper.ResponseAPI(c, res, rental)

}

func (r *RentalHandler) ReturnRental(c *gin.Context) {
	apiCallID := c.GetString(constant.RequestIDKey)
	rentalUUID := c.Param("rentalUUID")
	err := r.Validator.Var(rentalUUID, "uuid")
	if err != nil {
		helper.LogInfo(apiCallID, "Invalid rental UUID: "+err.Error())
		helper.ResponseAPI(c, constant.Res400InvalidRentalUUID)
		return
	}

	var payload request.ReturnRentalRequest
	if err := c.ShouldBind(&payload); err != nil {
		helper.ResponseAPI(c, constant.Res400InvalidPayload)
		return
	}

	err = r.Validator.Struct(payload)
	if err != nil {
		formattedErrors := helper.ErrorValidationFormatter(err.(validator.ValidationErrors))
		helper.ResponseAPI(c, constant.Res400InvalidPayload, formattedErrors)
		return
	}

	res := r.Service.ReturnRental(apiCallID, rentalUUID, payload)
	helper.ResponseAPI(c, res)
}

func (r *RentalHandler) RentalPayment(c *gin.Context) {
	apiCallID := c.GetString(constant.RequestIDKey)
	rentalUUID := c.Param("rentalUUID")
	err := r.Validator.Var(rentalUUID, "uuid")
	if err != nil {
		helper.LogInfo(apiCallID, "Invalid rental UUID: "+err.Error())
		helper.ResponseAPI(c, constant.Res400InvalidRentalUUID)
		return
	}

	var payload request.RentalPaymentRequest
	if err := c.ShouldBind(&payload); err != nil {
		helper.ResponseAPI(c, constant.Res400InvalidPayload)
		return
	}

	err = r.Validator.Struct(payload)
	if err != nil {
		formattedErrors := helper.ErrorValidationFormatter(err.(validator.ValidationErrors))
		helper.ResponseAPI(c, constant.Res400InvalidPayload, formattedErrors)
		return
	}

	res := r.Service.RentalPayment(apiCallID, rentalUUID, payload)
	helper.ResponseAPI(c, res)
}

func (r *RentalHandler) CancelRental(c *gin.Context) {
	apiCallID := c.GetString(constant.RequestIDKey)
	rentalUUID := c.Param("rentalUUID")
	err := r.Validator.Var(rentalUUID, "uuid")
	if err != nil {
		helper.LogInfo(apiCallID, "Invalid motorcycle UUID: "+err.Error())
		helper.ResponseAPI(c, constant.Res400InvalidRentalUUID)
		return
	}

	res := r.Service.CancelRental(apiCallID, rentalUUID)
	helper.ResponseAPI(c, res)
}

func (r *RentalHandler) GetListRental(c *gin.Context) {
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
	CustomerUUID := c.Query("customer_uuid")
	MotocycleUUID := c.Query("motorcycle_uuid")
	RentDateStart := c.Query("rent_date_start")
	RentDateEnd := c.Query("rent_date_end")
	Status := c.Query("status")

	param := helper.PaginationParam{
		Page:   page,
		Limit:  limit,
		Search: search,
	}

	filter := request.GetRentalListFilter{
		CustomerUUID:   CustomerUUID,
		MotorcycleUUID: MotocycleUUID,
		RentDateStart:  RentDateStart,
		RentDateEnd:    RentDateEnd,
		Status:         Status,
	}

	err = r.Validator.Struct(filter)
	if err != nil {
		formattedErrors := helper.ErrorValidationFormatter(err.(validator.ValidationErrors))
		helper.ResponseAPI(c, constant.Res400InvalidPayload, formattedErrors)
		return
	}

	rental, meta, res := r.Service.GetListRentalPagination(apiCallID, param, filter)
	helper.ResponseAPI(c, res, rental, meta)
}

func (r *RentalHandler) GetRentalDetail(c *gin.Context) {
	apiCallID := c.GetString(constant.RequestIDKey)
	rentalUUID := c.Param("rentalUUID")
	err := r.Validator.Var(rentalUUID, "uuid")
	if err != nil {
		helper.LogInfo(apiCallID, "Invalid customer UUID: "+err.Error())
		helper.ResponseAPI(c, constant.Res400InvalidRentalUUID)
		return
	}

	rental, res := r.Service.GetDetailRental(apiCallID, rentalUUID)
	helper.ResponseAPI(c, res, rental)
}
