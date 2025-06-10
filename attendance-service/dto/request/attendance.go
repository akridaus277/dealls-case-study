package requestDto

import "time"

type AttendanceSummaryRequest struct {
	PeriodStart time.Time `json:"periodStart" binding:"required"`
	PeriodEnd   time.Time `json:"periodEnd" binding:"required"`
}
