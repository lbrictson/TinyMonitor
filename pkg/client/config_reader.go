package client

import (
	"encoding/json"
	"os"
)

type Config struct {
	ServerURL string `json:"server_url"`
	Username  string `json:"username"`
	APIKey    string `json:"api_key"`
}

func ReadConfig() (*Config, error) {
	var config Config
	// Check if home directory /.tinymonitor/config.json exists
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	configPath := home + "/.tinymonitor/config.json"
	// Check if configPath exists
	_, err = os.Stat(configPath)
	if err == nil {
		data, err := os.ReadFile(configPath)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, &config)
		if err != nil {
			return nil, err
		}
	}
	// Read env vars
	if serverURL := os.Getenv("TINYMONITOR_SERVER_URL"); serverURL != "" {
		config.ServerURL = serverURL
	}
	if username := os.Getenv("TINYMONITOR_USERNAME"); username != "" {
		config.Username = username
	}
	if apiKey := os.Getenv("TINYMONITOR_API_KEY"); apiKey != "" {
		config.APIKey = apiKey
	}
	return &config, nil
}
