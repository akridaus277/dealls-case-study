package models

import "time"

type Attendance struct {
	Auditable
	Nip  string    `json:"nip"`
	Date time.Time `json:"date" gorm:"uniqueIndex:idx_user_date"`
}
