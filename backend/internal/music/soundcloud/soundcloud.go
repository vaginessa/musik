package soundcloud

import (
	"net/http"

	"github.com/imthaghost/musik/backend/internal/logger"
)

// SoundCloud struct
type SoundCloud struct {
	Client *http.Client
	Logger logger.Service
}

// NewService function
func NewService(l logger.Service) *SoundCloud {
	return &SoundCloud{
		Client: &http.Client{},
		Logger: l,
	}
}
