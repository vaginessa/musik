package api

import (
	"context"
	"github.com/imthaghost/musik/backend/internal/music"

	"github.com/labstack/echo/v4"

	"github.com/imthaghost/musik/backend/config"
	"github.com/imthaghost/musik/backend/internal/errors"
	"github.com/imthaghost/musik/backend/internal/logger"
)

// Server represents an API wrapper around the SoundCloud API
type Server struct {
	// Echo server
	Echo *echo.Echo
	// Config
	Config config.Config
	// Logger Service
	LoggerService logger.Service
	// ErrorMonitor service
	ErrorMonitor errors.Service
	// Music service
	MusicService music.Service
}

// NewServer returns a new server instance
func NewServer(cfg config.Config, log logger.Service, errorMonitor errors.Service, musicService music.Service) *Server {
	return &Server{
		Echo:          echo.New(),
		Config:        cfg,
		LoggerService: log,
		ErrorMonitor:  errorMonitor,
		MusicService:  musicService,
	}
}

// Start initializes the HTTP server
func (s *Server) Start() error {
	env := s.Config.General.AppEnv
	server := s.Echo

	// hide banner
	if env != "dev" {
		server.HideBanner = true
	}

	// register routes
	s.Routes()

	// start server
	err := server.Start(":" + "8080")
	if err != nil {

		return err
	}

	return nil
}

// Close gracefully shuts down server and closes all connections
func (s *Server) Close(ctx context.Context) error {

	// stop sentry
	s.ErrorMonitor.TearDown()

	return s.Echo.Shutdown(ctx)
}
