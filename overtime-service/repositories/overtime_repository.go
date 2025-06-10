package repositories

import (
	"fmt"
	"overtime-service/database"
	responseDto "overtime-service/dto/response"
	"overtime-service/models"
	"time"
)

func HasSubmittedOvertimeToday(nip string, date time.Time) (bool, error) {
	var count int64
	err := database.DB.Model(&models.Overtime{}).
		Where("nip = ? AND date = ?", nip, date.Format("2006-01-02")).
		Count(&count).Error
	return count > 0, err
}

func CreateOvertime(overtime *models.Overtime) error {
	return database.DB.Create(overtime).Error
}

func GetSummaryOvertime(periodStart, periodEnd time.Time) ([]responseDto.OvertimeSummaryResponse, error) {
	type RawOvertime struct {
		Nip   string    `gorm:"column:nip" json:"nip"`
		Date  time.Time `gorm:"column:date" json:"date"`
		Hours float64   `gorm:"column:hours" json:"hours"`
	}

	var rawData []RawOvertime

	// Extend periodEnd to end of day
	periodEnd = time.Date(periodEnd.Year(), periodEnd.Month(), periodEnd.Day(), 23, 59, 59, 0, periodEnd.Location())

	database.DB = database.DB.Debug()
	err := database.DB.
		Table("overtimes").
		Select("nip, date, hours").
		Where("date BETWEEN ? AND ?", periodStart, periodEnd).
		Order("nip, date").
		Scan(&rawData).Error

	if err != nil {
		return nil, fmt.Errorf("failed to get overtime details: %w", err)
	}

	// Grouping by NIP
	grouped := make(map[string][]responseDto.OvertimeDetail)
	for _, row := range rawData {
		grouped[row.Nip] = append(grouped[row.Nip], responseDto.OvertimeDetail{
			Date:  row.Date.Format("2006-01-02"),
			Hours: row.Hours,
		})
	}

	// Build response
	var result []responseDto.OvertimeSummaryResponse
	for nip, details := range grouped {
		result = append(result, responseDto.OvertimeSummaryResponse{
			Nip:            nip,
			OvertimeDetail: details,
		})
	}

	return result, nil
}
