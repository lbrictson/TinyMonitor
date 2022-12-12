package api

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"github.com/lbrictson/TinyMonitor/pkg/db"
	"github.com/sirupsen/logrus"
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
			doHTTPCheck(m, database, logger)
		}
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
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Warnf("Failed to perform HTTP request for monitor %s: %v", monitor.Name, err)
		now := time.Now()
		status := "Down"
		fCount := monitor.FailureCount + 1
		database.UpdateMonitor(context.Background(), monitor.Name, db.UpdateMonitorInput{
			Status:              &status,
			LastCheckedAt:       &now,
			StatusLastChangedAt: &now,
			Paused:              nil,
			Description:         nil,
			FailureCount:        &fCount,
		})
		return
	}
	defer resp.Body.Close()

}
