package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gosuri/uitable"
	"github.com/lbrictson/TinyMonitor/pkg/api"
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
					return c.ListMonitors(ops, context.String("output"))
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
				Name:        "apply",
				Description: "edit/create monitor",
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
	table := uitable.New()
	table.AddRow("Name", "Type", "Interval", "Status", "Last Checked")
	for _, x := range data {
		table.AddRow(x.Name, x.MonitorType, x.IntervalSeconds, x.Status, x.LastCheckedFriendly)
	}
	fmt.Println(table.String())
}

type ListMonitorOptions struct {
	Limit  *int
	Offset *int
	Type   *string
	Status *string
}

func (c *APIClient) ListMonitors(options ListMonitorOptions, outputFormat string) error {
	path := "/api/v1/monitor"
	if options.Limit != nil {
		path += fmt.Sprintf("?limit=%v", *options.Limit)
	} else {
		path += "?limit=1000"
	}
	if options.Offset != nil {
		path += fmt.Sprintf("&offset=%v", *options.Offset)
	}
	if options.Type != nil {
		path += fmt.Sprintf("&type=%v", *options.Type)
	}
	if options.Status != nil {
		path += fmt.Sprintf("&status=%v", *options.Status)
	}
	data, err := c.do(path, "GET", nil)
	if err != nil {
		return err
	}
	monitors := []api.MonitorModel{}
	err = json.Unmarshal(data, &monitors)
	if err != nil {
		return err
	}
	if outputFormat == "json" {
		return emitJSON(monitors)
	}
	printTextMonitorData(monitors)
	return nil
}

func (c *APIClient) GetMonitor(name string, outputformat string) error {
	data, err := c.do(fmt.Sprintf("/api/v1/monitor/%v", name), "GET", nil)
	if err != nil {
		return err
	}
	if outputformat == "json" {
		return emitJSON(data)
	}
	fmt.Println(data)
	return nil
}

func (c *APIClient) EditMonitor(name string) error {
	data, err := c.do(fmt.Sprintf("/api/v1/monitor/%v", name), "GET", nil)
	if err != nil {
		return err
	}
	monitor := api.MonitorModel{}
	err = json.Unmarshal(data, &monitor)
	if err != nil {
		return err
	}
	editedData, err := editStructInEditor(monitor)
	if err != nil {
		return err
	}
	editedMonitor := api.MonitorModel{}
	err = json.Unmarshal(editedData, &editedMonitor)
	if err != nil {
		return err
	}
	Updates := api.UpdateMonitorInput{
		IntervalSeconds: &editedMonitor.IntervalSeconds,
		Paused:          &editedMonitor.Paused,
		Config:          editedMonitor.Config,
		Description:     &editedMonitor.Description,
	}
	b, err := json.Marshal(Updates)
	if err != nil {
		return err
	}
	_, err = c.do(fmt.Sprintf("/api/v1/monitor/%v", name), "PATCH", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	fmt.Println("monitor updated")
	return nil
}

func (c *APIClient) ApplyMonitor(fileLocation string, outputformat string) error {
	// Read in the specified file
	file, err := os.ReadFile(fileLocation)
	if err != nil {
		return err
	}
	// Unmarshal the file into a monitor model
	monitor := api.MonitorModel{}
	err = json.Unmarshal(file, &monitor)
	if err != nil {
		return err
	}
	// Marshal the monitor model into a byte array
	b, err := json.Marshal(monitor)
	if err != nil {
		return err
	}
	// See if the monitor already exists
	_, err = c.do(fmt.Sprintf("/api/v1/monitor/%v", monitor.Name), "GET", nil)
	if err != nil {
		// If it doesn't exist, create it
		_, err = c.do("/api/v1/monitor", "POST", bytes.NewBuffer(b))
		if err != nil {
			return err
		}
		fmt.Println("monitor created")
		return nil
	}
	// Send the request to the API
	_, err = c.do(fmt.Sprintf("/api/v1/monitor/%v", monitor.Name), "PATCH", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	fmt.Println("monitor updated")
	return nil
}
