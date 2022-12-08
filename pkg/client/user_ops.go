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
		Subcommands: []*cli.Command{
			{
				Name:        "list",
				Description: "List all users",
				Usage:       "List all users",
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
				Description: "Get a user",
				Usage:       "Get a user",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "username",
						Aliases: []string{"u"},
						Usage:   "The username of the user",
					},
					&cli.IntFlag{
						Name:  "id",
						Usage: "The id of the user",
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
					if context.IsSet("username") {
						u := context.String("username")
						return c.GetUser(nil, &u, context.String("output"))
					}
					if context.IsSet("id") {
						i := context.Int("id")
						return c.GetUser(&i, nil, context.String("output"))
					}
					return fmt.Errorf("must specify either username or id")
				},
			},
			{
				Name:        "create",
				Description: "Create a user",
				Usage:       "Create a user",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "username",
						Aliases: []string{"u"},
						Usage:   "The username of the user",
						Required: true,
					},
					&cli.StringFlag{
						Name:  "role",
						Usage: "The role of the user (read_only, write, admin)",
						Aliases: []string{"r"},
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
					return c.CreateUser(context.String("username"), context.String("role"))
				},
			},
		},
	}
}

func printTextUserData(data []api.UserModel) {
	table := uitable.New()
	table.AddRow("ID", "Username", "Role", "Locked")
	for _, x := range data {
		table.AddRow(x.ID, x.Username, x.Role, x.LockedOut)
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

func (c *APIClient) GetUser(id *int, username *string, outputformat string) error {
	if id != nil {
		data, err := c.do(fmt.Sprintf("/api/v1/user/%v", *id), "GET", nil)
		if err != nil {
			return err
		}
		if outputformat == "json" {
			return emitJSON(data)
		}
		fmt.Println(data)
		return nil
	}
	if username != nil {
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
			if x.Username == *username {
				if outputformat == "json" {
					return emitJSON(x)
				}
				printTextUserData([]api.UserModel{x})
				return nil
			}
		}
		return errors.New("user not found")
	}
	return errors.New("id or username is required")
}

func (c *APIClient) CreateUser(username string, role string) error {
	apiInput := api.CreateUserRequest{
		Username: username,
		Role:     role,
	}
	b, err := json.Marshal(&apiInput)
	if err != nil {
		return err
	}
	_, err = c.do("/api/v1/user", "POST", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	fmt.Printf("User %v created successfully\n", username)
	return nil
}
