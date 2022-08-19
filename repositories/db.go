package repositories

import (
	"fmt"

	"github.com/rs/zerolog/log"

	_ "github.com/lib/pq"
	"github.com/nenodias/millenium/config"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func Init() {
	var err error
	username := config.GetEnv("POSTGRES_USER", "postgres")
	password := config.GetEnv("POSTGRES_PASSWORD", "postgres")
	host := config.GetEnv("POSTGRESQL_SERVICE_HOST", "localhost")
	port := config.GetEnv("POSTGRES_PORT", "5432")
	database := config.GetEnv("POSTGRES_DB", "carrit")
	urlConnection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)
	engine, err = xorm.NewEngine("postgres", urlConnection)
	if err != nil {
		log.Fatal().Stack().Err(err).Str("service", "database").Msgf("Cannot start connection with database")
	}
	engine.SetLogger(&CurrentLogger{logger: &log.Logger, showSQL: true})
}

type DatabaseEngine struct {
	xorm.Engine
}

func GetEngine() *xorm.Engine {
	return engine
}
