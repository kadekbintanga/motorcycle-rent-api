package config

import (
	"motorcycle-rent-api/app/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	server := gin.Default()
	server.Use(cors.Default())
	server.Use(middleware.RequestID())
	return server
}
