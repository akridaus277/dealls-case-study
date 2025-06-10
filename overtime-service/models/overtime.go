package models

import "time"

type Overtime struct {
	Auditable
	Nip   string    `json:"nip"`
	Date  time.Time `json:"date" gorm:"index"`             // Date without time part
	Hours float64   `json:"hours" gorm:"check:hours <= 3"` // Max 3 jam
}
