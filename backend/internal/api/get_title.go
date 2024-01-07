package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetTitle - 200 OK
func (s *Server) GetTitle(c echo.Context) error {
	// get url from query param
	url := c.QueryParam("url")

	// get title from music service
	title, err := s.MusicService.GetSongTitle(c.Request().Context(), url)
	if err != nil {
		s.ErrorMonitor.Report(err)
		return echo.ErrInternalServerError
	}

	// return title
	return c.JSON(http.StatusOK, title)
}
