package service

import (
	"audit-service/models"
	"audit-service/repositories"
)

type AuditService struct {
	repo *repositories.AuditRepository
}

func NewAuditService(repo *repositories.AuditRepository) *AuditService {
	return &AuditService{repo}
}

func (s *AuditService) Log(log *models.AuditLog) error {
	return s.repo.Save(log)
}
