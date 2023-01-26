package api

import (
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/TinyMonitor/pkg/db"
	"github.com/lbrictson/TinyMonitor/pkg/validators"
	"net/http"
	"strings"
	"time"
)

type BaseSink struct {
	Name      string                 `json:"name"`
	SinkType  string                 `json:"sink_type"`
	Config    map[string]interface{} `json:"config"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
}

func convertDBSinkToAPISink(dbSink *db.BaseSink) *BaseSink {
	if dbSink == nil {
		return nil
	}
	return &BaseSink{
		Name:      dbSink.Name,
		SinkType:  dbSink.SinkType,
		Config:    dbSink.Config,
		CreatedAt: dbSink.CreatedAt,
		UpdatedAt: dbSink.UpdatedAt,
	}
}

func (s *Server) getSink(c echo.Context) error {
	m, err := s.dbConnection.GetSink(c.Request().Context(), c.Param("id"))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return s.returnErrorResponse(c, http.StatusNotFound, errors.New("sink not found"))
		}
		return s.returnErrorResponse(c, http.StatusInternalServerError, err)
	}
	return s.returnSuccessResponse(c, http.StatusOK, convertDBSinkToAPISink(m))
}

func (s *Server) listSinks(c echo.Context) error {
	sinks, err := s.dbConnection.ListSinks(c.Request().Context())
	if err != nil {
		return s.returnErrorResponse(c, http.StatusInternalServerError, err)
	}
	apiSinks := make([]*BaseSink, 0)
	for _, m := range sinks {
		apiSinks = append(apiSinks, convertDBSinkToAPISink(m))
	}
	return s.returnSuccessResponse(c, http.StatusOK, apiSinks)
}

type CreateSinkInput struct {
	Name     string                 `json:"name"`
	SinkType string                 `json:"sink_type"`
	Config   map[string]interface{} `json:"config"`
}

func (c *CreateSinkInput) Validate() error {
	if c.Name == "" {
		return errors.New("name is required")
	}
	if c.SinkType == "" {
		return errors.New("sink_type is required")
	}
	if err := validators.ValidateSinkType(c.SinkType); err != nil {
		return err
	}
	if err := validators.ValidateName(c.Name); err != nil {
		return err
	}
	if c.Config == nil {
		return errors.New("config is required")
	}
	switch c.SinkType {
	case "influxdb-v1":
		b, err := json.Marshal(c.Config)
		if err != nil {
			return err
		}
		var config InfluxDBV1SinkConfig
		if err := json.Unmarshal(b, &config); err != nil {
			return err
		}
		if err := validateInfluxDBV1SinkConfig(config); err != nil {
			return err
		}
	case "timestream":
		b, err := json.Marshal(c.Config)
		if err != nil {
			return err
		}
		var config TimeStreamSinkConfig
		if err := json.Unmarshal(b, &config); err != nil {
			return err
		}
		if err := validateTimeStreamSinkConfig(config); err != nil {
			return err
		}
	case "cloudwatch":
		break
	default:
		return errors.New("unknown sink type")
	}
	return nil
}

func (s *Server) createSink(c echo.Context) error {
	input := &CreateSinkInput{}
	if err := c.Bind(input); err != nil {
		return s.returnErrorResponse(c, http.StatusBadRequest, err)
	}
	if err := input.Validate(); err != nil {
		return s.returnErrorResponse(c, http.StatusBadRequest, err)
	}
	sink, err := s.dbConnection.CreateSink(c.Request().Context(), db.CreateSinkInput{
		Name:     input.Name,
		SinkType: input.SinkType,
		Config:   input.Config,
	})
	if err != nil {
		return s.returnErrorResponse(c, http.StatusInternalServerError, err)
	}
	loadSingleSinkIntoMetrics(*convertDBSinkToAPISink(sink))
	return s.returnSuccessResponse(c, http.StatusCreated, convertDBSinkToAPISink(sink))
}

type UpdateSinkInput struct {
	Config map[string]interface{} `json:"config"`
}

func (u *UpdateSinkInput) Validate() error {
	if u.Config == nil {
		return errors.New("config is required")
	}
	return nil
}

func (s *Server) updateSink(c echo.Context) error {
	input := &UpdateSinkInput{}
	if err := c.Bind(input); err != nil {
		return s.returnErrorResponse(c, http.StatusBadRequest, err)
	}
	if err := input.Validate(); err != nil {
		return s.returnErrorResponse(c, http.StatusBadRequest, err)
	}
	sink, err := s.dbConnection.UpdateSink(c.Request().Context(), c.Param("id"), db.UpdateSinkInput{
		Config: input.Config,
	})
	if err != nil {
		return s.returnErrorResponse(c, http.StatusInternalServerError, err)
	}
	loadSingleSinkIntoMetrics(*convertDBSinkToAPISink(sink))
	return s.returnSuccessResponse(c, http.StatusOK, convertDBSinkToAPISink(sink))
}

func (s *Server) deleteSink(c echo.Context) error {
	if err := s.dbConnection.DeleteSink(c.Request().Context(), c.Param("id")); err != nil {
		return s.returnErrorResponse(c, http.StatusInternalServerError, err)
	}
	removeSinkFromMetrics(c.Param("id"))
	return s.returnSuccessResponse(c, http.StatusOK, nil)
}
