package responseDto

import "time"

type AttendancePeriodPayrollResponse struct {
	StartDate         time.Time `json:"startDate"`
	EndDate           time.Time `json:"endDate"`
	PayrollPeriodCode string    `json:"payrollPeriodCode"`
}
