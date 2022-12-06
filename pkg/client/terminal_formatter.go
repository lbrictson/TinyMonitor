package client

import (
	"encoding/json"
	"fmt"
)

func emitJSON(data interface{}) error {
	// Pretty print the JSON
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	fmt.Println(string(val))
	return nil
}
