package main

import (
	"context"
	"fmt"
	"github.com/lbrictson/TinyMonitor/ent"
	"github.com/lbrictson/TinyMonitor/pkg/api"
	"github.com/lbrictson/TinyMonitor/pkg/config"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
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
	dsn := fmt.Sprintf("host=%s port=%d user=%v password=%v dbname=%v sslmode=%v", conf.DBHost, conf.DBPort, conf.DBUser, conf.DBPass, conf.DBName, conf.DBSSLMode)
	dbClient, err := ent.Open("postgres", dsn)
	if err != nil {
		l.Fatalf("Failed to connect to DB: %v", err)
	}
	err = dbClient.Schema.Create(context.Background())
	if err != nil {
		l.Fatalf("Failed to perform migrations: %v", err)
	}
	s, err := api.NewServer(api.NewServerInput{
		Port:     conf.Port,
		AutoTLS:  conf.AutoTLS,
		Hostname: conf.Hostname,
		Logger:   l,
	})
	if err != nil {
		l.Fatalf("Failed to create server: %v", err)
	}
	s.Run()
}
