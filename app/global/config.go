package global

import "time"

var GlobalConfig *EnvConfig

type EnvConfig struct {
	AppStatus               string
	AppPort                 string
	AutoMigration           bool
	GormBatchSize           int
	DBUsername              string
	DBPassword              string
	DBHost                  string
	DBPort                  string
	DBName                  string
	DBSSLMode               string
	DBConnMaxIdleTime       time.Duration
	DBConnMaxLifeTime       time.Duration
	DBMaxOpenConn           int
	DBMaxIdleConn           int
	LogLevel                string
	LogFileLocation         string
	LogSTDOUT               bool
	JWTSecretAdmin          string
	JWTExpiredDurationAdmin time.Duration
	BlacklistDayLate        int
}
