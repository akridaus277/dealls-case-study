package integration_test

import (
	"authentication-service/controllers"
	"authentication-service/database"
	"authentication-service/internal/feignclient"
	"authentication-service/models"
	"authentication-service/seed"
	"authentication-service/services"
	"authentication-service/utils"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	// Inisialisasi client dan service
	employeeClient := &MockEmployeeClient{}
	authService := services.NewAuthService(employeeClient)
	authController := controllers.NewAuthController(authService)

	api := r.Group("/api")
	{
		api.POST("/register", authController.Register)
		api.POST("/login", authController.Login)

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
	// cleanupTestUsers()

	os.Exit(code)
}

func cleanupTestUsers() {
	database.DB.Exec("DELETE FROM users WHERE email = ?", "test@example.com")
}

func insertTestUser() {
	hashedPassword, err := utils.HashPassword("Dealls30!")
	if err != nil {
		return
	}
	user := models.User{
		Email:    "test@example.com",
		Username: "3040400",
		Password: string(hashedPassword),
	}
	log.Printf("hashedpassword : %v", hashedPassword)
	if err := database.DB.Create(&user).Error; err != nil {
		log.Fatal("Failed to create test user:", err)
	}
}

type MockEmployeeClient struct{}

func (m *MockEmployeeClient) GetEmployeeByNIP(nip string) (*feignclient.EmployeeResponse, error) {
	return &feignclient.EmployeeResponse{
		NIP:  nip,
		Name: "Mock User",
	}, nil
}

func TestLogin_Success(t *testing.T) {
	// Ganti client ke mock

	seed.SeedUsers()

	r := setupRouter() // router biasa, tidak perlu ubah handler

	body := map[string]string{
		"username": "3030300",
		"password": "Dealls30!",
	}
	payload, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	log.Printf("w body : %v", w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Successfully logged in")
}

func TestLogin_InvalidPassword(t *testing.T) {
	// Ganti client ke mock
	seed.SeedUsers()

	r := setupRouter() // router biasa, tidak perlu ubah handler

	body := map[string]string{
		"username": "3030300",
		"password": "Dealls40!",
	}
	payload, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	log.Printf("w body : %v", w.Body.String())
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid credentials")
}
