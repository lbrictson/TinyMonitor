package client

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/gosuri/uitable"
	"github.com/lbrictson/TinyMonitor/pkg/api"
	"github.com/urfave/cli/v2"
)

func (c *APIClient) LoadSecretCLICommands() *cli.Command {
	return &cli.Command{
		Name:        "secret",
		Description: "Manage secrets",
		Usage:       "secret -h",
		Subcommands: []*cli.Command{
			{
				Name:        "list",
				Description: "List secrets",
				Usage:       "secret list",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Value:   "text",
						Usage:   "Output format (text or json)",
					},
				},
				Action: func(context *cli.Context) error {
					return c.ListSecrets(context.String("output"))
				},
			},
			{
				Name:        "get",
				Description: "get secret",
				Usage:       "secret get $name",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Value:   "text",
						Usage:   "Output format (text or json)",
					},
				},
				Action: func(context *cli.Context) error {
					return c.GetSecret(context.Args().First(), context.String("output"))
				},
			},
			{
				Name:        "edit",
				Description: "edit secret",
				Usage:       "secret edit --name $name --value $newValue",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Value:   "text",
						Usage:   "Output format (text or json)",
					},
					&cli.StringFlag{
						Name:     "name",
						Aliases:  []string{"n"},
						Usage:    "name",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "value",
						Aliases:  []string{"v"},
						Usage:    "value",
						Required: true,
					},
				},
				Action: func(context *cli.Context) error {
					return c.EditSecret(context.String("name"), context.String("value"))
				},
			},
			{
				Name:        "delete",
				Description: "delete secret",
				Usage:       "secret delete $name",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Value:   "text",
						Usage:   "Output format (text or json)",
					},
				},
				Action: func(context *cli.Context) error {
					return c.DeleteSecret(context.Args().First(), context.String("output"))
				},
			},
			{
				Name:        "create",
				Description: "create secret",
				Usage:       "secret create --name $name --value $value",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Value:   "text",
						Usage:   "Output format (text or json)",
					},
					&cli.StringFlag{
						Name:     "name",
						Aliases:  []string{"n"},
						Usage:    "name",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "value",
						Aliases:  []string{"v"},
						Usage:    "value",
						Required: true,
					},
				},
				Action: func(context *cli.Context) error {
					return c.CreateSecret(context.String("name"), context.String("value"))
				},
			},
		},
	}
}

func printTextSecretData(data []api.SecretModel) {
	table := uitable.New()
	table.AddRow("Name", "Updated By", "Last Updated")
	for _, x := range data {
		table.AddRow(x.Name, x.LastUpdatedBy, humanize.Time(x.UpdatedAt))
	}
	fmt.Println(table.String())
}

func (c *APIClient) ListSecrets(output string) error {
	secrets, err := c.sdk.ListSecrets()
	if err != nil {
		return fmt.Errorf("error listing secrets: %w", err)
	}
	if output == "json" {
		return emitJSON(secrets)
	}
	printTextSecretData(secrets)
	return nil
}

func (c *APIClient) GetSecret(name string, output string) error {
	secret, err := c.sdk.GetSecret(name)
	if err != nil {
		return fmt.Errorf("error getting secret %v: %v", name, err)
	}
	if output == "json" {
		return emitJSON(secret)
	}
	printTextSecretData([]api.SecretModel{*secret})
	return nil
}

func (c *APIClient) EditSecret(name string, newValue string) error {
	err := c.sdk.UpdateSecret(name, api.UpdateSecretInput{Value: newValue})
	if err != nil {
		return fmt.Errorf("failed to update secret %v: %v", name, err)
	}
	fmt.Printf("Secret %s updated successfully\n", name)
	return nil
}

func (c *APIClient) DeleteSecret(name string, output string) error {
	err := c.sdk.DeleteSecret(name)
	if err != nil {
		return fmt.Errorf("fFailed to delete secret %v: %v", name, err)
	}
	fmt.Printf("Secret %s deleted successfully\n", name)
	return nil
}

func (c *APIClient) CreateSecret(name string, value string) error {
	err := c.sdk.CreateSecret(api.CreateSecretInput{Name: name, Value: value})
	if err != nil {
		return fmt.Errorf("failed to create secret %v: %v", name, err)
	}
	fmt.Printf("Secret %s created successfully\n", name)
	return nil
}
