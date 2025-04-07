package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type APIResponse struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       interface{}       `json:"body"`
	TimeMs     int64             `json:"timeMs"`
	Error      string            `json:"error,omitempty"`
}

type RequestConfig struct {
	Method      string            `json:"method"`
	URL         string            `json:"url"`
	Headers     map[string]string `json:"headers"`
	QueryParams map[string]string `json:"queryParams"`
	Body        string            `json:"body"`
}

type APIService struct{}

func NewAPIService() *APIService {
	return &APIService{}
}

func (a *APIService) SendRequest(config RequestConfig) (*APIResponse, error) {
	startTime := time.Now()

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

	var reqBody io.Reader
	if config.Body != "" {
		reqBody = bytes.NewBufferString(config.Body)
	}

	req, err := http.NewRequest(config.Method, reqURL, reqBody)
	if err != nil {
		return &APIResponse{Error: err.Error()}, nil
	}

	for key, value := range config.Headers {
		req.Header.Add(key, value)
	}

	if config.Method != "GET" && config.Method != "HEAD" && req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &APIResponse{Error: err.Error()}, nil
	}
	defer func() {
	    if closeErr := resp.Body.Close(); closeErr != nil {
	        log.Printf("Error closing response body: %v", closeErr)
	    }
	}()

	headers := make(map[string]string)
	for key, values := range resp.Header {
		if len(values) > 0 {
			headers[key] = values[0]
		}
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &APIResponse{Error: err.Error()}, nil
	}

	var bodyInterface interface{}
	if err := json.Unmarshal(respBody, &bodyInterface); err != nil {
		bodyInterface = string(respBody)
	}

	elapsedTime := time.Since(startTime).Milliseconds()

	return &APIResponse{
		StatusCode: resp.StatusCode,
		Headers:    headers,
		Body:       bodyInterface,
		TimeMs:     elapsedTime,
		Error:      "",
	}, nil
}
