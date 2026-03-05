package aws

import (
	"encoding/json"
	"net/http"
)

// Client represents an AWS integration client
type Client struct {
	APIKey string
	Endpoint string
}

// Data represents the structure of the data to be sent to AWS
type Data struct {
	CollectedData interface{} `json:"collected_data"`
}

// NewClient creates a new AWS client
func NewClient(apiKey, endpoint string) *Client {
	return &Client{APIKey: apiKey, Endpoint: endpoint}
}

// SendData sends collected data to AWS using JSON format
func (c *Client) SendData(data interface{}) error {
	payload := Data{CollectedData: data}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.Endpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", c.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send data: %s", resp.Status)
	}

	return nil
}