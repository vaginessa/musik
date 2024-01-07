package soundcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/k0kubun/pp/v3"
	"io"
	"math/rand"
	"net/url"
	"strings"
	"time"
)

// 1. https://api-v2.soundcloud.com/users/1319568996/playlists_without_albums?offset=10&limit=10&client_id=nUB9ZvnjRiqKF43CkKf3iu69D8bboyKY&app_version=1704451463&app_locale=en
// 2. https://api-v2.soundcloud.com/playlists/1714689219?representation=full&client_id=nUB9ZvnjRiqKF43CkKf3iu69D8bboyKY&app_version=1704451463&app_locale=en
// 3.https://api-v2.soundcloud.com/tracks?ids=1650758970%2C1650759159%2C1685290914%2C1686057453%2C1689877233%2C1700527512%2C1702996545%2C1702998747%2C1703000283%2C1704960288%2C1704961347%2C1704961983%2C1704962586%2C1704963903%2C1704964458&client_id=nUB9ZvnjRiqKF43CkKf3iu69D8bboyKY&%5Bobject%20Object%5D=&app_version=1704451463&app_locale=en

// GetRandomSongURL is a method on SoundCloud that returns a random song URL.
func (s *SoundCloud) GetRandomSongURL(ctx context.Context) (string, error) {
	rand.Seed(time.Now().UnixNano())

	// get client id
	clientID, err := s.GetClientID()
	if err != nil {
		return "", err
	}

	playlistURL := fmt.Sprintf("https://api-v2.soundcloud.com/users/1319568996/playlists_without_albums?offset=10&limit=10&client_id=%s&app_version=1704451463&app_locale=en", clientID)

	resp, err := s.Client.Get(playlistURL)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var playlistResponse PlaylistResponse
	err = json.Unmarshal(body, &playlistResponse)
	if err != nil {
		return "", err
	}

	randomPlaylist := rand.Intn(len(playlistResponse.Collection)-1-0+1) + 0

	pp.Println("Random Number: ", randomPlaylist)
	pp.Println(playlistResponse.Collection[randomPlaylist].Title)
	pp.Println(playlistResponse.Collection[randomPlaylist].PermalinkURL)
	pp.Println(playlistResponse.Collection[randomPlaylist].ID)

	playlistID := playlistResponse.Collection[randomPlaylist].ID

	albumURL := fmt.Sprintf("https://api-v2.soundcloud.com/playlists/%d?representation=full&client_id=%s&app_version=1704451463&app_locale=en", playlistID, clientID)

	pp.Println("Album URL: ", albumURL)

	resp, err = s.Client.Get(albumURL)

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var albumResponse AlbumResponse
	err = json.Unmarshal(body, &albumResponse)
	if err != nil {
		return "", err
	}

	pp.Println("Artwork URL: ", albumResponse.ArtworkURL)
	pp.Println("Title: ", albumResponse.Title)
	pp.Println("TrackCount: ", albumResponse.TrackCount)
	pp.Println("Track: ", albumResponse.Tracks[0].Title)
	pp.Println("Track ID: ", albumResponse.Tracks[0].ID)

	baseTrackURL := "https://api-v2.soundcloud.com/tracks"
	appVersion := "1704451463"
	appLocale := "en"

	// Convert the track IDs to a comma-separated string
	var sb strings.Builder
	for i, track := range albumResponse.Tracks {
		sb.WriteString(fmt.Sprintf("%d", track.ID))
		if i < len(albumResponse.Tracks)-1 {
			sb.WriteString(",")
		}
	}
	trackIDsString := sb.String()

	// Encode the parameters
	values := url.Values{}
	values.Add("ids", trackIDsString)
	values.Add("client_id", clientID)
	values.Add("app_version", appVersion)
	values.Add("app_locale", appLocale)

	tracksURL := fmt.Sprintf("%s?%s", baseTrackURL, values.Encode())

	resp, err = s.Client.Get(tracksURL)

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var trackResponse TrackResponse
	err = json.Unmarshal(body, &trackResponse)
	if err != nil {
		return "", err
	}

	pp.Println("Track Artwork: ", trackResponse[0].ArtworkURL)
	pp.Println("Track Title: ", trackResponse[0].Title)
	pp.Println("Track PermalinkURL: ", trackResponse[0].PermalinkURL)
	//pp.Println("Track Permalink: ", trackResponse[0].Permalink)
	//pp.Println("Is Public: ", trackResponse[0].Public)
	//pp.Println("Is Streamable: ", trackResponse[0].Streamable)
	//pp.Println("Track Stream URL: ", trackResponse[0].Media.Transcodings[0].URL)
	//pp.Println("Track Authorization: ", trackResponse[0].TrackAuthorization)
	//pp.Println("Stream Token: ", trackResponse[0].SecretToken)

	randomTrackURL := trackResponse[rand.Intn(len(trackResponse)-1-0+1)+0].PermalinkURL

	return randomTrackURL, nil

}
