package models

import (
	"time"

	"gorm.io/gorm"
)

type Auditable struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	CreatedBy string
	UpdatedBy string
	DeletedBy string
}
