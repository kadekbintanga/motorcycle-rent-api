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

func (m *MotorcycleHandler) GetListMotorcycles(c *gin.Context) {
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
	typeString := c.Query("type")
	status := c.Query("status")
	yearString := c.Query("year")
	year, err := strconv.Atoi(yearString)
	if err != nil && yearString != "" {
		helper.LogInfo(apiCallID, "Error converting year query param: "+err.Error()+". Ignoring year filter.")
		year = 0
	}

	param := helper.PaginationParam{
		Page:   page,
		Limit:  limit,
		Search: search,
	}

	filter := request.GetMotorcycleListFilter{
		Type:   typeString,
		Status: status,
		Year:   year,
	}

	err = m.Validator.Struct(filter)
	if err != nil {
		formattedErrors := helper.ErrorValidationFormatter(err.(validator.ValidationErrors))
		helper.ResponseAPI(c, constant.Res400InvalidPayload, formattedErrors)
		return
	}

	motorcycles, meta, res := m.Service.GetListMotorcyclesPagination(apiCallID, param, filter)
	helper.ResponseAPI(c, res, motorcycles, meta)

}

func (m *MotorcycleHandler) UpdateMotorcycleDetail(c *gin.Context) {
	apiCallID := c.GetString(constant.RequestIDKey)
	motorcycleUUID := c.Param("motorcycleUUID")
	err := m.Validator.Var(motorcycleUUID, "uuid")
	if err != nil {
		helper.LogInfo(apiCallID, "Invalid motorcycle UUID: "+err.Error())
		helper.ResponseAPI(c, constant.Res400InvalidMotorcycleUUID)
		return
	}

	var payload request.UpdateMotorcycleRequest
	if err := c.ShouldBind(&payload); err != nil {
		helper.ResponseAPI(c, constant.Res400InvalidPayload)
		return
	}

	err = m.Validator.Struct(payload)
	if err != nil {
		formattedErrors := helper.ErrorValidationFormatter(err.(validator.ValidationErrors))
		helper.ResponseAPI(c, constant.Res400InvalidPayload, formattedErrors)
		return
	}

	res := m.Service.UpdateMotorcycleDetail(apiCallID, motorcycleUUID, payload)
	helper.ResponseAPI(c, res)
}

func (m *MotorcycleHandler) UpdateMotorcycleStatus(c *gin.Context) {
	apiCallID := c.GetString(constant.RequestIDKey)
	motorcycleUUID := c.Param("motorcycleUUID")
	err := m.Validator.Var(motorcycleUUID, "uuid")
	if err != nil {
		helper.LogInfo(apiCallID, "Invalid motorcycle UUID: "+err.Error())
		helper.ResponseAPI(c, constant.Res400InvalidMotorcycleUUID)
		return
	}

	var payload request.UpdateMotorcycleStatusRequest
	if err := c.ShouldBind(&payload); err != nil {
		helper.ResponseAPI(c, constant.Res400InvalidPayload)
		return
	}

	helper.LogInfo(apiCallID, "status: "+string(payload.Status))

	err = m.Validator.Struct(payload)
	if err != nil {
		formattedErrors := helper.ErrorValidationFormatter(err.(validator.ValidationErrors))
		helper.ResponseAPI(c, constant.Res400InvalidPayload, formattedErrors)
		return
	}

	res := m.Service.UpdateMotorcycleStatus(apiCallID, motorcycleUUID, payload)
	helper.ResponseAPI(c, res)
}

func (m *MotorcycleHandler) GetMotocycleSummary(c *gin.Context) {
	apiCallID := c.GetString(constant.RequestIDKey)
	summary, res := m.Service.GetMotocycleSummary(apiCallID)
	helper.ResponseAPI(c, res, summary)
}
