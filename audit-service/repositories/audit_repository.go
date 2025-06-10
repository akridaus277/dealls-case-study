package repositories

import (
	"audit-service/models"

	"gorm.io/gorm"
)

type AuditRepository struct {
	db *gorm.DB
}

func NewAuditRepository(db *gorm.DB) *AuditRepository {
	return &AuditRepository{db}
}

func (r *AuditRepository) Save(log *models.AuditLog) error {
	return r.db.Create(log).Error
}
