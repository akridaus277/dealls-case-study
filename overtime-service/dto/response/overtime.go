package responseDto

type OvertimeDetail struct {
	Date  string  `json:"date"`
	Hours float64 `json:"hours" `
}

type OvertimeSummaryResponse struct {
	Nip            string           `json:"nip"`
	OvertimeDetail []OvertimeDetail `json:"overtimeDetails"`
}
