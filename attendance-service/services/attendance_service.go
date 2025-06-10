package services

import (
	"attendance-service/models"
	"attendance-service/repositories"
	"attendance-service/utils"
	"time"
)

func SubmitAttendance(username string) utils.Response {
	now := time.Now()
	weekday := now.Weekday()

	if weekday == time.Saturday || weekday == time.Sunday {
		return utils.NewBadRequestResponse("Attendance not allowed on weekends")
	}

	alreadyCheckedIn, err := repositories.HasCheckedIn(username, now)
	if err != nil {
		return utils.NewInternalServerErrorResponse("Failed to get attendance data")
	}

	if alreadyCheckedIn {
		return utils.NewBadRequestResponse("You have already submitted attendance today")
	}

	attendance := &models.Attendance{
		Nip:       username,
		Date:      now,
		Auditable: utils.NewAuditableCreate(username),
	}

	err = repositories.CreateAttendance(attendance)
	if err != nil {
		return utils.NewInternalServerErrorResponse("Failed to submit attendance")
	}

	return utils.NewSuccessResponse("Attendance submitted successfully", nil)
}

func GetAttendanceSummary(periodStart time.Time, periodEnd time.Time) utils.Response {
	summaries, err := repositories.GetSummaryAttendance(periodStart, periodEnd)
	if err != nil {
		return utils.NewInternalServerErrorResponse("Failed to get attendance details")
	}

	return utils.NewSuccessResponse("Attendance summaries retrieved successfully", summaries)
}
