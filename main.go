package main

import (
	"github.com/rs/zerolog/log"

	"github.com/nenodias/millenium/config"
	"github.com/nenodias/millenium/database"
)

func main() {
	config.Init()
	database.Init()
	log.Info().Msg("Hello World")
}
