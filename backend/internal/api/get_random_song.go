package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetRandomSong - 200 OK
func (s *Server) GetRandomSong(c echo.Context) error {
	song, err := s.MusicService.GetRandomSongURL(c.Request().Context())
	if err != nil {
		s.ErrorMonitor.Report(err)
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, song)
}
