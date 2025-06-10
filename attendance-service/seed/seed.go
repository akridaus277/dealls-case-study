package seed

import (
	"attendance-service/database"
	"attendance-service/models"
	"strconv"
	"strings"
	"time"
)

func SeedAttendancePeriod() {

	yearNow := strconv.Itoa(time.Now().Year())
	monthNow := strings.ToUpper(time.Now().AddDate(0, -1, 0).Month().String())
	monthPrev := strings.ToUpper(time.Now().Month().String())
	monthNext := strings.ToUpper(time.Now().AddDate(0, 1, 0).Month().String())

	codeNow := yearNow + "_" + monthNow
	codePrev := yearNow + "_" + monthPrev
	codeNext := yearNow + "_" + monthNext

	var countCodeNow int64
	database.DB.Model(&models.AttendancePeriod{}).
		Where("payroll_period_code = ?", codeNow).
		Count(&countCodeNow)
	if countCodeNow == 0 {
		firstOfMonth := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.Now().Location())
		lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
		atdPeriodNow := models.AttendancePeriod{
			PayrollPeriodCode: codeNow,
			StartDate:         firstOfMonth,
			EndDate:           lastOfMonth,
			Auditable: models.Auditable{
				CreatedBy: "SYSTEM",
				UpdatedBy: "SYSTEM",
			},
		}
		database.DB.Create(&atdPeriodNow)
	}

	var countCodePrev int64
	database.DB.Model(&models.AttendancePeriod{}).
		Where("payroll_period_code = ?", codePrev).
		Count(&countCodePrev)
	if countCodePrev == 0 {
		firstOfMonth := time.Date(time.Now().Year(), time.Now().AddDate(0, -1, 0).Month(), 1, 0, 0, 0, 0, time.Now().Location())
		lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
		atdPeriodPrev := models.AttendancePeriod{
			PayrollPeriodCode: codePrev,
			StartDate:         firstOfMonth,
			EndDate:           lastOfMonth,
			Auditable: models.Auditable{
				CreatedBy: "SYSTEM",
				UpdatedBy: "SYSTEM",
			},
		}
		database.DB.Create(&atdPeriodPrev)
	}

	var countCodeNext int64
	database.DB.Model(&models.AttendancePeriod{}).
		Where("payroll_period_code = ?", codeNext).
		Count(&countCodeNext)
	if countCodeNext == 0 {
		firstOfMonth := time.Date(time.Now().Year(), time.Now().AddDate(0, 1, 0).Month(), 1, 0, 0, 0, 0, time.Now().Location())
		lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
		atdPeriodNext := models.AttendancePeriod{
			PayrollPeriodCode: codeNext,
			StartDate:         firstOfMonth,
			EndDate:           lastOfMonth,
			Auditable: models.Auditable{
				CreatedBy: "SYSTEM",
				UpdatedBy: "SYSTEM",
			},
		}
		database.DB.Create(&atdPeriodNext)
	}

}
