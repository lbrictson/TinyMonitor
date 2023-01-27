package alert

import (
	"github.com/PagerDuty/go-pagerduty"
)

func init() {
	pdAlertKeys = make(map[string]string)
}

var pdAlertKeys = map[string]string{}

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
	id, err := pagerduty.CreateEvent(event)
	if err == nil {
		pdAlertKeys[monitorName] = id.IncidentKey
	}
	return err
}

func (a *PagerDutyAlerter) SendUp(monitorName string, message string) error {
	// Only attempt to resolve if the key exists
	if _, ok := pdAlertKeys[monitorName]; ok {
		event := pagerduty.Event{
			Type:        "resolve",
			ServiceKey:  a.serviceKey,
			Description: "TinyStatus Monitor Up: " + monitorName,
			Details:     message,
			IncidentKey: pdAlertKeys[monitorName],
		}
		_, err := pagerduty.CreateEvent(event)
		// Remove the key
		delete(pdAlertKeys, monitorName)
		return err
	}
	return nil
}
