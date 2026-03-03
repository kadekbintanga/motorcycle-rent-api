package handler

import (
	"motorcycle-rent-api/app/constant"
	"motorcycle-rent-api/app/helper"
	"motorcycle-rent-api/app/resource/request"
	"motorcycle-rent-api/app/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AdminHandler struct {
	Service   service.AdminServiceInterface
	Validator *validator.Validate
}

func NewAdminHandler(service service.AdminServiceInterface, validator *validator.Validate) *AdminHandler {
	return &AdminHandler{
		Service:   service,
		Validator: validator,
	}
}

func (a *AdminHandler) Login(c *gin.Context) {
	apiCallID := c.GetString(constant.RequestIDKey)
	var payload request.AdminLoginRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		helper.ResponseAPI(c, constant.Res400InvalidPayload)
		return
	}

	err := a.Validator.Struct(payload)
	if err != nil {
		formattedErrors := helper.ErrorValidationFormatter(err.(validator.ValidationErrors))
		helper.ResponseAPI(c, constant.Res400InvalidPayload, formattedErrors)
		return
	}

	data, res := a.Service.Login(apiCallID, payload)
	helper.ResponseAPI(c, res, data)

}
