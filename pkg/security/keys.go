package security

import (
	"github.com/google/uuid"
	"strings"
)

// GenerateAPIKey generates a new API key that is unique
func GenerateAPIKey() string {
	// TODO make this more secure
	return strings.Replace(uuid.New().String()+uuid.New().String(), "-", "", -1)
}
