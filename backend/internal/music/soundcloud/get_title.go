package soundcloud

import (
	"context"
	"fmt"
	"log"

	"github.com/antchfx/htmlquery"
)

// GetSongTitle returns the title of the video
func (s *SoundCloud) GetSongTitle(ctx context.Context, url string) (string, error) {
	resp, err := s.Client.Get(url)
	if err != nil {
		return "", err
	}

	// parse html
	doc, err := htmlquery.Parse(resp.Body)
	if err != nil {
		log.Println(err)
	}

	// XPath query
	titlePath := "//meta[@property='og:title']/@content"

	// Query the document for the title node
	nodes, err := htmlquery.QueryAll(doc, titlePath)
	if err != nil {
		fmt.Println("Error executing XPath query:", err)
		return "", err
	}

	// Check if any nodes were found
	if len(nodes) > 0 {
		// Extract the content from the first node
		title := htmlquery.InnerText(nodes[0])
		return title, nil
	}

	return "", nil
}
