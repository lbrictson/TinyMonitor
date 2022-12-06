package seeder

import (
	"context"
	"github.com/lbrictson/TinyMonitor/pkg/db"
	"github.com/lbrictson/TinyMonitor/pkg/security"
	"os"
)

func Run(dbConnection *db.DatabaseConnection) error {
	ctx := context.Background()
	initialAdminAPIKey := security.GenerateAPIKey()
	if os.Getenv("TINYSTATUS_TESTING") == "true" {
		initialAdminAPIKey = "aaaabbbbcccceeeedddd"
	}
	allUsers, err := dbConnection.ListUsers(ctx)
	if err != nil {
		return err
	}
	if len(allUsers) == 0 {
		// Create the initial admin user
		_, err = dbConnection.CreateUser(ctx, db.CreateUserInput{
			Username: "admin",
			APIKey:   initialAdminAPIKey,
			Role:     "admin",
		})
		if err != nil {
			return err
		}
	}
	return nil
}
