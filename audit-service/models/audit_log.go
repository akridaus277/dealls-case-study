package models

type AuditLog struct {
	Auditable
	RequestID string `json:"requestId" gorm:"index"`
	IPAddress string `json:"ipAddress" `
	Action    string `json:"action" `
	UserID    *uint  `json:"userId" gorm:"index"`
	Metadata  string `json:"metadata" `
	Service   string `json:"service" `
}
