package alert

import (
	"github.com/slack-go/slack"
)

type NewSlackAlerterInput struct {
	WebhookURL string
}

func NewSlackAlerter(input NewSlackAlerterInput) *SlackAlerter {
	return &SlackAlerter{
		webhookURL: input.WebhookURL,
	}
}

type SlackAlerter struct {
	webhookURL string
}

func (a *SlackAlerter) SendDown(monitorName string, message string) error {
	// Send to slack using webhook method
	attachment := slack.Attachment{
		Color:      "danger",
		Fallback:   "Monitor Down: " + monitorName,
		AuthorName: "[ALERT] Monitor Down: " + monitorName,
		Text:       message,
		Footer:     "TinyMonitor",
	}
	msg := slack.WebhookMessage{
		Attachments: []slack.Attachment{attachment},
	}

	return slack.PostWebhook(a.webhookURL, &msg)
}

func (a *SlackAlerter) SendUp(monitorName string, message string) error {
	// Send to slack using webhook method
	attachment := slack.Attachment{
		Color:      "good",
		Fallback:   "Monitor Up: " + monitorName,
		AuthorName: "[RESOLVED] Monitor Up: " + monitorName,
		Text:       message,
		Footer:     "TinyMonitor",
	}
	msg := slack.WebhookMessage{
		Attachments: []slack.Attachment{attachment},
	}

	return slack.PostWebhook(a.webhookURL, &msg)
}
