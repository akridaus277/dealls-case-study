package seed

import (
	"payroll-service/database"
	"payroll-service/models"
	"strconv"
	"strings"
	"time"
)

func SeedPayrolls() {

	yearNow := strconv.Itoa(time.Now().Year())
	monthNow := strings.ToUpper(time.Now().AddDate(0, -1, 0).Month().String())
	monthPrev := strings.ToUpper(time.Now().Month().String())
	monthNext := strings.ToUpper(time.Now().AddDate(0, 1, 0).Month().String())

	codeNow := yearNow + "_" + monthNow
	codePrev := yearNow + "_" + monthPrev
	codeNext := yearNow + "_" + monthNext

	var countCodeNow int64
	database.DB.Model(&models.Payroll{}).
		Where("code = ?", codeNow).
		Count(&countCodeNow)
	if countCodeNow == 0 {
		payrollNow := models.Payroll{
			Code:        codeNow,
			Processed:   false,
			ProcessedAt: nil,
			Auditable: models.Auditable{
				CreatedBy: "SYSTEM",
				UpdatedBy: "SYSTEM",
			},
		}
		database.DB.Create(&payrollNow)
	}

	var countCodePrev int64
	database.DB.Model(&models.Payroll{}).
		Where("code = ?", codePrev).
		Count(&countCodePrev)
	if countCodePrev == 0 {
		payrollPrev := models.Payroll{
			Code:        codePrev,
			Processed:   false,
			ProcessedAt: nil,
			Auditable: models.Auditable{
				CreatedBy: "SYSTEM",
				UpdatedBy: "SYSTEM",
			},
		}
		database.DB.Create(&payrollPrev)
	}

	var countCodeNext int64
	database.DB.Model(&models.Payroll{}).
		Where("code = ?", codeNext).
		Count(&countCodeNext)
	if countCodeNext == 0 {
		payrollNext := models.Payroll{
			Code:        codeNext,
			Processed:   false,
			ProcessedAt: nil,
			Auditable: models.Auditable{
				CreatedBy: "SYSTEM",
				UpdatedBy: "SYSTEM",
			},
		}
		database.DB.Create(&payrollNext)
	}

}
