package api

import (
	"encoding/json"
	"errors"
)

type BrowserMonitorConfig struct {
	URL                string `json:"url"`
	TimeoutMS          int    `json:"timeout_ms"`
	BodyContains       string `json:"expected_body_contains"`
	Browser            string `json:"browser"`
	ExpectResponseCode int    `json:"expect_response_code"`
}

func ConvertBrowserMonitorConfigToGeneric(config BrowserMonitorConfig) map[string]interface{} {
	return map[string]interface{}{
		"url":                    config.URL,
		"timeout_ms":             config.TimeoutMS,
		"expected_body_contains": config.BodyContains,
		"browser":                config.Browser,
		"expect_response_code":   config.ExpectResponseCode,
	}
}

func validateBrowserMonitorConfig(raw map[string]interface{}) error {
	// Marshall the raw config into a BrowserMonitorConfig
	var config BrowserMonitorConfig
	b, err := json.Marshal(raw)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &config)
	if err != nil {
		return err
	}
	// Validate the config
	if config.URL == "" {
		return errors.New("url is required")
	}
	if config.TimeoutMS == 0 {
		return errors.New("timeout_ms is required")
	}
	if config.ExpectResponseCode == 0 {
		return errors.New("expect_response_code is required")
	}
	validBrowsers := []string{"chrome", "firefox", "webkit"}
	for _, browser := range validBrowsers {
		if config.Browser == browser {
			return nil
		}
	}
	return errors.New("browser must be one of chrome, firefox, webkit")
}

type HTTPMonitorConfig struct {
	URL                string            `json:"url"`
	Method             string            `json:"method"`
	BodyContains       string            `json:"expected_body_contains"`
	TimeoutMS          int               `json:"timeout_ms"`
	ExpectResponseCode int               `json:"expect_response_code"`
	SkipTLSValidation  bool              `json:"skip_tls_validation"`
	RequestBody        string            `json:"request_body"`
	Headers            map[string]string `json:"headers"`
}

func ConvertHTTPMonitorConfigToGeneric(config HTTPMonitorConfig) map[string]interface{} {
	return map[string]interface{}{
		"url":                    config.URL,
		"method":                 config.Method,
		"expected_body_contains": config.BodyContains,
		"timeout_ms":             config.TimeoutMS,
		"expect_response_code":   config.ExpectResponseCode,
		"skip_tls_validation":    config.SkipTLSValidation,
		"request_body":           config.RequestBody,
		"headers":                config.Headers,
	}
}

func validateHTTPMonitorConfig(raw map[string]interface{}) error {
	// Marshall the raw config into a HTTPMonitorConfig
	var config HTTPMonitorConfig
	b, err := json.Marshal(raw)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &config)
	if err != nil {
		return err
	}
	// Validate the config
	if config.URL == "" {
		return errors.New("url is required")
	}
	if config.Method == "" {
		return errors.New("method is required")
	}
	switch config.Method {
	case "GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS", "PATCH":
	default:
		return errors.New("method must be one of GET, POST, PUT, DELETE, HEAD, OPTIONS, PATCH")
	}
	if config.TimeoutMS == 0 {
		return errors.New("timeout_ms is required")
	}
	if config.ExpectResponseCode == 0 {
		return errors.New("expect_response_code is required")
	}
	return nil
}
