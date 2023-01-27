package api

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/TinyMonitor/pkg/alert"
	"github.com/lbrictson/TinyMonitor/pkg/db"
	"github.com/lbrictson/TinyMonitor/pkg/validators"
	"net/http"
	"strings"
	"time"
)

type AlertChannelModel struct {
	Name             string                 `json:"name"`
	AlertChannelType string                 `json:"alert_channel_type"`
	Config           map[string]interface{} `json:"config"`
	CreatedAt        time.Time              `json:"created_at"`
	UpdatedAt        time.Time              `json:"updated_at"`
	Monitors         []string               `json:"monitors"`
}

func (a *AlertChannelModel) SendDown(monitorName string, message string, dbConn *db.DatabaseConnection) error {
	ctx := context.TODO()
	switch strings.ToLower(a.AlertChannelType) {
	case "email":
		emailConfig := AlertChannelEmailConfig{}
		b, err := json.Marshal(a.Config)
		if err != nil {
			return err
		}
		err = json.Unmarshal(b, &emailConfig)
		if err != nil {
			return err
		}
		for i, _ := range emailConfig.To {
			emailConfig.To[i] = injectSecretsIntoContent(ctx, dbConn, emailConfig.To[i])
		}
		for i, _ := range emailConfig.CC {
			emailConfig.CC[i] = injectSecretsIntoContent(ctx, dbConn, emailConfig.CC[i])
		}
		for i, _ := range emailConfig.BCC {
			emailConfig.BCC[i] = injectSecretsIntoContent(ctx, dbConn, emailConfig.BCC[i])
		}
		return alert.NewEmailAlerter(alert.NewEmailAlerterInput{
			From: injectSecretsIntoContent(ctx, dbConn, emailConfig.From),
			Host: injectSecretsIntoContent(ctx, dbConn, emailConfig.Host),
			Port: emailConfig.Port,
			User: injectSecretsIntoContent(ctx, dbConn, emailConfig.Username),
			Pass: injectSecretsIntoContent(ctx, dbConn, emailConfig.Password),
			To:   emailConfig.To,
			CC:   emailConfig.CC,
			BCC:  emailConfig.BCC,
		}).SendDown(monitorName, message)
	case "slack":
		slackConfig := AlertChannelSlackConfig{}
		b, err := json.Marshal(a.Config)
		if err != nil {
			return err
		}
		err = json.Unmarshal(b, &slackConfig)
		if err != nil {
			return err
		}
		return alert.NewSlackAlerter(alert.NewSlackAlerterInput{
			WebhookURL: injectSecretsIntoContent(ctx, dbConn, slackConfig.WebhookURL),
		}).SendDown(monitorName, message)
	case "pagerduty":
		pagerDutyConfig := AlertChannelPagerDutyConfig{}
		b, err := json.Marshal(a.Config)
		if err != nil {
			return err
		}
		err = json.Unmarshal(b, &pagerDutyConfig)
		if err != nil {
			return err
		}
		return alert.NewPagerDutyAlerter(alert.NewPagerDutyAlerterInput{
			ServiceKey: injectSecretsIntoContent(ctx, dbConn, pagerDutyConfig.ServiceKey),
		}).SendDown(monitorName, message)
	case "webhook":
		webhookConfig := AlertChannelWebhookConfig{}
		b, err := json.Marshal(a.Config)
		if err != nil {
			return err
		}
		err = json.Unmarshal(b, &webhookConfig)
		if err != nil {
			return err
		}
		return alert.NewWebhookAlerter(alert.NewWebhookAlerterInput{
			WebhookURL: injectSecretsIntoContent(ctx, dbConn, webhookConfig.URL),
		}).SendDown(monitorName, message)
	}
	return nil
}

func (a *AlertChannelModel) SendUp(monitorName string, message string, dbConn *db.DatabaseConnection) error {
	ctx := context.TODO()
	switch strings.ToLower(a.AlertChannelType) {
	case "email":
		emailConfig := AlertChannelEmailConfig{}
		b, err := json.Marshal(a.Config)
		if err != nil {
			return err
		}
		err = json.Unmarshal(b, &emailConfig)
		if err != nil {
			return err
		}
		for i, _ := range emailConfig.To {
			emailConfig.To[i] = injectSecretsIntoContent(ctx, dbConn, emailConfig.To[i])
		}
		for i, _ := range emailConfig.CC {
			emailConfig.CC[i] = injectSecretsIntoContent(ctx, dbConn, emailConfig.CC[i])
		}
		for i, _ := range emailConfig.BCC {
			emailConfig.BCC[i] = injectSecretsIntoContent(ctx, dbConn, emailConfig.BCC[i])
		}
		return alert.NewEmailAlerter(alert.NewEmailAlerterInput{
			From: injectSecretsIntoContent(ctx, dbConn, emailConfig.From),
			Host: injectSecretsIntoContent(ctx, dbConn, emailConfig.Host),
			Port: emailConfig.Port,
			User: injectSecretsIntoContent(ctx, dbConn, emailConfig.Username),
			Pass: injectSecretsIntoContent(ctx, dbConn, emailConfig.Password),
			To:   emailConfig.To,
			CC:   emailConfig.CC,
			BCC:  emailConfig.BCC,
		}).SendUp(monitorName, message)
	case "slack":
		slackConfig := AlertChannelSlackConfig{}
		b, err := json.Marshal(a.Config)
		if err != nil {
			return err
		}
		err = json.Unmarshal(b, &slackConfig)
		if err != nil {
			return err
		}
		return alert.NewSlackAlerter(alert.NewSlackAlerterInput{
			WebhookURL: injectSecretsIntoContent(ctx, dbConn, slackConfig.WebhookURL),
		}).SendUp(monitorName, message)
	case "pagerduty":
		pagerDutyConfig := AlertChannelPagerDutyConfig{}
		b, err := json.Marshal(a.Config)
		if err != nil {
			return err
		}
		err = json.Unmarshal(b, &pagerDutyConfig)
		if err != nil {
			return err
		}
		return alert.NewPagerDutyAlerter(alert.NewPagerDutyAlerterInput{
			ServiceKey: injectSecretsIntoContent(ctx, dbConn, pagerDutyConfig.ServiceKey),
		}).SendUp(monitorName, message)
	case "webhook":
		webhookConfig := AlertChannelWebhookConfig{}
		b, err := json.Marshal(a.Config)
		if err != nil {
			return err
		}
		err = json.Unmarshal(b, &webhookConfig)
		if err != nil {
			return err
		}
		return alert.NewWebhookAlerter(alert.NewWebhookAlerterInput{
			WebhookURL: injectSecretsIntoContent(ctx, dbConn, webhookConfig.URL),
		}).SendUp(monitorName, message)
	}
	return nil
}

func convertDBAlertChannelToAlertChannelModel(dbAlertChannel *db.BaseAlertChannel) *AlertChannelModel {
	if dbAlertChannel == nil {
		return nil
	}
	return &AlertChannelModel{
		Name:             dbAlertChannel.Name,
		AlertChannelType: dbAlertChannel.ChannelType,
		Config:           dbAlertChannel.Config,
		CreatedAt:        dbAlertChannel.CreatedAt,
		UpdatedAt:        dbAlertChannel.UpdatedAt,
		Monitors:         dbAlertChannel.Monitors,
	}
}

type CreateAlertChannelInput struct {
	Name             string                 `json:"name"`
	AlertChannelType string                 `json:"alert_channel_type"`
	Config           map[string]interface{} `json:"config"`
}

func (i *CreateAlertChannelInput) validate() error {
	if validators.ValidateName(i.Name) != nil {
		return errors.New("invalid name")
	}
	if validators.ValidateAlertChannelType(i.AlertChannelType) != nil {
		return errors.New("invalid alert channel type")
	}
	return nil
}

func (s *Server) createAlertChannel(c echo.Context) error {
	var input CreateAlertChannelInput
	if err := c.Bind(&input); err != nil {
		return s.returnErrorResponse(c, http.StatusBadRequest, err)
	}
	if err := input.validate(); err != nil {
		return s.returnErrorResponse(c, http.StatusBadRequest, err)
	}
	dbAlertChannel, err := s.dbConnection.CreateAlertChannel(c.Request().Context(), db.CreateAlertChannelInput{
		Name:        input.Name,
		ChannelType: input.AlertChannelType,
		Config:      input.Config,
	})
	if err != nil {
		return s.returnErrorResponse(c, http.StatusInternalServerError, err)
	}
	return s.returnSuccessResponse(c, http.StatusOK, convertDBAlertChannelToAlertChannelModel(dbAlertChannel))
}

func (s *Server) getAlertChannel(c echo.Context) error {
	name := c.Param("id")
	dbAlertChannel, err := s.dbConnection.GetAlertChannel(c.Request().Context(), name)
	if err != nil {
		return s.returnErrorResponse(c, http.StatusInternalServerError, err)
	}
	return s.returnSuccessResponse(c, http.StatusOK, convertDBAlertChannelToAlertChannelModel(dbAlertChannel))
}

func (s *Server) listAlertChannels(c echo.Context) error {
	dbAlertChannels, err := s.dbConnection.ListAlertChannels(c.Request().Context())
	if err != nil {
		return s.returnErrorResponse(c, http.StatusInternalServerError, err)
	}
	var alertChannels []AlertChannelModel
	for _, dbAlertChannel := range dbAlertChannels {
		alertChannels = append(alertChannels, *convertDBAlertChannelToAlertChannelModel(dbAlertChannel))
	}
	return s.returnSuccessResponse(c, http.StatusOK, alertChannels)
}

func (s *Server) deleteAlertChannel(c echo.Context) error {
	name := c.Param("id")
	err := s.dbConnection.DeleteAlertChannel(c.Request().Context(), name)
	if err != nil {
		return s.returnErrorResponse(c, http.StatusInternalServerError, err)
	}
	return s.returnSuccessResponse(c, http.StatusOK, nil)
}

type UpdateAlertChannelInput struct {
	Config map[string]interface{} `json:"config"`
}

func (i *UpdateAlertChannelInput) validate() error {
	return nil
}

func (s *Server) updateAlertChannel(c echo.Context) error {
	var input UpdateAlertChannelInput
	if err := c.Bind(&input); err != nil {
		return s.returnErrorResponse(c, http.StatusBadRequest, err)
	}
	if err := input.validate(); err != nil {
		return s.returnErrorResponse(c, http.StatusBadRequest, err)
	}
	name := c.Param("id")
	dbAlertChannel, err := s.dbConnection.UpdateAlertChannel(c.Request().Context(), name, db.UpdateAlertChannelInput{
		Config: input.Config,
	})
	if err != nil {
		return s.returnErrorResponse(c, http.StatusInternalServerError, err)
	}
	return s.returnSuccessResponse(c, http.StatusOK, convertDBAlertChannelToAlertChannelModel(dbAlertChannel))
}
