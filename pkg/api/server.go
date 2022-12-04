package api

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Port     int    // Port to listen on
	AutoTLS  bool   // If true, use Let's Encrypt to automatically get a TLS certificate
	Hostname string // Hostname must be set if using AutoTLS, also used when sending email alerts for callback URLs
	Logger   *logrus.Logger
}

// NewServerInput is the input for creating a Server.
type NewServerInput struct {
	Port     int            // Port to listen on
	AutoTLS  bool           // If true, use Let's Encrypt to automatically get a TLS certificate
	Hostname string         // Hostname must be set if using AutoTLS, also used when sending email alerts for callback URLs
	Logger   *logrus.Logger // Logger to use for logging
}

// NewServer creates a new api server.
func NewServer(input NewServerInput) (*Server, error) {
	if input.AutoTLS && input.Hostname == "" {
		return nil, errors.New("Hostname must be set if using AutoTLS")
	}
	if input.Port == 0 {
		return nil, errors.New("Port must be set")
	}
	if input.Logger == nil {
		return nil, errors.New("Logger must be set")
	}
	return &Server{
		Port:     input.Port,
		AutoTLS:  input.AutoTLS,
		Hostname: input.Hostname,
		Logger:   input.Logger,
	}, nil
}

// Run starts the api server, it will block until the server is stopped.
func (s *Server) Run() {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	s.Logger.Infof("Starting server on port %d", s.Port)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", s.Port)))
	return
}
