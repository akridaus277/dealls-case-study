package models

import "time"

type AttendancePeriod struct {
	StartDate         time.Time `json:"startDate"`
	EndDate           time.Time `json:"endDate"`
	PayrollPeriodCode string    `json:"payrollPeriodCode" gorm:"unique"`
	Auditable
}
