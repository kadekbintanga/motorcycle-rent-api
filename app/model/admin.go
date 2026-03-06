package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	ID        uint           `gorm:"column:id;primaryKey"`
	UUID      uuid.UUID      `gorm:"column:uuid;type:uuid;default:uuid_generate_v4();index:uuid_unique;unique;"`
	Name      string         `gorm:"column:name"`
	Email     string         `gorm:"column:email;unique;type:varchar(255)"`
	Password  string         `gorm:"column:password"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index;"`
}
