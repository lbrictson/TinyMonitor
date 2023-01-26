package api

import (
	"context"
	"github.com/lbrictson/TinyMonitor/pkg/db"
	"strings"
)

func injectSecretsIntoContent(ctx context.Context, dbConn *db.DatabaseConnection, content string) string {
	// Get all secrets from the database
	secrets, err := dbConn.ListSecrets(ctx)
	if err != nil {
		return content
	}
	// Replace all instances of {{secret_name}} with the secret value
	for _, secret := range secrets {
		content = strings.Replace(content, "${{"+secret.Name+"}}", secret.Value, -1)
	}
	return content
}
