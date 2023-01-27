package client

import (
	"encoding/json"
	"fmt"
	"github.com/gosuri/uitable"
	"github.com/lbrictson/TinyMonitor/pkg/api"
	"github.com/urfave/cli/v2"
	"os"
)

func (c *APIClient) LoadAlertChannelCLICommands() *cli.Command {
	return &cli.Command{
		Name:        "alert",
		Description: "Manage alert channels",
		Usage:       "alert -h",
		Subcommands: []*cli.Command{
			{
				Name:        "list",
				Description: "List alert channels",
				Usage:       "alert list",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Value:   "text",
						Usage:   "Output format (text or json)",
					},
				},
				Action: func(context *cli.Context) error {
					return c.ListAlertChannels(context.String("output"))
				},
			},
			{
				Name:        "get",
				Description: "get alert channel",
				Usage:       "alert get $name",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Value:   "text",
						Usage:   "Output format (text or json)",
					},
				},
				Action: func(context *cli.Context) error {
					return c.GetAlertChannel(context.Args().First(), context.String("output"))
				},
			},
			{
				Name:        "edit",
				Description: "edit alert channel",
				Usage:       "alert edit $name",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Value:   "text",
						Usage:   "Output format (text or json)",
					},
				},
				Action: func(context *cli.Context) error {
					return c.EditAlertChannel(context.Args().First())
				},
			},
			{
				Name:        "delete",
				Description: "delete alert channel",
				Usage:       "alert delete $name",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Value:   "text",
						Usage:   "Output format (text or json)",
					},
				},
				Action: func(context *cli.Context) error {
					return c.DeleteAlertChannel(context.Args().First(), context.String("output"))
				},
			},
			{
				Name:        "apply",
				Description: "edit/create alert channel via file",
				Usage:       "alert apply -f $fileLocation",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Value:   "text",
						Usage:   "Output format (text or json)",
					},
					&cli.StringFlag{
						Name:    "file",
						Aliases: []string{"f"},
						Usage:   "File location",
					},
				},
				Action: func(context *cli.Context) error {
					return c.ApplyAlertChannel(context.String("file"), context.String("output"))
				},
			},
		},
	}
}

func printTextAlertChannelData(data []api.AlertChannelModel) {
	table := uitable.New()
	table.AddRow("Name", "Type")
	for _, x := range data {
		table.AddRow(x.Name, x.AlertChannelType)
	}
	fmt.Println(table.String())
}

func (c *APIClient) ListAlertChannels(outputFormat string) error {
	data, err := c.sdk.ListAlertChannels()
	if err != nil {
		return fmt.Errorf("error listing alert channels: %v", err)
	}
	if outputFormat == "json" {
		return emitJSON(data)
	}
	printTextAlertChannelData(data)
	return nil
}

func (c *APIClient) GetAlertChannel(name string, outputformat string) error {
	alert, err := c.sdk.GetAlertChannel(name)
	if err != nil {
		return fmt.Errorf("error getting alert channel %v: %v", name, err)
	}
	if outputformat == "json" {
		return emitJSON(alert)
	}
	printTextAlertChannelData([]api.AlertChannelModel{*alert})
	return nil
}

func (c *APIClient) EditAlertChannel(name string) error {
	data, err := c.sdk.GetAlertChannel(name)
	if err != nil {
		return fmt.Errorf("error editing alert channel %v: %v", name, err)
	}
	alert := api.UpdateAlertChannelInput{
		Config: data.Config,
	}
	editedData, err := editStructInEditor(alert)
	if err != nil {
		return fmt.Errorf("error editing alert channel %v: %v", name, err)
	}
	editedAlert := api.UpdateAlertChannelInput{}
	err = json.Unmarshal(editedData, &editedAlert)
	if err != nil {
		return fmt.Errorf("error editing alert channel %v: %v", name, err)
	}
	Updates := api.UpdateAlertChannelInput{
		Config: editedAlert.Config,
	}
	err = c.sdk.UpdateAlertChannel(name, Updates)
	if err != nil {
		return fmt.Errorf("error updating alert channel %v: %v", name, err)
	}
	fmt.Println("alert channel updated")
	return nil
}

func (c *APIClient) ApplyAlertChannel(fileLocation string, outputformat string) error {
	// Read in the specified file
	file, err := os.ReadFile(fileLocation)
	if err != nil {
		return fmt.Errorf("error reading file %v: %v", fileLocation, err)
	}
	// Unmarshal the file into a alert model
	alert := api.AlertChannelModel{}
	err = json.Unmarshal(file, &alert)
	if err != nil {
		return fmt.Errorf("error unmarshalling file %v: %v", fileLocation, err)
	}
	// See if the alert already exists
	_, err = c.sdk.GetAlertChannel(alert.Name)
	if err != nil {
		// If it doesn't exist, create it
		err = c.sdk.CreateAlertChannel(api.CreateAlertChannelInput{
			Name:             alert.Name,
			Config:           alert.Config,
			AlertChannelType: alert.AlertChannelType,
		})
		if err != nil {
			return fmt.Errorf("error creating alert channel %v: %v", alert.Name, err)
		}
		fmt.Println("alert channel created")
		return nil
	}
	// Send the request to the API
	err = c.sdk.UpdateAlertChannel(alert.Name, api.UpdateAlertChannelInput{
		Config: alert.Config,
	})
	if err != nil {
		return fmt.Errorf("error updating alert channel %v: %v", alert.Name, err)
	}
	fmt.Println("alert channel updated")
	return nil
}

func (c *APIClient) DeleteAlertChannel(name string, outputformat string) error {
	err := c.sdk.DeleteAlertChannel(name)
	if err != nil {
		return fmt.Errorf("error deleting alert channel %v: %v", name, err)
	}
	fmt.Println("alert channel deleted")
	return nil
}
