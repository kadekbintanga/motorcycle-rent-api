package model

import (
	"motorcycle-rent-api/app/constant"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	ID         uint                   `gorm:"column:id;primaryKey"`
	UUID       uuid.UUID              `gorm:"column:uuid;type:uuid;default:uuid_generate_v4();index:uuid_unique;unique;"`
	RentalUUID uuid.UUID              `gorm:"column:rental_uuid;type:uuid;index"`
	Amount     float64                `gorm:"column:amount;type:decimal(12,2)"`
	Method     constant.PaymentMethod `gorm:"column:method;type:payment_method"`
	Type       constant.PaymentType   `gorm:"column:type;type:payment_type"`
	CreatedAt  time.Time              `gorm:"column:created_at"`
	UpdatedAt  time.Time              `gorm:"column:updated_at"`
	DeletedAt  gorm.DeletedAt         `gorm:"column:deleted_at"`
	Rental     Rental                 `gorm:"foreignKey:RentalUUID;references:UUID"`
}
