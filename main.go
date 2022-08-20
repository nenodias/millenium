package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nenodias/millenium/config"
	tecnicoHandlers "github.com/nenodias/millenium/handlers/tecnico"
	database "github.com/nenodias/millenium/repositories"
	"github.com/nenodias/millenium/repositories/models"
	"github.com/rs/zerolog/log"
)

func main() {
	config.Init()
	database.Init()
	engine := database.GetEngine()
	//var service tDomain.TecnicoService = models.NewTecnicoService(engine)
	service := models.NewTecnicoService(engine)
	controller := tecnicoHandlers.NewTecnicoController(&service)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/tecnico/", controller.FindMany).Methods("GET")
	router.HandleFunc("/api/tecnico/", controller.Save).Methods("POST")
	router.HandleFunc("/api/tecnico/{id}", controller.FindOne).Methods("GET")
	router.HandleFunc("/api/tecnico/{id}", controller.DeleteOne).Methods("DELETE")
	router.HandleFunc("/api/tecnico/{id}", controller.Update).Methods("PUT")

	srv := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8080",
	}

	log.Error().Msg(srv.ListenAndServe().Error())
}
