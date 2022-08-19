package tecnico

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	domain "github.com/nenodias/millenium/core/domain/tecnico"
)

type TecnicoController struct {
	Service *domain.TecnicoService
}

func NewTecnicoController(service *domain.TecnicoService) *TecnicoController {
	return &TecnicoController{
		Service: service,
	}
}

func (t *TecnicoController) FindMany(w http.ResponseWriter, r *http.Request) {
	//t.Service.FindMany()
}

func (t *TecnicoController) Save(w http.ResponseWriter, r *http.Request) {
	model := new(domain.Tecnico)
	err := json.NewDecoder(r.Body).Decode(model)
	if err != nil {
		log.Error().Msg(err.Error())
		w.WriteHeader(400)
	} else {
		success, err := (*t.Service).Save(model)
		if err != nil {
			log.Error().Msg(err.Error())
			w.WriteHeader(500)
		}
		if success {
			w.WriteHeader(200)
			err := json.NewEncoder(w).Encode(model)
			if err != nil {
				log.Error().Msg(err.Error())
				w.WriteHeader(500)
			}
		} else {
			w.WriteHeader(400)
		}
	}
}

func (t *TecnicoController) FindOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	txtId := params["id"]
	id, err := strconv.ParseInt(txtId, 10, 64)
	if err != nil {
		log.Error().Msg(err.Error())
		w.WriteHeader(400)
	} else {
		model, err := (*t.Service).FindOne(id)
		if err != nil {
			log.Error().Msg(err.Error())
			w.WriteHeader(500)
		}
		if model != nil {
			w.WriteHeader(200)
			err := json.NewEncoder(w).Encode(model)
			if err != nil {
				log.Error().Msg(err.Error())
				w.WriteHeader(500)
			}
		} else {
			w.WriteHeader(204)
		}
	}
}

func (t *TecnicoController) DeleteOne(w http.ResponseWriter, r *http.Request) {
	//t.Service.DeleteOne()
}

func (t *TecnicoController) Update(w http.ResponseWriter, r *http.Request) {
	//t.Service.FindOne()
	//t.Service.Save()
}
