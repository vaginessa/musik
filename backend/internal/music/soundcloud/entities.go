package soundcloud

import "time"

// PlaylistResponse is the response from the SoundCloud API
type PlaylistResponse struct {
	Collection []struct {
		ArtworkURL     string    `json:"artwork_url"`
		CreatedAt      time.Time `json:"created_at"`
		Description    string    `json:"description"`
		Duration       int       `json:"duration"`
		EmbeddableBy   string    `json:"embeddable_by"`
		Genre          string    `json:"genre"`
		ID             int       `json:"id"`
		Kind           string    `json:"kind"`
		LabelName      string    `json:"label_name"`
		LastModified   time.Time `json:"last_modified"`
		License        string    `json:"license"`
		LikesCount     int       `json:"likes_count"`
		ManagedByFeeds bool      `json:"managed_by_feeds"`
		Permalink      string    `json:"permalink"`
		PermalinkURL   string    `json:"permalink_url"`
		Public         bool      `json:"public"`
		PurchaseTitle  any       `json:"purchase_title"`
		PurchaseURL    any       `json:"purchase_url"`
		ReleaseDate    any       `json:"release_date"`
		RepostsCount   int       `json:"reposts_count"`
		SecretToken    any       `json:"secret_token"`
		Sharing        string    `json:"sharing"`
		TagList        string    `json:"tag_list"`
		Title          string    `json:"title"`
		URI            string    `json:"uri"`
		UserID         int       `json:"user_id"`
		SetType        string    `json:"set_type"`
		IsAlbum        bool      `json:"is_album"`
		PublishedAt    time.Time `json:"published_at"`
		DisplayDate    time.Time `json:"display_date"`
		//User           struct {
		//	AvatarURL      string    `json:"avatar_url"`
		//	FirstName      string    `json:"first_name"`
		//	FollowersCount int       `json:"followers_count"`
		//	FullName       string    `json:"full_name"`
		//	ID             int       `json:"id"`
		//	Kind           string    `json:"kind"`
		//	LastModified   time.Time `json:"last_modified"`
		//	LastName       string    `json:"last_name"`
		//	Permalink      string    `json:"permalink"`
		//	PermalinkURL   string    `json:"permalink_url"`
		//	URI            string    `json:"uri"`
		//	Urn            string    `json:"urn"`
		//	Username       string    `json:"username"`
		//	Verified       bool      `json:"verified"`
		//	City           string    `json:"city"`
		//	CountryCode    string    `json:"country_code"`
		//	Badges         struct {
		//		Pro          bool `json:"pro"`
		//		ProUnlimited bool `json:"pro_unlimited"`
		//		Verified     bool `json:"verified"`
		//	} `json:"badges"`
		//	StationUrn       string `json:"station_urn"`
		//	StationPermalink string `json:"station_permalink"`
		//} `json:"user"`
		//Tracks []struct {
		//	ArtworkURL        string    `json:"artwork_url,omitempty"`
		//	Caption           any       `json:"caption,omitempty"`
		//	Commentable       bool      `json:"commentable,omitempty"`
		//	CommentCount      int       `json:"comment_count,omitempty"`
		//	CreatedAt         time.Time `json:"created_at,omitempty"`
		//	Description       string    `json:"description,omitempty"`
		//	Downloadable      bool      `json:"downloadable,omitempty"`
		//	DownloadCount     int       `json:"download_count,omitempty"`
		//	Duration          int       `json:"duration,omitempty"`
		//	FullDuration      int       `json:"full_duration,omitempty"`
		//	EmbeddableBy      string    `json:"embeddable_by,omitempty"`
		//	Genre             string    `json:"genre,omitempty"`
		//	HasDownloadsLeft  bool      `json:"has_downloads_left,omitempty"`
		//	ID                int       `json:"id"`
		//	Kind              string    `json:"kind"`
		//	LabelName         any       `json:"label_name,omitempty"`
		//	LastModified      time.Time `json:"last_modified,omitempty"`
		//	License           string    `json:"license,omitempty"`
		//	LikesCount        int       `json:"likes_count,omitempty"`
		//	Permalink         string    `json:"permalink,omitempty"`
		//	PermalinkURL      string    `json:"permalink_url,omitempty"`
		//	PlaybackCount     int       `json:"playback_count,omitempty"`
		//	Public            bool      `json:"public,omitempty"`
		//	PublisherMetadata struct {
		//		ID             int    `json:"id"`
		//		Urn            string `json:"urn"`
		//		Artist         string `json:"artist"`
		//		ContainsMusic  bool   `json:"contains_music"`
		//		Isrc           string `json:"isrc"`
		//		Explicit       bool   `json:"explicit"`
		//		WriterComposer string `json:"writer_composer"`
		//	} `json:"publisher_metadata,omitempty"`
		//	PurchaseTitle any       `json:"purchase_title,omitempty"`
		//	PurchaseURL   any       `json:"purchase_url,omitempty"`
		//	ReleaseDate   any       `json:"release_date,omitempty"`
		//	RepostsCount  int       `json:"reposts_count,omitempty"`
		//	SecretToken   any       `json:"secret_token,omitempty"`
		//	Sharing       string    `json:"sharing,omitempty"`
		//	State         string    `json:"state,omitempty"`
		//	Streamable    bool      `json:"streamable,omitempty"`
		//	TagList       string    `json:"tag_list,omitempty"`
		//	Title         string    `json:"title,omitempty"`
		//	TrackFormat   string    `json:"track_format,omitempty"`
		//	URI           string    `json:"uri,omitempty"`
		//	Urn           string    `json:"urn,omitempty"`
		//	UserID        int       `json:"user_id,omitempty"`
		//	Visuals       any       `json:"visuals,omitempty"`
		//	WaveformURL   string    `json:"waveform_url,omitempty"`
		//	DisplayDate   time.Time `json:"display_date,omitempty"`
		//	Media         struct {
		//		Transcodings []struct {
		//			URL      string `json:"url"`
		//			Preset   string `json:"preset"`
		//			Duration int    `json:"duration"`
		//			Snipped  bool   `json:"snipped"`
		//			Format   struct {
		//				Protocol string `json:"protocol"`
		//				MimeType string `json:"mime_type"`
		//			} `json:"format"`
		//			Quality string `json:"quality"`
		//		} `json:"transcodings"`
		//	} `json:"media,omitempty"`
		//	StationUrn         string `json:"station_urn,omitempty"`
		//	StationPermalink   string `json:"station_permalink,omitempty"`
		//	TrackAuthorization string `json:"track_authorization,omitempty"`
		//	MonetizationModel  string `json:"monetization_model"`
		//	Policy             string `json:"policy"`
		//	User               struct {
		//		AvatarURL      string    `json:"avatar_url"`
		//		FirstName      string    `json:"first_name"`
		//		FollowersCount int       `json:"followers_count"`
		//		FullName       string    `json:"full_name"`
		//		ID             int       `json:"id"`
		//		Kind           string    `json:"kind"`
		//		LastModified   time.Time `json:"last_modified"`
		//		LastName       string    `json:"last_name"`
		//		Permalink      string    `json:"permalink"`
		//		PermalinkURL   string    `json:"permalink_url"`
		//		URI            string    `json:"uri"`
		//		Urn            string    `json:"urn"`
		//		Username       string    `json:"username"`
		//		Verified       bool      `json:"verified"`
		//		City           string    `json:"city"`
		//		CountryCode    string    `json:"country_code"`
		//		Badges         struct {
		//			Pro          bool `json:"pro"`
		//			ProUnlimited bool `json:"pro_unlimited"`
		//			Verified     bool `json:"verified"`
		//		} `json:"badges"`
		//		StationUrn       string `json:"station_urn"`
		//		StationPermalink string `json:"station_permalink"`
		//	} `json:"user,omitempty"`
		//	PublisherMetadata0 struct {
		//		ID              int    `json:"id"`
		//		Urn             string `json:"urn"`
		//		Artist          string `json:"artist"`
		//		AlbumTitle      string `json:"album_title"`
		//		ContainsMusic   bool   `json:"contains_music"`
		//		UpcOrEan        string `json:"upc_or_ean"`
		//		Isrc            string `json:"isrc"`
		//		Explicit        bool   `json:"explicit"`
		//		PLine           string `json:"p_line"`
		//		PLineForDisplay string `json:"p_line_for_display"`
		//		WriterComposer  string `json:"writer_composer"`
		//		ReleaseTitle    string `json:"release_title"`
		//	} `json:"publisher_metadata,omitempty"`
		//	PublisherMetadata1 struct {
		//		ID              int    `json:"id"`
		//		Urn             string `json:"urn"`
		//		Artist          string `json:"artist"`
		//		AlbumTitle      string `json:"album_title"`
		//		ContainsMusic   bool   `json:"contains_music"`
		//		UpcOrEan        string `json:"upc_or_ean"`
		//		Isrc            string `json:"isrc"`
		//		Explicit        bool   `json:"explicit"`
		//		PLine           string `json:"p_line"`
		//		PLineForDisplay string `json:"p_line_for_display"`
		//		CLine           string `json:"c_line"`
		//		CLineForDisplay string `json:"c_line_for_display"`
		//		ReleaseTitle    string `json:"release_title"`
		//	} `json:"publisher_metadata,omitempty"`
		//	PublisherMetadata2 struct {
		//		ID              int    `json:"id"`
		//		Urn             string `json:"urn"`
		//		Artist          string `json:"artist"`
		//		AlbumTitle      string `json:"album_title"`
		//		ContainsMusic   bool   `json:"contains_music"`
		//		UpcOrEan        string `json:"upc_or_ean"`
		//		Isrc            string `json:"isrc"`
		//		Explicit        bool   `json:"explicit"`
		//		PLine           string `json:"p_line"`
		//		PLineForDisplay string `json:"p_line_for_display"`
		//		WriterComposer  string `json:"writer_composer"`
		//		ReleaseTitle    string `json:"release_title"`
		//	} `json:"publisher_metadata,omitempty"`
		//	PublisherMetadata3 struct {
		//		ID            int    `json:"id"`
		//		Urn           string `json:"urn"`
		//		Artist        string `json:"artist"`
		//		ContainsMusic bool   `json:"contains_music"`
		//		Isrc          string `json:"isrc"`
		//	} `json:"publisher_metadata,omitempty"`
		//} `json:"tracks"`
		TrackCount int `json:"track_count"`
	} `json:"collection"`
	NextHref any `json:"next_href"`
	QueryUrn any `json:"query_urn"`
}

// AlbumResponse is the response from the SoundCloud API
type AlbumResponse struct {
	ArtworkURL     string    `json:"artwork_url"`
	CreatedAt      time.Time `json:"created_at"`
	Description    string    `json:"description"`
	Duration       int       `json:"duration"`
	EmbeddableBy   string    `json:"embeddable_by"`
	Genre          string    `json:"genre"`
	ID             int       `json:"id"`
	Kind           string    `json:"kind"`
	LabelName      string    `json:"label_name"`
	LastModified   time.Time `json:"last_modified"`
	License        string    `json:"license"`
	LikesCount     int       `json:"likes_count"`
	ManagedByFeeds bool      `json:"managed_by_feeds"`
	Permalink      string    `json:"permalink"`
	PermalinkURL   string    `json:"permalink_url"`
	Public         bool      `json:"public"`
	PurchaseTitle  any       `json:"purchase_title"`
	PurchaseURL    any       `json:"purchase_url"`
	ReleaseDate    any       `json:"release_date"`
	RepostsCount   int       `json:"reposts_count"`
	SecretToken    any       `json:"secret_token"`
	Sharing        string    `json:"sharing"`
	TagList        string    `json:"tag_list"`
	Title          string    `json:"title"`
	URI            string    `json:"uri"`
	UserID         int       `json:"user_id"`
	SetType        string    `json:"set_type"`
	IsAlbum        bool      `json:"is_album"`
	PublishedAt    time.Time `json:"published_at"`
	DisplayDate    time.Time `json:"display_date"`
	User           struct {
		AvatarURL      string    `json:"avatar_url"`
		FirstName      string    `json:"first_name"`
		FollowersCount int       `json:"followers_count"`
		FullName       string    `json:"full_name"`
		ID             int       `json:"id"`
		Kind           string    `json:"kind"`
		LastModified   time.Time `json:"last_modified"`
		LastName       string    `json:"last_name"`
		Permalink      string    `json:"permalink"`
		PermalinkURL   string    `json:"permalink_url"`
		URI            string    `json:"uri"`
		Urn            string    `json:"urn"`
		Username       string    `json:"username"`
		Verified       bool      `json:"verified"`
		City           string    `json:"city"`
		CountryCode    string    `json:"country_code"`
		Badges         struct {
			Pro          bool `json:"pro"`
			ProUnlimited bool `json:"pro_unlimited"`
			Verified     bool `json:"verified"`
		} `json:"badges"`
		StationUrn       string `json:"station_urn"`
		StationPermalink string `json:"station_permalink"`
	} `json:"user"`
	Tracks []struct {
		ArtworkURL        string    `json:"artwork_url,omitempty"`
		Caption           any       `json:"caption,omitempty"`
		Commentable       bool      `json:"commentable,omitempty"`
		CommentCount      int       `json:"comment_count,omitempty"`
		CreatedAt         time.Time `json:"created_at,omitempty"`
		Description       string    `json:"description,omitempty"`
		Downloadable      bool      `json:"downloadable,omitempty"`
		DownloadCount     int       `json:"download_count,omitempty"`
		Duration          int       `json:"duration,omitempty"`
		FullDuration      int       `json:"full_duration,omitempty"`
		EmbeddableBy      string    `json:"embeddable_by,omitempty"`
		Genre             string    `json:"genre,omitempty"`
		HasDownloadsLeft  bool      `json:"has_downloads_left,omitempty"`
		ID                int       `json:"id"`
		Kind              string    `json:"kind"`
		LabelName         any       `json:"label_name,omitempty"`
		LastModified      time.Time `json:"last_modified,omitempty"`
		License           string    `json:"license,omitempty"`
		LikesCount        int       `json:"likes_count,omitempty"`
		Permalink         string    `json:"permalink,omitempty"`
		PermalinkURL      string    `json:"permalink_url,omitempty"`
		PlaybackCount     int       `json:"playback_count,omitempty"`
		Public            bool      `json:"public,omitempty"`
		PublisherMetadata struct {
			ID             int    `json:"id"`
			Urn            string `json:"urn"`
			Artist         string `json:"artist"`
			ContainsMusic  bool   `json:"contains_music"`
			Isrc           string `json:"isrc"`
			Explicit       bool   `json:"explicit"`
			WriterComposer string `json:"writer_composer"`
		} `json:"publisher_metadata,omitempty"`
		PurchaseTitle any       `json:"purchase_title,omitempty"`
		PurchaseURL   any       `json:"purchase_url,omitempty"`
		ReleaseDate   any       `json:"release_date,omitempty"`
		RepostsCount  int       `json:"reposts_count,omitempty"`
		SecretToken   any       `json:"secret_token,omitempty"`
		Sharing       string    `json:"sharing,omitempty"`
		State         string    `json:"state,omitempty"`
		Streamable    bool      `json:"streamable,omitempty"`
		TagList       string    `json:"tag_list,omitempty"`
		Title         string    `json:"title,omitempty"`
		TrackFormat   string    `json:"track_format,omitempty"`
		URI           string    `json:"uri,omitempty"`
		Urn           string    `json:"urn,omitempty"`
		UserID        int       `json:"user_id,omitempty"`
		Visuals       any       `json:"visuals,omitempty"`
		WaveformURL   string    `json:"waveform_url,omitempty"`
		DisplayDate   time.Time `json:"display_date,omitempty"`
		Media         struct {
			Transcodings []struct {
				URL      string `json:"url"`
				Preset   string `json:"preset"`
				Duration int    `json:"duration"`
				Snipped  bool   `json:"snipped"`
				Format   struct {
					Protocol string `json:"protocol"`
					MimeType string `json:"mime_type"`
				} `json:"format"`
				Quality string `json:"quality"`
			} `json:"transcodings"`
		} `json:"media,omitempty"`
		StationUrn         string `json:"station_urn,omitempty"`
		StationPermalink   string `json:"station_permalink,omitempty"`
		TrackAuthorization string `json:"track_authorization,omitempty"`
		MonetizationModel  string `json:"monetization_model"`
		Policy             string `json:"policy"`
		User               struct {
			AvatarURL      string    `json:"avatar_url"`
			FirstName      string    `json:"first_name"`
			FollowersCount int       `json:"followers_count"`
			FullName       string    `json:"full_name"`
			ID             int       `json:"id"`
			Kind           string    `json:"kind"`
			LastModified   time.Time `json:"last_modified"`
			LastName       string    `json:"last_name"`
			Permalink      string    `json:"permalink"`
			PermalinkURL   string    `json:"permalink_url"`
			URI            string    `json:"uri"`
			Urn            string    `json:"urn"`
			Username       string    `json:"username"`
			Verified       bool      `json:"verified"`
			City           string    `json:"city"`
			CountryCode    string    `json:"country_code"`
			Badges         struct {
				Pro          bool `json:"pro"`
				ProUnlimited bool `json:"pro_unlimited"`
				Verified     bool `json:"verified"`
			} `json:"badges"`
			StationUrn       string `json:"station_urn"`
			StationPermalink string `json:"station_permalink"`
		} `json:"user,omitempty"`
		PublisherMetadata0 struct {
			ID              int    `json:"id"`
			Urn             string `json:"urn"`
			Artist          string `json:"artist"`
			AlbumTitle      string `json:"album_title"`
			ContainsMusic   bool   `json:"contains_music"`
			UpcOrEan        string `json:"upc_or_ean"`
			Isrc            string `json:"isrc"`
			Explicit        bool   `json:"explicit"`
			PLine           string `json:"p_line"`
			PLineForDisplay string `json:"p_line_for_display"`
			WriterComposer  string `json:"writer_composer"`
			ReleaseTitle    string `json:"release_title"`
		} `json:"publisher_metadata,omitempty"`
		PublisherMetadata1 struct {
			ID              int    `json:"id"`
			Urn             string `json:"urn"`
			Artist          string `json:"artist"`
			AlbumTitle      string `json:"album_title"`
			ContainsMusic   bool   `json:"contains_music"`
			UpcOrEan        string `json:"upc_or_ean"`
			Isrc            string `json:"isrc"`
			Explicit        bool   `json:"explicit"`
			PLine           string `json:"p_line"`
			PLineForDisplay string `json:"p_line_for_display"`
			CLine           string `json:"c_line"`
			CLineForDisplay string `json:"c_line_for_display"`
			ReleaseTitle    string `json:"release_title"`
		} `json:"publisher_metadata,omitempty"`
		PublisherMetadata2 struct {
			ID              int    `json:"id"`
			Urn             string `json:"urn"`
			Artist          string `json:"artist"`
			AlbumTitle      string `json:"album_title"`
			ContainsMusic   bool   `json:"contains_music"`
			UpcOrEan        string `json:"upc_or_ean"`
			Isrc            string `json:"isrc"`
			Explicit        bool   `json:"explicit"`
			PLine           string `json:"p_line"`
			PLineForDisplay string `json:"p_line_for_display"`
			WriterComposer  string `json:"writer_composer"`
			ReleaseTitle    string `json:"release_title"`
		} `json:"publisher_metadata,omitempty"`
		PublisherMetadata3 struct {
			ID            int    `json:"id"`
			Urn           string `json:"urn"`
			Artist        string `json:"artist"`
			ContainsMusic bool   `json:"contains_music"`
			Isrc          string `json:"isrc"`
		} `json:"publisher_metadata,omitempty"`
	} `json:"tracks"`
	TrackCount int `json:"track_count"`
}

// TrackResponse is the response from the SoundCloud API
type TrackResponse []struct {
	ArtworkURL        string    `json:"artwork_url"`
	Caption           any       `json:"caption"`
	Commentable       bool      `json:"commentable"`
	CommentCount      int       `json:"comment_count"`
	CreatedAt         time.Time `json:"created_at"`
	Description       string    `json:"description"`
	Downloadable      bool      `json:"downloadable"`
	DownloadCount     int       `json:"download_count"`
	Duration          int       `json:"duration"`
	FullDuration      int       `json:"full_duration"`
	EmbeddableBy      string    `json:"embeddable_by"`
	Genre             string    `json:"genre"`
	HasDownloadsLeft  bool      `json:"has_downloads_left"`
	ID                int       `json:"id"`
	Kind              string    `json:"kind"`
	LabelName         any       `json:"label_name"`
	LastModified      time.Time `json:"last_modified"`
	License           string    `json:"license"`
	LikesCount        int       `json:"likes_count"`
	Permalink         string    `json:"permalink"`
	PermalinkURL      string    `json:"permalink_url"`
	PlaybackCount     int       `json:"playback_count"`
	Public            bool      `json:"public"`
	PublisherMetadata struct {
		ID            int    `json:"id"`
		Urn           string `json:"urn"`
		Artist        string `json:"artist"`
		ContainsMusic bool   `json:"contains_music"`
		Isrc          string `json:"isrc"`
	} `json:"publisher_metadata,omitempty"`
	PurchaseTitle any       `json:"purchase_title"`
	PurchaseURL   any       `json:"purchase_url"`
	ReleaseDate   any       `json:"release_date"`
	RepostsCount  int       `json:"reposts_count"`
	SecretToken   any       `json:"secret_token"`
	Sharing       string    `json:"sharing"`
	State         string    `json:"state"`
	Streamable    bool      `json:"streamable"`
	TagList       string    `json:"tag_list"`
	Title         string    `json:"title"`
	TrackFormat   string    `json:"track_format"`
	URI           string    `json:"uri"`
	Urn           string    `json:"urn"`
	UserID        int       `json:"user_id"`
	Visuals       any       `json:"visuals"`
	WaveformURL   string    `json:"waveform_url"`
	DisplayDate   time.Time `json:"display_date"`
	Media         struct {
		Transcodings []struct {
			URL      string `json:"url"`
			Preset   string `json:"preset"`
			Duration int    `json:"duration"`
			Snipped  bool   `json:"snipped"`
			Format   struct {
				Protocol string `json:"protocol"`
				MimeType string `json:"mime_type"`
			} `json:"format"`
			Quality string `json:"quality"`
		} `json:"transcodings"`
	} `json:"media"`
	StationUrn         string `json:"station_urn"`
	StationPermalink   string `json:"station_permalink"`
	TrackAuthorization string `json:"track_authorization"`
	MonetizationModel  string `json:"monetization_model"`
	Policy             string `json:"policy"`
	User               struct {
		AvatarURL      string    `json:"avatar_url"`
		FirstName      string    `json:"first_name"`
		FollowersCount int       `json:"followers_count"`
		FullName       string    `json:"full_name"`
		ID             int       `json:"id"`
		Kind           string    `json:"kind"`
		LastModified   time.Time `json:"last_modified"`
		LastName       string    `json:"last_name"`
		Permalink      string    `json:"permalink"`
		PermalinkURL   string    `json:"permalink_url"`
		URI            string    `json:"uri"`
		Urn            string    `json:"urn"`
		Username       string    `json:"username"`
		Verified       bool      `json:"verified"`
		City           string    `json:"city"`
		CountryCode    any       `json:"country_code"`
		Badges         struct {
			Pro          bool `json:"pro"`
			ProUnlimited bool `json:"pro_unlimited"`
			Verified     bool `json:"verified"`
		} `json:"badges"`
		StationUrn       string `json:"station_urn"`
		StationPermalink string `json:"station_permalink"`
	} `json:"user"`
	PublisherMetadata0 struct {
		ID              int    `json:"id"`
		Urn             string `json:"urn"`
		Artist          string `json:"artist"`
		AlbumTitle      string `json:"album_title"`
		ContainsMusic   bool   `json:"contains_music"`
		UpcOrEan        string `json:"upc_or_ean"`
		Isrc            string `json:"isrc"`
		Explicit        bool   `json:"explicit"`
		PLine           string `json:"p_line"`
		PLineForDisplay string `json:"p_line_for_display"`
		CLine           string `json:"c_line"`
		CLineForDisplay string `json:"c_line_for_display"`
		ReleaseTitle    string `json:"release_title"`
	} `json:"publisher_metadata,omitempty"`
	PublisherMetadata1 struct {
		ID              int    `json:"id"`
		Urn             string `json:"urn"`
		Artist          string `json:"artist"`
		AlbumTitle      string `json:"album_title"`
		ContainsMusic   bool   `json:"contains_music"`
		UpcOrEan        string `json:"upc_or_ean"`
		Isrc            string `json:"isrc"`
		Explicit        bool   `json:"explicit"`
		PLine           string `json:"p_line"`
		PLineForDisplay string `json:"p_line_for_display"`
		CLine           string `json:"c_line"`
		CLineForDisplay string `json:"c_line_for_display"`
		ReleaseTitle    string `json:"release_title"`
	} `json:"publisher_metadata,omitempty"`
	PublisherMetadata2 struct {
		ID              int    `json:"id"`
		Urn             string `json:"urn"`
		Artist          string `json:"artist"`
		AlbumTitle      string `json:"album_title"`
		ContainsMusic   bool   `json:"contains_music"`
		UpcOrEan        string `json:"upc_or_ean"`
		Isrc            string `json:"isrc"`
		Explicit        bool   `json:"explicit"`
		PLine           string `json:"p_line"`
		PLineForDisplay string `json:"p_line_for_display"`
		CLine           string `json:"c_line"`
		CLineForDisplay string `json:"c_line_for_display"`
		ReleaseTitle    string `json:"release_title"`
	} `json:"publisher_metadata,omitempty"`
	PublisherMetadata3 struct {
		ID              int    `json:"id"`
		Urn             string `json:"urn"`
		Artist          string `json:"artist"`
		AlbumTitle      string `json:"album_title"`
		ContainsMusic   bool   `json:"contains_music"`
		UpcOrEan        string `json:"upc_or_ean"`
		Isrc            string `json:"isrc"`
		Explicit        bool   `json:"explicit"`
		PLine           string `json:"p_line"`
		PLineForDisplay string `json:"p_line_for_display"`
		CLine           string `json:"c_line"`
		CLineForDisplay string `json:"c_line_for_display"`
		ReleaseTitle    string `json:"release_title"`
	} `json:"publisher_metadata,omitempty"`
	PublisherMetadata4 struct {
		ID              int    `json:"id"`
		Urn             string `json:"urn"`
		Artist          string `json:"artist"`
		AlbumTitle      string `json:"album_title"`
		ContainsMusic   bool   `json:"contains_music"`
		UpcOrEan        string `json:"upc_or_ean"`
		Isrc            string `json:"isrc"`
		Explicit        bool   `json:"explicit"`
		PLine           string `json:"p_line"`
		PLineForDisplay string `json:"p_line_for_display"`
		CLine           string `json:"c_line"`
		CLineForDisplay string `json:"c_line_for_display"`
		ReleaseTitle    string `json:"release_title"`
	} `json:"publisher_metadata,omitempty"`
	PublisherMetadata5 struct {
		ID              int    `json:"id"`
		Urn             string `json:"urn"`
		Artist          string `json:"artist"`
		AlbumTitle      string `json:"album_title"`
		ContainsMusic   bool   `json:"contains_music"`
		UpcOrEan        string `json:"upc_or_ean"`
		Isrc            string `json:"isrc"`
		Explicit        bool   `json:"explicit"`
		PLine           string `json:"p_line"`
		PLineForDisplay string `json:"p_line_for_display"`
		CLine           string `json:"c_line"`
		CLineForDisplay string `json:"c_line_for_display"`
		ReleaseTitle    string `json:"release_title"`
	} `json:"publisher_metadata,omitempty"`
}
