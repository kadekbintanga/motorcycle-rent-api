package main

import (
	"fmt"
	"log"

	"motorcycle-rent-api/app/config"
	"motorcycle-rent-api/app/global"
)

func main() {
	config.InitEnvConfig()
	globalConfig := global.GlobalConfig

	db := config.InitDatabase(globalConfig)
	server := config.NewServer()
	validator := config.InitValidator()

	config.InitLogger(globalConfig)
	logger := global.Logger

	config.InitConfig(&config.BootstrapConfig{
		Config:    globalConfig,
		DB:        db,
		Server:    server,
		Validator: validator,
		Logger:    logger,
	})

	if globalConfig.AutoMigration == true {
		log.Println("Database migration is ACTIVATED")
		config.MigrateDatabase(db)
		config.SeedDatabase(db, globalConfig)
	} else {
		log.Println("Database migration is DISABLED")
	}

	defer func() {
		config.CloseLoggerFile()
		config.DisconnectDatabase(db)
	}()

	err := server.Run(fmt.Sprintf(":%s", globalConfig.AppPort))
	if err != nil {
		log.Fatalf("Server error : %v", err)
	}
}
