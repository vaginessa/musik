package api

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetStreamURL - 200 OK
func (s *Server) GetStreamURL(c echo.Context) error {
	// get url from query param
	url := c.QueryParam("url")

	// get stream url from music service
	streamURL, err := s.MusicService.GetStreamURL(context.Background(), url)
	if err != nil {
		s.ErrorMonitor.Report(err)
		return echo.ErrInternalServerError
	}

	// return stream url
	return c.JSON(http.StatusOK, streamURL)
}
