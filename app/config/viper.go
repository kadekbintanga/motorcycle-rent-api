package config

import (
	"fmt"

	"motorcycle-rent-api/app/constant"
	"motorcycle-rent-api/app/global"

	"github.com/spf13/viper"
)

func InitEnvConfig() {
	config := viper.New()

	config.AutomaticEnv()
	config.AddConfigPath("./")
	config.SetConfigName(".env")
	config.SetConfigType("env")

	if err := config.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	config.SetDefault("APP_STATUS", constant.DefaultAppStatus)

	global.GlobalConfig = &global.EnvConfig{
		AppStatus:               config.GetString("APP_STATUS"),
		AppPort:                 config.GetString("APP_PORT"),
		AutoMigration:           config.GetBool("AUTO_MIGRATION_SWITCH"),
		GormBatchSize:           config.GetInt("GORM_BATCH_SIZE"),
		DBUsername:              config.GetString("DB_USERNAME"),
		DBPassword:              config.GetString("DB_PASSWORD"),
		DBHost:                  config.GetString("DB_HOST"),
		DBPort:                  config.GetString("DB_PORT"),
		DBName:                  config.GetString("DB_NAME"),
		DBSSLMode:               config.GetString("DB_SSL_MODE"),
		DBConnMaxIdleTime:       config.GetDuration("DB_CONN_MAX_IDLE_TIME"),
		DBConnMaxLifeTime:       config.GetDuration("DB_CONN_MAX_LIFE_TIME"),
		DBMaxOpenConn:           config.GetInt("DB_MAX_OPEN_CONN"),
		DBMaxIdleConn:           config.GetInt("DB_MAX_IDLE_CONN"),
		LogLevel:                config.GetString("LOG_LEVEL"),
		LogFileLocation:         config.GetString("LOG_FILE_LOCATION"),
		LogSTDOUT:               config.GetBool("LOG_STDOUT"),
		JWTSecretAdmin:          config.GetString("JWT_SECRET_ADMIN"),
		JWTExpiredDurationAdmin: config.GetDuration("JWT_EXPIRED_DURATION_ADMIN"),
	}
}
