package global

import "time"

var GlobalConfig *EnvConfig

type EnvConfig struct {
	// AppStatus value from .env of APP_STATUS.
	AppStatus string
	// AppPort value from .env of APP_PORT.
	AppPort string

	// AutoMigration value from .env of AUTO_MIGRATION_SWITCH.
	AutoMigration bool
	// GormBatchSize value form .env of GORM_BATCH_SIZE
	GormBatchSize int

	// DBUsername value from .env of DB_USERNAME.
	DBUsername string
	// DBPassword value from .env of DB_PASSWORD.
	DBPassword string
	// DBHost value from .env of DB_HOST.
	DBHost string
	// DBPort value from .env of DB_PORT.
	DBPort string
	// DBName value from .env of DB_NAME.
	DBName string
	// DBSSLMode value from .env of DB_SSL_MODE.
	DBSSLMode string

	// DBConnMaxIdleTime value from .env of DB_CONN_MAX_IDLE_TIME.
	DBConnMaxIdleTime time.Duration
	// DBConnMaxLifeTime value from .env of DB_CONN_MAX_LIFE_TIME.
	DBConnMaxLifeTime time.Duration
	// DBMaxOpenConn value from .env of DB_MAX_OPEN_CONN.
	DBMaxOpenConn int
	// DBMaxIdleConn value from .env of DB_MAX_IDLE_CONN.
	DBMaxIdleConn int

	// LogLevel value from .env of LOG_LEVEL.
	LogLevel string
	// LogFileLocation value from .env of LOG_FILE_LOCATION.
	LogFileLocation string
	// LogSTDOUT value from .env of LOG_STDOUT.
	LogSTDOUT bool
	// JWTSecretAdmin value from .env JWT_SECRET_ADMIN
	JWTSecretAdmin string
	// JWTExpiredDurationAdmin value form .env JWT_EXPIRED_DURATION_ADMIN
	JWTExpiredDurationAdmin time.Duration
}
