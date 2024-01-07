package soundcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/antchfx/htmlquery"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// AudioLink struct for unmarshalling data
type AudioLink struct {
	URL string `json:"url"`
}

// GetStreamURL function
func (s *SoundCloud) GetStreamURL(ctx context.Context, url string) (string, error) {
	resp, err := s.Client.Get(url)
	if err != nil {
		return "", err
	}

	// parse html
	doc, err := htmlquery.Parse(resp.Body)
	if err != nil {
		log.Println(err)
	}

	// get client id
	clientID, err := s.GetClientID()
	if err != nil {
		return "", err
	}

	// get track authorization
	trackAuth, err := s.GetTrackAuthorization(doc)
	if err != nil {
		return "", err
	}

	// get hls stream url
	hlsStreamURL, err := s.GetHLSURL(doc)
	if err != nil {
		return "", err
	}
	log.Println("Track Auth: ", trackAuth)
	log.Println("Client ID: ", clientID)
	log.Println("HLS Stream URL: ", hlsStreamURL)

	trackID, streamToken, err := getTrackInfo(hlsStreamURL)

	// construct stream url
	baseURL := "https://api-v2.soundcloud.com/media/soundcloud:tracks:%s/%s/stream/hls?client_id=%s&track_authorization=%s"

	streamURL := fmt.Sprintf(baseURL, trackID, streamToken, clientID, trackAuth)

	// Get the response from the URL
	streamResp, err := http.Get(streamURL)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer streamResp.Body.Close()

	// Read the body of the response
	body, err := ioutil.ReadAll(streamResp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return "", err
	}

	// Unmarshal the JSON into the struct
	var audioResp AudioLink
	err = json.Unmarshal(body, &audioResp)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return "", err
	}
	log.Println("Track ID: ", trackID)
	log.Println("Stream Token: ", streamToken)

	log.Println("Stream URL: ", streamURL)

	return audioResp.URL, nil
}

func getTrackInfo(url string) (string, string, error) {
	// Split the URL by '/'
	parts := strings.Split(url, "/")

	// Check if the parts length is as expected
	if len(parts) < 8 {
		return "", "", fmt.Errorf("invalid URL format")
	}

	// Extract the part that contains 'soundcloud:tracks:TRACK_ID'
	trackIDPart := parts[4] // "soundcloud:tracks:373180994"

	// Further split the track ID part to get the actual track ID
	trackIDParts := strings.Split(trackIDPart, ":")
	if len(trackIDParts) < 3 {
		return "", "", fmt.Errorf("invalid track ID format in URL")
	}
	trackID := trackIDParts[2]

	// The stream token is the next part of the URL
	streamToken := parts[5] // "ca04c4eb-a299-4f4b-9ff1-ac20ed014fe5"

	return trackID, streamToken, nil
}
