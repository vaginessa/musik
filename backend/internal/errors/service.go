package errors

// Service represents an error monitoring service
type Service interface {
	// Setup will perform any configuration
	Setup()
	// Report ...
	Report(err error)
	// TearDown ...
	TearDown()
}
