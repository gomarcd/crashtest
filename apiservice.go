package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// APIResponse represents the response from an API call
type APIResponse struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       interface{}       `json:"body"`
	TimeMs     int64             `json:"timeMs"`
	Error      string            `json:"error,omitempty"`
}

// RequestConfig represents the configuration for an API request
type RequestConfig struct {
	Method      string            `json:"method"`
	URL         string            `json:"url"`
	Headers     map[string]string `json:"headers"`
	QueryParams map[string]string `json:"queryParams"`
	Body        string            `json:"body"`
}

// APIService handles API calls
type APIService struct{}

// NewAPIService creates a new API service
func NewAPIService() *APIService {
	return &APIService{}
}

// SendRequest sends an API request and returns the response
func (a *APIService) SendRequest(config RequestConfig) (*APIResponse, error) {
	startTime := time.Now()

	// Prepare the request URL with query parameters
	reqURL := config.URL
	if len(config.QueryParams) > 0 {
		req, err := http.NewRequest(config.Method, config.URL, nil)
		if err != nil {
			return &APIResponse{Error: err.Error()}, nil
		}

		q := req.URL.Query()
		for key, value := range config.QueryParams {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
		reqURL = req.URL.String()
	}

	// Create request body if necessary
	var reqBody io.Reader
	if config.Body != "" {
		reqBody = bytes.NewBufferString(config.Body)
	}

	// Create the request
	req, err := http.NewRequest(config.Method, reqURL, reqBody)
	if err != nil {
		return &APIResponse{Error: err.Error()}, nil
	}

	// Add headers
	for key, value := range config.Headers {
		req.Header.Add(key, value)
	}

	// Set default content type if not specified
	if config.Method != "GET" && config.Method != "HEAD" && req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &APIResponse{Error: err.Error()}, nil
	}
	defer resp.Body.Close()

	// Extract headers
	headers := make(map[string]string)
	for key, values := range resp.Header {
		if len(values) > 0 {
			headers[key] = values[0]
		}
	}

	// Read body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &APIResponse{Error: err.Error()}, nil
	}

	// Try to parse as JSON, otherwise return as string
	var bodyInterface interface{}
	if err := json.Unmarshal(respBody, &bodyInterface); err != nil {
		bodyInterface = string(respBody)
	}

	elapsedTime := time.Since(startTime).Milliseconds()

	// Return the response
	return &APIResponse{
		StatusCode: resp.StatusCode,
		Headers:    headers,
		Body:       bodyInterface,
		TimeMs:     elapsedTime,
		Error:      "",
	}, nil
}
