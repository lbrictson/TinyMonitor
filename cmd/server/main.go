package main

import (
	"github.com/lbrictson/TinyMonitor/pkg/api"
	"github.com/lbrictson/TinyMonitor/pkg/config"
	"github.com/lbrictson/TinyMonitor/pkg/db"
	"github.com/lbrictson/TinyMonitor/pkg/seeder"
	_ "github.com/lib/pq"
	"github.com/playwright-community/playwright-go"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	if os.Getenv("SETUP_PLAYWRIGHT") == "true" {
		playwright.Install()
		return
	}
	// Setup browser automation testing
	err := playwright.Install()
	if err != nil {
		panic(err)
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
		Port:         conf.Port,
		AutoTLS:      conf.AutoTLS,
		Hostname:     conf.Hostname,
		Logger:       l,
		DBConnection: dbConn,
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
