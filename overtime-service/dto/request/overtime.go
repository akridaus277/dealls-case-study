package requestDto

import "time"

type OvertimeRequest struct {
	Hours float64 `json:"hours" binding:"required"`
}

type OvertimeSummaryRequest struct {
	PeriodStart time.Time `json:"periodStart" binding:"required"`
	PeriodEnd   time.Time `json:"periodEnd" binding:"required"`
}
