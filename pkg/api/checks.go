package api

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/lbrictson/TinyMonitor/pkg/db"
	"github.com/lbrictson/TinyMonitor/pkg/sink"
	"github.com/playwright-community/playwright-go"
	probing "github.com/prometheus-community/pro-bing"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"runtime"
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
				if !allowBrowserMonitors {
					logger.Warnf("Browser monitoring is currently disabled, skipping check on monitor %s", name)
					continue
				}
				// Perform a browser check
				doBrowserCheck(m, database, logger)
			case "ping":
				doPingCheck(m, database, logger)
			}
		}
	}
}

func processMonitorCheckDownResult(failureExplanation string, monitor *db.BaseMonitor, database *db.DatabaseConnection, logger *logrus.Logger, latencyMS float64) {
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
		for k, v := range metrics {
			senderErr := v.SendMetric([]sink.SendMetricInput{
				{
					Tags:        map[string]string{"monitor": monitor.Name, "monitor-type": monitor.MonitorType},
					MetricName:  "status",
					MetricValue: 0,
					MetricUnit:  "None",
					MetricTime:  time.Now(),
					MonitorName: monitor.Name,
				},
				{
					Tags:        map[string]string{"monitor": monitor.Name, "monitor-type": monitor.MonitorType},
					MetricName:  "latency-ms",
					MetricValue: latencyMS,
					MetricUnit:  "None",
					MetricTime:  time.Now(),
					MonitorName: monitor.Name,
				},
			})
			if senderErr != nil {
				logger.Warnf("Failed to send metric to %v: %v", k, senderErr)
			}
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
			fireAlertDown(monitor, failureExplanation, database, logger)
		}
		_, err := database.UpdateMonitor(context.Background(), monitor.Name, changes)
		if err != nil {
			logger.Warnf("Failed to update monitor %s: %v", monitor.Name, err)
		}
	}
	for k, v := range metrics {
		senderErr := v.SendMetric([]sink.SendMetricInput{
			{
				Tags:        map[string]string{"monitor": monitor.Name, "monitor-type": monitor.MonitorType},
				MetricName:  "status",
				MetricValue: 0,
				MetricUnit:  "None",
				MetricTime:  time.Now(),
				MonitorName: monitor.Name,
			},
			{
				Tags:        map[string]string{"monitor": monitor.Name, "monitor-type": monitor.MonitorType},
				MetricName:  "latency-ms",
				MetricValue: latencyMS,
				MetricUnit:  "None",
				MetricTime:  time.Now(),
				MonitorName: monitor.Name,
			},
		})
		if senderErr != nil {
			logger.Warnf("Failed to send metric to %v: %v", k, senderErr)
		}
	}
}

func processMonitorOkResult(monitor *db.BaseMonitor, database *db.DatabaseConnection, logger *logrus.Logger, latencyMS float64) {
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
	if strings.ToLower(monitor.Status) == "down" || strings.ToLower(monitor.Status) == "initializing" {
		// Only update if success threshold has been reached
		if successCount >= monitor.SuccessThreshold {
			up := "Up"
			changes.Status = &up
			changes.StatusLastChangedAt = &now
			// Fire up alert if the monitor was originally down
			if strings.ToLower(monitor.Status) == "down" {
				fireAlertUp(monitor, database, logger)
			}
		}
	}
	_, err := database.UpdateMonitor(context.Background(), monitor.Name, changes)
	if err != nil {
		logger.Warnf("Failed to update monitor %s: %v", monitor.Name, err)
	}
	for k, v := range metrics {
		senderErr := v.SendMetric([]sink.SendMetricInput{
			{
				Tags:        map[string]string{"monitor": monitor.Name, "monitor-type": monitor.MonitorType},
				MetricName:  "status",
				MetricValue: 1.0,
				MetricUnit:  "None",
				MetricTime:  time.Now(),
				MonitorName: monitor.Name,
			},
			{
				Tags:        map[string]string{"monitor": monitor.Name, "monitor-type": monitor.MonitorType},
				MetricName:  "latency-ms",
				MetricValue: latencyMS,
				MetricUnit:  "None",
				MetricTime:  time.Now(),
				MonitorName: monitor.Name,
			},
		})
		if senderErr != nil {
			logger.Warnf("Failed to send metric to %v: %v", k, senderErr)
		}
	}
}

func doPingCheck(monitor *db.BaseMonitor, database *db.DatabaseConnection, logger *logrus.Logger) {
	specificConfig := PingMonitorConfig{}
	b, err := json.Marshal(monitor.Config)
	if err != nil {
		logger.Warnf("Failed to marshal monitor config: %v", err)
		return
	}
	err = json.Unmarshal(b, &specificConfig)
	if err != nil {
		logger.Warnf("Failed to unmarshal monitor config: %v", err)
		return
	}
	ping, err := probing.NewPinger(specificConfig.Host)
	if err != nil {
		logger.Warnf("Failed to create ping process for monitor %v: %v", monitor.Name, err)
		return
	}
	// If windows set privileged to true
	if runtime.GOOS == "windows" {
		ping.SetPrivileged(true)
	}
	if runtime.GOOS == "linux" {
		ping.SetPrivileged(false)
	}
	ping.Count = 1
	ping.Timeout = time.Duration(specificConfig.TimeoutMS) * time.Millisecond
	err = ping.Run()
	if err != nil {
		logger.Warnf("Failed to run ping process for monitor %v: %v", monitor.Name, err)
		return
	}
	stats := ping.Statistics()
	if stats.PacketsRecv == 0 {
		logger.Warnf("Failed to receive ping response for monitor %v", monitor.Name)
		processMonitorCheckDownResult("No packets received", monitor, database, logger, 0)
		return
	}
	latencyMS := float64(stats.AvgRtt.Milliseconds())
	if latencyMS > float64(specificConfig.TimeoutMS) {
		processMonitorCheckDownResult("Timeout exceeded", monitor, database, logger, latencyMS)
		return
	}
	processMonitorOkResult(monitor, database, logger, latencyMS)
}

func doHTTPCheck(monitor *db.BaseMonitor, database *db.DatabaseConnection, logger *logrus.Logger) {
	ctx := context.TODO()
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
	req, err := http.NewRequest(strings.ToUpper(specificConfig.Method), specificConfig.URL, strings.NewReader(injectSecretsIntoContent(ctx, database, specificConfig.RequestBody)))
	if err != nil {
		logger.Warnf("Failed to create request for monitor %s: %v", monitor.Name, err)
		processMonitorCheckDownResult(err.Error(), monitor, database, logger, 0)
		return
	}
	for k, v := range specificConfig.Headers {
		req.Header.Set(injectSecretsIntoContent(ctx, database, k), injectSecretsIntoContent(ctx, database, v))
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
		processMonitorCheckDownResult(err.Error(), monitor, database, logger, 0)
		return
	}
	latencyMSFloat64 := float64(time.Since(start).Milliseconds())
	defer resp.Body.Close()
	// Make sure request didn't breach timeout expectation
	if time.Since(start) > time.Duration(specificConfig.TimeoutMS)*time.Millisecond {
		logger.Warnf("HTTP request for monitor %s took longer than expected (%v vs limit of %v", monitor.Name, time.Since(start), time.Duration(specificConfig.TimeoutMS)*time.Millisecond)
		processMonitorCheckDownResult(fmt.Sprintf("Request breached timeout value of %vms", specificConfig.TimeoutMS), monitor, database, logger, latencyMSFloat64)
		return
	}
	// Check the response code
	if resp.StatusCode != specificConfig.ExpectResponseCode {
		logger.Warnf("HTTP request for monitor %s returned unexpected response code", monitor.Name)
		processMonitorCheckDownResult(fmt.Sprintf("Expected response code %v, got %v", specificConfig.ExpectResponseCode, resp.StatusCode), monitor, database, logger, latencyMSFloat64)
		return
	}
	// Check the body if need be
	if specificConfig.BodyContains != "" {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.Warnf("Failed to read response body for monitor %s: %v", monitor.Name, err)
			processMonitorCheckDownResult(err.Error(), monitor, database, logger, latencyMSFloat64)
			return
		}
		if !strings.Contains(string(body), specificConfig.BodyContains) {
			logger.Warnf("HTTP request for monitor %s returned unexpected response body", monitor.Name)
			processMonitorCheckDownResult(fmt.Sprintf("Expected body to contain %v, got %v", specificConfig.BodyContains, string(body)), monitor, database, logger, latencyMSFloat64)
			return
		}
	}
	// If we got here, the monitor is up
	processMonitorOkResult(monitor, database, logger, latencyMSFloat64)
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
		processMonitorCheckDownResult(launchErr.Error(), monitor, database, logger, 0)
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
			processMonitorCheckDownResult(err.Error(), monitor, database, logger, 0)
			return
		}
	case "firefox":
		browser.browser, err = pw.Firefox.Launch()
		if err != nil {
			logger.Warnf("Failed to launch Firefox for monitor %s: %v", monitor.Name, err)
			processMonitorCheckDownResult(err.Error(), monitor, database, logger, 0)
			return
		}
	case "webkit":
		browser.browser, err = pw.WebKit.Launch()
		if err != nil {
			logger.Warnf("Failed to launch WebKit for monitor %s: %v", monitor.Name, err)
			processMonitorCheckDownResult(err.Error(), monitor, database, logger, 0)
			return
		}
	default:
		logger.Warnf("Failed to launch browser for monitor %s: %v", monitor.Name, specificConfig.Browser)
		processMonitorCheckDownResult(fmt.Sprintf("Failed to launch browser for monitor %s: %v", monitor.Name, specificConfig.Browser), monitor, database, logger, 0)
		return
	}
	defer browser.browser.Close()
	page, err := browser.browser.NewPage()
	if err != nil {
		logger.Warnf("Failed to create page for monitor %s: %v", monitor.Name, err)
		processMonitorCheckDownResult(err.Error(), monitor, database, logger, 0)
		return
	}
	start := time.Now()
	resp, err := page.Goto(specificConfig.URL, playwright.PageGotoOptions{WaitUntil: playwright.WaitUntilStateNetworkidle})
	latencyMSFloat64 := float64(time.Since(start).Milliseconds())
	if err != nil {
		logger.Warnf("Failed to navigate to URL for monitor %s: %v", monitor.Name, err)
		processMonitorCheckDownResult(fmt.Sprintf("Failed to navigate to URL for monitor %s: %v", monitor.Name, err), monitor, database, logger, latencyMSFloat64)
		return
	}
	if resp.Status() != specificConfig.ExpectResponseCode {
		logger.Warnf("Response code for monitor %s was unexpected", monitor.Name)
		processMonitorCheckDownResult(fmt.Sprintf("Expected response code %v, got %v", specificConfig.ExpectResponseCode, resp.Status()), monitor, database, logger, latencyMSFloat64)
		return
	}
	if time.Since(start) > time.Duration(specificConfig.TimeoutMS)*time.Millisecond {
		logger.Warnf("Request for monitor %s took longer than expected (%v vs limit %v)", monitor.Name, time.Since(start), time.Duration(specificConfig.TimeoutMS)*time.Millisecond)
		processMonitorCheckDownResult(fmt.Sprintf("Request breached timeout value of %vms", specificConfig.TimeoutMS), monitor, database, logger, latencyMSFloat64)
		return
	}
	if specificConfig.BodyContains != "" {
		body, err := page.Content()
		if err != nil {
			logger.Warnf("Failed to get page content for monitor %s: %v", monitor.Name, err)
			processMonitorCheckDownResult(err.Error(), monitor, database, logger, latencyMSFloat64)
			return
		}
		if !strings.Contains(body, specificConfig.BodyContains) {
			logger.Warnf("Page content for monitor %s was unexpected", monitor.Name)
			processMonitorCheckDownResult(fmt.Sprintf("Expected body to contain %v, got %v", specificConfig.BodyContains, body), monitor, database, logger, latencyMSFloat64)
			return
		}
	}
	// If we got here, the monitor is up
	processMonitorOkResult(monitor, database, logger, latencyMSFloat64)
}
