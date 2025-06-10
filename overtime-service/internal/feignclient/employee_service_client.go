package feignclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// Response body jika ditemukan
type EmployeeResponse struct {
	ID   uint   `json:"id"`
	NIP  string `json:"nip"`
	Name string `json:"name"`
	// Tambahkan field lain sesuai struktur response employee-service
}

// Interface-like struct
type EmployeeClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

// Constructor (bisa dipanggil dari main atau inject lewat service)
func NewEmployeeClient() *EmployeeClient {
	baseURL := os.Getenv("EMPLOYEE_SERVICE_BASE_URL")

	return &EmployeeClient{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *EmployeeClient) GetEmployeeByNIP(nip string) (*EmployeeResponse, error) {
	url := fmt.Sprintf("%s/services/employee/get-by-nip", c.BaseURL)

	// Buat request body dalam bentuk JSON
	reqBody := map[string]string{
		"nip": nip,
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to encode request body: %w", err)
	}

	// Buat POST request dengan body
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call employee service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("nip not found")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("employee service error: %s", resp.Status)
	}

	var emp EmployeeResponse
	if err := json.NewDecoder(resp.Body).Decode(&emp); err != nil {
		return nil, fmt.Errorf("invalid response format from employee service: %w", err)
	}

	return &emp, nil
}
