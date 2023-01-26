package client

import (
	"encoding/json"
	"fmt"
	"github.com/gosuri/uitable"
	"github.com/lbrictson/TinyMonitor/pkg/api"
	"github.com/urfave/cli/v2"
	"os"
)

func (c *APIClient) LoadSinkCLICommands() *cli.Command {
	return &cli.Command{
		Name:        "sink",
		Description: "Manage Metric Sinks",
		Usage:       "sink -h",
		Subcommands: []*cli.Command{
			{
				Name:        "list",
				Description: "List sinks",
				Usage:       "sink list",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Value:   "text",
						Usage:   "Output format (text or json)",
					},
				},
				Action: func(context *cli.Context) error {
					return c.ListSinks(context.String("output"))
				},
			},
			{
				Name:        "get",
				Description: "get sink",
				Usage:       "sink get $name",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Value:   "text",
						Usage:   "Output format (text or json)",
					},
				},
				Action: func(context *cli.Context) error {
					return c.GetSink(context.Args().First(), context.String("output"))
				},
			},
			{
				Name:        "edit",
				Description: "edit sink",
				Usage:       "sink edit $name",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Value:   "text",
						Usage:   "Output format (text or json)",
					},
				},
				Action: func(context *cli.Context) error {
					return c.EditSink(context.Args().First())
				},
			},
			{
				Name:        "delete",
				Description: "delete sink",
				Usage:       "sink delete $name",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Value:   "text",
						Usage:   "Output format (text or json)",
					},
				},
				Action: func(context *cli.Context) error {
					return c.DeleteSink(context.Args().First(), context.String("output"))
				},
			},
			{
				Name:        "apply",
				Description: "edit/create sink via file",
				Usage:       "sink apply -f $fileLocation",
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
					return c.ApplySink(context.String("file"), context.String("output"))
				},
			},
		},
	}
}

func printTextSinkData(data []api.BaseSink) {
	table := uitable.New()
	table.AddRow("Name", "Type")
	for _, x := range data {
		table.AddRow(x.Name, x.SinkType)
	}
	fmt.Println(table.String())
}

func (c *APIClient) ListSinks(outputFormat string) error {
	data, err := c.sdk.ListSinks()
	if err != nil {
		return fmt.Errorf("error listing sinks: %v", err)
	}
	if outputFormat == "json" {
		return emitJSON(data)
	}
	printTextSinkData(data)
	return nil
}

func (c *APIClient) GetSink(name string, outputformat string) error {
	s, err := c.sdk.GetSink(name)
	if err != nil {
		return fmt.Errorf("error getting monitor %v: %v", name, err)
	}
	if outputformat == "json" {
		return emitJSON(s)
	}
	printTextSinkData([]api.BaseSink{*s})
	return nil
}

func (c *APIClient) EditSink(name string) error {
	data, err := c.sdk.GetSink(name)
	if err != nil {
		return fmt.Errorf("error editing sink %v: %v", name, err)
	}
	s := api.UpdateSinkInput{
		Config: data.Config,
	}
	editedData, err := editStructInEditor(s)
	if err != nil {
		return fmt.Errorf("error editing sink %v: %v", name, err)
	}
	editedSink := api.UpdateSinkInput{}
	err = json.Unmarshal(editedData, &editedSink)
	if err != nil {
		return fmt.Errorf("error editing sink %v: %v", name, err)
	}
	Updates := api.UpdateSinkInput{
		Config: editedSink.Config,
	}
	err = c.sdk.UpdateSink(name, Updates)
	if err != nil {
		return fmt.Errorf("error updating sink %v: %v", name, err)
	}
	fmt.Println("sink updated")
	return nil
}

func (c *APIClient) ApplySink(fileLocation string, outputformat string) error {
	// Read in the specified file
	file, err := os.ReadFile(fileLocation)
	if err != nil {
		return fmt.Errorf("error reading file %v: %v", fileLocation, err)
	}
	// Unmarshal the file into a sink model
	sink := api.BaseSink{}
	err = json.Unmarshal(file, &sink)
	if err != nil {
		return fmt.Errorf("error unmarshalling file %v: %v", fileLocation, err)
	}
	// See if the sink already exists
	_, err = c.sdk.GetSink(sink.Name)
	if err != nil {
		// If it doesn't exist, create it
		err = c.sdk.CreateSink(api.CreateSinkInput{
			Name:     sink.Name,
			Config:   sink.Config,
			SinkType: sink.SinkType,
		})
		if err != nil {
			return fmt.Errorf("error creating sink %v: %v", sink.Name, err)
		}
		fmt.Println("sink created")
		return nil
	}
	// Send the request to the API
	err = c.sdk.UpdateSink(sink.Name, api.UpdateSinkInput{
		Config: sink.Config,
	})
	if err != nil {
		return fmt.Errorf("error updating sink %v: %v", sink.Name, err)
	}
	fmt.Println("sink updated")
	return nil
}

func (c *APIClient) DeleteSink(name string, outputformat string) error {
	err := c.sdk.DeleteSink(name)
	if err != nil {
		return fmt.Errorf("error deleting sink %v: %v", name, err)
	}
	fmt.Println("sink deleted")
	return nil
}
