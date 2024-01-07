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
		DSN: "https://9a47c172155cbcfc50fb6aa4deea00a5@o4506205650681856.ingest.sentry.io/4506205651337216",
	}

	if env == STAGING {
		config.DSN = "https://9a47c172155cbcfc50fb6aa4deea00a5@o4506205650681856.ingest.sentry.io/4506205651337216"
	}

	if env == PROD {
		config.DSN = "https://9a47c172155cbcfc50fb6aa4deea00a5@o4506205650681856.ingest.sentry.io/4506205651337216"
	}

	return config
}
