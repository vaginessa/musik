package api

// Routes - define all routes for the API
func (s *Server) Routes() {
	server := s.Echo
	// group for REST API functions
	api := server.Group("/api/v1")

	// Health
	api.GET("/health", s.HealthCheck)

	// Music
	api.GET("/music/title", s.GetTitle)
	api.GET("/music/artwork", s.GetArtworkURL)
	api.GET("/music/stream", s.GetStreamURL)
	api.GET("/music/random", s.GetRandomSong)

}
