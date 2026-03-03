package router

import (
	"motorcycle-rent-api/app/global"
	"motorcycle-rent-api/app/handler"
	"motorcycle-rent-api/app/middleware"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Config struct {
	Server        *gin.Engine
	Logger        *logrus.Logger
	Config        *global.EnvConfig
	DB            *gorm.DB
	HealthHandler *handler.HealthHandler
	AdminHandler  *handler.AdminHandler
}

func (c *Config) Init() {
	c.Server.GET("/health", middleware.InboundLogger(c.Logger), c.HealthHandler.HealthCheck)

	api := c.Server.Group("/api/v1")

	noLoggerGroup := api.Group("/")
	noLoggerGroup.POST("/login", c.AdminHandler.Login)

}
