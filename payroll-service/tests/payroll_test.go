package integration_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"payroll-service/controllers"
	"payroll-service/database"
	"payroll-service/internal/feignclient"
	"payroll-service/middlewares"
	"payroll-service/seed"
	"payroll-service/services"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	employeeClient := &MockEmployeeClient{}
	attendanceClient := &MockAttendanceClient{}
	overtimeClient := &MockOvertimeClient{}
	reimbursementClient := &MockReimbursementClient{}
	payrollService := services.NewPayrollService(employeeClient, attendanceClient, overtimeClient, reimbursementClient)
	payrollController := controllers.NewPayrollController(payrollService)

	admin := r.Group("/admin", middlewares.JWTAuthAdminMiddleware())
	{
		admin.POST("/payroll/run", payrollController.RunPayroll)
		admin.POST("/payslip/report-excel", payrollController.GetPayslipAllEmployee)
	}
	api := r.Group("/api", middlewares.JWTAuthMiddleware())
	{
		api.POST("/payslip", payrollController.GetPayslip)
	}
	return r
}

func TestMain(m *testing.M) {
	// Load .env.test
	err := godotenv.Load(".env.test")
	if err != nil {
		log.Fatal("Failed to load .env.test:", err)
	}

	// Init database
	database.Init()

	seed.SeedPayrolls()

	code := m.Run()

	// Bersihkan test DB
	// cleanupTestUsers()

	os.Exit(code)
}

type MockEmployeeClient struct{}

func (m *MockEmployeeClient) GetEmployeeByNIP(nip string) (*feignclient.EmployeeResponse, error) {
	return &feignclient.EmployeeResponse{NIP: "3030300", Name: "Mock User"}, nil
}

func (m *MockEmployeeClient) GetAllEmployee() (*[]feignclient.EmployeeResponse, error) {
	mockData := []feignclient.EmployeeResponse{
		{NIP: "3030300", Name: "Harun"},
		{NIP: "3030302", Name: "Haran"},
	}

	return &mockData, nil
}

type MockAttendanceClient struct{}

func (m *MockAttendanceClient) GetAttendancePeriodPayroll(payrollCode string) (*feignclient.AttendancePeriodPayrollResponse, error) {
	startDate, _ := time.Parse(time.RFC3339, "2025-06-01T07:00:00+07:00")
	endDate, _ := time.Parse(time.RFC3339, "2025-06-30T07:00:00+07:00")
	return &feignclient.AttendancePeriodPayrollResponse{PayrollPeriodCode: "2025_JUNE", StartDate: startDate, EndDate: endDate}, nil
}

func (m *MockAttendanceClient) GetAttendanceSummary(periodStart, periodEnd time.Time) (*[]feignclient.AttendanceSummaryResponse, error) {
	mockData := []feignclient.AttendanceSummaryResponse{
		{
			Nip: "3030300",
			AttendanceDetail: []feignclient.AttendanceDetail{
				{
					Date: "2025-06-02",
				},
			},
		},
		{
			Nip: "3030302",
			AttendanceDetail: []feignclient.AttendanceDetail{
				{Date: "2025-06-02"}, {Date: "2025-06-03"},
			},
		},
	}

	return &mockData, nil
}

type MockOvertimeClient struct{}

func (m *MockOvertimeClient) GetOvertimeSummary(periodStart, periodEnd time.Time) (*[]feignclient.OvertimeSummaryResponse, error) {
	mockData := []feignclient.OvertimeSummaryResponse{
		{
			Nip: "3030300",
			OvertimeDetail: []feignclient.OvertimeDetail{
				{
					Date:  "2025-06-02",
					Hours: 2.5,
				},
			},
		},
		{
			Nip: "3030302",
			OvertimeDetail: []feignclient.OvertimeDetail{
				{
					Date:  "2025-06-02",
					Hours: 2.0,
				},
			},
		},
	}

	return &mockData, nil
}

type MockReimbursementClient struct{}

func (m *MockReimbursementClient) GetReimbursementSummary(periodStart, periodEnd time.Time) (*[]feignclient.ReimbursementSummaryResponse, error) {
	mockData := []feignclient.ReimbursementSummaryResponse{
		{
			Nip: "3030300",
			ReimbursementDetail: []feignclient.ReimbursementDetail{
				{
					Date:        "2025-06-02",
					Amount:      100000,
					Description: "parkir bulanan",
				},
			},
		},
		{
			Nip: "3030302",
			ReimbursementDetail: []feignclient.ReimbursementDetail{
				{
					Date:        "2025-06-02",
					Amount:      200000,
					Description: "parkir bulanan mobil",
				},
			},
		},
	}

	return &mockData, nil
}

func cleanupTestData() {
	database.DB.Exec("DELETE FROM payslips WHERE created_at BETWEEN ? AND ?", time.Now().Add(time.Duration(-1)*time.Minute), time.Now().Add(time.Duration(1)*time.Minute))
}

var bearerToken string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDk1NDY0ODQsInJvbGVzIjpbIkFETUlOX0hSX1BBWVJPTEwiLCJFTVBMT1lFRSJdLCJ1c2VyIjoiMzAzMDMwMCJ9.18JRvZxtntHMuY0v-cdgoATv-WvQ1hApJV6veK6skGQ"

func TestRunPayroll_Success(t *testing.T) {
	// Ganti client ke mock

	defer cleanupTestData()

	r := setupRouter()

	body := map[string]string{
		"payrollCode": "2025_JUNE",
	}
	payload, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/admin/payroll/run", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", bearerToken))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	log.Printf("w body : %v", w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Payroll ran successfully")
}

func TestRunPayroll_Failed(t *testing.T) {
	// Ganti client ke mock

	defer cleanupTestData()

	r := setupRouter()

	body := map[string]string{
		"payrollCode": "2025_DECEMBER",
	}
	payload, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/admin/payroll/run", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", bearerToken))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	log.Printf("w body : %v", w.Body.String())
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Payroll for this period not found")
}

func TestGetPayslip_Success(t *testing.T) {
	// Ganti client ke mock
	defer cleanupTestData()

	r := setupRouter()

	body := map[string]string{
		"payrollCode": "2025_JUNE",
	}
	payload, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/admin/payroll/run", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", bearerToken))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	log.Printf("w body : %v", w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Payroll ran successfully")

	body = map[string]string{
		"payrollCode": "2025_JUNE",
	}
	payload, _ = json.Marshal(body)

	req, _ = http.NewRequest("POST", "/api/payslip", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", bearerToken))

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	log.Printf("w body : %v", w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetPayslip_Failed(t *testing.T) {
	// Ganti client ke mock
	defer cleanupTestData()

	r := setupRouter()

	body := map[string]string{
		"payrollCode": "2025_JUNE",
	}
	payload, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/admin/payroll/run", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", bearerToken))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	log.Printf("w body : %v", w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Payroll ran successfully")

	body = map[string]string{
		"payrollCode": "2025_DECEMBER",
	}
	payload, _ = json.Marshal(body)

	req, _ = http.NewRequest("POST", "/api/payslip", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", bearerToken))

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	log.Printf("w body : %v", w.Body.String())
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetPayslipReport_Success(t *testing.T) {
	// Ganti client ke mock
	defer cleanupTestData()

	r := setupRouter()

	body := map[string]string{
		"payrollCode": "2025_JUNE",
	}
	payload, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/admin/payroll/run", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", bearerToken))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	log.Printf("w body : %v", w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Payroll ran successfully")

	body = map[string]string{
		"payrollCode": "2025_JUNE",
	}
	payload, _ = json.Marshal(body)

	req, _ = http.NewRequest("POST", "/admin/payslip/report-excel", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", bearerToken))

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	log.Printf("w body : %v", w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetPayslipReport_Failed(t *testing.T) {
	// Ganti client ke mock
	defer cleanupTestData()

	r := setupRouter()

	body := map[string]string{
		"payrollCode": "2025_JUNE",
	}
	payload, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/admin/payroll/run", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", bearerToken))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	log.Printf("w body : %v", w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Payroll ran successfully")

	body = map[string]string{
		"payrollCode": "2025_DECEMBER",
	}
	payload, _ = json.Marshal(body)

	req, _ = http.NewRequest("POST", "/admin/payslip/report-excel", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", bearerToken))

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	log.Printf("w body : %v", w.Body.String())
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
