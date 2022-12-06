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
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to decode error response",
		}, err
	}
	return resp, nil
}

func (c *APIClient) decodeSuccessResponse(body []byte) (api.SuccessResponse, error) {
	var resp api.SuccessResponse
	err := json.Unmarshal(body, &resp)
	if err != nil {
		return api.SuccessResponse{}, err
	}
	return resp, nil
}

func (c *APIClient) wasRequestSuccessful(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}

func (c *APIClient) do(path string, method string, payload io.Reader) (api.SuccessResponse, api.ErrorResponse, error) {
	req, err := http.NewRequest(method, c.formatURL(path), payload)
	if err != nil {
		return api.SuccessResponse{}, api.ErrorResponse{Message: err.Error()}, err
	}
	c.addAuthHeaders(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return api.SuccessResponse{}, api.ErrorResponse{Message: err.Error()}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return api.SuccessResponse{}, api.ErrorResponse{Message: err.Error()}, err
	}
	if !c.wasRequestSuccessful(resp.StatusCode) {
		errResp, err := c.decodeErrorResponse(body)
		if err != nil {
			return api.SuccessResponse{}, api.ErrorResponse{Message: err.Error()}, err
		}
		return api.SuccessResponse{}, errResp, errors.New(errResp.Message)
	}
	successResp, err := c.decodeSuccessResponse(body)
	if err != nil {
		return api.SuccessResponse{}, api.ErrorResponse{Message: err.Error()}, err
	}
	return successResp, api.ErrorResponse{}, nil
}
