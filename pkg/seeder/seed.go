package seeder

import (
	"context"
	"github.com/lbrictson/TinyMonitor/pkg/db"
	"github.com/lbrictson/TinyMonitor/pkg/security"
)

func Run(dbConnection *db.DatabaseConnection) error {
	ctx := context.Background()
	allUsers, err := dbConnection.ListUsers(ctx)
	if err != nil {
		return err
	}
	if len(allUsers) == 0 {
		// Create the initial admin user
		_, err = dbConnection.CreateUser(ctx, db.CreateUserInput{
			Username: "admin",
			APIKey:   security.GenerateAPIKey(),
			Role:     "admin",
		})
		if err != nil {
			return err
		}
	}
	return nil
}
