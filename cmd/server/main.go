package main

import (
	"github.com/lbrictson/TinyMonitor/pkg/api"
	"github.com/lbrictson/TinyMonitor/pkg/config"
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
	s, err := api.NewServer(api.NewServerInput{
		Port:     conf.Port,
		AutoTLS:  conf.AutoTLS,
		Hostname: conf.Hostname,
		Logger:   l,
	})
	if err != nil {
		panic(err)
	}
	s.Run()
}
