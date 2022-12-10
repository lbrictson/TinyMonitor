package client

import (
	"encoding/json"
	"fmt"
	"github.com/gosuri/uitable"
	"github.com/lbrictson/TinyMonitor/pkg/api"
	"github.com/urfave/cli/v2"
)

func (c *APIClient) LoadMonitorCLICommands() *cli.Command {
	return &cli.Command{
		Name:        "monitor",
		Description: "Manage monitors",
		Subcommands: []*cli.Command{
			{
				Name:        "list",
				Description: "List monitors",
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
		},
	}
}

func printTextMonitorData(data []api.MonitorModel) {
	table := uitable.New()
	table.AddRow("ID", "Name", "Type", "Interval", "Status", "Last Checked")
	for _, x := range data {
		table.AddRow(x.ID, x.Name, x.MonitorType, x.IntervalSeconds, x.Status, x.LastCheckedFriendly)
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

func (c *APIClient) GetMonitor(id *int, name *string, outputformat string) error {
	if id != nil {
		data, err := c.do(fmt.Sprintf("/api/v1/monitor/%v", *id), "GET", nil)
		if err != nil {
			return err
		}
		if outputformat == "json" {
			return emitJSON(data)
		}
		fmt.Println(data)
		return nil
	}
	if name != nil {
		monitorsAPIData, err := c.do("/api/v1/monitor", "GET", nil)
		if err != nil {
			return err
		}
		monitors := []api.MonitorModel{}
		err = json.Unmarshal(monitorsAPIData, &monitors)
		if err != nil {
			return err
		}
		for _, x := range monitors {
			if x.Name == *name {
				if outputformat == "json" {
					return emitJSON(x)
				}
				printTextMonitorData([]api.MonitorModel{x})
				return nil
			}
		}
		return fmt.Errorf("monitor not found")
	}
	return fmt.Errorf("must specify either id or name")
}
