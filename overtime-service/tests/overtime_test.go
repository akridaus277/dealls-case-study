package integration_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"overtime-service/controllers"
	"overtime-service/database"
	"overtime-service/middlewares"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	api := r.Group("/api", middlewares.JWTAuthMiddleware())
	{
		api.POST("/overtime", controllers.SubmitOvertime)
	}

	apiMicroservice := r.Group("/services")
	{
		apiMicroservice.POST("/get-overtime-summary", controllers.GetOvertimeSummary)
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

	code := m.Run()

	// Bersihkan test DB
	cleanupTestData()

	os.Exit(code)
}

func cleanupTestData() {
	database.DB.Debug()
	database.DB.Exec("DELETE FROM overtimes WHERE nip = ? AND date = NOW()::date", "3030300")
}

var bearerToken string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDk1NDY0ODQsInJvbGVzIjpbIkFETUlOX0hSX1BBWVJPTEwiLCJFTVBMT1lFRSJdLCJ1c2VyIjoiMzAzMDMwMCJ9.18JRvZxtntHMuY0v-cdgoATv-WvQ1hApJV6veK6skGQ"

func TestEmployeeSubmitOvertime_Failed(t *testing.T) {
	// insertTestUser()
	defer cleanupTestData()

	r := setupRouter()

	body := map[string]any{
		"hours": 3.5,
	}
	payload, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/api/overtime", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", bearerToken))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	log.Printf("w body : %v", w.Body.String())
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Overtime must be between 1 and 3 hours")
}

func TestEmployeeSubmitOvertime_Success(t *testing.T) {
	// insertTestUser()
	defer cleanupTestData()

	r := setupRouter()

	body := map[string]any{
		"hours": 2.5,
	}
	payload, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/api/overtime", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", bearerToken))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	log.Printf("w body : %v", w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Overtime submitted successfully")
}
