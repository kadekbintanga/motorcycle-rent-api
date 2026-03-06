package repository

import (
	"motorcycle-rent-api/app/model"
	"strings"

	"gorm.io/gorm"
)

type AdminRepositoryInterface interface {
	GetAdminByEmail(db *gorm.DB, email string) (*model.Admin, error)
}

type AdminRepository struct{}

func NewAdminRepository() AdminRepositoryInterface {
	return &AdminRepository{}
}

func (a *AdminRepository) GetAdminByEmail(db *gorm.DB, email string) (*model.Admin, error) {
	var adminFound model.Admin
	err := db.Where("email = ?", strings.ToLower(email)).First(&adminFound).Error
	if err != nil {
		return nil, err
	}

	return &adminFound, nil
}
