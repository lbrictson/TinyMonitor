package seeder

import (
	"context"
	"fmt"
	"github.com/lbrictson/TinyMonitor/pkg/api"
	"github.com/lbrictson/TinyMonitor/pkg/db"
	"github.com/lbrictson/TinyMonitor/pkg/security"
	"os"
)

func Run(dbConnection *db.DatabaseConnection) error {
	ctx := context.Background()
	initialAdminAPIKey := security.GenerateAPIKey()
	if os.Getenv("TINYMONITOR_TESTING") == "true" {
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
		fmt.Println("Initial Admin user created - this is the only time you will see this API Key - Save it somewhere safe!")
		fmt.Println("API Key: " + initialAdminAPIKey)
		fmt.Println("Username: admin")
		fmt.Println("You should use this initial admin user to create your own admin user and then delete the seeded admin user.")
	}
	// Create an initial monitor if there are none
	allMonitors, err := dbConnection.ListMonitors(ctx, db.ListMonitorOptions{})
	if err != nil {
		return err
	}
	if len(allMonitors) == 0 {
		_, err = dbConnection.CreateMonitor(ctx, db.CreateMonitorInput{
			Name:             "Server-Heartbeat-Check",
			IntervalSeconds:  10,
			MonitorType:      "http",
			FailureThreshold: 1,
			SuccessThreshold: 1,
			Config: api.ConvertHTTPMonitorConfigToGeneric(api.HTTPMonitorConfig{
				URL:                "http://127.0.0.1:8080/api/v1/health",
				Method:             "GET",
				BodyContains:       "",
				TimeoutMS:          200,
				ExpectResponseCode: 200,
				SkipTLSValidation:  true,
			}),
		})
		if err != nil {
			return err
		}
	}
	return nil
}
