package sdk

import (
	"bytes"
	"encoding/json"
	"github.com/lbrictson/TinyMonitor/pkg/api"
)

func (c *Client) ListSecrets() ([]api.SecretModel, error) {
	path := "/api/v1/secret"
	data, err := c.do(path, "GET", nil)
	if err != nil {
		return nil, err
	}
	secrets := []api.SecretModel{}
	err = json.Unmarshal(data, &secrets)
	if err != nil {
		return nil, err
	}
	return secrets, nil
}

func (c *Client) GetSecret(name string) (*api.SecretModel, error) {
	data, err := c.do("/api/v1/secret/"+name, "GET", nil)
	if err != nil {
		return nil, err
	}
	secret := api.SecretModel{}
	err = json.Unmarshal(data, &secret)
	if err != nil {
		return nil, err
	}
	return &secret, nil
}

func (c *Client) CreateSecret(secret api.CreateSecretInput) error {
	b, err := json.Marshal(secret)
	if err != nil {
		return err
	}
	data, err := c.do("/api/v1/secret", "POST", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	secretModel := api.SecretModel{}
	err = json.Unmarshal(data, &secretModel)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteSecret(name string) error {
	_, err := c.do("/api/v1/secret/"+name, "DELETE", nil)
	return err
}

func (c *Client) UpdateSecret(name string, secret api.UpdateSecretInput) error {
	b, err := json.Marshal(secret)
	if err != nil {
		return err
	}
	_, err = c.do("/api/v1/secret/"+name, "PATCH", bytes.NewBuffer(b))
	return err
}
