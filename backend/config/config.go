// Config

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	DEV     = "dev"
	STAGING = "staging"
	PROD    = "prod"
)

// New ...
type New struct{}

// Load will load the config
func (n *New) Load() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
}

// Get will return the default config
func (n *New) Get() Config {
	// get the config
	return Config{
		General:      getGeneralConfig(),
		SentryConfig: getSentryConfig(),
	}
}

func getGeneralConfig() GeneralConfig {
	// default
	config := GeneralConfig{
		AppEnv: os.Getenv("APP_ENV"),
	}

	return config
}

// getSentryConfig returns the Sentry config
func getSentryConfig() SentryConfig {
	env := getGeneralConfig().AppEnv

	config := SentryConfig{
		DSN: "https://c3f03af62f043fb1e70a8ea8c951dd2d@o4506542909358080.ingest.sentry.io/4506542910865408",
	}

	if env == PROD {
		config.DSN = "https://e60551da90863c28355cad718c84902a@o4506542909358080.ingest.sentry.io/4506542917484544"
	}

	return config
}
