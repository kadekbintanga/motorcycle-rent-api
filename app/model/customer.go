package model

import (
	"motorcycle-rent-api/app/constant"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	ID              uint                    `gorm:"column:id;primaryKey"`
	UUID            uuid.UUID               `gorm:"column:uuid;type:uuid;default:uuid_generate_v4();index:uuid_unique;unique;"`
	Name            string                  `gorm:"column:name;type:varchar(255)"`
	IDNumber        string                  `gorm:"column:id_number;unique;type:varchar(255)"`
	SIMNumber       string                  `gorm:"column:sim_number;unique;type:varchar(255)"`
	Phone           string                  `gorm:"column:phone;type:varchar(255)"`
	Address         string                  `gorm:"column:address"`
	Status          constant.CustomerStatus `gorm:"column:status;type:customer_status;default:'ACTIVE';"`
	BlacklistReason string                  `gorm:"column:blacklist_reason"`
	CreatedAt       time.Time               `gorm:"column:created_at"`
	UpdatedAt       time.Time               `gorm:"column:updated_at"`
	DeletedAt       gorm.DeletedAt          `gorm:"column:deleted_at"`
}
