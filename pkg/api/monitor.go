package api

import (
	"errors"
	"github.com/dustin/go-humanize"
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/TinyMonitor/pkg/db"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type MonitorModel struct {
	Name                      string                 `json:"name"`
	Description               string                 `json:"description"`
	CurrentOutageReason       string                 `json:"current_outage_reason"`
	IntervalSeconds           int                    `json:"interval_seconds"`
	Status                    string                 `json:"status"`
	LastCheckedAt             *time.Time             `json:"last_checked_at"`
	LastCheckedFriendly       string                 `json:"last_checked_friendly"`
	StatusLastChangedAt       time.Time              `json:"status_last_changed_at"`
	StatusLastChangedFriendly string                 `json:"status_last_changed_friendly"`
	CreatedAt                 time.Time              `json:"created_at"`
	UpdatedAt                 time.Time              `json:"updated_at"`
	MonitorType               string                 `json:"monitor_type"`
	Config                    map[string]interface{} `json:"config"`
	Paused                    bool                   `json:"paused"`
	FailureCount              int                    `json:"failure_count"`
	SuccessThreshold          int                    `json:"success_threshold"`
	FailureThreshold          int                    `json:"failure_threshold"`
}

func convertDBMonitorToAPIMonitor(dbMonitor *db.BaseMonitor) *MonitorModel {
	if dbMonitor == nil {
		return nil
	}
	friendlyLastChecked := "Never"
	if dbMonitor.LastCheckedAt != nil {
		friendlyLastChecked = humanize.Time(*dbMonitor.LastCheckedAt)
	}
	return &MonitorModel{
		Name:                      dbMonitor.Name,
		Description:               dbMonitor.Description,
		CurrentOutageReason:       dbMonitor.CurrentOutageReason,
		IntervalSeconds:           dbMonitor.IntervalSeconds,
		Status:                    dbMonitor.Status,
		LastCheckedAt:             dbMonitor.LastCheckedAt,
		LastCheckedFriendly:       friendlyLastChecked,
		StatusLastChangedAt:       dbMonitor.StatusLastChangedAt,
		StatusLastChangedFriendly: humanize.Time(dbMonitor.StatusLastChangedAt),
		CreatedAt:                 dbMonitor.CreatedAt,
		UpdatedAt:                 dbMonitor.UpdatedAt,
		MonitorType:               dbMonitor.MonitorType,
		Config:                    dbMonitor.Config,
		Paused:                    dbMonitor.Paused,
		FailureCount:              dbMonitor.FailureCount,
		SuccessThreshold:          dbMonitor.SuccessThreshold,
		FailureThreshold:          dbMonitor.FailureThreshold,
	}
}

func (s *Server) listMonitors(c echo.Context) error {
	ops := db.ListMonitorOptions{
		Limit:         nil,
		Offset:        nil,
		MonitorType:   nil,
		MonitorStatus: nil,
	}
	if c.QueryParam("limit") != "" {
		limit := c.QueryParam("limit")
		// Convert limit to int
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			return s.returnErrorResponse(c, http.StatusBadRequest, err)
		}
		ops.Limit = &limitInt
	}
	if c.QueryParam("offset") != "" {
		offset := c.QueryParam("offset")
		// Convert offset to int
		offsetInt, err := strconv.Atoi(offset)
		if err != nil {
			return s.returnErrorResponse(c, http.StatusBadRequest, err)
		}
		ops.Offset = &offsetInt
	}
	if c.QueryParam("type") != "" {
		monitorType := c.QueryParam("type")
		ops.MonitorType = &monitorType
	}
	if c.QueryParam("status") != "" {
		monitorStatus := c.QueryParam("status")
		ops.MonitorStatus = &monitorStatus
	}
	monitors, err := s.dbConnection.ListMonitors(c.Request().Context(), ops)
	if err != nil {
		return s.returnErrorResponse(c, http.StatusInternalServerError, err)
	}
	var apiMonitors []*MonitorModel
	for _, m := range monitors {
		apiMonitors = append(apiMonitors, convertDBMonitorToAPIMonitor(m))
	}
	return s.returnSuccessResponse(c, http.StatusOK, apiMonitors)
}

func (s *Server) getMonitor(c echo.Context) error {
	monitor, err := s.dbConnection.GetMonitorByName(c.Request().Context(), c.Param("id"))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return s.returnErrorResponse(c, http.StatusNotFound, errors.New("monitor not found"))
		}
		return s.returnErrorResponse(c, http.StatusInternalServerError, err)
	}
	return s.returnSuccessResponse(c, http.StatusOK, convertDBMonitorToAPIMonitor(monitor))
}

type CreateMonitorInput struct {
	Name            string                 `json:"name"`
	Description     string                 `json:"description"`
	IntervalSeconds int                    `json:"interval_seconds"`
	MonitorType     string                 `json:"monitor_type"`
	Config          map[string]interface{} `json:"config"`
}

func (c *CreateMonitorInput) Validate() error {
	if c.Name == "" {
		return errors.New("name is required")
	}
	if c.IntervalSeconds == 0 {
		c.IntervalSeconds = 60
	}
	if c.MonitorType == "" {
		return errors.New("monitor_type is required")
	}
	if c.Config == nil {
		return errors.New("config is required")
	}
	return nil
}

func (s *Server) createMonitor(c echo.Context) error {
	var input CreateMonitorInput
	err := c.Bind(&input)
	if err != nil {
		return s.returnErrorResponse(c, http.StatusBadRequest, err)
	}
	err = input.Validate()
	if err != nil {
		return s.returnErrorResponse(c, http.StatusBadRequest, err)
	}
	switch input.MonitorType {
	case "http":
		// Validate config
		err = validateHTTPMonitorConfig(input.Config)
		if err != nil {
			return s.returnErrorResponse(c, http.StatusBadRequest, err)
		}
	default:
		return s.returnErrorResponse(c, http.StatusBadRequest, errors.New("invalid monitor_type: expected one of [http]"))
	}
	monitor, err := s.dbConnection.CreateMonitor(c.Request().Context(), db.CreateMonitorInput{
		Name:            input.Name,
		IntervalSeconds: input.IntervalSeconds,
		MonitorType:     input.MonitorType,
		Config:          input.Config,
		Description:     input.Description,
	})
	if err != nil {
		return s.returnErrorResponse(c, http.StatusInternalServerError, err)
	}
	return s.returnSuccessResponse(c, http.StatusCreated, convertDBMonitorToAPIMonitor(monitor))
}

type UpdateMonitorInput struct {
	IntervalSeconds *int                   `json:"interval_seconds"`
	Paused          *bool                  `json:"paused"`
	Config          map[string]interface{} `json:"config"`
	Description     *string                `json:"description"`
}

func (c *UpdateMonitorInput) Validate() error {
	if c.Config == nil {
		return errors.New("config is required")
	}
	return nil
}

func (s *Server) updateMonitor(c echo.Context) error {
	var input UpdateMonitorInput
	err := c.Bind(&input)
	if err != nil {
		return s.returnErrorResponse(c, http.StatusBadRequest, err)
	}
	err = input.Validate()
	if err != nil {
		return s.returnErrorResponse(c, http.StatusBadRequest, err)
	}
	// First get the monitor from the database
	monitor, err := s.dbConnection.GetMonitorByName(c.Request().Context(), c.Param("id"))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return s.returnErrorResponse(c, http.StatusNotFound, errors.New("monitor not found"))
		}
		return s.returnErrorResponse(c, http.StatusInternalServerError, err)
	}
	switch monitor.MonitorType {
	case "http":
		// Validate config
		err = validateHTTPMonitorConfig(input.Config)
		if err != nil {
			return s.returnErrorResponse(c, http.StatusBadRequest, err)
		}
	default:
		return s.returnErrorResponse(c, http.StatusBadRequest, errors.New("invalid monitor_type: expected one of [http]"))
	}
	monitor, err = s.dbConnection.UpdateMonitor(c.Request().Context(), monitor.Name, db.UpdateMonitorInput{
		IntervalSeconds: input.IntervalSeconds,
		Config:          input.Config,
		Paused:          input.Paused,
		Description:     input.Description,
	})
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return s.returnErrorResponse(c, http.StatusNotFound, errors.New("monitor not found"))
		}
		return s.returnErrorResponse(c, http.StatusInternalServerError, err)
	}
	return s.returnSuccessResponse(c, http.StatusOK, convertDBMonitorToAPIMonitor(monitor))
}
