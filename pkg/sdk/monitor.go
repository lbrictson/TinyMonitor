package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lbrictson/TinyMonitor/pkg/api"
)

type ListMonitorOptions struct {
	Limit  *int
	Offset *int
	Type   *string
	Status *string
}

func (c *Client) ListMonitors(opts ListMonitorOptions) ([]api.MonitorModel, error) {
	path := "/api/v1/monitor"
	if opts.Limit != nil {
		path += fmt.Sprintf("?limit=%v", *opts.Limit)
	} else {
		path += "?limit=1000"
	}
	if opts.Offset != nil {
		path += fmt.Sprintf("&offset=%v", *opts.Offset)
	}
	if opts.Type != nil {
		path += fmt.Sprintf("&type=%v", *opts.Type)
	}
	if opts.Status != nil {
		path += fmt.Sprintf("&status=%v", *opts.Status)
	}
	data, err := c.do(path, "GET", nil)
	if err != nil {
		return nil, err
	}
	monitors := []api.MonitorModel{}
	err = json.Unmarshal(data, &monitors)
	if err != nil {
		return nil, err
	}
	return monitors, nil
}

func (c *Client) GetMonitor(name string) (*api.MonitorModel, error) {
	data, err := c.do(fmt.Sprintf("/api/v1/monitor/%v", name), "GET", nil)
	if err != nil {
		return nil, err
	}
	monitor := api.MonitorModel{}
	err = json.Unmarshal(data, &monitor)
	if err != nil {
		return nil, err
	}
	return &monitor, nil
}

func (c *Client) UpdateMonitor(name string, updates api.UpdateMonitorInput) (*api.MonitorModel, error) {
	b, err := json.Marshal(updates)
	if err != nil {
		return nil, err
	}
	data, err := c.do(fmt.Sprintf("/api/v1/monitor/%v", name), "PATCH", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	monitor := api.MonitorModel{}
	err = json.Unmarshal(data, &monitor)
	if err != nil {
		return nil, err
	}
	return &monitor, nil
}

func (c *Client) CreateMonitor(monitor api.CreateMonitorInput) (*api.MonitorModel, error) {
	b, err := json.Marshal(monitor)
	if err != nil {
		return nil, err
	}
	data, err := c.do("/api/v1/monitor", "POST", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	monitorModel := api.MonitorModel{}
	err = json.Unmarshal(data, &monitorModel)
	if err != nil {
		return nil, err
	}
	return &monitorModel, nil
}

func (c *Client) DeleteMonitor(name string) error {
	_, err := c.do(fmt.Sprintf("/api/v1/monitor/%v", name), "DELETE", nil)
	return err
}
