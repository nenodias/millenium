package main

import (
	"github.com/rs/zerolog/log"

	"github.com/nenodias/millenium/config"
	"github.com/nenodias/millenium/database"
	"github.com/nenodias/millenium/database/models"
)

func main() {
	config.Init()
	database.Init()
	engine := database.GetEngine()
	p := &models.Peca{Valor: 10.0, Descricao: "Peça Nova"}
	exists, err := engine.ID(6).Get(p)
	if err != nil {
		log.Error().Msg(err.Error())
	} else {
		log.Info().Msgf("Registro encontrado: %v", exists)
		log.Info().Msgf("Registro: %v", p)
	}
	/*
		ret, err := engine.Insert(p)
		if err != nil {
			log.Error().Msg(err.Error())
		} else {
			log.Info().Msgf("Linhas inseridas: %d", ret)
			log.Info().Msgf("Id inserido: %d", p.Id)
		}
	*/
	log.Info().Msg("Hello World")
}
