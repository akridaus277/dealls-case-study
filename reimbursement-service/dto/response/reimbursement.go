package responseDto

type ReimbursementDetail struct {
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Date        string  `json:"date"`
}

type ReimbursementSummaryResponse struct {
	Nip                 string                `json:"nip"`
	ReimbursementDetail []ReimbursementDetail `json:"reimbursementDetails"`
}
