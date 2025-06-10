package requestDto

type AttendancePeriodRequest struct {
	StartDate         string `json:"startDate" binding:"required"`
	EndDate           string `json:"endDate" binding:"required"`
	PayrollPeriodCode string `json:"payrollPeriodCode" binding:"required"`
}

type AttendancePeriodPayrollRequest struct {
	PayrollPeriodCode string `json:"payrollPeriodCode" binding:"required"`
}
