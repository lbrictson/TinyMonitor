package db

import (
	"context"
	"errors"
	"github.com/lbrictson/TinyMonitor/ent"
	"github.com/lbrictson/TinyMonitor/ent/monitor"
	"strings"
	"time"
)

type BaseMonitor struct {
	Name                string                 `json:"name"`
	Description         string                 `json:"description"`
	CurrentOutageReason string                 `json:"current_outage_reason"`
	IntervalSeconds     int                    `json:"interval_seconds"`
	Status              string                 `json:"status"`
	LastCheckedAt       *time.Time             `json:"last_checked_at"`
	StatusLastChangedAt time.Time              `json:"status_last_changed_at"`
	CreatedAt           time.Time              `json:"created_at"`
	UpdatedAt           time.Time              `json:"updated_at"`
	MonitorType         string                 `json:"monitor_type"`
	Config              map[string]interface{} `json:"config"`
	Paused              bool                   `json:"paused"`
	FailureCount        int                    `json:"failure_count"`
	SuccessCount        int                    `json:"success_count"`
	SuccessThreshold    int                    `json:"success_threshold"`
	FailureThreshold    int                    `json:"failure_threshold"`
}

func convertEntMonitorToDBMonitor(entMonitor *ent.Monitor) *BaseMonitor {
	if entMonitor == nil {
		return nil
	}
	return &BaseMonitor{
		Name:                entMonitor.ID,
		Description:         entMonitor.Description,
		CurrentOutageReason: entMonitor.CurrentDownReason,
		IntervalSeconds:     entMonitor.IntervalSeconds,
		Status:              entMonitor.Status,
		LastCheckedAt:       entMonitor.LastCheckedAt,
		StatusLastChangedAt: entMonitor.StatusLastChangedAt,
		CreatedAt:           entMonitor.CreatedAt,
		UpdatedAt:           entMonitor.UpdatedAt,
		MonitorType:         entMonitor.MonitorType,
		Config:              entMonitor.Config,
		Paused:              entMonitor.Paused,
		FailureCount:        entMonitor.FailureCount,
		SuccessCount:        entMonitor.SuccessCount,
		SuccessThreshold:    entMonitor.SuccessThreshold,
		FailureThreshold:    entMonitor.FailureThreshold,
	}
}

type ListMonitorOptions struct {
	Limit         *int
	Offset        *int
	MonitorType   *string
	MonitorStatus *string
}

func (db *DatabaseConnection) ListMonitors(ctx context.Context, options ListMonitorOptions) ([]*BaseMonitor, error) {
	q := db.client.Monitor.Query()
	if options.Limit != nil {
		q = q.Limit(*options.Limit)
	}
	if options.Offset != nil {
		q = q.Offset(*options.Offset)
	}
	if options.MonitorType != nil {
		q = q.Where(monitor.MonitorType(*options.MonitorType))
	}
	if options.MonitorStatus != nil {
		q = q.Where(monitor.Status(*options.MonitorStatus))
	}
	q.Order(ent.Asc(monitor.FieldStatus))
	monitors, err := q.All(ctx)
	if err != nil {
		return nil, err
	}
	var dbMonitors []*BaseMonitor
	for _, m := range monitors {
		dbMonitors = append(dbMonitors, convertEntMonitorToDBMonitor(m))
	}
	return dbMonitors, nil
}

func (db *DatabaseConnection) DeleteMonitor(ctx context.Context, name string) error {
	return db.client.Monitor.DeleteOneID(name).Exec(ctx)
}

func (db *DatabaseConnection) GetMonitorByName(ctx context.Context, name string) (*BaseMonitor, error) {
	m, err := db.client.Monitor.Get(ctx, name)
	return convertEntMonitorToDBMonitor(m), err
}

type CreateMonitorInput struct {
	Name             string                 `json:"name"`
	IntervalSeconds  int                    `json:"interval_seconds"`
	MonitorType      string                 `json:"monitor_type"`
	Config           map[string]interface{} `json:"config"`
	Description      string                 `json:"description"`
	SuccessThreshold int                    `json:"success_threshold"`
	FailureThreshold int                    `json:"failure_threshold"`
}

func (i *CreateMonitorInput) validate() error {
	if i.Name == "" {
		return errors.New("name is required")
	}
	if strings.Contains(i.Name, " ") {
		return errors.New("name cannot contain spaces")
	}
	// Validate name  only contains letters, numbers, and dashes
	if !strings.ContainsAny(i.Name, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-") {
		return errors.New("name can only contain letters, numbers, and dashes")
	}
	if i.IntervalSeconds <= 0 {
		return errors.New("interval_seconds must be greater than 0")
	}
	if i.MonitorType == "" {
		return errors.New("monitor_type is required")
	}
	if i.FailureThreshold <= 0 {
		i.FailureThreshold = 1
	}
	if i.SuccessThreshold <= 0 {
		i.SuccessThreshold = 1
	}
	return nil
}

func (db *DatabaseConnection) CreateMonitor(ctx context.Context, input CreateMonitorInput) (*BaseMonitor, error) {
	m, err := db.client.Monitor.Create().
		SetID(input.Name).
		SetIntervalSeconds(input.IntervalSeconds).
		SetMonitorType(input.MonitorType).
		SetConfig(input.Config).
		SetDescription(input.Description).
		SetStatus("initializing").
		SetSuccessThreshold(input.SuccessThreshold).
		SetFailureThreshold(input.FailureThreshold).
		Save(ctx)
	return convertEntMonitorToDBMonitor(m), err
}

type UpdateMonitorInput struct {
	IntervalSeconds     *int                   `json:"interval_seconds"`
	Config              map[string]interface{} `json:"config"`
	Status              *string                `json:"status"`
	LastCheckedAt       *time.Time             `json:"last_checked_at"`
	StatusLastChangedAt *time.Time             `json:"status_last_changed_at"`
	Paused              *bool                  `json:"paused"`
	Description         *string                `json:"description"`
	FailureCount        *int                   `json:"failure_count"`
	SuccessCount        *int                   `json:"success_count"`
	SuccessThreshold    *int                   `json:"success_threshold"`
	FailureThreshold    *int                   `json:"failure_threshold"`
	CurrentOutageReason *string                `json:"current_outage_reason"`
}

func (db *DatabaseConnection) UpdateMonitor(ctx context.Context, name string, input UpdateMonitorInput) (*BaseMonitor, error) {
	update := db.client.Monitor.UpdateOneID(name)
	if input.IntervalSeconds != nil {
		update = update.SetIntervalSeconds(*input.IntervalSeconds)
	}
	if input.Config != nil {
		update = update.SetConfig(input.Config)
	}
	if input.Status != nil {
		update = update.SetStatus(*input.Status)
	}
	if input.LastCheckedAt != nil {
		update = update.SetLastCheckedAt(*input.LastCheckedAt)
	}
	if input.StatusLastChangedAt != nil {
		update = update.SetStatusLastChangedAt(*input.StatusLastChangedAt)
	}
	if input.Paused != nil {
		update = update.SetPaused(*input.Paused)
	}
	if input.Description != nil {
		update = update.SetDescription(*input.Description)
	}
	if input.FailureCount != nil {
		update = update.SetFailureCount(*input.FailureCount)
	}
	if input.SuccessCount != nil {
		update = update.SetSuccessCount(*input.SuccessCount)
	}
	if input.SuccessThreshold != nil {
		update = update.SetSuccessThreshold(*input.SuccessThreshold)
	}
	if input.FailureThreshold != nil {
		update = update.SetFailureThreshold(*input.FailureThreshold)
	}
	if input.CurrentOutageReason != nil {
		update = update.SetCurrentDownReason(*input.CurrentOutageReason)
	}
	m, err := update.Save(ctx)
	return convertEntMonitorToDBMonitor(m), err
}
