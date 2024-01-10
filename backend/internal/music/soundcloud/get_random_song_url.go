package soundcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/url"
	"regexp"
	"strings"
	"time"
)

// 1. https://api-v2.soundcloud.com/stream/users/1319605599?client_id=ltFvoc9UP9uzxTZwn95RphwV0n7zbbNN&limit=20&offset=0&linked_partitioning=1&app_version=1704798544&app_locale=en
// 2. https://api-v2.soundcloud.com/playlists/1714689219?representation=full&client_id=nUB9ZvnjRiqKF43CkKf3iu69D8bboyKY&app_version=1704451463&app_locale=en
// 3.https://api-v2.soundcloud.com/tracks?ids=1650758970%2C1650759159%2C1685290914%2C1686057453%2C1689877233%2C1700527512%2C1702996545%2C1702998747%2C1703000283%2C1704960288%2C1704961347%2C1704961983%2C1704962586%2C1704963903%2C1704964458&client_id=nUB9ZvnjRiqKF43CkKf3iu69D8bboyKY&%5Bobject%20Object%5D=&app_version=1704451463&app_locale=en

// GetRandomSong GetRandomSongURL is a method on SoundCloud that returns a random song and its details.
func (s *SoundCloud) GetRandomSong(ctx context.Context) (string, string, string, string, error) {

	// get client id
	clientID, err := s.GetClientID()
	if err != nil {
		return "", "", "", "", err
	}

	// TODO: add more users
	usersIDs := []string{"1319605599", "1319568996", "1284333535", "193", "603451386"}

	userURL := fmt.Sprintf("https://api-v2.soundcloud.com/stream/users/%s?client_id=%s&limit=20&offset=0&linked_partitioning=1&app_version=1704798544&app_locale=en", usersIDs[generateRandomNumber(0, len(usersIDs)-1)], clientID)

	// get user's playlists
	resp, err := s.Client.Get(userURL)
	if err != nil {
		return "", "", "", "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", "", "", err
	}

	// unmarshal user response
	var userResponse UserResponse
	err = json.Unmarshal(body, &userResponse)
	if err != nil {
		return "", "", "", "", err
	}

	for _, playlist := range userResponse.Collection {
		log.Println(playlist.Playlist.ID)
	}
	playlistID := userResponse.Collection[generateRandomNumber(0, len(userResponse.Collection)-1)].Playlist.ID
	if playlistID == 0 {
		log.Println("Playlist ID is 0, getting another playlist...")
		for {
			playlistID = userResponse.Collection[generateRandomNumber(0, len(userResponse.Collection)-1)].Playlist.ID
			if playlistID != 0 {
				break
			}
		}
	}
	// generate random playlist id until we find one that doesn't have a 0 id

	albumURL := fmt.Sprintf("https://api-v2.soundcloud.com/playlists/%d?representation=full&client_id=%s&app_version=1704451463&app_locale=en", playlistID, clientID)

	// get album
	resp, err = s.Client.Get(albumURL)
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		return "", "", "", "", err
	}

	log.Println("Getting album from albumURL: ", albumURL)

	// unmarshal album response
	var albumResponse AlbumResponse
	err = json.Unmarshal(body, &albumResponse)
	if err != nil {
		log.Print(err)
		return "", "", "", "", err
	}

	// if the track length is greater than 50, we need to remove the difference
	// if the track length is greater than 50, we need to limit it to 50
	if len(albumResponse.Tracks) > 50 {
		// limiting the slice to the first 50 elements
		albumResponse.Tracks = albumResponse.Tracks[:50]
	}

	baseTrackURL := "https://api-v2.soundcloud.com/tracks"
	appVersion := "1704451463"
	appLocale := "en"

	// convert the track IDs to a comma-separated string
	var sb strings.Builder
	for i, track := range albumResponse.Tracks {
		sb.WriteString(fmt.Sprintf("%d", track.ID))
		if i < len(albumResponse.Tracks)-1 {
			sb.WriteString(",")
		}
	}

	trackIDsString := sb.String()

	// encode the parameters
	values := url.Values{}
	values.Add("ids", trackIDsString)
	values.Add("client_id", clientID)
	values.Add("app_version", appVersion)
	values.Add("app_locale", appLocale)

	tracksURL := fmt.Sprintf("%s?%s", baseTrackURL, values.Encode())

	// get tracks
	resp, err = s.Client.Get(tracksURL)
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", "", "", "", err
	}

	log.Println("Getting tracks from tracksURL: ", tracksURL)

	// unmarshal tracks response
	var trackResponse TrackResponse
	err = json.Unmarshal(body, &trackResponse)
	if err != nil {
		log.Println(err)
		return "", "", "", "", err
	}

	//var track string
	// get random track
	track := trackResponse[generateRandomNumber(0, len(trackResponse)-1)]
	if strings.Contains(track.Media.Transcodings[0].URL, "preview") {
		s.Logger.Msg("Track is a preview, getting another track...")
		// Keep picking a random track until we find one without "preview" in the URL
		for {
			track = trackResponse[generateRandomNumber(0, len(trackResponse)-1)]
			if !strings.Contains(track.Media.Transcodings[0].URL, "preview") {
				break
			}
		}
	}

	log.Println("Track Stream URL: ", track.Media.Transcodings[0].URL)

	// extract track id and stream token
	trackID, streamToken, err := extractStreamTokenAndTrackID(track.Media.Transcodings[0].URL)
	if err != nil {
		log.Println(err)
		return "", "", "", "", err
	}

	// create stream url
	streamReq := fmt.Sprintf("https://api-v2.soundcloud.com/media/soundcloud:tracks:%s/%s/stream/hls?client_id=%s&track_authorization=%s", trackID, streamToken, clientID, track.TrackAuthorization)

	resp, err = s.Client.Get(streamReq)
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return "", "", "", "", err
	}

	// unmarshal stream response
	var streamResponse StreamResponse
	err = json.Unmarshal(body, &streamResponse)
	if err != nil {
		return "", "", "", "", err
	}

	streamURL := streamResponse.URL

	// logging information for dev environment
	if s.Config.General.AppEnv == "dev" {
		//s.Logger.Msg(fmt.Sprintf("Track Title: %s", track.Title))
		//s.Logger.Msg(fmt.Sprintf("Track Artwork: %s", track.ArtworkURL))
		//s.Logger.Msg(fmt.Sprintf("Track PermalinkURL: %s ", track.PermalinkURL))
		//s.Logger.Msg(fmt.Sprintf("Is Streamable: %t", track.Streamable))
		//s.Logger.Msg(fmt.Sprintf("Monetization Model: %s", track.MonetizationModel))
		//s.Logger.Msg(fmt.Sprintf("Streaming URL: %s", track.Media.Transcodings[0].URL))
		//s.Logger.Msg(fmt.Sprintf("Track Authorization Token: %s", track.TrackAuthorization))
		//s.Logger.Msg(fmt.Sprintf("Stream URL: %s", streamURL))
	}

	// return track url, artwork url, track title, and stream url
	return track.PermalinkURL, track.ArtworkURL, track.Title, streamURL, nil

}

// generateRandomNumber is a helper function that generates a random number.
func generateRandomNumber(min, max int) int {
	if min >= max {
		return 0
	}

	// seed the random number generator.
	rand.NewSource(time.Now().UnixNano())

	// generate a random number in the range [min, max].
	randNum := rand.Intn(max-min+1) + min

	return randNum
}

// extractStreamTokenAndTrackID is a helper function that extracts the stream token and track ID from the stream URL.
func extractStreamTokenAndTrackID(url string) (string, string, error) {
	// define a regular expression pattern to extract track ID and stream token
	re := regexp.MustCompile(`tracks:(\d+)/([a-zA-Z0-9-]+)/stream/hls`)

	// find match
	matches := re.FindStringSubmatch(url)
	if matches == nil || len(matches) < 3 {
		return "", "", fmt.Errorf("no matches found in the URL")
	}

	trackID := matches[1]
	streamToken := matches[2]

	// matches[1] will be the track ID, and matches[2] will be the stream token
	return trackID, streamToken, nil
}
