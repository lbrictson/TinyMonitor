package client

import (
	"encoding/json"
	"errors"
	"github.com/lbrictson/TinyMonitor/pkg/api"
	"io"
	"net/http"
)

type APIClient struct {
	username  string
	apiKey    string
	serverURL string
}

func NewAPIClient(server, username, apiKey string) *APIClient {
	return &APIClient{
		username:  username,
		apiKey:    apiKey,
		serverURL: server,
	}
}

func (c *APIClient) addAuthHeaders(req *http.Request) {
	req.Header.Add("X-Username", c.username)
	req.Header.Add("X-Api-Key", c.apiKey)
	req.Header.Add("Content-Type", "application/json")
}

func (c *APIClient) formatURL(path string) string {
	return c.serverURL + path
}

func (c *APIClient) decodeErrorResponse(body []byte) (api.ErrorResponse, error) {
	var resp api.ErrorResponse
	err := json.Unmarshal(body, &resp)
	if err != nil {
		return api.ErrorResponse{
			Error: "failed to decode error response",
		}, err
	}
	return resp, nil
}

func (c *APIClient) wasRequestSuccessful(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}

func (c *APIClient) do(path string, method string, payload io.Reader) ([]byte, error) {
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
