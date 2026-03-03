package seeder

import (
	"motorcycle-rent-api/app/constant"
	"motorcycle-rent-api/app/global"
	"motorcycle-rent-api/app/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedConfigs(db *gorm.DB, config *global.EnvConfig) {
	configs := []model.Config{
		{
			ID:    1,
			Key:   constant.AppStatusKey,
			Value: config.AppStatus,
		},
	}

	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"key", "value"}),
	}).Create(&configs)
}
