package models

import "time"

type Reimbursement struct {
	Auditable
	Nip         string    `json:"nip" gorm:"index"`
	Amount      float64   `json:"amount" gorm:"not null"`
	Description string    `json:"description" gorm:"type:text"`
	Date        time.Time `json:"date" gorm:"index"` // Date without time part
}
