package middleware

import (
	"motorcycle-rent-api/app/constant"
	"motorcycle-rent-api/app/helper"

	"github.com/gin-gonic/gin"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiID := helper.GenerateApiCallID()
		c.Set(constant.RequestIDKey, apiID)

		c.Next()
	}
}
