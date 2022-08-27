package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nenodias/millenium/config"
	"github.com/nenodias/millenium/handlers"
	modeloHandlers "github.com/nenodias/millenium/handlers/modelo"
	tecnicoHandlers "github.com/nenodias/millenium/handlers/tecnico"
	veiculoHandlers "github.com/nenodias/millenium/handlers/veiculo"
	database "github.com/nenodias/millenium/repositories"
	modeloModels "github.com/nenodias/millenium/repositories/models/modelo"
	tecnicoModels "github.com/nenodias/millenium/repositories/models/tecnico"
	veiculoModels "github.com/nenodias/millenium/repositories/models/veiculo"
	"github.com/rs/zerolog/log"
)

func main() {
	config.Init()
	database.Init()
	engine := database.GetEngine()

	modeloService := modeloModels.NewTecnicoService(engine)
	modeloController := modeloHandlers.NewModeloController(&modeloService)

	tecnicoService := tecnicoModels.NewTecnicoService(engine)
	tecnicoController := tecnicoHandlers.NewTecnicoController(&tecnicoService)

	veiculoService := veiculoModels.NewVeiculoService(engine)
	veiculoController := veiculoHandlers.NewVeiculoController(&veiculoService)

	router := mux.NewRouter().StrictSlash(true)
	MappingApi(router, "tecnico", tecnicoController)
	MappingApi(router, "modelo", modeloController)
	MappingApi(router, "veiculo", veiculoController)

	srv := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8080",
	}
	log.Info().Msg("Listening on port :8080")
	log.Error().Msg(srv.ListenAndServe().Error())
}

func MappingApi(router *mux.Router, context string, controller handlers.CrudAPI) {
	router.HandleFunc("/api/"+context+"/", controller.FindMany).Methods("GET")
	router.HandleFunc("/api/"+context+"/", controller.Save).Methods("POST")
	router.HandleFunc("/api/"+context+"/{id}", controller.FindOne).Methods("GET")
	router.HandleFunc("/api/"+context+"/{id}", controller.DeleteOne).Methods("DELETE")
	router.HandleFunc("/api/"+context+"/{id}", controller.Update).Methods("PUT")
}
