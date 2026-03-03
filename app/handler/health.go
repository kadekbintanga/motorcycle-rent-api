package handler

import (
	"motorcycle-rent-api/app/constant"
	"motorcycle-rent-api/app/helper"
	"motorcycle-rent-api/app/service"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	HealthService service.HealthServiceInterface
}

func NewHealthHandler(healthService service.HealthServiceInterface) *HealthHandler {
	return &HealthHandler{
		HealthService: healthService,
	}
}

func (h *HealthHandler) HealthCheck(ctx *gin.Context) {
	apiCallID := ctx.GetString(constant.RequestIDKey)
	data, res := h.HealthService.Health(apiCallID)
	helper.ResponseAPI(ctx, res, data)
}
