package feignclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Response body jika ditemukan
type EmployeeResponse struct {
	ID     uint   `json:"id"`
	NIP    string `json:"nip"`
	Name   string `json:"name"`
	Salary int    `json:"salary"`
	// Tambahkan field lain sesuai struktur response employee-service
}

// Interface-like struct
type EmployeeClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

type IEmployeeClient interface {
	GetEmployeeByNIP(nip string) (*EmployeeResponse, error)
	GetAllEmployee() (*[]EmployeeResponse, error)
}

// Constructor (bisa dipanggil dari main atau inject lewat service)
func NewEmployeeClient() IEmployeeClient {
	baseURL := os.Getenv("EMPLOYEE_SERVICE_BASE_URL")
	if baseURL == "" {
		log.Println("Warning: EMPLOYEE_SERVICE_BASE_URL not set")
	}
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

	var respBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, fmt.Errorf("invalid response format from employee service: %w", err)
	}

	data, ok := respBody["data"]
	if !ok {
		return nil, fmt.Errorf("field 'data' not found in response")
	}

	// Step 1: Marshal (encode) data map ke JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Step 2: Unmarshal ke struct EmployeeResponse
	var empResp EmployeeResponse
	err = json.Unmarshal(jsonData, &empResp)
	if err != nil {
		return nil, err
	}
	log.Printf("isi emp response : %v", &empResp)
	return &empResp, nil
}

func (c *EmployeeClient) GetAllEmployee() (*[]EmployeeResponse, error) {
	url := fmt.Sprintf("%s/services/employee/get-all-employee", c.BaseURL)

	// Buat request body dalam bentuk JSON

	// Buat GET request dengan body
	req, err := http.NewRequest("GET", url, nil)
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

	var respBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, fmt.Errorf("invalid response format from employee service: %w", err)
	}

	data, ok := respBody["data"]
	if !ok {
		return nil, fmt.Errorf("field 'data' not found in response")
	}

	// Step 1: Marshal (encode) data map ke JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Step 2: Unmarshal ke struct EmployeeResponse
	var empResp []EmployeeResponse
	err = json.Unmarshal(jsonData, &empResp)
	if err != nil {
		return nil, err
	}
	log.Printf("isi emp response : %v", &empResp)
	return &empResp, nil
}
