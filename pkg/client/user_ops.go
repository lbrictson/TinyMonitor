package client

import (
	"errors"
	"fmt"
)

func (c *APIClient) ListUsers(outputFormat string) error {
	data, errResp, err := c.do("/api/v1/user", "GET", nil)
	if err != nil {
		fmt.Println(errResp.Message)
		return errors.New(errResp.Message)
	}
	if outputFormat == "json" {
		return emitJSON(data.Data)
	}
	fmt.Println(data.Data)
	return nil
}
