package model

import (
	"motorcycle-rent-api/app/constant"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Motorcycle struct {
	ID          uint                      `gorm:"column:id;primaryKey"`
	UUID        uuid.UUID                 `gorm:"column:uuid;type:uuid;default:uuid_generate_v4();index:uuid_unique;unique;"`
	PlateNumber string                    `gorm:"column:plate_number;unique"`
	Brand       string                    `gorm:"column:brand"`
	Type        constant.MotorcycleType   `gorm:"column:type;type:motorcycle_type;default:'MATIC';"`
	Year        int                       `gorm:"column:year"`
	Status      constant.MotorcycleStatus `gorm:"column:status;type:motorcycle_status;default:'AVAILABLE';"`
	PricePerDay float64                   `gorm:"column:price_per_day"`
	CreatedAt   time.Time                 `gorm:"column:created_at"`
	UpdatedAt   time.Time                 `gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt            `gorm:"column:deleted_at"`
}
