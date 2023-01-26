package sdk

import (
	"bytes"
	"encoding/json"
	"github.com/lbrictson/TinyMonitor/pkg/api"
)

func (c *Client) ListSinks() ([]api.BaseSink, error) {
	path := "/api/v1/sink"
	data, err := c.do(path, "GET", nil)
	if err != nil {
		return nil, err
	}
	sinks := []api.BaseSink{}
	err = json.Unmarshal(data, &sinks)
	if err != nil {
		return nil, err
	}
	return sinks, nil
}

func (c *Client) GetSink(name string) (*api.BaseSink, error) {
	data, err := c.do("/api/v1/sink/"+name, "GET", nil)
	if err != nil {
		return nil, err
	}
	sink := api.BaseSink{}
	err = json.Unmarshal(data, &sink)
	if err != nil {
		return nil, err
	}
	return &sink, nil
}

func (c *Client) CreateSink(sink api.CreateSinkInput) error {
	b, err := json.Marshal(sink)
	if err != nil {
		return err
	}
	data, err := c.do("/api/v1/sink", "POST", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	sinkModel := api.BaseSink{}
	err = json.Unmarshal(data, &sinkModel)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteSink(name string) error {
	_, err := c.do("/api/v1/sink/"+name, "DELETE", nil)
	return err
}

func (c *Client) UpdateSink(name string, sink api.UpdateSinkInput) error {
	b, err := json.Marshal(sink)
	if err != nil {
		return err
	}
	_, err = c.do("/api/v1/sink/"+name, "PATCH", bytes.NewBuffer(b))
	return err
}
