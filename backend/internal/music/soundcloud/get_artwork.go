package soundcloud

import (
	"context"
	"fmt"
	"log"

	"github.com/antchfx/htmlquery"
)

// GetArtworkURL returns the artwork for a given SoundCloud ID
func (s *SoundCloud) GetArtworkURL(ctx context.Context, url string) (string, error) {
	resp, err := s.Client.Get(url)

	log.Println("URL:", url)
	if err != nil {
		return "", err
	}

	// parse html
	doc, err := htmlquery.Parse(resp.Body)
	if err != nil {
		log.Println(err)
	}

	// artworkXPath
	artworkPath := "//meta[@property='og:image']/@content"

	// Query the document for the artwork node
	nodes, err := htmlquery.QueryAll(doc, artworkPath)
	if err != nil {
		fmt.Println("Error executing XPath query:", err)
		return "", err
	}

	// Check if any nodes were found
	if len(nodes) > 0 {
		// Extract the content from the first node
		artwork := htmlquery.InnerText(nodes[0])

		return artwork, nil
	}

	return "", nil
}
