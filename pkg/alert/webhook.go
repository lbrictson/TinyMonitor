package alert

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type WebhookPayload struct {
	Name      string    `json:"name,omitempty"`
	Reason    string    `json:"reason,omitempty"`
	Status    string    `json:"status,omitempty"`
	EventTime time.Time `json:"event_time"`
}

type WebhookAlerter struct {
	webhookURL string
}

type NewWebhookAlerterInput struct {
	WebhookURL string
}

func NewWebhookAlerter(input NewWebhookAlerterInput) *WebhookAlerter {
	return &WebhookAlerter{
		webhookURL: input.WebhookURL,
	}
}

func (a *WebhookAlerter) SendDown(monitorName string, message string) error {
	// Make an HTTP POST to the specified webhook URL
	content := WebhookPayload{
		Name:      monitorName,
		Reason:    message,
		Status:    "Down",
		EventTime: time.Now(),
	}
	b, err := json.Marshal(content)
	if err != nil {
		return err
	}
	_, err = http.DefaultClient.Post(a.webhookURL, "application/json", bytes.NewReader(b))
	return err
}

func (a *WebhookAlerter) SendUp(monitorName string, message string) error {
	// Make an HTTP POST to the specified webhook URL
	content := WebhookPayload{
		Name:      monitorName,
		Reason:    message,
		Status:    "Up",
		EventTime: time.Now(),
	}
	b, err := json.Marshal(content)
	if err != nil {
		return err
	}
	_, err = http.DefaultClient.Post(a.webhookURL, "application/json", bytes.NewReader(b))
	return err
}
