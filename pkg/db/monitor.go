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
	ID                  int                    `json:"id"`
	Name                string                 `json:"name"`
	Description         string                 `json:"description"`
	IntervalSeconds     int                    `json:"interval_seconds"`
	Status              string                 `json:"status"`
	LastCheckedAt       *time.Time             `json:"last_checked_at"`
	StatusLastChangedAt time.Time              `json:"status_last_changed_at"`
	CreatedAt           time.Time              `json:"created_at"`
	UpdatedAt           time.Time              `json:"updated_at"`
	MonitorType         string                 `json:"monitor_type"`
	Config              map[string]interface{} `json:"config"`
	Paused              bool                   `json:"paused"`
}

func convertEntMonitorToDBMonitor(entMonitor *ent.Monitor) *BaseMonitor {
	if entMonitor == nil {
		return nil
	}
	return &BaseMonitor{
		ID:                  entMonitor.ID,
		Name:                entMonitor.Name,
		Description:         entMonitor.Description,
		IntervalSeconds:     entMonitor.IntervalSeconds,
		Status:              entMonitor.Status,
		LastCheckedAt:       entMonitor.LastCheckedAt,
		StatusLastChangedAt: entMonitor.StatusLastChangedAt,
		CreatedAt:           entMonitor.CreatedAt,
		UpdatedAt:           entMonitor.UpdatedAt,
		MonitorType:         entMonitor.MonitorType,
		Config:              entMonitor.Config,
		Paused:              entMonitor.Paused,
	}
}

func (db *DatabaseConnection) GetMonitorByID(ctx context.Context, id int) (*BaseMonitor, error) {
	m, err := db.client.Monitor.Query().Where(monitor.ID(id)).First(ctx)
	return convertEntMonitorToDBMonitor(m), err
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

func (db *DatabaseConnection) DeleteMonitor(ctx context.Context, id int) error {
	return db.client.Monitor.DeleteOneID(id).Exec(ctx)
}

func (db *DatabaseConnection) GetMonitorByName(ctx context.Context, name string) (*BaseMonitor, error) {
	m, err := db.client.Monitor.Query().Where(monitor.Name(name)).First(ctx)
	return convertEntMonitorToDBMonitor(m), err
}

type CreateMonitorInput struct {
	Name            string                 `json:"name"`
	IntervalSeconds int                    `json:"interval_seconds"`
	MonitorType     string                 `json:"monitor_type"`
	Config          map[string]interface{} `json:"config"`
	Description     string                 `json:"description"`
}

func (i *CreateMonitorInput) validate() error {
	if i.Name == "" {
		return errors.New("name is required")
	}
	if strings.Contains(i.Name, " ") {
		return errors.New("name cannot contain spaces")
	}
	if i.IntervalSeconds <= 0 {
		return errors.New("interval_seconds must be greater than 0")
	}
	if i.MonitorType == "" {
		return errors.New("monitor_type is required")
	}
	return nil
}

func (db *DatabaseConnection) CreateMonitor(ctx context.Context, input CreateMonitorInput) (*BaseMonitor, error) {
	m, err := db.client.Monitor.Create().
		SetName(input.Name).
		SetIntervalSeconds(input.IntervalSeconds).
		SetMonitorType(input.MonitorType).
		SetConfig(input.Config).
		SetDescription(input.Description).
		SetStatus("initializing").
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
}

func (db *DatabaseConnection) UpdateMonitor(ctx context.Context, id int, input UpdateMonitorInput) (*BaseMonitor, error) {
	update := db.client.Monitor.UpdateOneID(id)
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
	m, err := update.Save(ctx)
	return convertEntMonitorToDBMonitor(m), err
}
