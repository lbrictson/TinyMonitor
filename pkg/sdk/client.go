package sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lbrictson/TinyMonitor/pkg/api"
	"io"
	"net/http"
)

type Client struct {
	username  string
	apiKey    string
	serverURL string
}

type NewClientOpts struct {
	Username  string
	APIKey    string
	ServerURL string
}

func NewClient(opts NewClientOpts) (*Client, error) {
	if opts.Username == "" {
		return nil, fmt.Errorf("username is required")
	}
	if opts.APIKey == "" {
		return nil, fmt.Errorf("api key is required")
	}
	if opts.ServerURL == "" {
		return nil, fmt.Errorf("server url is required")
	}
	return &Client{
		username:  opts.Username,
		apiKey:    opts.APIKey,
		serverURL: opts.ServerURL,
	}, nil
}
func (c *Client) addAuthHeaders(req *http.Request) {
	req.Header.Add("X-Username", c.username)
	req.Header.Add("X-Api-Key", c.apiKey)
	req.Header.Add("Content-Type", "application/json")
}

func (c *Client) formatURL(path string) string {
	return c.serverURL + path
}

func (c *Client) decodeErrorResponse(body []byte) (api.ErrorResponse, error) {
	var resp api.ErrorResponse
	err := json.Unmarshal(body, &resp)
	if err != nil {
		return api.ErrorResponse{
			Error: "failed to decode error response",
		}, err
	}
	return resp, nil
}

func (c *Client) wasRequestSuccessful(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}

func (c *Client) do(path string, method string, payload io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, c.formatURL(path), payload)
	if err != nil {
		return nil, err
	}
	c.addAuthHeaders(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if !c.wasRequestSuccessful(resp.StatusCode) {
		errResp, err := c.decodeErrorResponse(body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(errResp.Error)
	}
	return body, nil
}
