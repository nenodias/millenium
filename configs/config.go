package configs

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/joho/godotenv"
)

const (
	SERVER_PORT   = "SERVER_PORT"
	SERVER_SECRET = "SERVER_SECRET"
	ALLOW_ORIGIN  = "ALLOW_ORIGIN"

	REPORT_COMPANY_NAME      = "REPORT_COMPANY_NAME"
	REPORT_COMPANY_ADDRESS   = "REPORT_COMPANY_ADDRESS"
	REPORT_COMPANY_PHONE     = "REPORT_COMPANY_PHONE"
	REPORT_COMPANY_CELLPHONE = "REPORT_COMPANY_CELLPHONE"
	REPORT_COMPANY_EMAIL     = "REPORT_COMPANY_EMAIL"

	POSTGRES_USER           = "POSTGRES_USER"
	POSTGRES_PASSWORD       = "POSTGRES_PASSWORD"
	POSTGRESQL_SERVICE_HOST = "POSTGRESQL_SERVICE_HOST"
	POSTGRES_PORT           = "POSTGRES_PORT"
	POSTGRES_DB             = "POSTGRES_DB"

	USER_DEFAULT = "USER_DEFAULT"
	PASS_DEFAULT = "PASS_DEFAULT"
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
