package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lbrictson/TinyMonitor/pkg/api"
)

func (c *Client) ListAlertChannels() ([]api.AlertChannelModel, error) {
	data, err := c.do("/api/v1/alert-channel", "GET", nil)
	if err != nil {
		return nil, err
	}
	alertChannels := []api.AlertChannelModel{}
	err = json.Unmarshal(data, &alertChannels)
	if err != nil {
		return nil, err
	}
	return alertChannels, nil
}

func (c *Client) GetAlertChannel(name string) (*api.AlertChannelModel, error) {
	data, err := c.do(fmt.Sprintf("/api/v1/alert-channel/%v", name), "GET", nil)
	if err != nil {
		return nil, err
	}
	alertChannel := api.AlertChannelModel{}
	err = json.Unmarshal(data, &alertChannel)
	if err != nil {
		return nil, err
	}
	return &alertChannel, nil
}

func (c *Client) CreateAlertChannel(alertChannel api.CreateAlertChannelInput) error {
	b, err := json.Marshal(alertChannel)
	if err != nil {
		return err
	}
	_, err = c.do("/api/v1/alert-channel", "POST", bytes.NewReader(b))
	return err
}

func (c *Client) DeleteAlertChannel(name string) error {
	_, err := c.do(fmt.Sprintf("/api/v1/alert-channel/%v", name), "DELETE", nil)
	return err
}

func (c *Client) UpdateAlertChannel(name string, alertChannel api.UpdateAlertChannelInput) error {
	b, err := json.Marshal(alertChannel)
	if err != nil {
		return err
	}
	_, err = c.do(fmt.Sprintf("/api/v1/alert-channel/%v", name), "PATCH", bytes.NewReader(b))
	return err
}
