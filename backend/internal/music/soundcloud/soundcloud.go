package soundcloud

import "net/http"

// SoundCloud struct
type SoundCloud struct {
	Client *http.Client
}

// NewService function
func NewService() *SoundCloud {
	return &SoundCloud{
		Client: &http.Client{},
	}
}
