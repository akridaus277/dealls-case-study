package requestDto

type AuditLog struct {
	RequestID string `json:"requestId" gorm:"index"`
	IPAddress string `json:"ipAddress" `
	Action    string `json:"action" `
	Username  string `json:"username" gorm:"index"`
	Metadata  string `json:"metadata" `
	Service   string `json:"service" `
}
