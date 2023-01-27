package alert

import (
	"github.com/PagerDuty/go-pagerduty"
)

type PagerDutyAlerter struct {
	serviceKey string
}

type NewPagerDutyAlerterInput struct {
	ServiceKey string
}

func NewPagerDutyAlerter(input NewPagerDutyAlerterInput) *PagerDutyAlerter {
	return &PagerDutyAlerter{
		serviceKey: input.ServiceKey,
	}
}

func (a *PagerDutyAlerter) SendDown(monitorName string, message string) error {
	event := pagerduty.Event{
		Type:        "trigger",
		ServiceKey:  a.serviceKey,
		Description: "TinyStatus Monitor Down: " + monitorName,
		Details:     message,
	}
	_, err := pagerduty.CreateEvent(event)
	return err
}

func (a *PagerDutyAlerter) SendUp(monitorName string, message string) error {
	event := pagerduty.Event{
		Type:        "resolve",
		ServiceKey:  a.serviceKey,
		Description: "TinyStatus Monitor Up: " + monitorName,
		Details:     message,
	}
	_, err := pagerduty.CreateEvent(event)
	return err
}
