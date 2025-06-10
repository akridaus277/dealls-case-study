package services

import (
	"reimbursement-service/models"
	"reimbursement-service/repositories"
	"reimbursement-service/utils"
	"time"
)

func SubmitReimbursement(nip string, amount float64, description string) utils.Response {
	if amount <= 0 {
		return utils.NewBadRequestResponse("Amount must be greater than zero")
	}
	if len(description) == 0 {
		return utils.NewBadRequestResponse("Description cannot be empty")
	}

	now := time.Now()
	date := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	reimb := &models.Reimbursement{
		Nip:         nip,
		Amount:      amount,
		Description: description,
		Date:        date,
		Auditable:   utils.NewAuditableCreate(nip),
	}

	err := repositories.CreateReimbursement(reimb)
	if err != nil {
		return utils.NewInternalServerErrorResponse("Failed to submit reimbursement")
	}

	return utils.NewSuccessResponse("Reimbursement submitted successfully", nil)
}

func GetReimbursementSummary(periodStart time.Time, periodEnd time.Time) utils.Response {
	summaries, err := repositories.GetSummaryReimbursement(periodStart, periodEnd)
	if err != nil {
		return utils.NewInternalServerErrorResponse("Failed to get reimbursement details")
	}

	return utils.NewSuccessResponse("Attendance summaries retrieved successfully", summaries)
}
