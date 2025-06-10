package repositories

import (
	"log"
	"payroll-service/database"
	"payroll-service/models"
)

func CreatePayroll(p *models.Payroll) error {
	return database.DB.Create(p).Error
}

func GetPayrollByPeriod(start, end string) (*models.Payroll, error) {
	var payroll models.Payroll
	err := database.DB.Where("period_start = ? AND period_end = ?", start, end).First(&payroll).Error
	return &payroll, err
}

func GetPayrollById(payrollId int) (*models.Payroll, error) {
	var payroll models.Payroll
	database.DB = database.DB.Debug()
	err := database.DB.Where("id = ? ", payrollId).First(&payroll).Error
	log.Printf("isi payroll : %v", payroll)
	return &payroll, err
}

func GetPayrollByCode(payrollCode string) (*models.Payroll, error) {
	var payroll models.Payroll
	database.DB = database.DB.Debug()
	err := database.DB.Where("code = ? ", payrollCode).First(&payroll).Error
	log.Printf("isi payroll : %v", payroll)
	return &payroll, err
}

func SavePayslip(p *models.Payslip) error {
	return database.DB.Create(p).Error
}

func GetPayslipsByPayroll(payrollID uint) ([]models.Payslip, error) {
	var payslips []models.Payslip
	database.DB.Debug()
	err := database.DB.Where("payroll_id = ?", payrollID).Find(&payslips).Error
	return payslips, err
}

func GetPayslipForEmployee(payrollID uint, nip string) (*models.Payslip, error) {
	var payslip models.Payslip
	err := database.DB.Where("payroll_id = ? AND nip = ?", payrollID, nip).First(&payslip).Error
	return &payslip, err
}
