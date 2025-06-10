package services

import (
	"attendance-service/models"
	"attendance-service/repositories"
	"attendance-service/utils"
	"time"
)

func AddAttendancePeriod(startDate, endDate string, payrollCode string, createdBy string) utils.Response {
	start, _ := time.Parse("2006-01-02", startDate)
	end, _ := time.Parse("2006-01-02", endDate)

	period := &models.AttendancePeriod{
		StartDate:         start,
		EndDate:           end,
		PayrollPeriodCode: payrollCode,
		Auditable:         utils.NewAuditableCreate(createdBy),
	}

	if err := repositories.CreateAttendancePeriod(period); err != nil {
		return utils.NewInternalServerErrorResponse("Failed to create attendance period")
	}
	return utils.NewSuccessResponse("Attendance period created successfully", nil)
}

func GetAttendancePeriodPayroll(payrollCode string) utils.Response {

	atdPeriod, err := repositories.GetAttendancePeriodPayroll(payrollCode)
	if err != nil {
		return utils.NewInternalServerErrorResponse("Failed to get attendance period")
	}
	return utils.NewSuccessResponse("Attendance period created successfully", atdPeriod)
}
