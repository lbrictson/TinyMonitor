package client

import (
	"github.com/lbrictson/TinyMonitor/pkg/sdk"
)

type APIClient struct {
	username  string
	apiKey    string
	serverURL string
	sdk       *sdk.Client
}

func NewAPIClient(server, username, apiKey string) *APIClient {
	s, err := sdk.NewClient(sdk.NewClientOpts{
		Username:  username,
		APIKey:    apiKey,
		ServerURL: server,
	})
	if err != nil {
		panic(err)
	}
	return &APIClient{
		username:  username,
		apiKey:    apiKey,
		serverURL: server,
		sdk:       s,
	}
}
