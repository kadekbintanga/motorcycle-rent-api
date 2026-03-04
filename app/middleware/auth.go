package middleware

import (
	"motorcycle-rent-api/app/constant"
	"motorcycle-rent-api/app/global"
	"motorcycle-rent-api/app/helper"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	BearerSchema = "Bearer "
	ClaimsKey    = "admin_claims"
)

func AdminAuthorized(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiID := c.GetString(constant.RequestIDKey)
		header := c.GetHeader("Authorization")

		if header == "" {
			helper.LogError(apiID, "Unauthorized Access : Authorization header is missing")
			res := constant.Res401Unauthorized
			response := helper.Response{
				ApiID:   apiID,
				Status:  res.Status,
				Message: res.Message,
				Data:    nil,
				Meta:    nil,
			}
			c.AbortWithStatusJSON(res.Code, response)
			return
		}

		tokenString := strings.TrimPrefix(header, BearerSchema)
		if tokenString == "" {
			helper.LogError(apiID, "Unauthorized Access : Token is missing")
			res := constant.Res401Unauthorized
			response := helper.Response{
				ApiID:   apiID,
				Status:  res.Status,
				Message: res.Message,
				Data:    nil,
				Meta:    nil,
			}
			c.AbortWithStatusJSON(res.Code, response)
			return
		}

		claims, err := helper.ValidateJWTAdmin(tokenString, global.GlobalConfig.JWTSecretAdmin, db)
		if err != nil {
			helper.LogError(apiID, "Unauthorized Access : "+err.Error())
			res := constant.Res401Unauthorized
			response := helper.Response{
				ApiID:   apiID,
				Status:  res.Status,
				Message: res.Message,
				Data:    nil,
				Meta:    nil,
			}
			c.AbortWithStatusJSON(res.Code, response)
			return
		}

		c.Set(ClaimsKey, claims)
		c.Next()
	}
}
