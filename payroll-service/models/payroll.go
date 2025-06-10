package models

import "time"

type Payroll struct {
	Auditable
	Code        string     `json:"code" gorm:"unique"`
	Processed   bool       `json:"processed" gorm:"default:false"`
	ProcessedAt *time.Time `json:"processedAt"`
}
