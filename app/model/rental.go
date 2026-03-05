package model

import (
	"motorcycle-rent-api/app/constant"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Rental struct {
	ID                            uint                  `gorm:"column:id;primaryKey"`
	UUID                          uuid.UUID             `gorm:"column:uuid;type:uuid;default:uuid_generate_v4();index:uuid_unique;unique;"`
	CustomerUUID                  uuid.UUID             `gorm:"column:customer_uuid;type:uuid;index"`
	MotorcycleUUID                uuid.UUID             `gorm:"column:motorcycle_uuid;type:uuid;index"`
	RentDate                      time.Time             `gorm:"column:rent_date;type:date"`
	ReturnDatePlan                time.Time             `gorm:"column:return_date_plan;type:date"`
	ReturnDateActual              *time.Time            `gorm:"column:return_date_actual;type:date"`
	LateDay                       int                   `gorm:"column:late_day"`
	PricePerDayCaptured           float64               `gorm:"column:price_per_day_captured;type:decimal(12,2)"`
	CustomerNameCaptured          string                `gorm:"column:customer_name_captured;type:varchar(255)"`
	CustomerIDNumberCaptured      string                `gorm:"column:customer_id_number_captured;type:varchar(255)"`
	CustomerSIMNumberCaptured     string                `gorm:"column:customer_sim_number_captured;type:varchar(255)"`
	CustomerPhoneCaptured         string                `gorm:"column:customer_phone_captured;type:varchar(255)"`
	MotorcyclePlateNumberCaptured string                `gorm:"column:motorcycle_plate_number_captured;type:varchar(255)"`
	RentPrice                     float64               `gorm:"column:rent_price;type:decimal(12,2)"`
	PenaltyPrice                  float64               `gorm:"column:penalty_price;type:decimal(12,2)"`
	Status                        constant.RentalStatus `gorm:"column:status;type:rental_status;default:'ONGOING'"`
	CreatedAt                     time.Time             `gorm:"column:created_at"`
	UpdatedAt                     time.Time             `gorm:"column:updated_at"`
	DeletedAt                     gorm.DeletedAt        `gorm:"column:deleted_at"`
	Customer                      Customer              `gorm:"foreignKey:CustomerUUID;references:UUID"`
	Motorcycle                    Motorcycle            `gorm:"foreignKey:MotorcycleUUID;references:UUID"`
	Payment                       []Payment             `gorm:"foreignKey:RentalUUID;references:UUID;"`
}
