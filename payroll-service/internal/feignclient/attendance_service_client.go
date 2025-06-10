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
type AttendancePeriodPayrollResponse struct {
	StartDate         time.Time `json:"startDate"`
	EndDate           time.Time `json:"endDate"`
	PayrollPeriodCode string    `json:"payrollPeriodCode"`
}

type AttendanceDetail struct {
	Date string `json:"date"`
}

type AttendanceSummaryResponse struct {
	Nip              string             `json:"nip"`
	AttendanceDetail []AttendanceDetail `json:"attendanceDetails"`
}

// Interface-like struct
type AttendanceClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

type IAttendanceClient interface {
	GetAttendancePeriodPayroll(payrollCode string) (*AttendancePeriodPayrollResponse, error)
	GetAttendanceSummary(periodStart, periodEnd time.Time) (*[]AttendanceSummaryResponse, error)
}

// Constructor (bisa dipanggil dari main atau inject lewat service)
func NewAttendanceClient() IAttendanceClient {
	baseURL := os.Getenv("ATTENDANCE_SERVICE_BASE_URL")
	if baseURL == "" {
		log.Println("Warning: ATTENDANCE_SERVICE_BASE_URL not set")
	}
	return &AttendanceClient{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *AttendanceClient) GetAttendancePeriodPayroll(payrollCode string) (*AttendancePeriodPayrollResponse, error) {
	url := fmt.Sprintf("%s/services/get-attendance-period", c.BaseURL)

	// Buat request body dalam bentuk JSON
	reqBody := map[string]any{
		"payrollPeriodCode": payrollCode,
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
	var empResp AttendancePeriodPayrollResponse
	err = json.Unmarshal(jsonData, &empResp)
	if err != nil {
		return nil, err
	}
	log.Printf("isi emp response : %v", &empResp)
	return &empResp, nil
}

func (c *AttendanceClient) GetAttendanceSummary(periodStart, periodEnd time.Time) (*[]AttendanceSummaryResponse, error) {
	url := fmt.Sprintf("%s/services/get-attendance-summary", c.BaseURL)

	// Buat request body dalam bentuk JSON
	reqBody := map[string]any{
		"periodStart": periodStart,
		"periodEnd":   periodEnd,
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
	var atdSumResp []AttendanceSummaryResponse
	err = json.Unmarshal(jsonData, &atdSumResp)
	if err != nil {
		return nil, err
	}
	log.Printf("isi atdSumResp : %v", &atdSumResp)
	return &atdSumResp, nil
}
