package client

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/gosuri/uitable"
	"github.com/lbrictson/TinyMonitor/pkg/api"
	"github.com/lbrictson/TinyMonitor/pkg/sdk"
	"github.com/urfave/cli/v2"
	"os"
)

func (c *APIClient) LoadMonitorCLICommands() *cli.Command {
	return &cli.Command{
		Name:        "monitor",
		Description: "Manage monitors",
		Usage:       "monitor -h",
		Subcommands: []*cli.Command{
			{
				Name:        "list",
				Description: "List monitors",
				Usage:       "monitor list",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Value:   "text",
						Usage:   "Output format (text or json)",
					},
					&cli.IntFlag{
						Name:  "limit",
						Usage: "Limit number of returned monitors (Default 1000)",
						Value: 1000,
					},
					&cli.IntFlag{
						Name:  "offset",
						Usage: "Offset number of returned monitors (Default 0)",
						Value: 0,
					},
					&cli.StringFlag{
						Name:  "type",
						Usage: "Filter by monitor type",
					},
					&cli.StringFlag{
						Name:  "status",
						Usage: "Filter by monitor status",
					},
					&cli.StringFlag{
						Name:  "tag",
						Usage: "Filter by tag",
					},
					&cli.BoolFlag{
						Name:  "silenced",
						Usage: "Filter by silenced",
					},
				},
				Action: func(context *cli.Context) error {
					ops := ListMonitorOptions{}
					if context.IsSet("limit") {
						l := context.Int("limit")
						ops.Limit = &l
					}
					if context.IsSet("offset") {
						o := context.Int("offset")
						ops.Offset = &o
					}
					if context.IsSet("type") {
						t := context.String("type")
						ops.Type = &t
					}
					if context.IsSet("status") {
						s := context.String("status")
						ops.Status = &s
					}
					if context.IsSet("tag") {
						t := context.String("tag")
						ops.Tag = &t
					}
					if context.IsSet("silenced") {
						s := context.Bool("silenced")
						ops.Silenced = &s
					}
					return c.ListMonitors(ops, context.String("output"))
				},
			},
			{
				Name:        "get",
				Description: "get monitor",
				Usage:       "monitor get $name",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Value:   "text",
						Usage:   "Output format (text or json)",
					},
				},
				Action: func(context *cli.Context) error {
					return c.GetMonitor(context.Args().First(), context.String("output"))
				},
			},
			{
				Name:        "edit",
				Description: "edit monitor",
				Usage:       "monitor edit $name",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Value:   "text",
						Usage:   "Output format (text or json)",
					},
				},
				Action: func(context *cli.Context) error {
					return c.EditMonitor(context.Args().First())
				},
			},
			{
				Name:        "delete",
				Description: "delete monitor",
				Usage:       "monitor delete $name",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Value:   "text",
						Usage:   "Output format (text or json)",
					},
				},
				Action: func(context *cli.Context) error {
					return c.DeleteMonitor(context.Args().First(), context.String("output"))
				},
			},
			{
				Name:        "pause",
				Description: "pause monitor from being checked",
				Usage:       "monitor pause $name",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Value:   "text",
						Usage:   "Output format (text or json)",
					},
				},
				Action: func(context *cli.Context) error {
					return c.PauseMonitor(context.Args().First(), context.String("output"))
				},
			},
			{
				Name:        "unpause",
				Description: "unpause monitor",
				Usage:       "monitor unpause $name",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Value:   "text",
						Usage:   "Output format (text or json)",
					},
				},
				Action: func(context *cli.Context) error {
					return c.UnPauseMonitor(context.Args().First(), context.String("output"))
				},
			},
			{
				Name:        "silence",
				Description: "silence monitor so that it no longer sends alerts",
				Usage:       "monitor silence $name",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Value:   "text",
						Usage:   "Output format (text or json)",
					},
				},
				Action: func(context *cli.Context) error {
					return c.SilenceMonitor(context.Args().First(), context.String("output"))
				},
			},
			{
				Name:        "unsilence",
				Description: "unsilence monitor so that it sends alerts again",
				Usage:       "monitor unsilence $name",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Value:   "text",
						Usage:   "Output format (text or json)",
					},
				},
				Action: func(context *cli.Context) error {
					return c.UnSilenceMonitor(context.Args().First(), context.String("output"))
				},
			},
			{
				Name:        "apply",
				Description: "edit/create monitor via file",
				Usage:       "monitor apply -f $fileLocation",
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
					return c.ApplyMonitor(context.String("file"), context.String("output"))
				},
			},
		},
	}
}

func printTextMonitorData(data []api.MonitorModel) {
	red := color.New(color.FgRed)
	table := uitable.New()
	table.AddRow("Name", "Type", "Interval", "Status", "State", "Last Checked")
	for _, x := range data {
		state := "Active"
		if x.Paused {
			state = "Paused"
		}
		if x.Status == "Down" {
			x.Status = red.Sprint(x.Status)
		}
		if x.Silenced {
			x.Status = x.Status + " (Silenced)"
		}
		table.AddRow(x.Name, x.MonitorType, x.IntervalSeconds, fmt.Sprintf("%v (%v)", x.Status, x.StatusLastChangedFriendly), state, x.LastCheckedFriendly)
	}
	fmt.Println(table.String())
}

type ListMonitorOptions struct {
	Limit    *int
	Offset   *int
	Type     *string
	Status   *string
	Tag      *string
	Silenced *bool
}

func (c *APIClient) ListMonitors(options ListMonitorOptions, outputFormat string) error {
	data, err := c.sdk.ListMonitors(sdk.ListMonitorOptions{
		Limit:    options.Limit,
		Offset:   options.Offset,
		Type:     options.Type,
		Status:   options.Status,
		Tag:      options.Tag,
		Silenced: options.Silenced,
	})
	if err != nil {
		return fmt.Errorf("error listing monitors: %v", err)
	}
	if outputFormat == "json" {
		return emitJSON(data)
	}
	printTextMonitorData(data)
	return nil
}

func (c *APIClient) GetMonitor(name string, outputformat string) error {
	monitor, err := c.sdk.GetMonitor(name)
	if err != nil {
		return fmt.Errorf("error getting monitor %v: %v", name, err)
	}
	if outputformat == "json" {
		return emitJSON(monitor)
	}
	printTextMonitorData([]api.MonitorModel{*monitor})
	return nil
}

func (c *APIClient) EditMonitor(name string) error {
	data, err := c.sdk.GetMonitor(name)
	if err != nil {
		return fmt.Errorf("error editing monitor %v: %v", name, err)
	}
	if data.AlertChannels == nil {
		data.AlertChannels = []string{}
	}
	if data.Tags == nil {
		data.Tags = []string{}
	}
	monitor := api.UpdateMonitorInput{
		IntervalSeconds:  &data.IntervalSeconds,
		Paused:           &data.Paused,
		Config:           data.Config,
		Description:      &data.Description,
		SuccessThreshold: &data.SuccessThreshold,
		FailureThreshold: &data.FailureThreshold,
		Tags:             &data.Tags,
		AlertChannels:    &data.AlertChannels,
		Silenced:         &data.Silenced,
	}
	editedData, err := editStructInEditor(monitor)
	if err != nil {
		return fmt.Errorf("error editing monitor %v: %v", name, err)
	}
	editedMonitor := api.UpdateMonitorInput{}
	err = json.Unmarshal(editedData, &editedMonitor)
	if err != nil {
		return fmt.Errorf("error editing monitor %v: %v", name, err)
	}
	Updates := api.UpdateMonitorInput{
		IntervalSeconds:  editedMonitor.IntervalSeconds,
		Paused:           editedMonitor.Paused,
		Config:           editedMonitor.Config,
		Description:      editedMonitor.Description,
		Tags:             editedMonitor.Tags,
		SuccessThreshold: editedMonitor.SuccessThreshold,
		FailureThreshold: editedMonitor.FailureThreshold,
		AlertChannels:    editedMonitor.AlertChannels,
		Silenced:         editedMonitor.Silenced,
	}
	_, err = c.sdk.UpdateMonitor(name, Updates)
	if err != nil {
		return fmt.Errorf("error updating monitor %v: %v", name, err)
	}
	fmt.Println("monitor updated")
	return nil
}

func (c *APIClient) ApplyMonitor(fileLocation string, outputformat string) error {
	// Read in the specified file
	file, err := os.ReadFile(fileLocation)
	if err != nil {
		return fmt.Errorf("error reading file %v: %v", fileLocation, err)
	}
	// Unmarshal the file into a monitor model
	monitor := api.MonitorModel{}
	err = json.Unmarshal(file, &monitor)
	if err != nil {
		return fmt.Errorf("error unmarshalling file %v: %v", fileLocation, err)
	}
	// See if the monitor already exists
	_, err = c.sdk.GetMonitor(monitor.Name)
	if err != nil {
		// If it doesn't exist, create it
		_, err = c.sdk.CreateMonitor(api.CreateMonitorInput{
			Name:             monitor.Name,
			Description:      monitor.Description,
			IntervalSeconds:  monitor.IntervalSeconds,
			MonitorType:      monitor.MonitorType,
			Config:           monitor.Config,
			SuccessThreshold: monitor.SuccessThreshold,
			FailureThreshold: monitor.FailureThreshold,
			Tags:             monitor.Tags,
			AlertChannels:    monitor.AlertChannels,
			Silenced:         monitor.Silenced,
		})
		if err != nil {
			return fmt.Errorf("error creating monitor %v: %v", monitor.Name, err)
		}
		fmt.Println("monitor created")
		return nil
	}
	// Send the request to the API
	_, err = c.sdk.UpdateMonitor(monitor.Name, api.UpdateMonitorInput{
		IntervalSeconds:  &monitor.IntervalSeconds,
		Paused:           &monitor.Paused,
		Config:           monitor.Config,
		Description:      &monitor.Description,
		SuccessThreshold: &monitor.SuccessThreshold,
		FailureThreshold: &monitor.FailureThreshold,
		Tags:             &monitor.Tags,
		AlertChannels:    &monitor.AlertChannels,
		Silenced:         &monitor.Silenced,
	})
	if err != nil {
		return fmt.Errorf("error updating monitor %v: %v", monitor.Name, err)
	}
	fmt.Println("monitor updated")
	return nil
}

func (c *APIClient) DeleteMonitor(name string, outputformat string) error {
	err := c.sdk.DeleteMonitor(name)
	if err != nil {
		return fmt.Errorf("error deleting monitor %v: %v", name, err)
	}
	fmt.Println("monitor deleted")
	return nil
}

func (c *APIClient) SilenceMonitor(name string, outputformat string) error {
	silence := true
	_, err := c.sdk.UpdateMonitor(name, api.UpdateMonitorInput{
		Silenced: &silence,
	})
	if err != nil {
		return fmt.Errorf("error silencing monitor %v: %v", name, err)
	}
	fmt.Printf("monitor %v silenced\n", name)
	return nil
}

func (c *APIClient) UnSilenceMonitor(name string, outputformat string) error {
	silence := false
	_, err := c.sdk.UpdateMonitor(name, api.UpdateMonitorInput{
		Silenced: &silence,
	})
	if err != nil {
		return fmt.Errorf("error unsilenced monitor %v: %v", name, err)
	}
	fmt.Printf("monitor %v unsilenced\n", name)
	return nil
}

func (c *APIClient) PauseMonitor(name string, outputformat string) error {
	pause := true
	_, err := c.sdk.UpdateMonitor(name, api.UpdateMonitorInput{
		Paused: &pause,
	})
	if err != nil {
		return fmt.Errorf("error pausing monitor %v: %v", name, err)
	}
	fmt.Printf("monitor %v paused\n", name)
	return nil
}

func (c *APIClient) UnPauseMonitor(name string, outputformat string) error {
	pause := false
	_, err := c.sdk.UpdateMonitor(name, api.UpdateMonitorInput{
		Paused: &pause,
	})
	if err != nil {
		return fmt.Errorf("error unpausing monitor %v: %v", name, err)
	}
	fmt.Printf("monitor %v unpaused\n", name)
	return nil
}
