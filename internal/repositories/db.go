package repositories

import (
	"fmt"

	"github.com/rs/zerolog/log"

	_ "github.com/lib/pq"
	"github.com/nenodias/millenium/configs"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func Init() {
	var err error
	username := configs.GetEnv(configs.POSTGRES_USER, "postgres")
	password := configs.GetEnv(configs.POSTGRES_PASSWORD, "postgres")
	host := configs.GetEnv(configs.POSTGRESQL_SERVICE_HOST, "localhost")
	port := configs.GetEnv(configs.POSTGRES_PORT, "5432")
	database := configs.GetEnv(configs.POSTGRES_DB, "carrit")
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

func GetEngine() *DatabaseEngine {
	return &DatabaseEngine{*engine}
}
