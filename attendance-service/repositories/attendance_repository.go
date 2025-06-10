package repositories

import (
	"attendance-service/database"
	responseDto "attendance-service/dto/response"
	"attendance-service/models"
	"fmt"
	"time"
)

func HasCheckedIn(nip string, date time.Time) (bool, error) {
	var count int64
	database.DB = database.DB.Debug()
	err := database.DB.Model(&models.Attendance{}).
		Where("nip = ? AND to_char(date, 'yyyy-mm-dd') = ?", nip, date.Format("2006-01-02")).
		Count(&count).Error
	return count > 0, err
}

func CreateAttendance(att *models.Attendance) error {
	return database.DB.Create(att).Error
}

func GetSummaryAttendance(periodStart, periodEnd time.Time) ([]responseDto.AttendanceSummaryResponse, error) {
	type RawAttendance struct {
		Nip  string    `gorm:"column:nip" json:"nip"`
		Date time.Time `gorm:"column:date" json:"date"`
	}

	var rawData []RawAttendance

	// Extend periodEnd to end of day
	periodEnd = time.Date(periodEnd.Year(), periodEnd.Month(), periodEnd.Day(), 23, 59, 59, 0, periodEnd.Location())

	database.DB = database.DB.Debug()
	err := database.DB.
		Table("attendances").
		Select("nip, date").
		Where("date BETWEEN ? AND ?", periodStart, periodEnd).
		Order("nip, date").
		Scan(&rawData).Error

	if err != nil {
		return nil, fmt.Errorf("failed to get attendance details: %w", err)
	}

	// Grouping by NIP
	grouped := make(map[string][]responseDto.AttendanceDetail)
	for _, row := range rawData {
		grouped[row.Nip] = append(grouped[row.Nip], responseDto.AttendanceDetail{
			Date: row.Date.Format("2006-01-02"),
		})
	}

	// Build response
	var result []responseDto.AttendanceSummaryResponse
	for nip, details := range grouped {
		result = append(result, responseDto.AttendanceSummaryResponse{
			NIP:              nip,
			AttendanceDetail: details,
		})
	}

	return result, nil
}
