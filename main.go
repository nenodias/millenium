package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nenodias/millenium/config"
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

	router.HandleFunc("/api/modelo/", modeloController.FindMany).Methods("GET")
	router.HandleFunc("/api/modelo/", modeloController.Save).Methods("POST")
	router.HandleFunc("/api/modelo/{id}", modeloController.FindOne).Methods("GET")
	router.HandleFunc("/api/modelo/{id}", modeloController.DeleteOne).Methods("DELETE")
	router.HandleFunc("/api/modelo/{id}", modeloController.Update).Methods("PUT")

	router.HandleFunc("/api/tecnico/", tecnicoController.FindMany).Methods("GET")
	router.HandleFunc("/api/tecnico/", tecnicoController.Save).Methods("POST")
	router.HandleFunc("/api/tecnico/{id}", tecnicoController.FindOne).Methods("GET")
	router.HandleFunc("/api/tecnico/{id}", tecnicoController.DeleteOne).Methods("DELETE")
	router.HandleFunc("/api/tecnico/{id}", tecnicoController.Update).Methods("PUT")

	router.HandleFunc("/api/veiculo/", veiculoController.FindMany).Methods("GET")
	router.HandleFunc("/api/veiculo/", veiculoController.Save).Methods("POST")
	router.HandleFunc("/api/veiculo/{id}", veiculoController.FindOne).Methods("GET")
	router.HandleFunc("/api/veiculo/{id}", veiculoController.DeleteOne).Methods("DELETE")
	router.HandleFunc("/api/veiculo/{id}", veiculoController.Update).Methods("PUT")

	srv := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8080",
	}
	log.Info().Msg("Listening on port :8080")
	log.Error().Msg(srv.ListenAndServe().Error())
}
