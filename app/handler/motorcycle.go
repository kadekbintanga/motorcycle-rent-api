package handler

import (
	"motorcycle-rent-api/app/constant"
	"motorcycle-rent-api/app/helper"
	"motorcycle-rent-api/app/resource/request"
	"motorcycle-rent-api/app/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MotorcycleHandler struct {
	Service   service.MotorcycleServiceInterface
	Validator *validator.Validate
}

func NewMotorcycleHandler(service service.MotorcycleServiceInterface, validator *validator.Validate) *MotorcycleHandler {
	return &MotorcycleHandler{
		Service:   service,
		Validator: validator,
	}
}

func (m *MotorcycleHandler) CreateMotorcycle(c *gin.Context) {
	apiCallID := c.GetString(constant.RequestIDKey)
	var payload request.CreateMotorcycleRequest

	if err := c.ShouldBind(&payload); err != nil {
		helper.ResponseAPI(c, constant.Res400InvalidPayload)
		return
	}

	err := m.Validator.Struct(payload)
	if err != nil {
		formattedErrors := helper.ErrorValidationFormatter(err.(validator.ValidationErrors))
		helper.ResponseAPI(c, constant.Res400InvalidPayload, formattedErrors)
		return
	}

	res := m.Service.CreateMotorcycle(apiCallID, payload)
	helper.ResponseAPI(c, res)
}
