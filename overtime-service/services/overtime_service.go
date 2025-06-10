package services

import (
	"overtime-service/models"
	"overtime-service/repositories"
	"overtime-service/utils"
	"time"
)

func SubmitOvertime(nip string, hours float64) utils.Response {
	if hours == 0 || hours > 3 {
		return utils.NewBadRequestResponse("Overtime must be between 1 and 3 hours")
	}

	now := time.Now()
	date := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	exists, err := repositories.HasSubmittedOvertimeToday(nip, date)
	if err != nil {
		return utils.NewBadRequestResponse("Failed to check overtime data")
	}

	if exists {
		return utils.NewBadRequestResponse("You have already submitted overtime today")
	}

	overtime := &models.Overtime{
		Nip:       nip,
		Date:      date,
		Hours:     hours,
		Auditable: utils.NewAuditableCreate(nip),
	}

	err = repositories.CreateOvertime(overtime)
	if err != nil {
		return utils.NewInternalServerErrorResponse("Failed to submit overtime")
	}

	return utils.NewSuccessResponse("Overtime submitted successfully", nil)
}

func GetOvertimeSummary(periodStart time.Time, periodEnd time.Time) utils.Response {
	summaries, err := repositories.GetSummaryOvertime(periodStart, periodEnd)
	if err != nil {
		return utils.NewInternalServerErrorResponse("Failed to get overtime details")
	}

	return utils.NewSuccessResponse("Attendance summaries retrieved successfully", summaries)
}
