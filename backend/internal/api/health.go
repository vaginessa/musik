package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HealthCheck - 200 OK
func (s *Server) HealthCheck(c echo.Context) error {

	return c.NoContent(http.StatusOK)
}
