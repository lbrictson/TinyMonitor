package validators

import (
	"errors"
	"strings"
)

func ValidateName(name string) error {
	if strings.Contains(name, " ") {
		return errors.New("name cannot contain spaces")
	}
	// Validate name  only contains letters, numbers, and dashes
	if !strings.ContainsAny(name, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-") {
		return errors.New("name can only contain letters, numbers, and dashes")
	}
	return nil
}

func ValidateSinkType(sinkType string) error {
	if sinkType == "" {
		return errors.New("sink_type is required")
	}
	available := []string{"influxdb-v1", "timestream", "cloudwatch"}
	for _, s := range available {
		if sinkType == s {
			return nil
		}
	}
	return nil
}
