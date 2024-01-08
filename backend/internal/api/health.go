package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HealthCheck - 200 OK
func (s *Server) HealthCheck(c echo.Context) error {
	s.LoggerService.Msg("Health check")
	return c.NoContent(http.StatusOK)
}
