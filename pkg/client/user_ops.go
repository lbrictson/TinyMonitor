package client

import (
	"bytes"
	"encoding/json"
	"errors"
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
	data, err := c.do("/api/v1/user", "GET", nil)
	if err != nil {
		return err
	}
	user := []api.UserModel{}
	err = json.Unmarshal(data, &user)
	if err != nil {
		return err
	}
	if outputFormat == "json" {
		return emitJSON(user)
	}
	printTextUserData(user)
	return nil
}

func (c *APIClient) GetUser(username string, outputformat string) error {
	usersAPIData, err := c.do("/api/v1/user", "GET", nil)
	if err != nil {
		return err
	}
	users := []api.UserModel{}
	err = json.Unmarshal(usersAPIData, &users)
	if err != nil {
		return err
	}
	for _, x := range users {
		if x.Username == username {
			if outputformat == "json" {
				return emitJSON(x)
			}
			printTextUserData([]api.UserModel{x})
			return nil
		}
	}
	return errors.New("user not found")
}

func (c *APIClient) CreateUser(username string, role string, outputformat string) error {
	apiInput := api.CreateUserRequest{
		Username: username,
		Role:     role,
	}
	b, err := json.Marshal(&apiInput)
	if err != nil {
		return err
	}
	resp, err := c.do("/api/v1/user", "POST", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	data := api.CreateUserResponse{}
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return err
	}
	if outputformat == "json" {
		return emitJSON(data)
	}
	fmt.Printf("User %v created successfully, API key is: %v\n", username, data.APIKey)
	return nil
}

func (c *APIClient) DeleteUser(username string, outputformat string) error {
	_, err := c.do(fmt.Sprintf("/api/v1/user/%v", username), "DELETE", nil)
	if err != nil {
		return err
	}
	fmt.Printf("User %v deleted successfully\n", username)
	return nil
}

func (c *APIClient) ChangeUserRole(username string, newRole string, outputformat string) error {
	apiInput := api.UpdateUserRequest{
		Role:      &newRole,
		LockedOut: nil,
	}
	b, err := json.Marshal(&apiInput)
	if err != nil {
		return err
	}
	_, err = c.do(fmt.Sprintf("/api/v1/user/%v", username), "PATCH", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	fmt.Printf("User %v role changed successfully to %v\n", username, newRole)
	return nil
}
