package repositories

import (
	"attendance-service/database"
	responseDto "attendance-service/dto/response"
	"attendance-service/models"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func CreateAttendancePeriod(period *models.AttendancePeriod) error {
	return database.DB.Create(period).Error
}

func GetAttendancePeriodPayroll(payrollCode string) (*responseDto.AttendancePeriodPayrollResponse, error) {
	var result responseDto.AttendancePeriodPayrollResponse

	db := database.DB

	err := db.
		Table("attendance_periods").
		Where("payroll_period_code = ?", payrollCode).
		Scan(&result).Error

	if err != nil {
		log.Printf("attendance period err: %+v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("attendance period not found for payroll ID %s", payrollCode)
		}
		return nil, fmt.Errorf("failed to get attendance period: %w", err)
	}

	log.Printf("attendance period: %+v", result)
	return &result, nil
}
