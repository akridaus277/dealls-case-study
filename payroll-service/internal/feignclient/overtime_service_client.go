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
type OvertimeDetail struct {
	Date  string  `json:"date"`
	Hours float64 `json:"hours" `
}

type OvertimeSummaryResponse struct {
	Nip            string           `json:"nip"`
	OvertimeDetail []OvertimeDetail `json:"overtimeDetails"`
}

// Interface-like struct
type OvertimeClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

type IOvertimeClient interface {
	GetOvertimeSummary(periodStart, periodEnd time.Time) (*[]OvertimeSummaryResponse, error)
}

// Constructor (bisa dipanggil dari main atau inject lewat service)
func NewOvertimeClient() IOvertimeClient {
	baseURL := os.Getenv("OVERTIME_SERVICE_BASE_URL")
	if baseURL == "" {
		log.Println("Warning: OVERTIME_SERVICE_BASE_URL not set")
	}
	return &OvertimeClient{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *OvertimeClient) GetOvertimeSummary(periodStart, periodEnd time.Time) (*[]OvertimeSummaryResponse, error) {
	url := fmt.Sprintf("%s/services/get-overtime-summary", c.BaseURL)

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
		return nil, fmt.Errorf("failed to call overtime service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("nip not found")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("overtime service error: %s", resp.Status)
	}

	var respBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, fmt.Errorf("invalid response format from overtime service: %w", err)
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
	var ovtSumResp []OvertimeSummaryResponse
	err = json.Unmarshal(jsonData, &ovtSumResp)
	if err != nil {
		return nil, err
	}
	log.Printf("isi ovtSumResp : %v", &ovtSumResp)
	return &ovtSumResp, nil
}
