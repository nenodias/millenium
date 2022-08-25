package modelo

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	core "github.com/nenodias/millenium/core/domain"
	"github.com/nenodias/millenium/core/domain/utils"
	"github.com/rs/zerolog/log"

	domain "github.com/nenodias/millenium/core/domain/modelo"
)

type ModeloController struct {
	Service *domain.ModeloService
}

func NewModeloController(service *domain.ModeloService) *ModeloController {
	return &ModeloController{
		Service: service,
	}
}

func (t *ModeloController) FindMany(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	page := utils.StringToInt(query.Get("page"), utils.DEFAULT_PAGE)
	size := utils.StringToInt(query.Get("size"), utils.DEFAULT_SIZE)
	sortColumn := utils.StringNormalized(query.Get("sortColumn"), "id")
	sortDirection := core.GetSortDirection(query.Get("sortDirection"))
	nome := utils.StringNormalized(query.Get("nome"), "")
	idModelo := utils.StringToInt64(query.Get("id_monta"), 0)
	filter := domain.ModeloFilter{
		Nome:     nome,
		IdModelo: idModelo,
		Pageable: core.Pageable{
			PageSize: size, PageNumber: page,
			Sort: core.SortRequest{
				SortColumn: sortColumn, SortDirection: sortDirection,
			},
		},
	}
	response, err := (*t.Service).FindMany(&filter)
	if err != nil {
		log.Error().Msg(err.Error())
		w.WriteHeader(500)
	} else {
		utils.WriteJson(response, w, 200, 500)
	}
}

func (t *ModeloController) Save(w http.ResponseWriter, r *http.Request) {
	model := new(domain.Modelo)
	err := json.NewDecoder(r.Body).Decode(model)
	if err != nil {
		log.Error().Msg(err.Error())
		w.WriteHeader(400)
		return
	}
	success, err := (*t.Service).Save(model)
	if err != nil {
		log.Error().Msg(err.Error())
		w.WriteHeader(500)
		return
	}
	if success {
		utils.WriteJson(model, w, 200, 500)
	} else {
		w.WriteHeader(400)
	}
}

func (t *ModeloController) FindOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	txtId := params["id"]
	id, err := strconv.ParseInt(txtId, 10, 64)
	if err != nil {
		log.Error().Msg(err.Error())
		w.WriteHeader(400)
		return
	}
	model, err := (*t.Service).FindOne(id)
	if err != nil {
		log.Error().Msg(err.Error())
		w.WriteHeader(500)
		return
	}
	if model != nil {
		utils.WriteJson(model, w, 200, 500)
	} else {
		w.WriteHeader(204)
	}
}

func (t *ModeloController) DeleteOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	txtId := params["id"]
	id, err := strconv.ParseInt(txtId, 10, 64)
	if err != nil {
		log.Error().Msg(err.Error())
		w.WriteHeader(400)
		return
	}

	_, err = (*t.Service).DeleteOne(id)
	if err != nil {
		log.Error().Msg(err.Error())
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(204)
}

func (t *ModeloController) Update(w http.ResponseWriter, r *http.Request) {
	model := new(domain.Modelo)
	params := mux.Vars(r)
	txtId := params["id"]
	id, err := strconv.ParseInt(txtId, 10, 64)
	if err != nil {
		log.Error().Msg(err.Error())
		w.WriteHeader(400)
		return
	}
	record, err := (*t.Service).FindOne(id)
	if err != nil {
		log.Error().Msg(err.Error())
		w.WriteHeader(500)
		return
	}
	if record != nil {
		err := json.NewDecoder(r.Body).Decode(model)
		if err != nil {
			log.Error().Msg(err.Error())
			w.WriteHeader(400)
			return
		}
		success, err := (*t.Service).Save(model)
		if err != nil {
			log.Error().Msg(err.Error())
			w.WriteHeader(500)
			return
		}
		if success {
			utils.WriteJson(model, w, 200, 500)
		} else {
			w.WriteHeader(400)
		}
	} else {
		w.WriteHeader(404)
	}
}
