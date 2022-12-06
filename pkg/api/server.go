package api

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/TinyMonitor/pkg/db"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
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
	// User API
	e.GET("/api/v1/user/:id", s.getUser, s.userAuthRequired)
	e.GET("/api/v1/user", s.listUsers, s.userAuthRequired)
	e.DELETE("/api/v1/user/:id", s.deleteUser, s.userAuthRequired, s.adminRequired)
	e.POST("/api/v1/user", s.createUser, s.userAuthRequired, s.adminRequired)
	e.PATCH("/api/v1/user/:id", s.updateUser, s.userAuthRequired, s.adminRequired)
	s.logger.Infof("Starting server on port %d", s.port)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", s.port)))
	return
}
