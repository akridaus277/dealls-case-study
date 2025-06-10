package responseDto

type AttendanceDetail struct {
	Date string `json:"date"`
}

type AttendanceSummaryResponse struct {
	NIP              string             `json:"nip"`
	AttendanceDetail []AttendanceDetail `json:"attendanceDetails"`
}
