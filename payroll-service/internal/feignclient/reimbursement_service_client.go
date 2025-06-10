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
type ReimbursementDetail struct {
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Date        string  `json:"date"`
}

type ReimbursementSummaryResponse struct {
	Nip                 string                `json:"nip"`
	ReimbursementDetail []ReimbursementDetail `json:"reimbursementDetails"`
}

// Interface-like struct
type ReimbursementClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

type IReimbursementClient interface {
	GetReimbursementSummary(periodStart, periodEnd time.Time) (*[]ReimbursementSummaryResponse, error)
}

// Constructor (bisa dipanggil dari main atau inject lewat service)
func NewReimbursementClient() IReimbursementClient {
	baseURL := os.Getenv("REIMBURSEMENT_SERVICE_BASE_URL")
	if baseURL == "" {
		log.Println("Warning: REIMBURSEMENT_SERVICE_BASE_URL not set")
	}
	return &ReimbursementClient{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *ReimbursementClient) GetReimbursementSummary(periodStart, periodEnd time.Time) (*[]ReimbursementSummaryResponse, error) {
	url := fmt.Sprintf("%s/services/get-reimbursement-summary", c.BaseURL)

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
	var rmbSumResp []ReimbursementSummaryResponse
	err = json.Unmarshal(jsonData, &rmbSumResp)
	if err != nil {
		return nil, err
	}
	log.Printf("isi rmbSumResp : %v", &rmbSumResp)
	return &rmbSumResp, nil
}
