package sentry

import (
	"log"
	"time"

	"github.com/imthaghost/musik/backend/config"
	"github.com/imthaghost/musik/backend/internal/errors"

	"github.com/getsentry/sentry-go"
)

// Sentry error logging service
type Sentry struct {
	Config *config.Config
}

// Setup will start the Sentry service
func (s *Sentry) Setup() {
	env := s.Config.General.AppEnv
	// set debug mode
	debug := false

	// if local, set debug mode true
	if env == "dev" {
		debug = true
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn: s.Config.SentryConfig.DSN,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		SampleRate:       1.0,
		TracesSampleRate: 1.0,
		Debug:            debug,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
}

// TearDown will do any final cleanup, flushing of buffers
func (s *Sentry) TearDown() {
	// flush the buffer before the program exists
	sentry.Flush(3 * time.Second)
}

// Report will report a new error to Sentry
func (s *Sentry) Report(err error) {
	sentry.CaptureException(err)
}

// NewService will create a new Sentry service
func NewService(cfg *config.Config) errors.Service {
	monitor := &Sentry{
		Config: cfg,
	}
	// configure
	monitor.Setup()

	return monitor
}
