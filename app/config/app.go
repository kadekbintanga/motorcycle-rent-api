package config

import (
	"motorcycle-rent-api/app/global"
	"motorcycle-rent-api/app/handler"
	"motorcycle-rent-api/app/repository"
	"motorcycle-rent-api/app/router"
	"motorcycle-rent-api/app/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	Config    *global.EnvConfig
	DB        *gorm.DB
	Server    *gin.Engine
	Validator *validator.Validate
	Logger    *logrus.Logger
}

func InitConfig(config *BootstrapConfig) {
	// REPOSITORY
	configRepo := repository.NewConfigRepository()
	adminRepo := repository.NewAdminRepository()
	motorcycleRepo := repository.NewMotorcycleRepository()
	customerRepo := repository.NewCustomerRepository()
	rentalRepo := repository.NewRentalRepository()
	paymentRepo := repository.NewPaymentRepository()

	// SERVICE
	healthService := service.NewHealthService(config.DB, configRepo)
	adminService := service.NewAdminService(config.DB, config.Config, adminRepo)
	motorcycleService := service.NewMotorcycleService(config.DB, motorcycleRepo)
	customerService := service.NewCustomerService(config.DB, customerRepo)
	rentalService := service.NewRentalService(config.DB, rentalRepo, customerRepo, motorcycleRepo, paymentRepo)

	// HANDLER
	healthHandler := handler.NewHealthHandler(healthService)
	adminHandler := handler.NewAdminHandler(adminService, config.Validator)
	motorcycleHandler := handler.NewMotorcycleHandler(motorcycleService, config.Validator)
	customerHandler := handler.NewCustomerHandler(customerService, config.Validator)
	rentalHandler := handler.NewRentalHandler(rentalService, config.Validator)

	// ROUTERS
	routeConfig := router.Config{
		Server:            config.Server,
		Logger:            config.Logger,
		Config:            config.Config,
		DB:                config.DB,
		HealthHandler:     healthHandler,
		AdminHandler:      adminHandler,
		MotorcycleHandler: motorcycleHandler,
		CustomerHandler:   customerHandler,
		RentalHandler:     rentalHandler,
	}

	routeConfig.Init()
}
