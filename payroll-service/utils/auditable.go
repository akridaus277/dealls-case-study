package utils

import "payroll-service/models"

func NewAuditableCreate(user string) models.Auditable {
	return models.Auditable{
		CreatedBy: user,
		UpdatedBy: user,
	}
}

func NewAuditableUpdate(user string) models.Auditable {
	return models.Auditable{
		UpdatedBy: user,
	}
}

func NewAuditableDelete(user string) models.Auditable {
	return models.Auditable{
		UpdatedBy: user,
		DeletedBy: user,
	}
}
