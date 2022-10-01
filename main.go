package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nenodias/millenium/config"
	auth "github.com/nenodias/millenium/core/domain/auth"
	"github.com/nenodias/millenium/handlers"
	authHandlers "github.com/nenodias/millenium/handlers/auth"
	clienteHandlers "github.com/nenodias/millenium/handlers/cliente"
	falhaHandlers "github.com/nenodias/millenium/handlers/falha"
	historicoHandlers "github.com/nenodias/millenium/handlers/historico"
	lembreteHandlers "github.com/nenodias/millenium/handlers/lembrete"
	modeloHandlers "github.com/nenodias/millenium/handlers/modelo"
	montadoraHandlers "github.com/nenodias/millenium/handlers/montadora"
	pecaHandlers "github.com/nenodias/millenium/handlers/peca"
	servicoHandlers "github.com/nenodias/millenium/handlers/servico"
	tecnicoHandlers "github.com/nenodias/millenium/handlers/tecnico"
	veiculoHandlers "github.com/nenodias/millenium/handlers/veiculo"
	database "github.com/nenodias/millenium/repositories"
	clienteModels "github.com/nenodias/millenium/repositories/models/cliente"
	falhaModels "github.com/nenodias/millenium/repositories/models/falha"
	historicoModels "github.com/nenodias/millenium/repositories/models/historico"
	lembreteModels "github.com/nenodias/millenium/repositories/models/lembrete"
	modeloModels "github.com/nenodias/millenium/repositories/models/modelo"
	montadoraModels "github.com/nenodias/millenium/repositories/models/montadora"
	pecaModels "github.com/nenodias/millenium/repositories/models/peca"
	servicoModels "github.com/nenodias/millenium/repositories/models/servico"
	tecnicoModels "github.com/nenodias/millenium/repositories/models/tecnico"
	veiculoModels "github.com/nenodias/millenium/repositories/models/veiculo"
	"github.com/rs/zerolog/log"
)

func main() {
	config.Init()
	auth.Init()
	database.Init()
	engine := database.GetEngine()
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/auth/", authHandlers.Authenticate).Methods("POST")

	clienteService := clienteModels.NewService(engine)
	clienteController := clienteHandlers.NewController(&clienteService)
	MappingApi(router, "cliente", clienteController)

	falhaService := falhaModels.NewService(engine)
	falhaController := falhaHandlers.NewController(&falhaService)
	MappingApi(router, "falha", falhaController)

	historicoService := historicoModels.NewService(engine)
	historicoController := historicoHandlers.NewController(&historicoService)
	MappingApi(router, "historico", historicoController)
	router.HandleFunc("/api/historico/report/{id}", historicoController.GetReport).Methods("GET")

	lembreteService := lembreteModels.NewService(engine)
	lembreteController := lembreteHandlers.NewController(&lembreteService)
	MappingApi(router, "lembrete", lembreteController)

	modeloService := modeloModels.NewService(engine)
	modeloController := modeloHandlers.NewController(&modeloService)
	MappingApi(router, "modelo", modeloController)

	montadoraService := montadoraModels.NewService(engine)
	montadoraController := montadoraHandlers.NewController(&montadoraService)
	MappingApi(router, "montadora", montadoraController)

	pecaService := pecaModels.NewService(engine)
	pecaController := pecaHandlers.NewController(&pecaService)
	MappingApi(router, "peca", pecaController)

	servicoService := servicoModels.NewService(engine)
	servicoController := servicoHandlers.NewController(&servicoService)
	MappingApi(router, "servico", servicoController)

	tecnicoService := tecnicoModels.NewService(engine)
	tecnicoController := tecnicoHandlers.NewController(&tecnicoService)
	MappingApi(router, "tecnico", tecnicoController)

	veiculoService := veiculoModels.NewService(engine)
	veiculoController := veiculoHandlers.NewController(&veiculoService)
	MappingApi(router, "veiculo", veiculoController)

	srv := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8080",
	}
	log.Info().Msg("Listening on port :8080")
	log.Error().Msg(srv.ListenAndServe().Error())
}

func MappingApi(router *mux.Router, context string, controller handlers.CrudAPI) {
	router.HandleFunc("/api/"+context+"/", authHandlers.Middleware(controller.FindMany)).Methods("GET")
	router.HandleFunc("/api/"+context+"/", authHandlers.Middleware(controller.Save)).Methods("POST")
	router.HandleFunc("/api/"+context+"/{id}", authHandlers.Middleware(controller.FindOne)).Methods("GET")
	router.HandleFunc("/api/"+context+"/{id}", authHandlers.Middleware(controller.DeleteOne)).Methods("DELETE")
	router.HandleFunc("/api/"+context+"/{id}", authHandlers.Middleware(controller.Update)).Methods("PUT")
}
