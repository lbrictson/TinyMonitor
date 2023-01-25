package sdk

import (
	"bytes"
	"encoding/json"
	"github.com/lbrictson/TinyMonitor/pkg/api"
)

func (c *Client) ListUsers() ([]api.UserModel, error) {
	data, err := c.do("/api/v1/user", "GET", nil)
	if err != nil {
		return nil, err
	}
	users := []api.UserModel{}
	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (c *Client) GetUser(username string) (*api.UserModel, error) {
	data, err := c.do("/api/v1/user/"+username, "GET", nil)
	if err != nil {
		return nil, err
	}
	user := api.UserModel{}
	err = json.Unmarshal(data, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *Client) DeleteUser(username string) error {
	_, err := c.do("/api/v1/user/"+username, "DELETE", nil)
	return err
}

// CreateUser will return the users api key if successful
func (c *Client) CreateUser(username string, role string) (*string, error) {
	apiInput := api.CreateUserRequest{
		Username: username,
		Role:     role,
	}
	b, err := json.Marshal(&apiInput)
	if err != nil {
		return nil, err
	}
	resp, err := c.do("/api/v1/user", "POST", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	data := api.CreateUserResponse{}
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return &data.APIKey, nil
}

func (c *Client) UpdateUser(username string, role string) (*api.UserModel, error) {
	apiInput := api.UpdateUserRequest{
		Role: &role,
	}
	b, err := json.Marshal(&apiInput)
	if err != nil {
		return nil, err
	}
	resp, err := c.do("/api/v1/user/"+username, "PUT", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	data := api.UserModel{}
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
