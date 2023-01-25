package api

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/lbrictson/TinyMonitor/pkg/db"
	"github.com/playwright-community/playwright-go"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// performMonitoringChecks performs an HTTP check on the given monitor, endlessly
func performMonitoringChecks(name string, database *db.DatabaseConnection, logger *logrus.Logger) {
	// Get the monitor from the database
	ctx := context.Background()
	m, err := database.GetMonitorByName(ctx, name)
	if err != nil {
		logger.Warnf("Failed to get monitor %s from database: %v, exiting monitoring thread", name, err)
		return
	}
	// Create a ticket with the interval of the monitor
	interval := m.IntervalSeconds
	t := time.NewTicker(time.Duration(interval) * time.Second)
	for {
		select {
		case <-t.C:
			// Fetch the monitor from the database to see if anything has changed
			m, err := database.GetMonitorByName(ctx, name)
			if err != nil {
				logger.Warnf("Failed to get monitor %s from database: %v, exiting monitoring thread", name, err)
				return
			}
			// If the monitor is paused, skip this check
			if m.Paused {
				continue
			}
			if m.IntervalSeconds != interval {
				// If the interval has changed, update the ticker
				interval = m.IntervalSeconds
				t.Stop()
				t = time.NewTicker(time.Duration(interval) * time.Second)
			}
			// Perform the check
			logger.Infof("Performing check on monitor %s", name)
			switch strings.ToLower(m.MonitorType) {
			case "http":
				// Perform an HTTP check
				doHTTPCheck(m, database, logger)
			case "browser":
				// Perform a browser check
				doBrowserCheck(m, database, logger)
			}
		}
	}
}

func processMonitorCheckDownResult(failureExplanation string, monitor *db.BaseMonitor, database *db.DatabaseConnection, logger *logrus.Logger) {
	// Increment the failure count
	now := time.Now()
	fCount := monitor.FailureCount + 1
	// If the monitor is already down, do nothing
	if monitor.Status == "Down" {
		_, err := database.UpdateMonitor(context.Background(), monitor.Name, db.UpdateMonitorInput{
			LastCheckedAt:       &now,
			StatusLastChangedAt: nil,
			Paused:              nil,
			Description:         nil,
			FailureCount:        &fCount,
			CurrentOutageReason: &failureExplanation,
		})
		if err != nil {
			logger.Warnf("Failed to update monitor %s: %v", monitor.Name, err)
		}
		return
	} else {
		down := "Down"
		changes := db.UpdateMonitorInput{
			Status:              nil,
			LastCheckedAt:       &now,
			StatusLastChangedAt: nil,
			Paused:              nil,
			Description:         nil,
			FailureCount:        &fCount,
		}
		// If the monitor is up, update the database
		if fCount >= monitor.FailureThreshold {
			// Only update to down if the failure threshold has been reached or breached
			changes.Status = &down
			changes.StatusLastChangedAt = &now
			changes.CurrentOutageReason = &failureExplanation
		}
		_, err := database.UpdateMonitor(context.Background(), monitor.Name, changes)
		if err != nil {
			logger.Warnf("Failed to update monitor %s: %v", monitor.Name, err)
		}
	}
}

func processMonitorOkResult(monitor *db.BaseMonitor, database *db.DatabaseConnection, logger *logrus.Logger) {
	// Reset the failure count
	now := time.Now()
	zero := 0
	successCount := monitor.SuccessCount + 1
	changes := db.UpdateMonitorInput{
		LastCheckedAt:       &now,
		StatusLastChangedAt: nil,
		Paused:              nil,
		Description:         nil,
		FailureCount:        &zero,
		CurrentOutageReason: nil,
		SuccessCount:        &successCount,
	}
	// If the monitor is down or Initializing, update the database
	if monitor.Status == "Down" || monitor.Status == "Initializing" {
		// Only update if success threshold has been reached
		if successCount >= monitor.SuccessThreshold {
			up := "Up"
			changes.Status = &up
			changes.StatusLastChangedAt = &now
		}
	}
	_, err := database.UpdateMonitor(context.Background(), monitor.Name, changes)
	if err != nil {
		logger.Warnf("Failed to update monitor %s: %v", monitor.Name, err)
	}
}

func doHTTPCheck(monitor *db.BaseMonitor, database *db.DatabaseConnection, logger *logrus.Logger) {
	specificConfig := HTTPMonitorConfig{}
	b, err := json.Marshal(monitor.Config)
	if err != nil {
		logger.Warnf("Failed to marshal monitor %s config: %v", monitor.Name, err)
		return
	}
	err = json.Unmarshal(b, &specificConfig)
	if err != nil {
		logger.Warnf("Failed to unmarshal monitor %s config: %v", monitor.Name, err)
		return
	}
	req, err := http.NewRequest(strings.ToUpper(specificConfig.Method), specificConfig.URL, strings.NewReader(specificConfig.RequestBody))
	if err != nil {
		logger.Warnf("Failed to create request for monitor %s: %v", monitor.Name, err)
		processMonitorCheckDownResult(err.Error(), monitor, database, logger)
		return
	}
	for k, v := range specificConfig.Headers {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	if specificConfig.SkipTLSValidation {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	client.Timeout = time.Duration(specificConfig.TimeoutMS) * time.Millisecond
	start := time.Now()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Warnf("Failed to perform HTTP request for monitor %s: %v", monitor.Name, err)
		processMonitorCheckDownResult(err.Error(), monitor, database, logger)
		return
	}
	defer resp.Body.Close()
	// Make sure request didn't breach timeout expectation
	if time.Since(start) > time.Duration(specificConfig.TimeoutMS)*time.Millisecond {
		logger.Warnf("HTTP request for monitor %s took longer than expected", monitor.Name)
		processMonitorCheckDownResult(fmt.Sprintf("Request breached timeout value of %v", specificConfig.TimeoutMS), monitor, database, logger)
		return
	}
	// Check the response code
	if resp.StatusCode != specificConfig.ExpectResponseCode {
		logger.Warnf("HTTP request for monitor %s returned unexpected response code", monitor.Name)
		processMonitorCheckDownResult(fmt.Sprintf("Expected response code %v, got %v", specificConfig.ExpectResponseCode, resp.StatusCode), monitor, database, logger)
		return
	}
	// Check the body if need be
	if specificConfig.BodyContains != "" {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.Warnf("Failed to read response body for monitor %s: %v", monitor.Name, err)
			processMonitorCheckDownResult(err.Error(), monitor, database, logger)
			return
		}
		if !strings.Contains(string(body), specificConfig.BodyContains) {
			logger.Warnf("HTTP request for monitor %s returned unexpected response body", monitor.Name)
			processMonitorCheckDownResult(fmt.Sprintf("Expected body to contain %v, got %v", specificConfig.BodyContains, string(body)), monitor, database, logger)
			return
		}
	}
	// If we got here, the monitor is up
	processMonitorOkResult(monitor, database, logger)
	return
}

func doBrowserCheck(monitor *db.BaseMonitor, database *db.DatabaseConnection, logger *logrus.Logger) {
	specificConfig := BrowserMonitorConfig{}
	b, err := json.Marshal(monitor.Config)
	if err != nil {
		logger.Warnf("Failed to marshal monitor %s config: %v", monitor.Name, err)
		return
	}
	err = json.Unmarshal(b, &specificConfig)
	if err != nil {
		logger.Warnf("Failed to unmarshal monitor %s config: %v", monitor.Name, err)
		return
	}
	// Create a new browser instance
	pw, launchErr := playwright.Run()
	if launchErr != nil {
		logger.Warnf("Failed to launch browser for monitor %s: %v", monitor.Name, launchErr)
		processMonitorCheckDownResult(launchErr.Error(), monitor, database, logger)
		return
	}
	defer pw.Stop()
	type pwBrowser struct {
		browser playwright.Browser
	}
	browser := pwBrowser{}
	switch strings.ToLower(specificConfig.Browser) {
	case "chrome":
		browser.browser, err = pw.Chromium.Launch()
		if err != nil {
			logger.Warnf("Failed to launch Chromium for monitor %s: %v", monitor.Name, err)
			processMonitorCheckDownResult(err.Error(), monitor, database, logger)
			return
		}
	case "firefox":
		browser.browser, err = pw.Firefox.Launch()
		if err != nil {
			logger.Warnf("Failed to launch Firefox for monitor %s: %v", monitor.Name, err)
			processMonitorCheckDownResult(err.Error(), monitor, database, logger)
			return
		}
	case "webkit":
		browser.browser, err = pw.WebKit.Launch()
		if err != nil {
			logger.Warnf("Failed to launch WebKit for monitor %s: %v", monitor.Name, err)
			processMonitorCheckDownResult(err.Error(), monitor, database, logger)
			return
		}
	default:
		logger.Warnf("Failed to launch browser for monitor %s: %v", monitor.Name, specificConfig.Browser)
		processMonitorCheckDownResult(fmt.Sprintf("Failed to launch browser for monitor %s: %v", monitor.Name, specificConfig.Browser), monitor, database, logger)
		return
	}
	defer browser.browser.Close()
	page, err := browser.browser.NewPage()
	if err != nil {
		logger.Warnf("Failed to create page for monitor %s: %v", monitor.Name, err)
		processMonitorCheckDownResult(err.Error(), monitor, database, logger)
		return
	}
	start := time.Now()
	resp, err := page.Goto(specificConfig.URL, playwright.PageGotoOptions{WaitUntil: playwright.WaitUntilStateNetworkidle})
	if err != nil {
		logger.Warnf("Failed to navigate to URL for monitor %s: %v", monitor.Name, err)
		processMonitorCheckDownResult(fmt.Sprintf("Failed to navigate to URL for monitor %s: %v", monitor.Name, err), monitor, database, logger)
		return
	}
	if resp.Status() != specificConfig.ExpectResponseCode {
		logger.Warnf("Response code for monitor %s was unexpected", monitor.Name)
		processMonitorCheckDownResult(fmt.Sprintf("Expected response code %v, got %v", specificConfig.ExpectResponseCode, resp.Status()), monitor, database, logger)
		return
	}
	if time.Since(start) > time.Duration(specificConfig.TimeoutMS)*time.Millisecond {
		logger.Warnf("Request for monitor %s took longer than expected", monitor.Name)
		processMonitorCheckDownResult(fmt.Sprintf("Request breached timeout value of %v", specificConfig.TimeoutMS), monitor, database, logger)
		return
	}
	if specificConfig.BodyContains != "" {
		body, err := page.Content()
		if err != nil {
			logger.Warnf("Failed to get page content for monitor %s: %v", monitor.Name, err)
			processMonitorCheckDownResult(err.Error(), monitor, database, logger)
			return
		}
		if !strings.Contains(body, specificConfig.BodyContains) {
			logger.Warnf("Page content for monitor %s was unexpected", monitor.Name)
			processMonitorCheckDownResult(fmt.Sprintf("Expected body to contain %v, got %v", specificConfig.BodyContains, body), monitor, database, logger)
			return
		}
	}
	// If we got here, the monitor is up
	processMonitorOkResult(monitor, database, logger)
}
