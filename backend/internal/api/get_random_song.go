package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type RandomSongResponse struct {
	StreamURL  string `json:"stream_url"`
	ArtworkURL string `json:"artwork_url"`
	SongTitle  string `json:"song_title"`
	TrackURL   string `json:"track_url"`
}

// GetRandomSong - 200 OK
func (s *Server) GetRandomSong(c echo.Context) error {
	// get random song and its details
	trackURL, artworkURL, songTitle, streamURL, err := s.MusicService.GetRandomSong(c.Request().Context())
	if err != nil {
		s.ErrorMonitor.Report(err)
		return echo.ErrInternalServerError
	}

	// create response
	song := RandomSongResponse{
		StreamURL:  streamURL,
		ArtworkURL: artworkURL,
		SongTitle:  songTitle,
		TrackURL:   trackURL,
	}

	return c.JSON(http.StatusOK, song)
}
