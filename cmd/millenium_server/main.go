package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nenodias/millenium/internal/configs"
	auth "github.com/nenodias/millenium/internal/core/domain/auth"
	appHandlers "github.com/nenodias/millenium/internal/handlers"
	authHandlers "github.com/nenodias/millenium/internal/handlers/auth"
	clienteHandlers "github.com/nenodias/millenium/internal/handlers/cliente"
	falhaHandlers "github.com/nenodias/millenium/internal/handlers/falha"
	historicoHandlers "github.com/nenodias/millenium/internal/handlers/historico"
	lembreteHandlers "github.com/nenodias/millenium/internal/handlers/lembrete"
	modeloHandlers "github.com/nenodias/millenium/internal/handlers/modelo"
	montadoraHandlers "github.com/nenodias/millenium/internal/handlers/montadora"
	pecaHandlers "github.com/nenodias/millenium/internal/handlers/peca"
	servicoHandlers "github.com/nenodias/millenium/internal/handlers/servico"
	tecnicoHandlers "github.com/nenodias/millenium/internal/handlers/tecnico"
	veiculoHandlers "github.com/nenodias/millenium/internal/handlers/veiculo"
	database "github.com/nenodias/millenium/internal/repositories"
	clienteModels "github.com/nenodias/millenium/internal/repositories/models/cliente"
	falhaModels "github.com/nenodias/millenium/internal/repositories/models/falha"
	historicoModels "github.com/nenodias/millenium/internal/repositories/models/historico"
	lembreteModels "github.com/nenodias/millenium/internal/repositories/models/lembrete"
	modeloModels "github.com/nenodias/millenium/internal/repositories/models/modelo"
	montadoraModels "github.com/nenodias/millenium/internal/repositories/models/montadora"
	pecaModels "github.com/nenodias/millenium/internal/repositories/models/peca"
	servicoModels "github.com/nenodias/millenium/internal/repositories/models/servico"
	tecnicoModels "github.com/nenodias/millenium/internal/repositories/models/tecnico"
	veiculoModels "github.com/nenodias/millenium/internal/repositories/models/veiculo"
	"github.com/rs/zerolog/log"
)

func main() {
	configs.Init()
	auth.Init()
	database.Init()
	engine := database.GetEngine()
	router := chi.NewRouter()

	router.Post("/api/auth/", authHandlers.Authenticate)

	clienteService := clienteModels.NewService(engine)
	clienteController := clienteHandlers.NewController(&clienteService)
	MappingApi(router, "cliente", clienteController)

	falhaService := falhaModels.NewService(engine)
	falhaController := falhaHandlers.NewController(&falhaService)
	MappingApi(router, "falha", falhaController)

	historicoService := historicoModels.NewService(engine)
	historicoController := historicoHandlers.NewController(&historicoService)
	MappingApi(router, "historico", historicoController)
	router.Get("/api/historico/report/{id}", historicoController.GetReport)

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

	port := configs.GetEnv("SERVER_PORT", "8080")
	handler := appHandlers.CORSHandler{Inner: router, Origin: configs.GetEnv("ALLOW_ORIGIN", "*")}

	srv := &http.Server{
		Handler: handler,
		Addr:    "0.0.0.0:" + port,
	}
	log.Info().Msg("Listening on port :" + port)
	log.Error().Msg(srv.ListenAndServe().Error())
}

func MappingApi(router *chi.Mux, context string, controller appHandlers.CrudAPI) {
	router.Route("/api/"+context, func(r chi.Router) {
		r.Get("/", authHandlers.Middleware(controller.FindMany))
		r.Post("/", authHandlers.Middleware(controller.Save))
		r.Get("/{id}", authHandlers.Middleware(controller.FindOne))
		r.Delete("/{id}", authHandlers.Middleware(controller.DeleteOne))
		r.Put("/{id}", authHandlers.Middleware(controller.Update))
	})
}
