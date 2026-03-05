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
	Server            *gin.Engine
	Logger            *logrus.Logger
	Config            *global.EnvConfig
	DB                *gorm.DB
	HealthHandler     *handler.HealthHandler
	AdminHandler      *handler.AdminHandler
	MotorcycleHandler *handler.MotorcycleHandler
	CustomerHandler   *handler.CustomerHandler
	RentalHandler     *handler.RentalHandler
	PaymentHandler    *handler.PaymentHandler
}

func (c *Config) Init() {
	// c.Server.GET("/health", middleware.InboundLogger(c.Logger), c.HealthHandler.HealthCheck)

	api := c.Server.Group("/api/v1")

	noLoggerGroup := api.Group("/")
	noLoggerGroup.POST("/login", c.AdminHandler.Login)

	loggerGroup := api.Group("/", middleware.InboundLogger(c.Logger))
	loggerGroup.GET("/health", c.HealthHandler.HealthCheck)

	// Motorcycle routes
	motorcycleGroup := loggerGroup.Group("/motorcycles", middleware.AdminAuthorized())
	motorcycleGroup.POST("", c.MotorcycleHandler.CreateMotorcycle)
	motorcycleGroup.GET("", c.MotorcycleHandler.GetListMotorcycles)
	motorcycleGroup.PUT("/:motorcycleUUID", c.MotorcycleHandler.UpdateMotorcycleDetail)
	motorcycleGroup.PUT("/:motorcycleUUID/status", c.MotorcycleHandler.UpdateMotorcycleStatus)

	// Customer routes
	customerGroup := loggerGroup.Group("/customers", middleware.AdminAuthorized())
	customerGroup.POST("/", c.CustomerHandler.CreateCustomer)
	customerGroup.PUT("/:customerUUID", c.CustomerHandler.UpdateCustomerDetail)
	customerGroup.PUT("/:customerUUID/status", c.CustomerHandler.UpdateCustomerStatus)
	customerGroup.GET("/", c.CustomerHandler.GetListCustomers)
	customerGroup.GET("/:customerUUID", c.CustomerHandler.GetCustomerDetail)

	// Rental router
	rentalGroup := loggerGroup.Group("/rentals", middleware.AdminAuthorized())
	rentalGroup.POST("", c.RentalHandler.CreateRental)
	rentalGroup.POST("/:rentalUUID/return", c.RentalHandler.ReturnRental)
	rentalGroup.POST("/:rentalUUID/payment", c.RentalHandler.RentalPayment)
	rentalGroup.PUT("/:rentalUUID/cancel", c.RentalHandler.CancelRental)
	rentalGroup.GET("", c.RentalHandler.GetListRental)
	rentalGroup.GET("/:rentalUUID", c.RentalHandler.GetRentalDetail)

	// Payment router
	paymentGroup := loggerGroup.Group("/payments", middleware.AdminAuthorized())
	paymentGroup.GET("", c.PaymentHandler.GetListPayment)
	paymentGroup.GET("/:paymentUUID", c.PaymentHandler.GetPaymentDetail)

	// Summary router
	summaryGroup := loggerGroup.Group("/summaries", middleware.AdminAuthorized())
	summaryGroup.GET("/payments", c.PaymentHandler.GetPaymentSummary)
	summaryGroup.GET("/rentals", c.RentalHandler.GetRentalSummary)
	summaryGroup.GET("/customers", c.CustomerHandler.GetCustomerSummary)
	summaryGroup.GET("/motorcycles", c.MotorcycleHandler.GetMotocycleSummary)
}
