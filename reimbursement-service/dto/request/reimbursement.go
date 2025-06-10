package requestDto

import "time"

type ReimbursementRequest struct {
	Amount      float64 `json:"amount" binding:"required"`
	Description string  `json:"description" binding:"required"`
}

type ReimbursementSummaryRequest struct {
	PeriodStart time.Time `json:"periodStart" binding:"required"`
	PeriodEnd   time.Time `json:"periodEnd" binding:"required"`
}
