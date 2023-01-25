package client

import (
	"fmt"
	"github.com/gosuri/uitable"
	"github.com/lbrictson/TinyMonitor/pkg/api"
	"github.com/urfave/cli/v2"
)

func (c *APIClient) LoadUserCLICommands() *cli.Command {
	return &cli.Command{
		Name:        "user",
		Description: "Manage users",
		Usage:       "user -h",
		Subcommands: []*cli.Command{
			{
				Name:        "list",
				Description: "List all users",
				Usage:       "user list",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "output",
						Aliases:  []string{"o"},
						Value:    "text",
						Usage:    "Output format. One of: text, json",
						Required: false,
					},
				},
				Action: func(context *cli.Context) error {
					return c.ListUsers(context.String("output"))
				},
			},
			{
				Name:        "get",
				Description: "Get a user $username",
				Usage:       "user get $username",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "output",
						Aliases:  []string{"o"},
						Value:    "text",
						Usage:    "Output format. One of: text, json",
						Required: false,
					},
				},
				Action: func(context *cli.Context) error {
					return c.GetUser(context.Args().First(), context.String("output"))

				},
			},
			{
				Name:        "edit",
				Description: "Edit a user",
				Usage:       "user edit $username -r admin",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "role",
						Aliases:  []string{"r"},
						Usage:    "The role of the user",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "output",
						Aliases:  []string{"o"},
						Value:    "text",
						Usage:    "Output format. One of: text, json",
						Required: false,
					},
				},
				Action: func(context *cli.Context) error {
					return c.ChangeUserRole(context.Args().First(), context.String("role"), context.String("output"))
				},
			},
			{
				Name:        "delete",
				Description: "Delete a user",
				Usage:       "user delete $username",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "output",
						Aliases:  []string{"o"},
						Value:    "text",
						Usage:    "Output format. One of: text, json",
						Required: false,
					},
				},
				Action: func(context *cli.Context) error {
					return c.DeleteUser(context.Args().First(), context.String("output"))
				},
			},
			{
				Name:        "create",
				Description: "Create a user",
				Usage:       "user create -u $username -r admin",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "username",
						Aliases:  []string{"u"},
						Usage:    "The username of the user",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "role",
						Usage:    "The role of the user (read_only, write, admin)",
						Aliases:  []string{"r"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "output",
						Aliases:  []string{"o"},
						Value:    "text",
						Usage:    "Output format. One of: text, json",
						Required: false,
					},
				},
				Action: func(context *cli.Context) error {
					return c.CreateUser(context.String("username"), context.String("role"), context.String("output"))
				},
			},
		},
	}
}

func printTextUserData(data []api.UserModel) {
	table := uitable.New()
	table.AddRow("Username", "Role", "Locked")
	for _, x := range data {
		table.AddRow(x.Username, x.Role, x.LockedOut)
	}
	fmt.Println(table.String())
}

func (c *APIClient) ListUsers(outputFormat string) error {
	users, err := c.sdk.ListUsers()
	if err != nil {
		return err
	}
	if outputFormat == "json" {
		return emitJSON(users)
	}
	printTextUserData(users)
	return nil
}

func (c *APIClient) GetUser(username string, outputformat string) error {
	user, err := c.sdk.GetUser(username)
	if err != nil {
		return err
	}
	if outputformat == "json" {
		return emitJSON(user)
	}
	printTextUserData([]api.UserModel{*user})
	return nil
}

func (c *APIClient) CreateUser(username string, role string, outputformat string) error {
	user, err := c.sdk.CreateUser(username, role)
	if err != nil {
		return err
	}
	fmt.Printf("User %v created successfully, API key is: %v\n", username, *user)
	return nil
}

func (c *APIClient) DeleteUser(username string, outputformat string) error {
	err := c.sdk.DeleteUser(username)
	if err != nil {
		return err
	}
	fmt.Printf("User %v deleted successfully\n", username)
	return nil
}

func (c *APIClient) ChangeUserRole(username string, newRole string, outputformat string) error {
	_, err := c.sdk.UpdateUser(username, newRole)
	if err != nil {
		return err
	}
	fmt.Printf("User %v role changed successfully to %v\n", username, newRole)
	return nil
}
