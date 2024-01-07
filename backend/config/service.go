package config

// Service is an interface that defines the functions needed to implement a Config Service.
type Service interface {
	// Load will do any config setup (like load env vars)
	Load()
	// Get will get the config
	Get() Config
}

// Config is a service that is designed to provide various configuration to the rest of the application.
type Config struct {
	General      GeneralConfig
	SentryConfig SentryConfig
}

// GeneralConfig contains general information that the service needs to run.
type GeneralConfig struct {
	AppEnv string // the environment that the application is running in (dev, prod, etc)
}

// SentryConfig contains information that allows us to interact with the Sentry API
type SentryConfig struct {
	DSN string
}
