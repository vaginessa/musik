package music

import "context"

// Service is the interface that provides music methods.
type Service interface {
	GetArtworkURL(ctx context.Context, url string) (string, error)
	GetSongTitle(ctx context.Context, url string) (string, error)
	GetStreamURL(ctx context.Context, url string) (string, error)
}
