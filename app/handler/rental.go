package handler

import (
	"motorcycle-rent-api/app/constant"
	"motorcycle-rent-api/app/helper"
	"motorcycle-rent-api/app/resource/request"
	"motorcycle-rent-api/app/service"

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
		helper.LogInfo(apiCallID, "Invalid motorcycle UUID: "+err.Error())
		helper.ResponseAPI(c, constant.Res400InvalidMotorcycleUUID)
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
		helper.LogInfo(apiCallID, "Invalid motorcycle UUID: "+err.Error())
		helper.ResponseAPI(c, constant.Res400InvalidMotorcycleUUID)
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
