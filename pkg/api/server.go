package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/TinyMonitor/pkg/db"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type Server struct {
	port         int    // Port to listen on
	autoTLS      bool   // If true, use Let's Encrypt to automatically get a TLS certificate
	hostname     string // Hostname must be set if using AutoTLS, also used when sending email alerts for callback URLs
	logger       *logrus.Logger
	dbConnection *db.DatabaseConnection
	memoryCache  *cache.Cache
}

// NewServerInput is the input for creating a Server.
type NewServerInput struct {
	Port         int            // Port to listen on
	AutoTLS      bool           // If true, use Let's Encrypt to automatically get a TLS certificate
	Hostname     string         // Hostname must be set if using AutoTLS, also used when sending email alerts for callback URLs
	Logger       *logrus.Logger // Logger to use for logging
	DBConnection *db.DatabaseConnection
}

// NewServer creates a new api server.
func NewServer(input NewServerInput) (*Server, error) {
	if input.AutoTLS && input.Hostname == "" {
		return nil, errors.New("Hostname must be set if using AutoTLS")
	}
	if input.DBConnection == nil {
		return nil, errors.New("DBConnection must be set")
	}
	if input.Port == 0 {
		return nil, errors.New("Port must be set")
	}
	if input.Logger == nil {
		return nil, errors.New("Logger must be set")
	}
	return &Server{
		port:         input.Port,
		autoTLS:      input.AutoTLS,
		hostname:     input.Hostname,
		logger:       input.Logger,
		dbConnection: input.DBConnection,
		memoryCache:  cache.New(5*time.Minute, 10*time.Minute),
	}, nil
}

// Run starts the api server, it will block until the server is stopped.
func (s *Server) Run() {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	// Monitor API
	e.GET("/api/v1/monitor", s.listMonitors, s.userAuthRequired)
	e.GET("/api/v1/monitor/:id", s.getMonitor, s.userAuthRequired)
	e.POST("/api/v1/monitor", s.createMonitor, s.userAuthRequired, s.writeRequired)
	e.PATCH("/api/v1/monitor/:id", s.updateMonitor, s.userAuthRequired, s.writeRequired)
	e.DELETE("/api/v1/monitor/:id", s.deleteMonitor, s.userAuthRequired, s.writeRequired)

	// User API
	e.GET("/api/v1/user/:id", s.getUser, s.userAuthRequired)
	e.GET("/api/v1/user", s.listUsers, s.userAuthRequired)
	e.DELETE("/api/v1/user/:id", s.deleteUser, s.userAuthRequired, s.adminRequired)
	e.POST("/api/v1/user", s.createUser, s.userAuthRequired, s.adminRequired)
	e.PATCH("/api/v1/user/:id", s.updateUser, s.userAuthRequired, s.adminRequired)
	// Utility routes
	e.GET("/api/v1/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})
	s.logger.Infof("Starting server on port %d", s.port)
	if os.Getenv("TINYMONITOR_TESTING") == "true" {
		s.logger.Warn("Running in testing mode")
	}
	go s.initialStartMonitors()
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", s.port)))
	return
}

func (s *Server) initialStartMonitors() {
	// Start initial monitors
	monitors, err := s.dbConnection.ListMonitors(context.Background(), db.ListMonitorOptions{})
	if err != nil {
		s.logger.Fatalf("Error listing monitors: %v", err)
	}
	for _, monitor := range monitors {
		time.Sleep(1 * time.Second)
		go performMonitoringChecks(monitor.Name, s.dbConnection, s.logger)
	}
}
