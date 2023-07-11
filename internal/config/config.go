package config

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/joho/godotenv"
)

func logStartConfig() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.With().Str("app", "millenium").Caller().Logger()
	debug := GetEnv("DEBUG", "false")
	if debug == "true" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Debug().Msg("Debug enabled.")
	}
}

func Init() {
	logStartConfig()
	err := godotenv.Load()
	if err != nil {
		log.Info().Msg("Error loading .env file")
	}
}

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
