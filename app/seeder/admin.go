package seeder

import (
	"motorcycle-rent-api/app/helper"
	"motorcycle-rent-api/app/model"

	"gorm.io/gorm"
)

func SeedSuperAdmin(db *gorm.DB) {
	hashedPassword, _ := helper.Hash("Password123!")

	superAdmin := model.Admin{
		Name:     "Admin 1",
		Email:    "admin1@mailinator.com",
		Password: string(hashedPassword),
	}

	db.FirstOrCreate(&superAdmin)
}
