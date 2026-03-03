package config

import (
	"fmt"
	"log"
	"strings"

	"motorcycle-rent-api/app/constant"
	"motorcycle-rent-api/app/global"
	"motorcycle-rent-api/app/model"
	"motorcycle-rent-api/app/seeder"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(config *global.EnvConfig) *gorm.DB {
	log.Println("» Trying to connect database")
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", config.DBHost, config.DBPort, config.DBUsername, config.DBName, config.DBPassword, config.DBSSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("» Failed to connect database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("» Failed to get database instance: %v", err)
	}

	sqlDB.SetConnMaxIdleTime(config.DBConnMaxIdleTime)
	sqlDB.SetConnMaxLifetime(config.DBConnMaxLifeTime)
	sqlDB.SetMaxOpenConns(config.DBMaxOpenConn)
	sqlDB.SetMaxIdleConns(config.DBMaxIdleConn)

	log.Println("» Database Successfully connected ")

	return db
}

func DisconnectDatabase(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to kill connection from database")
	}
	err = dbSQL.Close()
	if err != nil {
		panic("Failed to kill connection from database")
	}
}

func MigrateEnum(db *gorm.DB) {
	var err error
	// Motorcycle Status
	if err = MigrateEnumType(db, "motorcycle_status", []string{
		string(constant.MotorcycleStatusAvailable),
		string(constant.MotorcycleStatusRented),
		string(constant.MotorcycleStatusMaintenance),
		string(constant.MotorcycleStatusInactive),
	}); err != nil {
		log.Fatalf("Failed to migrate enum motorcycle_status type: %v", err)
	}

	// Motorcycle Type
	if err = MigrateEnumType(db, "motorcycle_type", []string{
		string(constant.MotorcycleTypeMatic),
		string(constant.MotorcycleTypeManual),
	}); err != nil {
		log.Fatalf("Failed to migrate enum motorcycle_type type: %v", err)
	}
}

func MigrateTable(db *gorm.DB) {
	var err error

	models := []interface{}{
		&model.Config{},
		&model.Admin{},
		&model.Motorcycle{},
	}

	for _, modelMigrate := range models {
		if err = db.AutoMigrate(modelMigrate); err != nil {
			log.Printf("[ERROR] Error migrating %T: %v", modelMigrate, err)
		}
	}
}

func MigrateDatabase(db *gorm.DB) {
	InitUUID(db)

	MigrateEnum(db)

	MigrateTable(db)
}

func SeedDatabase(db *gorm.DB, config *global.EnvConfig) {
	seeder.SeedConfigs(db, config)
	seeder.SeedSuperAdmin(db)
	seeder.SeedMotorcycles(db)
}

func InitUUID(db *gorm.DB) {
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
}

func DropColumn(dst interface{}, column string, db *gorm.DB) {
	if db.Migrator().HasColumn(dst, column) {
		err := db.Migrator().DropColumn(dst, column)
		if err != nil {
			log.Println("[ERROR] Error drop " + column + " column: " + err.Error())
		}
	}
}

func DropIndex(dst interface{}, index string, db *gorm.DB) {
	if db.Migrator().HasIndex(dst, index) {
		err := db.Migrator().DropIndex(dst, index)
		if err != nil {
			log.Println("[ERROR] Error drop index " + index + ": " + err.Error())
		}
	}
}

func CreateJsonIndex(indexName, table, column string, db *gorm.DB) {
	err := db.Exec(fmt.Sprintf("CREATE INDEX IF NOT EXISTS %s ON %s USING gin (%s)", indexName, table, column)).Error
	if err != nil {
		log.Println("[ERROR] Error create Json index " + indexName + " for " + column + " column in " + table + " table : " + err.Error())
	}
}

func CreateLowerCaseIndex(indexName, table, column string, db *gorm.DB) {
	err := db.Exec(fmt.Sprintf("CREATE INDEX IF NOT EXISTS %s ON %s (LOWER(%s))", indexName, table, column)).Error
	if err != nil {
		log.Println("[ERROR] Error create LowerCase index " + indexName + " for " + column + " column in " + table + " table : " + err.Error())
	}
}

func MigrateEnumType(db *gorm.DB, enumTypeName string, enumValues []string) error {
	if len(enumValues) == 0 {
		return fmt.Errorf("enum values cannot be empty")
	}

	var quotedValues []string
	for _, val := range enumValues {
		escaped := strings.ReplaceAll(val, "'", "''")
		quotedValues = append(quotedValues, fmt.Sprintf("'%s'", escaped))
	}

	createQuery := fmt.Sprintf(`
		DO $$
		BEGIN
			IF NOT EXISTS (
				SELECT 1 FROM pg_type WHERE typname = '%s'
			) THEN
				CREATE TYPE %s AS ENUM (%s);
			END IF;
		END $$;
	`, enumTypeName, enumTypeName, strings.Join(quotedValues, ", "))

	if err := db.Exec(createQuery).Error; err != nil {
		return fmt.Errorf("failed to create enum type: %w", err)
	}

	return nil
}

func AddEnumValueIfNotExists(db *gorm.DB, enumName string, newValue string) error {
	query := fmt.Sprintf(`
        DO $$
        BEGIN
            IF NOT EXISTS (
                SELECT 1
                FROM pg_enum
                JOIN pg_type ON pg_enum.enumtypid = pg_type.oid
                WHERE pg_type.typname = '%s' AND enumlabel = '%s'
            ) THEN
                ALTER TYPE %s ADD VALUE '%s';
            END IF;
        END $$;
    `, enumName, newValue, enumName, newValue)

	if err := db.Exec(query).Error; err != nil {
		return fmt.Errorf("Failed to add value '%s' to enum %s: %v\n", newValue, enumName, err)
	}

	return nil
}
