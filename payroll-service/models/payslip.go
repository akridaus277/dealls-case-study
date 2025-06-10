package models

type Payslip struct {
	Auditable
	PayrollID      uint    `json:"payrollId" gorm:"index"`
	Nip            string  `json:"nip" gorm:"index"`
	BaseSalary     float64 `json:"baseSalary"`
	ProratedSalary float64 `json:"proratedSalary"`
	OvertimeHours  float64 `json:"overtimeHours"`
	OvertimePay    float64 `json:"overtimePay"`
	Reimbursement  float64 `json:"reimbursement"`
	TakeHomePay    float64 `json:"takeHomePay"`
	DataJSON       string  `gorm:"type:json" json:"data_json"`
}
