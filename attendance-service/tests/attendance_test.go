package integration_test

import (
	"attendance-service/controllers"
	"attendance-service/database"
	"attendance-service/middlewares"
	"attendance-service/seed"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	admin := r.Group("/admin", middlewares.JWTAuthAdminMiddleware())
	{
		admin.POST("/attendance-period", controllers.AddAttendancePeriod)
	}

	api := r.Group("/api", middlewares.JWTAuthMiddleware())
	{
		api.POST("/attendance", controllers.SubmitAttendance)
	}

	apiMicroservice := r.Group("/services")
	{
		apiMicroservice.POST("/get-attendance-summary", controllers.GetAttendanceSummary)
		apiMicroservice.POST("/get-attendance-period", controllers.GetAttendancePeriodPayroll)
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

	seed.SeedAttendancePeriod()

	code := m.Run()

	// Bersihkan test DB
	cleanupTestData()

	os.Exit(code)
}

func cleanupTestData() {
	database.DB.Exec("DELETE FROM attendance_periods WHERE payroll_period_code = ?", "3025_JUNE")
	database.DB.Exec("DELETE FROM attendances WHERE date BETWEEN ? AND ?", time.Now().Add(time.Duration(-1)*time.Minute), time.Now().Add(time.Duration(1)*time.Minute))
}

var bearerToken string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDk1NDY0ODQsInJvbGVzIjpbIkFETUlOX0hSX1BBWVJPTEwiLCJFTVBMT1lFRSJdLCJ1c2VyIjoiMzAzMDMwMCJ9.18JRvZxtntHMuY0v-cdgoATv-WvQ1hApJV6veK6skGQ"

func TestAdminCreateAttendancePeriod_Success(t *testing.T) {
	// insertTestUser()
	defer cleanupTestData()

	r := setupRouter()

	body := map[string]any{
		"startDate":         "3025-06-01",
		"endDate":           "3025-06-30",
		"payrollPeriodCode": "3025_JUNE",
	}
	payload, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/admin/attendance-period", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", bearerToken))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	log.Printf("w body : %v", w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Attendance period created successfully")
}

func TestEmployeeSubmitAttendance_Success(t *testing.T) {
	// insertTestUser()
	defer cleanupTestData()

	r := setupRouter()

	req, _ := http.NewRequest("POST", "/api/attendance", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", bearerToken))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	log.Printf("w body : %v", w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Attendance submitted successfully")
}

func TestEmployeeSubmitAttendance_Failed(t *testing.T) {
	// insertTestUser()
	defer cleanupTestData()

	r := setupRouter()

	req, _ := http.NewRequest("POST", "/api/attendance", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", bearerToken))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	log.Printf("w body : %v", w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Attendance submitted successfully")

	w2 := httptest.NewRecorder()
	r.ServeHTTP(w2, req)

	log.Printf("w2 body : %v", w2.Body.String())
	assert.Equal(t, http.StatusBadRequest, w2.Code)
	assert.Contains(t, w2.Body.String(), "You have already submitted attendance today")
}
