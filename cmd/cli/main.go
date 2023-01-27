package main

import (
	"fmt"
	"github.com/lbrictson/TinyMonitor/pkg/client"
	"github.com/urfave/cli/v2"
	"os"
)

var c *client.APIClient

func main() {
	config, err := client.ReadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	c = client.NewAPIClient(config.ServerURL, config.Username, config.APIKey)
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "print only the version",
	}
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:     "output",
			Aliases:  []string{"o"},
			Value:    "text",
			Usage:    "Output format. One of: text, json",
			Required: false,
		},
	}
	app := &cli.App{
		Name:     "tiny-monitor",
		Usage:    "The TinyMonitor CLI interface",
		Version:  "0.0.1",
		Flags:    flags,
		Commands: []*cli.Command{c.LoadUserCLICommands(), c.LoadMonitorCLICommands(), c.LoadSecretCLICommands(), c.LoadSinkCLICommands(), c.LoadAlertChannelCLICommands()},
	}
	app.Flags = flags
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
