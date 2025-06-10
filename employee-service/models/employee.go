package models

type Employee struct {
	Auditable
	Name   string `json:"name"`
	Nip    string `json:"nip" gorm:"unique"`
	Salary int    `json:"salary"`
}
