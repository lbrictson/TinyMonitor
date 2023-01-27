package main

import (
	"fmt"
	"github.com/lbrictson/TinyMonitor/pkg/api"
	"github.com/lbrictson/TinyMonitor/pkg/client"
	"github.com/lbrictson/TinyMonitor/pkg/config"
	"github.com/lbrictson/TinyMonitor/pkg/db"
	"github.com/lbrictson/TinyMonitor/pkg/seeder"
	_ "github.com/lib/pq"
	"github.com/playwright-community/playwright-go"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

func main() {
	arg := ""
	if len(os.Args) >= 2 {
		arg = os.Args[1]
	}
	if strings.ToLower(arg) == "server" {
		runServer()
	} else {
		runCLI()
	}
}

func runServer() {
	if os.Getenv("SETUP_PLAYWRIGHT") == "true" {
		playwright.Install()
		return
	}
	browserEnabled := true
	// Setup browser automation testing
	if os.Getenv("SKIP_PLAYWRIGHT") != "true" {
		err := playwright.Install()
		if err != nil {
			browserEnabled = false
			fmt.Printf("Error install playwright headless browser packages %v, browser based monitors will be disabled\n", err)
		}
	}
	conf, err := config.ReadServerConfig()
	if err != nil {
		panic(err)
	}
	// Configure logger
	l := logrus.New()
	if conf.LogFormat == "json" {
		l.SetFormatter(&logrus.JSONFormatter{})
	}
	switch conf.LogLevel {
	case "debug":
		l.SetLevel(logrus.DebugLevel)
	case "info":
		l.SetLevel(logrus.InfoLevel)
	case "warn":
		l.SetLevel(logrus.WarnLevel)
	case "error":
		l.SetLevel(logrus.ErrorLevel)
	case "fatal":
		l.SetLevel(logrus.FatalLevel)
	case "panic":
		l.SetLevel(logrus.PanicLevel)
	default:
		l.SetLevel(logrus.InfoLevel)
	}
	// Connect to DB
	dbConn, err := db.NewDatabaseConnection(db.NewDatabaseConnectionInput{
		InMemory: false,
		Location: conf.DBLocation,
	})
	if err != nil {
		l.Fatalf("Error connecting to database: %v", err)
	}
	s, err := api.NewServer(api.NewServerInput{
		Port:                     conf.Port,
		AutoTLS:                  conf.AutoTLS,
		Hostname:                 conf.Hostname,
		Logger:                   l,
		DBConnection:             dbConn,
		BrowserMonitoringEnabled: browserEnabled,
	})
	if err != nil {
		l.Fatalf("Failed to create server: %v", err)
	}
	err = seeder.Run(dbConn)
	if err != nil {
		l.Fatalf("Failed to seed DB: %v", err)
	}
	s.Run()
}

var c *client.APIClient

func runCLI() {
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
