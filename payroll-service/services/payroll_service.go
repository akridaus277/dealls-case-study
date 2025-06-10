package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"payroll-service/internal/feignclient"
	"payroll-service/models"
	"payroll-service/repositories"
	"payroll-service/utils"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"github.com/xuri/excelize/v2"
)

const OvertimeMultiplier = 2.0 // 2x gaji per jam

type MergedEmployee struct {
	Nip                 string                            `json:"nip"`
	Name                string                            `json:"name"`
	Salary              int                               `json:"salary"`
	AttendanceDetail    []feignclient.AttendanceDetail    `json:"attendanceDetails"`
	OvertimeDetail      []feignclient.OvertimeDetail      `json:"overtimeDetails"`
	ReimbursementDetail []feignclient.ReimbursementDetail `json:"reimbursementDetails"`
}

type PayrollService struct {
	EmployeeClient      feignclient.IEmployeeClient
	AttendanceClient    feignclient.IAttendanceClient
	OvertimeClient      feignclient.IOvertimeClient
	ReimbursementClient feignclient.IReimbursementClient
}

func NewPayrollService(employeeClient feignclient.IEmployeeClient, attendanceClient feignclient.IAttendanceClient, overtimeClient feignclient.IOvertimeClient, reimbursementClient feignclient.IReimbursementClient) *PayrollService {
	return &PayrollService{
		EmployeeClient:      employeeClient,
		AttendanceClient:    attendanceClient,
		OvertimeClient:      overtimeClient,
		ReimbursementClient: reimbursementClient,
	}
}

func (s *PayrollService) RunPayroll(payrollCode string) utils.Response {
	// Cek apakah payroll sudah pernah diproses
	payroll, err := repositories.GetPayrollByCode(payrollCode)
	if err != nil {
		return utils.NewBadRequestResponse("Payroll for this period not found")
	}

	_, err = repositories.GetPayslipsByPayroll(payroll.ID)
	if err != nil {
		return utils.NewBadRequestResponse("Payroll for this period has been generated")
	}

	atdPeriod, err := s.GetAttendancePeriodPayroll(payroll.Code)
	if err != nil {
		return utils.NewInternalServerErrorResponse("Failed to get attendance period for this payroll")
	}
	fmt.Printf("atdPeriod : %v", atdPeriod)
	periodStart := atdPeriod.StartDate
	periodEnd := atdPeriod.EndDate

	emps, err := s.GetAllEmployees()
	if err != nil {
		return utils.NewInternalServerErrorResponse("Failed to get employees")
	}
	fmt.Printf("emps : %v", emps)

	atdSum, err := s.GetAttendanceSummary(periodStart, periodEnd)
	if err != nil {
		return utils.NewInternalServerErrorResponse("Failed to get attendance summary")
	}
	fmt.Printf("atdSum : %v", atdSum)

	ovtSum, err := s.GetOvertimeSummary(periodStart, periodEnd)
	if err != nil {
		return utils.NewInternalServerErrorResponse("Failed to get overtime summary")
	}
	fmt.Printf("ovtSum : %v", ovtSum)

	rmbSum, err := s.GetReimbursementSummary(periodStart, periodEnd)
	if err != nil {
		return utils.NewInternalServerErrorResponse("Failed to get reimbursement summary")
	}
	fmt.Printf("rmbSum : %v", rmbSum)

	mergedEmployees, err := s.generatePayrollPayslip(payroll.ID, periodStart, periodEnd, *emps, *atdSum, *ovtSum, *rmbSum)
	if err != nil {
		return utils.NewInternalServerErrorResponse("Failed to save payslip")
	}
	fmt.Printf("mergedEmployees : %v", mergedEmployees)

	return utils.NewSuccessResponse("Payroll ran successfully", mergedEmployees)
}

func (s *PayrollService) GetAllEmployees() (*[]feignclient.EmployeeResponse, error) {
	emps, err := s.EmployeeClient.GetAllEmployee()
	if err != nil {
		return nil, err
	}

	return emps, nil
}

func (s *PayrollService) GetAttendancePeriodPayroll(payrollCode string) (*feignclient.AttendancePeriodPayrollResponse, error) {
	atdPeriod, err := s.AttendanceClient.GetAttendancePeriodPayroll(payrollCode)
	if err != nil {
		return nil, err
	}

	return atdPeriod, nil
}

func (s *PayrollService) GetOvertimeSummary(periodStart, periodEnd time.Time) (*[]feignclient.OvertimeSummaryResponse, error) {
	atdSum, err := s.OvertimeClient.GetOvertimeSummary(periodStart, periodEnd)
	if err != nil {
		return nil, err
	}

	return atdSum, nil
}

func (s *PayrollService) GetAttendanceSummary(periodStart, periodEnd time.Time) (*[]feignclient.AttendanceSummaryResponse, error) {
	atdSum, err := s.AttendanceClient.GetAttendanceSummary(periodStart, periodEnd)
	if err != nil {
		return nil, err
	}

	return atdSum, nil
}

func (s *PayrollService) GetReimbursementSummary(periodStart, periodEnd time.Time) (*[]feignclient.ReimbursementSummaryResponse, error) {
	rmbSum, err := s.ReimbursementClient.GetReimbursementSummary(periodStart, periodEnd)
	if err != nil {
		return nil, err
	}

	return rmbSum, nil
}

// Gabungkan dua array di memory
func (s *PayrollService) generatePayrollPayslip(payrollId uint, periodStart time.Time, periodEnd time.Time, employees []feignclient.EmployeeResponse, attendancesSummary []feignclient.AttendanceSummaryResponse, overtimeSummary []feignclient.OvertimeSummaryResponse, reimbursementSummary []feignclient.ReimbursementSummaryResponse) ([]MergedEmployee, error) {

	// Buat map untuk mempercepat lookup attendance berdasarkan NIP
	attendanceMap := make(map[string][]feignclient.AttendanceDetail)
	for _, att := range attendancesSummary {
		attendanceMap[att.Nip] = att.AttendanceDetail
	}
	overtimeMap := make(map[string][]feignclient.OvertimeDetail)
	for _, att := range overtimeSummary {
		overtimeMap[att.Nip] = att.OvertimeDetail
	}
	reimbursementMap := make(map[string][]feignclient.ReimbursementDetail)
	for _, att := range reimbursementSummary {
		reimbursementMap[att.Nip] = att.ReimbursementDetail
	}

	// Gabungkan berdasarkan NIP
	var result []MergedEmployee
	for _, emp := range employees {
		dataJson := MergedEmployee{
			Nip:                 emp.NIP,
			Name:                emp.Name,
			Salary:              emp.Salary,
			AttendanceDetail:    attendanceMap[emp.NIP],
			OvertimeDetail:      overtimeMap[emp.NIP],
			ReimbursementDetail: reimbursementMap[emp.NIP],
		}
		result = append(result, dataJson)
		totalWeekday := s.CountWeekdays(periodStart, periodEnd)
		dailySalary := float64(emp.Salary / totalWeekday)
		proratedSalary := float64(len(attendanceMap[emp.NIP])) * dailySalary
		var overtimeHours float64
		for _, details := range overtimeMap {
			for _, d := range details {
				overtimeHours += d.Hours
			}
		}
		var reimbursementAmount float64
		for _, details := range reimbursementMap {
			for _, d := range details {
				reimbursementAmount += d.Amount
			}
		}
		overtimePay := float64(overtimeHours) * (dailySalary / 8.0) * OvertimeMultiplier
		takeHome := proratedSalary + overtimePay + reimbursementAmount

		// Marshal kembali ke string JSON
		updatedDataJSONBytes, err := json.Marshal(dataJson)
		if err != nil {
			return nil, errors.New("")
		}

		// Konversi ke string
		updatedDataJSONString := string(updatedDataJSONBytes)

		payslip := &models.Payslip{
			PayrollID:      payrollId,
			Nip:            emp.NIP,
			BaseSalary:     float64(emp.Salary),
			ProratedSalary: proratedSalary,
			OvertimeHours:  overtimeHours,
			OvertimePay:    overtimePay,
			Reimbursement:  reimbursementAmount,
			TakeHomePay:    takeHome,
			DataJSON:       updatedDataJSONString,
		}
		repositories.SavePayslip(payslip)
	}
	return result, nil
}

func (s *PayrollService) CountWeekdays(start, end time.Time) int {
	if start.After(end) {
		start, end = end, start
	}

	count := 0
	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		weekday := d.Weekday()
		if weekday >= time.Monday && weekday <= time.Friday {
			count++
		}
	}
	return count
}

func (s *PayrollService) GeneratePayslipPDF(c *gin.Context, payrollCode string, nip string) {
	// Cek apakah payroll ada
	payroll, err := repositories.GetPayrollByCode(payrollCode)
	if err != nil {
		utils.SendResponse(c, utils.NewBadRequestResponse("Payroll for this period not found"))
	}
	// 1. Ambil payslip dari DB
	payslip, err := repositories.GetPayslipForEmployee(payroll.ID, nip)
	if err != nil {
		utils.SendResponse(c, utils.NewBadRequestResponse("Payslip not found"))

	}

	// 2. Decode DataJSON
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(payslip.DataJSON), &data); err != nil {
		utils.SendResponse(c, utils.NewBadRequestResponse("Failed to process payslip data"))

	}

	// 3. Inisialisasi PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 14)

	// Header
	pdf.Cell(40, 10, "Payslip")
	pdf.Ln(12)

	// Employee Info
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, fmt.Sprintf("NIP: %s", payslip.Nip))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Name: %v", data["name"]))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Base Salary: Rp %.2f", payslip.BaseSalary))
	pdf.Ln(12)

	// Attendance Breakdown
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(40, 10, "Attendance Details")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 11)
	if details, ok := data["attendanceDetails"].([]interface{}); ok {
		for _, d := range details {
			row := d.(map[string]interface{})
			pdf.Cell(40, 8, fmt.Sprintf("- Date: %s", row["date"]))
			pdf.Ln(6)
		}
	}
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Prorated Salary: Rp %.2f", payslip.ProratedSalary))
	pdf.Ln(12)

	// Overtime Breakdown
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(40, 10, "Overtime Details")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 11)
	if details, ok := data["overtimeDetails"].([]interface{}); ok {
		for _, d := range details {
			row := d.(map[string]interface{})
			pdf.Cell(60, 8, fmt.Sprintf("- Date: %s, Hours: %.2f", row["date"], row["hours"]))
			pdf.Ln(6)
		}
	}
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Overtime Pay: Rp %.2f", payslip.OvertimePay))
	pdf.Ln(12)

	// Reimbursements
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(40, 10, "Reimbursement Details")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 11)
	if details, ok := data["reimbursementDetails"].([]interface{}); ok {
		for _, d := range details {
			row := d.(map[string]interface{})
			pdf.Cell(100, 8, fmt.Sprintf("- %s: Rp %.0f (Date: %s)", row["description"], row["amount"], row["date"]))
			pdf.Ln(6)
		}
	}
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Total Reimbursement: Rp %.2f", payslip.Reimbursement))
	pdf.Ln(12)

	// Total Take-Home
	pdf.SetFont("Arial", "B", 13)
	pdf.Cell(40, 10, fmt.Sprintf("Total Take-Home Pay: Rp %.2f", payslip.TakeHomePay))

	// 4. Output PDF
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", fmt.Sprintf("inline; filename=payslip_%s.pdf", payslip.Nip))
	if err := pdf.Output(c.Writer); err != nil {
		utils.SendResponse(c, utils.NewInternalServerErrorResponse("Failed to generate PDF"))
	}
}

// Handler: GET /payslip/export-excel
func (s *PayrollService) ExportPayslipExcel(c *gin.Context, payrollCode string) {

	payroll, err := repositories.GetPayrollByCode(payrollCode)
	if err != nil {
		utils.SendResponse(c, utils.NewBadRequestResponse("Payroll for this period not found"))
	}
	// 1. Ambil payslip dari DB
	payslips, err := repositories.GetPayslipsByPayroll(payroll.ID)
	if err != nil {
		utils.SendResponse(c, utils.NewBadRequestResponse("Payslip not found"))

	}

	// Buat Excel baru
	f := excelize.NewFile()
	sheet := "Payslip Summary"
	index, err := f.NewSheet(sheet)
	if err != nil {
		utils.SendResponse(c, utils.NewInternalServerErrorResponse("Failed to create sheet"))
	}

	// Header kolom
	headers := []string{"NIP", "Name", "Base Salary", "Prorated Salary", "Overtime Hours", "Overtime Pay", "Reimbursement", "Take Home Pay"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	// Isi data
	totalTHP := 0.0
	for i, payslip := range payslips {
		// Decode DataJSON
		var data map[string]interface{}
		if err := json.Unmarshal([]byte(payslip.DataJSON), &data); err != nil {
			continue // skip jika gagal parsing
		}

		row := i + 2
		values := []interface{}{
			payslip.Nip,
			data["name"],
			payslip.BaseSalary,
			payslip.ProratedSalary,
			payslip.OvertimeHours,
			payslip.OvertimePay,
			payslip.Reimbursement,
			payslip.TakeHomePay,
		}

		// Tulis ke Excel
		for j, val := range values {
			cell, _ := excelize.CoordinatesToCellName(j+1, row)
			f.SetCellValue(sheet, cell, val)
		}

		// Akumulasi total
		totalTHP += payslip.TakeHomePay
	}

	// Tambahkan baris total di bawah
	totalRow := len(payslips) + 3
	f.SetCellValue(sheet, fmt.Sprintf("G%d", totalRow), "TOTAL THP")
	f.SetCellValue(sheet, fmt.Sprintf("H%d", totalRow), totalTHP)

	f.SetActiveSheet(index)

	// Kirim ke response sebagai file download
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=payslip_summary.xlsx")
	c.Header("File-Name", "payslip_summary.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")
	if err := f.Write(c.Writer); err != nil {
		utils.SendResponse(c, utils.NewInternalServerErrorResponse("Failed to generate Excel"))
	}
}
