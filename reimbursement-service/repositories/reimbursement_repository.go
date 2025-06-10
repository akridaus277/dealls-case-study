package repositories

import (
	"fmt"
	"reimbursement-service/database"
	responseDto "reimbursement-service/dto/response"
	"reimbursement-service/models"
	"time"
)

func CreateReimbursement(reimb *models.Reimbursement) error {
	return database.DB.Create(reimb).Error
}

func GetSummaryReimbursement(periodStart, periodEnd time.Time) ([]responseDto.ReimbursementSummaryResponse, error) {
	type RawReimbursement struct {
		Nip         string    `gorm:"column:nip" json:"nip"`
		Date        time.Time `gorm:"column:date" json:"date"`
		Amount      float64   `gorm:"column:amount" json:"amount"`
		Description string    `gorm:"column:description" json:"description"`
	}

	var rawData []RawReimbursement

	// Extend periodEnd to end of day
	periodEnd = time.Date(periodEnd.Year(), periodEnd.Month(), periodEnd.Day(), 23, 59, 59, 0, periodEnd.Location())

	database.DB = database.DB.Debug()
	err := database.DB.
		Table("reimbursements").
		Select("nip, date, amount, description").
		Where("date BETWEEN ? AND ?", periodStart, periodEnd).
		Order("nip, date").
		Scan(&rawData).Error

	if err != nil {
		return nil, fmt.Errorf("failed to get overtime details: %w", err)
	}

	// Grouping by NIP
	grouped := make(map[string][]responseDto.ReimbursementDetail)
	for _, row := range rawData {
		grouped[row.Nip] = append(grouped[row.Nip], responseDto.ReimbursementDetail{
			Date:        row.Date.Format("2006-01-02"),
			Amount:      row.Amount,
			Description: row.Description,
		})
	}

	// Build response
	var result []responseDto.ReimbursementSummaryResponse
	for nip, details := range grouped {
		result = append(result, responseDto.ReimbursementSummaryResponse{
			Nip:                 nip,
			ReimbursementDetail: details,
		})
	}

	return result, nil
}
