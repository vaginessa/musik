package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetArtworkURL - 200 OK
func (s *Server) GetArtworkURL(c echo.Context) error {
	artworkURL, err := s.MusicService.GetArtworkURL(c.Request().Context(), c.QueryParam("url"))
	if err != nil {
		s.ErrorMonitor.Report(err)
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, artworkURL)
}
